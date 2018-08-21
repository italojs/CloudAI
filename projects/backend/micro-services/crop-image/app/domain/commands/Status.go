package command

// Status is the json that will be trafficked
// between http responses
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
