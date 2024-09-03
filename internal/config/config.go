package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ENV *config

type config struct {
	environmentRoute string
	dbName           string
	dbUser           string
	dbPassword       string
	dbHost           string
	dbPort           string
	dbSSL            string
	jwtKey           string //config jwt
}

func NewConfig(environmentRoute string) *config {
	return &config{
		environmentRoute: environmentRoute,
	}
}

func (c *config) SetConfigFile() {

	err := godotenv.Load(c.environmentRoute)
	if err != nil {
		log.Fatalf("Can't read environment variables from file %v\n", c.environmentRoute)
	}
	c.dbName = getenv("DB_NAME", "postgres")
	c.dbUser = getenv("DB_USER", "postgres")
	c.dbPassword = getenv("DB_PASSWORD", "1234567")
	c.dbHost = getenv("DB_HOST", "localhost")
	c.dbPort = getenv("DB_PORT", "5432")
	c.dbSSL = getenv("DB_SSL", "disable")
	c.jwtKey = getenv("JWT", "")
	ENV = c
}

func (c *config) GetPostgresConfig() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", c.dbName, c.dbPassword, c.dbHost, c.dbName, c.dbSSL)
}

func (c *config) GetJWTKey() string {
	return c.jwtKey
}

func getenv(key string, alrt string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return alrt
}
