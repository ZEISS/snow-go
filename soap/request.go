package soap

import (
	"bytes"
	"encoding/xml"
	"net/http"

	snowgo "github.com/zeiss/snow-go"
)

var (
	_ snowgo.Request  = (*Request)(nil)
	_ snowgo.Response = (*Response)(nil)
)

// SOAPUrl is the URL for the ServiceNow SOAP API.
type SOAPUrl string

// String returns the string representation of the SOAP URL.
func (s SOAPUrl) String() string {
	return string(s)
}

// SOAPAction is the HTTP header for the SOAP action.
type SOAPAction string

// String returns the string representation of the SOAP action.
func (s SOAPAction) String() string {
	return string(s)
}

const (
	Aggregate      SOAPAction = "aggregate"
	DeleteMultiple SOAPAction = "deleteMultiple"
	DeleteRecord   SOAPAction = "deleteRecord"
	Get            SOAPAction = "get"
	GetKeys        SOAPAction = "getKeys"
	GetRecords     SOAPAction = "getRecords"
	Insert         SOAPAction = "insert"
	InsertMultiple SOAPAction = "insertMultiple"
	Update         SOAPAction = "update"
)

// Request represents a ServiceNow SOAP API request.
type Request struct {
	action SOAPAction
	url    SOAPUrl

	body  interface{}
	fault interface{}
}

// Marshal returns a new HTTP request for the ServiceNow SOAP API.
func (r *Request) Marshal() (*http.Request, error) {
	envelop := NewEnvelope(r.body)

	buf, err := xml.Marshal(envelop)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.url.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", r.action.String())

	return req, nil
}

// Response represents a ServiceNow SOAP API response.
type Response struct{}

// Unmarshal unmarshals an HTTP response into a ServiceNow SOAP API response.
func (r *Response) Unmarshal(*http.Response) error {
	return nil
}

// NewRequest returns a new ServiceNow SOAP API request.
func NewRequest(url SOAPUrl, action SOAPAction, body interface{}, fault interface{}) *Request {
	req := &Request{
		action: action,
		url:    url,
		body:   body,
		fault:  fault,
	}

	return req
}
