package soap

import (
	"encoding/xml"
	"fmt"
)

// Fault represents a ServiceNow SOAP API fault.
type Fault struct {
	// XMLName is the XML name of the fault.
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	// Code is the fault code.
	Code string `xml:"faultcode"`
	// String is the fault string.
	String string `xml:"faultstring"`
	// Detail is the fault detail.
	Detail string `xml:"detail"`
}

// NewFault returns a new ServiceNow SOAP API fault.
func NewFault() *Fault {
	return &Fault{}
}

// Error returns the string representation of the fault.
func (f *Fault) Error() string {
	return fmt.Sprintf("soap fault: %s (%s)", f.Code, f.String)
}
