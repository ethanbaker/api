// This file contains the configuration for the MySQL database connection. This is a custom implementation and can be changed as needed!
package config

import (
	"log"
	"time"

	mysql_driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/** GLOBALS: Exported to external packages and can be used across files/modules */

var db *gorm.DB

/** INIT FUNCTIONS: Called to initialize globals */

func todo_init() {
	// Get DSN credentials from environment variables
	user, ok := GetEnv("DB_USER")
	if !ok {
		log.Fatal("[CONFIG_ERR]: 'DB_USER' environment variable is not set")
	}

	passwd, ok := GetEnv("DB_PASSWD")
	if !ok {
		log.Fatal("[CONFIG_ERR]: 'DB_PASSWD' environment variable is not set")
	}

	net, ok := GetEnv("DB_NET")
	if !ok {
		log.Fatal("[CONFIG_ERR]: 'DB_NET' environment variable is not set")
	}

	addr, ok := GetEnv("DB_ADDR")
	if !ok {
		log.Fatal("[CONFIG_ERR]: 'DB_ADDR' environment variable is not set")
	}

	dbname, ok := GetEnv("DB_NAME")
	if !ok {
		log.Fatal("[CONFIG_ERR]: 'DB_NAME' environment variable is not set")
	}

	// Create the mysql DSN
	dsn := mysql_driver.Config{
		User:      user,
		Passwd:    passwd,
		Net:       net,
		Addr:      addr,
		DBName:    dbname,
		ParseTime: true,
		Loc:       time.Local,
	}

	// Open a new connection to the database
	var err error

	db, err = gorm.Open(mysql.Open(dsn.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("[CONFIG_ERR]: could not open gorm db (%v)\n", err)
	}
}

/** FUNCTIONS: Accessible from external packages */

// GetDB returns the global database connection
func GetDB() *gorm.DB {
	return db
}
