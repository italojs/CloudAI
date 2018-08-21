package interfaces

type AppConfig interface {
	GetEnv(string) string
}
