package config

// Arango represent config for arangodb
type Arango struct {
	Host         string
	Port         string
	DatabaseName string
	GraphName    string
}

// Config struct
type Config struct {
	Arango Arango
}

// DefaultConfig return config with default value
func DefaultConfig() *Config {
	arango := Arango{
		Host:         "127.0.0.1",
		Port:         "8529",
		DatabaseName: "ims",
		GraphName:    "img_graph",
	}

	return &Config{
		Arango: arango,
	}
}

// LoadConfig return config
func LoadConfig() *Config {
	return DefaultConfig()
}
