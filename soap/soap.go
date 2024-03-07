package soap

import (
	"encoding/xml"
	"fmt"
)

// SOAPEnvelope envelope
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	// These are generic namespaces used by all messages.
	XMLNSXsd string `xml:"xmlns:xsd,attr,omitempty"`
	XMLNSXsi string `xml:"xmlns:xsi,attr,omitempty"`

	Header *SOAPHeader `xml:",omitempty"`
	Body   *SOAPBody   `xml:",omitempty"`
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPBody body
type SOAPBody struct {
	// XMLName is the serialized name of this object.
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	// XMLNSWsu is the SOAP WS-Security utility namespace.
	XMLNSWsu string `xml:"xmlns:wsu,attr,omitempty"`
	// ID is a body ID used during WS-Security signing.
	ID string `xml:"wsu:Id,attr,omitempty"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// SOAPFault represents a ServiceNow SOAP API fault.
type SOAPFault struct {
	// XMLName is the XML name of the fault.
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	// Code is the fault code.
	Code string `xml:"faultcode"`
	// String is the fault string.
	String string `xml:"faultstring"`
	// Detail is the fault detail.
	Detail string `xml:"detail"`
}

// Error returns the string representation of the fault.
func (f *SOAPFault) Error() string {
	return fmt.Sprintf("soap fault: %s (%s)", f.Code, f.String)
}

// NewEnvelope creates a new SOAP envelope.
func NewEnvelope() *SOAPEnvelope {
	return &SOAPEnvelope{}
}
