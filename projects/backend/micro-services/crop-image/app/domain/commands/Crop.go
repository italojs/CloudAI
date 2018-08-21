package command

// Crop is the json that will be trafficked
// between http requests/responses
type Crop struct {
	From XY `json:"from"`
	To   XY `json:"to"`
}
