// Package jwt_auth provides an extendable interface for JWT authentication in a Gin-based API,
// leveraging the github.com/appleboy/gin-jwt middleware. It enables API developers to quickly set up
// stateless authentication mechanisms using JSON Web Tokens while offering hooks for customization
//
// Core Capabilities:
//
//   - SetSecretKey initializes the JWT secret key used for signing tokens
//   - Dynamically generates middleware parameters via `GenerateParams` from configurable inputs
//   - Registers standard JWT routes (`/login`, `/refresh`, `/logout`) with optional prefixing
//   - Provides a reusable middleware handler function for protecting API routes
//   - Offers convenience methods for standardized unauthorized responses and fallback NoRoute handling
//
// Usage:
//
// To use the package:
//
// ```go
//
//		// Step 1: Initialize JWT settings
//		err := jwt_auth.SetSecretKey(yourJwtSecret)
//		if err != nil {
//		    log.Fatalf("JWT initialization failed: %v", err)
//		}
//
//		// Step 2: Generate middleware parameters (example parameters shown below)
//		middleware, err := jwt_auth.GenerateParams[MyUserStruct](jwt_auth.JwtParams{
//		    SigningAlgorithm: "HS256",
//		    Timeout:          "1h",
//		    MaxRefresh:       "1h",
//		    TokenLookup:      "header: Authorization",
//		    TokenHeadName:    "Bearer",
//	     IdentityKey:      "username",
//		})
//		middleware.IdentityHandler = yourIdentityHandler // Reference 'example/users/jwt.go' for implementation
//		middleware.Authenticator = yourAuthenticator
//		middleware.Authorizator = yourAuthorizator
//
//		// Step 3: Create Gin jwt middleware instance
//		middleware, err := jwt.New(params)
//		if err != nil {
//		  log.Fatalf("[ERR]: error creating JWT middleware\n")
//		}
//
//		// Step 4: Register routes and middleware
//		mwh, err := auth.MiddlewareHandler(middleware)
//		if err != nil {
//		  log.Fatalf("[ERR]: error creating JWT middleware handler (%v)\n", err)
//		}
//		group.Use(mwh)
//		auth.RegisterRoute(group, middleware)
//
//		// Step 4: Use middleware for protected routes
//		api.GET("/protected", middleware.MiddlewareFunc(), func(c *gin.Context) {
//		    // Access claims via jwt.ExtractClaims(c)
//		})
//
// ```
//
// Design Decisions:
//
//   - **User Interface:** Users must implement a `user` interface containing `GetUsername()`
//   - **Generic Support:** `GenerateParams` uses generics, allowing flexible user object types
//   - **Pluggable Logic:** Core JWT functions like `Authenticator`, `Authorizator`, and `IdentityHandler` are left
//     undefined to allow for custom `user` struct implementations
//
// Types:
//
//   - JwtParams: Configuration for token duration, signing algorithm, lookup method, etc
//   - JwtRequest / JwtResponse: Structs for token exchange requests and responses
//   - User: Interface for user models to expose a username/identity key
//
// This package provides a clean foundation for stateless authentication while staying modular enough to support
// advanced needs or unconventional flows
package jwt_auth
