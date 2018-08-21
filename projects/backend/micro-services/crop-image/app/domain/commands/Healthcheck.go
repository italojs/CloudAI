package command

// Healthcheck will store information about its name and version
type Healthcheck struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
}
