package key

// NOTE: This service is a placeholder to demonstrate API key validation. An actual
// service would have more complex logic to validate the API key, such as
// checking a database or external service

// Validate the API key
func validateAPIKey(key string) bool {
	return key == "1234567890"
}
