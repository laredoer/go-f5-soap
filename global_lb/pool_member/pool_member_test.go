package pool_member

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.1.101.101/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestPoolMember_GetRatio(t *testing.T) {

	p := New(newClient(t))

	res, err := p.GetRatio([]string{"/Common/pool2"}, [][]common.IPPortDefinition{
		{
			{Address: "1.2.3.4", Port: 22},
			{Address: "10.2.5.5", Port: 0},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", res)
}

func TestPoolMember_GetObjectStatus(t *testing.T) {
	p := New(newClient(t))

	res, err := p.GetObjectStatus([]string{"/Common/pool2"}, [][]common.IPPortDefinition{
		{
			{Address: "1.2.3.4", Port: 22},
			{Address: "10.2.5.5", Port: 0},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", res)
}
