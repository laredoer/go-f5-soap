package topology

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.1.101.101/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestTopology_GetOrder(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetOrder([]TopologyRecord{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestTopology_GetList(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}
