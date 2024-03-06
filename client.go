package snowgo

import (
	"github.com/zeiss/snow-go/apis"
)

const (
	// DefaultSuffix is the default ServiceNow REST API URL suffix
	DefaultSuffix = ".service-now.com/api"
)

// New returns a new ServiceNow REST API client
func New(server string, opts ...apis.ClientOption) (*apis.Client, error) {
	return apis.NewClient(server, opts...)
}
