package virtual_server

import (
	"crypto/tls"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"

	"testing"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.2.0.44/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestVirtualServer_GetList(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestVirtualServer_GetMonitorAssociation(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetMonitorAssociation([]global_lb.VirtualServerDefinition{
		{Name: "vs_84_147_225_211_3306", Address: "84.147.225.211", Port: 3306},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestVirtualServer_GetServer(t *testing.T) {
	p := New(newClient(t))
	arr, err := p.GetServer([]global_lb.VirtualServerDefinition{
		{Name: "vs_76_33_110_80_2000", Address: "76.33.110.80", Port: 2000},
		{Name: "vs_76_33_110_80_80", Address: "76.33.110.80", Port: 80},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}
