package domain

import "fmt"

// Restriction represents a single compatibility rule between two parts in the system.
// It defines what one part in the system, the principal, imposese on another, the resource.
//
// Example Restriction:
//
//	– Principal = “mountain-wheels”, Resource = “Frame type:full-suspension”, Effect = Only
//	– Principal = “fat-bike-wheels”, Resource = “Rim Color:red”,        Effect = Forbid
type Restriction struct {
	UUID        string            `json:"uuid"`
	PrincipalID string            `json:"principal_id"`
	ResourceID  string            `json:"resource_id"`
	Effect      RestrictionEffect `json:"effect"`
}

// RestrictionEffect enumerates the possible effects a Principal can impose on a Resource
type RestrictionEffect string

func ParseRestrictionEffect(s string) (RestrictionEffect, error) {
	e := RestrictionEffect(s)
	switch e {
	case Allow, Forbid, Only:
		return e, nil
	default:
		return "", fmt.Errorf("invalid RestrictionEffect %q", s)
	}
}

const (
	// Allow indicates that the Resource is explicitly permitted when Principal is chosen
	//
	// By design this is the default interaction of all parts that belong in a category
	Allow RestrictionEffect = "allow"

	// Forbid indicates that the Resource is explicitly disallowed when Principal is chosen
	Forbid RestrictionEffect = "forbid"

	// Only indicates that the Resource is the exclusive choice allowed when Principal is chosen
	Only RestrictionEffect = "only"
)
