package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConfig_Parse(t *testing.T) {
	// Set up test environment
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Failed to load environment variables: %v", err)
	}

	// Create a new Config instance
	c := &Config{}

	// Call the Parse method
	err = c.Parse()
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	// Assert the expected values
	expectedPort := os.Getenv("POSTGRES_PORT")
	if c.PostgresPort != expectedPort {
		t.Errorf("Unexpected PostgresPort value. Expected: %s, Got: %s", expectedPort, c.PostgresPort)
	}

	expectedHost := os.Getenv("POSTGRES_HOST")
	if c.PostgresHost != expectedHost {
		t.Errorf("Unexpected PostgresHost value. Expected: %s, Got: %s", expectedHost, c.PostgresHost)
	}

	expectedUser := os.Getenv("POSTGRES_USER")
	if c.PostgresUser != expectedUser {
		t.Errorf("Unexpected PostgresUser value. Expected: %s, Got: %s", expectedUser, c.PostgresUser)
	}

	expectedDBName := os.Getenv("POSTGRES_DB")
	if c.PostgresDBName != expectedDBName {
		t.Errorf("Unexpected PostgresDBName value. Expected: %s, Got: %s", expectedDBName, c.PostgresDBName)
	}

	expectedPassword := os.Getenv("POSTGRES_PASSWORD")
	if c.PostgresPasword != expectedPassword {
		t.Errorf("Unexpected PostgresPasword value. Expected: %s, Got: %s", expectedPassword, c.PostgresPasword)
	}
}
