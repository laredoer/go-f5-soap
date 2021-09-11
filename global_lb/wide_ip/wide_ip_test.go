package wide_ip

import (
	"crypto/tls"
	"encoding/json"
	"testing"

	soap "github.com/wule61/go-f5-soap"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.2.0.44/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
		soap.WithDebug(),
	)
}

func TestWideIP_GetList(t *testing.T) {

	p := New(newClient(t))

	res, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(res))

	data, _ := json.Marshal(&res)

	t.Logf("%s", data)
}

func TestWideIP_GetWideIpPool(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(list))

	res, err := p.GetWideIpPool(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(res))

	t.Log(res)
}

func TestWideIP_GetLBMethod(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(list))

	res, err := p.GetLBMethod(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(res))

	t.Log(res)
}

func TestWideIP_GetObjectStatus(t *testing.T) {
	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	res, err := p.GetObjectStatus(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestWideIP_GetEnabledState(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(list))

	res, err := p.GetEnabledState(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(res))

	t.Log(res)
}
