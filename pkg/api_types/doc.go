// Package api_types defines reusable types and standard response structures
// for consistent API communications across the service
//
// This package centralizes common response formats to enforce uniformity in
// how success, failure, and error messages are conveyed in JSON responses,
// making it easier for users and client applications to handle and parse API replies
//
// Core Types:
//
//	StatusType: Enumerated string constants representing response status categories:
//	  - StatusSuccess: "success" for successful operations
//	  - StatusFail: "fail" for client-side (4xx) errors
//	  - StatusError: "error" for server-side (5xx) errors
//
//	ApiResponse: The canonical API response structure containing:
//	  - Status: the status of the response ("success", "fail", or "error")
//	  - Code: HTTP status code
//	  - Message: human-readable summary
//	  - Data: optional payload on success
//	  - Error: optional error details on failure or error
//
// Convenience Constructors:
//
//   - NewSuccessResponse(message string, data any) ApiResponse
//     Create a success response with an optional data payload
//
//   - NewFailResponse(code int, message string) ApiResponse
//     Create a client error response with a status code and message
//
//   - NewErrorResponse(code int, message string) ApiResponse
//     Create a server error response with a status code and message
//
// Integration:
//
//   - ApiResponse implements JSON marshaling for seamless encoding
//   - `AsGinResponse()` returns the response in a format suitable for Gin handlers,
//     i.e., HTTP status code and JSON body
//
// Usage Example:
//
//	c.JSON(api_types.NewSuccessResponse("User created successfully", userData).AsGinResponse())
//
// Design Rationale:
//
//   - Standardizes response formatting to simplify client-side parsing and debugging
//   - Reduces repetitive response construction across modules
//   - Distinguishes failure types clearly, aiding automated monitoring and logging
//
// Customization:
//
//   - Users may embed or extend ApiResponse with additional fields as needed
//   - StatusType constants can be extended if new categories arise
//
// This package is fundamental for building APIs that are consistent,
// predictable, and easy to integrate with from external clients or frontends
package api_types
