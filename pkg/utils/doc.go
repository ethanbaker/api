// Package utils provides common utility functions used across the API service
// These are generic, stateless helpers designed to simplify repetitive tasks and
// enforce consistent behaviors across multiple modules of the application
//
// Key Features:
//   - Standard NoRoute HTTP handler that returns a JSON 404 response
//   - Reusable helper logic that doesn’t belong to any specific domain (e.g., string manipulation, error formatting)
//
// The utils package is intended to be lightweight and static, meaning it can be imported and used
// without configuration or initialization. This makes it ideal for logic that is infrastructural
// rather than business-specific
//
// Example Use:
//
//	// Attach NoRoute handler to a router like Gin
//	router.NoRoute(utils.NoRouteHandler)
//
// The goal is to reduce code duplication and standardize how edge cases are handled,
// particularly in routing and error responses
//
// Typical Use Cases:
//   - Providing a fallback for unmatched API routes
//   - Centralizing reusable helper logic to keep other packages focused
//
// Note: Since the utils package aims to be generic and broadly applicable, it should not contain
// logic that introduces dependencies on business logic, middleware, or external services
package utils
