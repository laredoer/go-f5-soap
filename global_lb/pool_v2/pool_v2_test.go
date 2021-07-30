package pool_v2

import (
	"crypto/tls"
	go_f5_soap "github.com/wule61/go-f5-soap"
	"testing"
)


func newClient(t *testing.T) *go_f5_soap.Client {
	return go_f5_soap.NewClient("https://202.173.9.33/iControl/iControlPortal.cgi",
		go_f5_soap.WithBasicAuth("admin","admin"),
		go_f5_soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestPoolV2_GetMember(t *testing.T) {
	p := NewPoolV2(newClient(t))

	arr, err := p.GetMember([]PoolID{{"/Common/pool2","GTM_QUERY_TYPE_A"}})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v",arr)
}
