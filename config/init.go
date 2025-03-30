// Create the first init function to be executed in the package
package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Global environment variable map
var env map[string]string = _init()

// Read environment variables and set them in the global map
func _init() map[string]string {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("[CONFIG_ERR]: error loading .env file")
	}

	// Read environment variables into the global map
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("[CONFIG_ERR]: error reading environment variables")
	}

	return env
}
