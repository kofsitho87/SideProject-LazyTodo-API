package config

import (
	"fmt"
	"os"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT", "3000")

	//DB INFO
	DB          = getEnv("DB", "postgres")
	DB_PORT     = getEnv("DB_PORT", "5432")
	DB_USER     = getEnv("DB_USER", "postgres")
	DB_PASSWORD = getEnv("DB_PASSWORD", "changeme")

	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "test_key")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP = getEnv("TOKEN_EXP", "100h")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
