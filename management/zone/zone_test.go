package zone

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.1.107.243/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestZone_GetZoneName(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetZoneName([]string{"external"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
