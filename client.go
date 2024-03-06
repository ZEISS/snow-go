package snowgo

import (
	"net/http"
	"time"
)

const (
	// DefaultSuffix is the default ServiceNow REST API URL suffix
	DefaultSuffix = ".service-now.com/api"
)

// Credentials is a struct that represents ServiceNow REST API credentials
type Credentials struct {
	Username string
	Password string
}

// NewCredentials returns a new ServiceNow REST API credentials
func NewCredentials(username, password string) *Credentials {
	return &Credentials{
		Username: username,
		Password: password,
	}
}

// Client is a struct that represents a ServiceNow REST API client
type Client struct {
	BaseURL     string
	Credentials *Credentials

	client http.Client
}

// Opt is a function that sets a client option
type Opt func(*Client)

// WithBaseURL sets the ServiceNow instance URL
func WithBaseURL(baseURL string) Opt {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// WithHTTPClient sets the HTTP client
func WithHTTPClient(client http.Client) Opt {
	return func(c *Client) {
		c.client = client
	}
}

// WithCredentials sets the ServiceNow REST API credentials
func WithCredentials(cred *Credentials) Opt {
	return func(c *Client) {
		c.Credentials = cred
	}
}

// New returns a new ServiceNow REST API client
func New(opts ...Opt) *Client {
	c := new(Client)
	c.client = http.Client{
		Timeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
