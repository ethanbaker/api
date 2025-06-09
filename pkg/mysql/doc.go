// Package mysql provides a centralized and reusable interface for establishing
// and accessing a MySQL database connection using GORM
//
// This package is designed for API projects that rely on GORM for ORM functionality,
// offering a standardized way to configure and retrieve a database instance
//
// Current Features:
//
//   - Connect(dsn mysql_driver.Config) error
//     Establishes a GORM connection to a MySQL database using the provided DSN configuration
//     This should be called once during application startup
//
//   - Get() *gorm.DB
//     Returns the active GORM database instance for use across the application
//
// Example Usage:
//
//	// In main or initialization logic
//	err := mysql.Connect(yourDSNConfig)
//	if err != nil {
//	    log.Fatalf("failed to connect to DB: %v", err)
//	}
//
//	// In your services or repositories
//	db := mysql.Get()
//	db.Find(&users)
//
// Design Notes:
//
//   - The package abstracts away connection handling to keep DB setup logic
//     isolated from business logic and service layers
//   - The use of a DSN config struct allows compatibility with different environments
//     and simplifies connection string management
//
// Future Plans:
//
// This package is designed to be expanded in future versions to support:
//   - Loading specific GORM models automatically on startup
//   - Running predefined or dynamic custom SQL queries
//   - Connection pooling configuration and performance tuning
//   - Utility functions for transactions, migrations, and testing
//
// This package aims to reduce boilerplate and enforce consistency in how
// database access is configured and performed across the codebase
package mysql
