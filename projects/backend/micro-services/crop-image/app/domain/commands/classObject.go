package command

// ClassObject is the json that will be trafficked
// between http requests/responses
type ClassObject struct {
	Label string `json:"label"`
	Crop  Crop   `json:"crop"`
}
