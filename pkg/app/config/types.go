package config

type Config struct {
	Maintenance bool `env:"MAINTENANCE"`

	Server Server

	Database Database
}

type Server struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`

	SessionStorage string `env:"SESSION_STORAGE"`

	Resource string `env:"RESOURCE" default:"views"`
}

type Database struct {
	Driver     string `env:"DB_DRIVER"`
	Host       string `env:"DB_HOST"`
	Port       int    `env:"DB_PORT"`
	Name       string `env:"DB_NAME"`
	User       string `env:"DB_USER"`
	Password   string `env:"DB_PASSWORD"`
	RawOptions string `env:"DB_RAW_OPTIONS"`
}
