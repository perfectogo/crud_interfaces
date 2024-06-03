package config

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     int
	DbName     string
}

func NewConfig() Config {

	cfg := Config{}

	cfg.DbUser = "postgres"
	cfg.DbPassword = "postgres"
	cfg.DbHost = "0.0.0.0"
	cfg.DbPort = 5432
	cfg.DbName = "postgres"

	return cfg
}
