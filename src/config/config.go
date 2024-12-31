package config

type Db struct {
	Host     string `json:"host" env:"HOST"`
	Port     int    `json:"port" env:"PORT"`
	Username string `json:"username" env:"USERNAME"`
	Password string `json:"password" env:"PASSWORD"`
	Name     string `json:"name" env:"NAME"`
}

type Server struct {
	Port int `json:"port"`
}

type Metrics struct {
	Port    int  `json:"port"`
	Enabled bool `json:"enabled"`
}

type Config struct {
	Db      `json:"db" envPrefix:"DB_"`
	Server  `json:"server"`
	Metrics `json:"metrics"`
}
