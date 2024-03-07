package soap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/zeiss/snow-go/soap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalXML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  *soap.Fault
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
			out: &soap.Fault{
				Code:   "SOAP-ENV:Server",
				String: "Unable to parse SOAP document",
				Detail: "Error completing SOAP request",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := soap.NewFault()

			dec := xml.NewDecoder(bytes.NewBufferString(tt.in))
			err := dec.Decode(f)
			require.NoError(t, err)

			assert.Equal(t, tt.out, f)
		})
	}
}
