package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	snowgo "github.com/zeiss/snow-go"
)

// DefaultPushConnectorSource is the default source for the ServiceNow Push Connector API.
const DefaultPushConnectorSource = "GenericJson"

// PushConnectorUrl is the URL for the ServiceNow Push Connector API.
type PushConnectorUrl struct {
	Url    string
	Source string
}

// String returns the string representation of the Push Connector URL.
func (s PushConnectorUrl) String() string {
	return fmt.Sprintf("%s?source=%s", s.Url, s.Source)
}

// SetSource sets the source of the Push Connector URL.
func (s *PushConnectorUrl) SetSource(source string) {
	s.Source = source
}

// SetUrl sets the URL of the Push Connector URL.
func (s *PushConnectorUrl) SetUrl(url string) {
	s.Url = url
}

// NewPushConnectorUrl returns a new Push Connector URL.
func NewPushConnectorUrl(instance string, source string) PushConnectorUrl {
	url := fmt.Sprintf("https://%s/api/sn_em_connector/em/inbound_event", instance)

	if source == "" {
		source = DefaultPushConnectorSource
	}

	return PushConnectorUrl{
		Url:    url,
		Source: source,
	}
}

var _ snowgo.Request = (*Request)(nil)

// Request represents a ServiceNow Push Connector API request.
type Request struct {
	event cloudevents.Event
	url   PushConnectorUrl
}

// Marshal returns a new HTTP request for the ServiceNow Push Connector API.
func (r *Request) Marshal() (*http.Request, error) {
	buf, err := r.event.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	return req, nil
}

var _ snowgo.Response = (*Response)(nil)

// Response represents a ServiceNow Push Connector API response.
type Response struct {
	Result map[string]interface{} `json:"result"`
}

// Unmarshal reads the ServiceNow Push Connector API response from the HTTP response.
func (r *Response) Unmarshal(res *http.Response) error {
	err := json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return err
	}

	return nil
}

// NewRequest returns a new ServiceNow Push Connector API request.
func NewRequest(url PushConnectorUrl, event cloudevents.Event) *Request {
	return &Request{
		event: event,
		url:   url,
	}
}
