// This implementation uses GORM
package mysql

import (
	"fmt"
	"log"

	mysql_driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global database connection
var db *gorm.DB

// Connect initializes the MySQL database connection using a given DSN object
func Connect(dsn mysql_driver.Config) error {
	// Get DSN credentials from provided object and check for errors
	if dsn.User == "" {
		return fmt.Errorf("no user provided in dsn")
	}
	if dsn.Passwd == "" {
		return fmt.Errorf("no password provided in dsn")
	}
	if dsn.Addr == "" {
		return fmt.Errorf("no address provided in dsn")
	}
	if dsn.DBName == "" {
		return fmt.Errorf("no database name provided in dsn")
	}

	log.Printf("[MYSQL]: connecting to MySQL database at %s as user %s\n", dsn.Addr, dsn.User)

	// Open a new connection to the database
	var err error

	db, err = gorm.Open(mysql.Open(dsn.FormatDSN()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Printf("[MYSQL]: connected to MySQL database successfully\n")
	return nil
}

// Get returns the global database connection
func Get() *gorm.DB {
	return db
}
