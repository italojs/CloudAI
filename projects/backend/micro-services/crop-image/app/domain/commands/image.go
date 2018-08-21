package command

// Image is the json that will be trafficked
// between http requests/responses
type Image struct {
	User         string        `json:"user"`
	Base64       string        `json:"base64"`
	Length       XY            `json:"length"`
	Label        string        `json:"label"`
	ClassObjects []ClassObject `json:"ClassObjects"`
}
