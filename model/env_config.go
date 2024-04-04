package model

type AppConfig struct {
	AppPort     string
	APIPath     string
	AppHost     string
	ForceSSL    string
	PSQLEnabled string
	URLPrefix   string
}

type PostgresConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
}
