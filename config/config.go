package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// the address to listen on
	PostgresPort string
	// the address of the database
	PostgresHost string
	// the name of the database
	PostgresUser string
	// the user to connect to the database
	PostgresDBName string
	// the password to connect to the database
	PostgresPasword string
}

// create a new config struct with default values
func NewConfig() *Config {
	return &Config{}
}

// verify if the configuration is valid
func (c *Config) Validate() error {
	if c.PostgresPort == "" {
		return ErrMissingConfig
	}
	if c.PostgresHost == "" {
		return ErrMissingConfig
	}
	if c.PostgresUser == "" {
		return ErrMissingConfig
	}
	if c.PostgresDBName == "" {
		return ErrMissingConfig
	}
	if c.PostgresPasword == "" {
		return ErrMissingConfig
	}
	return nil
}

// parse the configuration from a .env file and returns an error if it fails
func (c *Config) Parse() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.PostgresPort = os.Getenv("POSTGRES_PORT")
	c.PostgresHost = os.Getenv("POSTGRES_HOST")
	c.PostgresUser = os.Getenv("POSTGRES_USER")
	c.PostgresDBName = os.Getenv("POSTGRES_DB")
	c.PostgresPasword = os.Getenv("POSTGRES_PASSWORD")
	log.Print("Gahmen API Configuration loaded")
	return nil
}

var ErrMissingConfig = errors.New("missing configuration")
