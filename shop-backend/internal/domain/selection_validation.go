package domain

// PartsSelection represents a group of Parts that is being asked of us to validate against our restriction rules
type PartsSelection []string

// Violation describes a single business-rule violation detected when validating
// a customer’s part selections. It captures the “principal” of the rule and which “resource” part is in conflict
type Violation struct {
	IfPart   string            `json:"principal"`
	WithPart string            `json:"resource_id"`
	Effect   RestrictionEffect `json:"effect"`
	Message  string            `json:"message"`
	Allowed  []Part            `json:"allowed"`
}
