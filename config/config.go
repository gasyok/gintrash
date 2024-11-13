package config

type Configuration struct {
	LogLevel string
	Port     int
}

// There's NDA libs
func NewConfig(port int) *Configuration {
	return &Configuration{
		Port:     port,
		LogLevel: "Info",
	}
}
