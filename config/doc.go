package config

/**
# Package config

Config package is used to handle environment variable configuration and other custom global services.

## Environment Variables

You can use this package to read environment variables, set them in a global map, and access them from
anywhere in your application. This is useful for storing sensitive information like API keys, database
credentials, or any other static configuration values.

## Custom Configuration Services

You can also use this package to create custom configuration services. For example, you can create a
custom configuration service for your MySQL database connection. This service can be used to open a new
connection to the database, set the connection parameters, and return the connection object.

Each custom configuration service is split into three parts:
* **Globals**: Cache the connection object and other configuration values
* **Init**: Initialize globals
* **Public Functions**: Exported functions that can be used by other packages

You can theoretically make a global variable public to be accessed by other packages, but it is not
recommended. Instead, you should create a public function that returns the global variable. This way,
you can control how the global variable is accessed and prevent any unwanted changes.
*/
