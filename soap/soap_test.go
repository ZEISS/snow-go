package soap_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/zeiss/snow-go/soap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFault(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  *soap.SOAPEnvelope
	}{
		{
			name: "fault",
			in: `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
			<SOAP-ENV:Header/>
			<SOAP-ENV:Body>
				<SOAP-ENV:Fault>
					<faultcode>SOAP-ENV:Server</faultcode>
					<faultstring>Unable to parse SOAP document</faultstring>
					<detail>Error completing SOAP request</detail>
				</SOAP-ENV:Fault>
			</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>`,
			out: &soap.SOAPEnvelope{
				XMLName: xml.Name{
					Space: "http://schemas.xmlsoap.org/soap/envelope/",
					Local: "Envelope",
				},
				XMLNSXsd: "",
				XMLNSXsi: "",
				Header: &soap.SOAPHeader{
					XMLName: xml.Name{
						Space: "http://schemas.xmlsoap.org/soap/envelope/",
						Local: "Header",
					},
				},
				Body: &soap.SOAPBody{
					XMLName: xml.Name{
						Space: "http://schemas.xmlsoap.org/soap/envelope/",
						Local: "Body",
					},
					Fault: &soap.SOAPFault{
						XMLName: xml.Name{
							Space: "http://schemas.xmlsoap.org/soap/envelope/",
							Local: "Fault",
						},
						Code:   "SOAP-ENV:Server",
						String: "Unable to parse SOAP document",
						Detail: "Error completing SOAP request",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := &soap.SOAPEnvelope{}

			dec := xml.NewDecoder(bytes.NewBufferString(tt.in))
			err := dec.Decode(f)
			require.NoError(t, err)

			assert.Equal(t, tt.out, f)
		})
	}
}

func TestRequest_Marshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		url    soap.SOAPUrl
		action soap.SOAPAction
		body   interface{}
		fault  interface{}
	}{
		{
			name:   "nil body",
			url:    soap.SOAPUrl("https://example.com"),
			action: soap.Get,
			body:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := soap.NewRequest(tt.url, tt.action, tt.body, tt.fault)

			req, err := r.Marshal()
			require.NoError(t, err)

			assert.Equal(t, tt.action.String(), req.Header.Get("Soapaction"))
			assert.Equal(t, "POST", req.Method)

			env := soap.NewEnvelope(tt.body)
			buf, err := xml.Marshal(env)
			require.NoError(t, err)

			b, err := io.ReadAll(req.Body)
			require.NoError(t, err)

			fmt.Println(string(b))

			assert.Equal(t, buf, b)
		})
	}
}
