// Package api_key provides middleware utilities for securing API endpoints
// using custom API key validation. This is useful for services that require
// a simple authentication mechanism without full JWT or OAuth implementations
//
// Overview:
//
// This package enables developers to attach middleware to routes that require
// an API key to access. Keys can be extracted either from HTTP headers or query
// parameters, and are validated using a user-supplied validation function
//
// Middleware Functions:
//
//   - APIKeyHeaderHandler(validate func(string) bool) gin.HandlerFunc
//     Validates API keys supplied via the 'X-API-KEY' header. The developer must
//     provide a `validate` function that returns true if the key is accepted
//
//   - APIKeyQueryHandler(param string, validate func(string) bool) gin.HandlerFunc
//     Validates API keys passed as query parameters. The name of the query parameter
//     and the `validate` function are both required
//
// Response Behavior:
//
//   - If the API key is missing or invalid, the middleware responds with a
//     403 Forbidden status and a structured failure response using the `api_types` package
//   - If the key is valid, the request is passed to the next handler
//
// Usage Example:
//
//	router.GET("/secure-data", api_key.APIKeyHeaderHandler(func(key string) bool {
//	    return key == myApiKey
//	}), dataHandler)
//
// Design Intent:
//
//   - Provide a lightweight, pluggable security layer for public or internal APIs
//   - Keep authentication logic decoupled and testable
//   - Offer consistent failure responses via standardized types
//
// This package is ideal for services needing lightweight security for
// webhook endpoints, internal tooling, or machine-to-machine APIs
package api_key
