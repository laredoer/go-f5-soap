package data_center

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

func TestDataCenter_GetList(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestDataCenter_GetServer(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetServer([]string{"/Common/SH", "/Common/JD", "/Common/BJ"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
