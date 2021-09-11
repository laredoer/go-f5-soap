package resource_record

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/management"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://202.173.9.33/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
		//soap.WithDebug(),
	)
}

func TestResourceRecord_GetRRS(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetRRS([]management.ViewZone{
		{ViewName: "external", ZoneName: "nbcb.com."},
		{ViewName: "external", ZoneName: "nbcb.com.cn."},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestResourceRecord_GetRRSDetailed(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetRRSDetailed([]management.ViewZone{
		{ViewName: "external", ZoneName: "nbcb.com."},
		{ViewName: "external", ZoneName: "nbcb.com.cn."},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
