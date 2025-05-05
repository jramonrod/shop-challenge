package services

import (
	"fmt"
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/repositories"
)

// RestrictionsService coordinates all business logic around product restrictions
type RestrictionsService struct {
	RestrictionsRepo repositories.RestrictionsRepo
	PartsRepo        repositories.PartsRepo
}

// NewRestrictionsService constructs a service with its repository dependency
func NewRestrictionsService(cr repositories.RestrictionsRepo, pr repositories.PartsRepo) *RestrictionsService {
	return &RestrictionsService{cr, pr}
}

// ValidateSelections checks a set of selected part IDs against defined restrictions for a given product. It returns a
// list of violations where the selections conflict with rules such as "forbid" and "only".
//
// Restrictions are fetched from the RestrictionsRepo and matched by their principal part ID. The method de-duplicates
// selected parts to avoid redundant validations and uses the part-to-category map to validate "Only" rules within categories.
//
// It is finally good to note that we make an extra DB call in this method for clarity and maintainability.
func (rs *RestrictionsService) ValidateSelections(
	productID string,
	selections domain.PartsSelection,
) ([]domain.Violation, error) {
	rulesByPrincipal, err := rs.RestrictionsRepo.GetRestrictionsForParts(productID, selections)
	if err != nil {
		return nil, err
	}

	violations := make([]domain.Violation, 0)
	partCategoryMap, err := rs.PartsRepo.GetCategoryMap(productID)
	if err != nil {
		return nil, err
	}

	// this makes sure the selection only appears once as we only need to validate the rules once
	selSet := make(map[string]struct{}, len(selections))
	for _, pid := range selections {
		selSet[pid] = struct{}{}
	}

	for principal, rules := range rulesByPrincipal {
		for _, rule := range rules {
			switch domain.RestrictionEffect(rule.Effect) {
			case domain.Forbid:
				// if the forbidden resource is within the group of selections
				// then we have a violation
				if _, chosen := selSet[rule.ResourceID]; chosen {
					violations = append(violations, domain.Violation{
						IfPart:   principal,
						WithPart: rule.ResourceID,
						Effect:   domain.Forbid,
						Message:  "Selection of " + rs.PartsRepo.GetNameForPart(principal) + " forbids " + rs.PartsRepo.GetNameForPart(rule.ResourceID),
					})
				}

			case domain.Only:
				// for this principal, in the resource’s category, only rule.ResourceID may be selected
				for selected := range selSet {
					// skip the allowed ones
					if selected == rule.ResourceID || partCategoryMap[selected] != partCategoryMap[rule.ResourceID] {
						continue
					}
					violations = append(violations, domain.Violation{
						IfPart:   principal,
						WithPart: selected,
						Effect:   domain.Only,
						Message:  "When selecting " + rs.PartsRepo.GetNameForPart(principal) + " you may only choose " + rs.PartsRepo.GetNameForPart(rule.ResourceID),
					})
				}

			case domain.Allow:
				// no-op: always permitted

			default:
				return nil, fmt.Errorf("critical error rule with effect besides allow, only or forbid")
			}
		}
	}

	return violations, nil
}
