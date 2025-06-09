// Package config provides utility functions for accessing environment variables
// in a structured and consistent manner across the API application. This uses
// the `github.com/joho/godotenv` package to load environment variables from a `.env`
// file
//
// This package acts as a centralized interface for reading from the runtime
// environment (e.g., from `.env` files, containerized environments, or CI/CD systems),
// and ensures defaults and key presence are handled gracefully
//
// Core Functions:
//
//   - GetEnv() map[string]string
//     Returns all available environment variables as a map for inspection or iteration
//
//   - GetEnvValue(key string) string
//     Retrieves the value of a specific environment variable. Returns an empty string if not set
//
//   - GetEnvWithDefault(key, defaultValue string) string
//     Returns the value of the specified environment variable, or a provided default if the variable is not set
//
//   - KeyExists(key string) bool
//     Checks whether a specific environment variable is present in the runtime environment
//
// Usage Example:
//
// ```go
//
//	dbHost := config.GetEnvWithDefault("DB_HOST", "localhost")
//
//	if !config.KeyExists("JWT_SECRET") {
//	    log.Fatal("JWT_SECRET is required but not set")
//	}
//
// ```
//
// Design Intent:
//
//   - Keep all environment access logic isolated to a single, testable package
//   - Reduce boilerplate in service-level configuration code
//   - Provide safe access patterns that avoid panics or misconfiguration
//
// This package is meant to be minimal and dependency-free, giving API templates a predictable
// and consistent way to interact with system-level configuration at runtime
package config
