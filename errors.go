package snowgo

import "fmt"

// Exception represents an exception message from the ServiceNow REST API
type Exception struct {
	// Detail is the exception detail message
	Detail string `json:"detail"`
	// Message is the exception message
	Message string `json:"message"`
}

// Error represents an error response from the ServiceNow REST API
type Error struct {
	// Exception is the exception message
	Exception Exception `json:"error"`
	// Status is the HTTP status code
	Status string `json:"status"`
}

// Error returns the error message
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Exception.Message, e.Exception.Detail)
}
