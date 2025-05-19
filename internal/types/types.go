package common

// Status represents the status of an API response
type Status string

const (
	StatusSuccess Status = "success" // Successful API response
	StatusError   Status = "error"   // Error API response
)

// Response is a generic struct for API responses.  It includes a status,
// message, and optional data.  The data field uses an any to
// allow it to hold any type of data
type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse is a specific type for error responses.  It embeds the
// generic Response and adds an optional Code field for more specific
// error identification.  It's often useful to have error codes that
// your frontend can use to handle errors in a specific way
type ErrorResponse struct {
	Response     // Embed generic Response
	Code     int `json:"code,omitempty"`
}

// NewSuccessResponse creates a new success response with the given message
// and optional data.  It sets the HTTP status code to 200 OK
func NewSuccessResponse(message string, data any) Response {
	return Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response with the given message,
// optional data, and HTTP status code.  It uses the ErrorResponse struct
func NewErrorResponse(statusCode int, message string, data any, code int) ErrorResponse {
	return ErrorResponse{
		Response: Response{
			Status:  StatusError,
			Message: message,
			Data:    data,
		},
		Code: code,
	}
}
