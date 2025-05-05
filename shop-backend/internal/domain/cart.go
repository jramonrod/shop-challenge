package domain

// Cart represents a singular cart
//
// Note to keep the scope of the challenge down, there is no user relation but
// this would of course be the next steps in the evolution of our cart model
type Cart struct {
	UUID     string            `json:"uuid"`
	Status   string            `json:"status"`
	Products map[string][]Part `json:"products"`
}
