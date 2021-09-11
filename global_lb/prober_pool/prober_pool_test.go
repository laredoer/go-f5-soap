package prober_pool

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://202.173.9.33/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestProberPool_GetList(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(list)
}

func TestProberPool_GetMember(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	members, err := p.GetMember(list)
	if err != nil {
		t.Fatal(err)
	}

	orders, err := p.GetMemberOrder(list, members)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(orders)
}
