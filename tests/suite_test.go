package tests

import (
	"os"

	"github.com/stretchr/testify/suite"
)

type TestsSuite struct {
	suite.Suite
}

// SetupTest will be run before each test in the suite
func (suite *TestsSuite) SetupTestEnvironmentals() {
	os.Setenv("APP_PORT", "8000")
	os.Setenv("APP_HOST", "localhost")
	os.Setenv("DEBUG_MODE", "true")
	os.Setenv("API_PATH", "/api/v1/")
	os.Setenv("FORCE_TLS", "false")
	os.Setenv("PSQL_ENABLED", "false")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_DB", "postgres")
}
