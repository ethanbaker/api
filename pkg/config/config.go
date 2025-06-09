package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Global environment variable map (initialized before init)
var env map[string]string = func() map[string]string {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("[CONFIG|ERR]: error loading .env file")
	}

	// Read environment variables into the global map
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("[CONFIG|ERR]: error reading environment variables")
	}

	return env
}()

// GetEnv returns the value and existance status of the environment variable with the given key
func GetEnv(key string) (string, bool) {
	if _, ok := env[key]; !ok {
		return "", false
	}

	return env[key], true
}

// GetEnvValue returns the value of the environment variable with the given key. If no such key exists, an empty string is returned
func GetEnvValue(key string) string {
	if _, ok := env[key]; !ok {
		return ""
	}

	return env[key]
}

// GetEnvWithDefault returns the value of the environment variable with the given key. If no such key exists, the default value is returned
func GetEnvWithDefault(key, defaultValue string) string {
	if _, ok := env[key]; !ok {
		return defaultValue
	}

	return env[key]
}

// KeyExists returns true if the environment variable with the given key exists, false otherwise
func KeyExists(key string) bool {
	_, ok := env[key]
	return ok
}
