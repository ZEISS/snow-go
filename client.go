package snowgo

import (
	"context"
	"io"
	"net/http"
	"strings"
)

const (
	// DefaultSuffix is the default ServiceNow REST API URL suffix
	DefaultSuffix = ".service-now.com/api"
)

// Request is an interface for ServiceNow REST API requests
type Request interface {
	Marshal() (*http.Request, error)
}

// Response is an interface for ServiceNow REST API responses
type Response interface {
	Unmarshal(*http.Response) error
}

// RequestEditorFn is a function that can be used to modify an HTTP request
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Client is an opaque type that holds the client configuration
type Client struct {
	http            *http.Client
	server          string
	requestEditorFn []RequestEditorFn
}

// Opt is a functional option type for configuring the client
type Opt func(*Client)

// WithHTTPClient sets the HTTP client to use
func WithHTTPClient(http *http.Client) Opt {
	return func(c *Client) {
		c.http = http
	}
}

// WithRequestEditorFn sets the request editor function to use
func WithRequestEditorFn(fn ...RequestEditorFn) Opt {
	return func(c *Client) {
		c.requestEditorFn = append(c.requestEditorFn, fn...)
	}
}

// New returns a new ServiceNow client
func New(server string, opts ...Opt) *Client {
	c := &Client{
		http:   http.DefaultClient,
		server: server,
	}

	if !strings.HasSuffix(c.server, "/") {
		c.server += "/"
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(ctx context.Context, req Request, resp Response) error {
	httpReq, err := req.Marshal()
	if err != nil {
		return err
	}

	for _, fn := range c.requestEditorFn {
		err := fn(ctx, httpReq)
		if err != nil {
			return err
		}
	}

	httpResp, err := c.http.Do(httpReq.WithContext(ctx))
	if err != nil {
		return err
	}

	err = resp.Unmarshal(httpResp)

	_, err = io.Copy(io.Discard, httpResp.Body)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	return nil
}
