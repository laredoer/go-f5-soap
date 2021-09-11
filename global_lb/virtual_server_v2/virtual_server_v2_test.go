package virtual_server_v2

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.2.0.44/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestVirtualServerV2_GetAddress(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetAddress([]VirtualServerID{
		{"vs_76_33_110_80_2000", "/Common/A1_AMC-ENTEGOR_AMC-APP-ENT"},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}
