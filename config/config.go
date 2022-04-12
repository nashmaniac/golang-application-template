package config

type Config struct {
	Database struct {
		Name     string
		Username string
		Password string
		Host     string
		Port     string
	}
	Application struct {
		Port int
	}
}
