package envmodel

type AppEnv struct {
	AppPort     string
	APIPath     string
	AppHost     string
	ForceSSL    string
	PSQLEnabled string
	URLPrefix   string
}

type PostgresEnv struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
}
