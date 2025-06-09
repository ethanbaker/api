package api_types

import "encoding/json"

// StatusType defines the type for status messages in API responses
type StatusType string

const (
	StatusSuccess StatusType = "success" // For successful operations
	StatusFail    StatusType = "fail"    // For 4xx client errors
	StatusError   StatusType = "error"   // For 5xx server errors
)

// ApiResponse represents a standard API response structure
type ApiResponse struct {
	Status  StatusType `json:"status"`          // Status message
	Code    int        `json:"code"`            // Status code
	Message string     `json:"message"`         // Human-readable message
	Data    any        `json:"data,omitempty"`  // Optional data field for successful responses
	Error   any        `json:"error,omitempty"` // Optional errors field for error responses
}

// Return the ApiResponse as a marshalable JSON object
func (r ApiResponse) AsJson() ([]byte, error) {
	return json.Marshal(r)
}

// Return the ApiResponse in a format to provide to Gin Context
func (r ApiResponse) AsGinResponse() (int, any) {
	return r.Code, r
}

// Create a new success response
func NewSuccessResponse(message string, data any) ApiResponse {
	return ApiResponse{
		Status:  StatusSuccess,
		Code:    200,
		Message: message,
		Data:    data,
	}
}

// Create a new fail response
func NewFailResponse(code int, message string) ApiResponse {
	return ApiResponse{
		Status: StatusFail,
		Code:   code,
		Error:  message,
	}
}

// Create a new error response
func NewErrorResponse(code int, message string) ApiResponse {
	return ApiResponse{
		Status: StatusError,
		Code:   code,
		Error:  message,
	}
}
