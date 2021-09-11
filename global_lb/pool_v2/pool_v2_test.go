package pool_v2

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.1.101.101/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestPoolV2_GetMember(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetMember([]PoolID{
		{"/Common/pool1", "GTM_QUERY_TYPE_A"},
		{"/Common/pool2", "GTM_QUERY_TYPE_A"},
		{"/Common/cname", "GTM_QUERY_TYPE_CNAME"},
		{"/Common/aaaa", "GTM_QUERY_TYPE_AAAA"},
		{"/Common/pool1", "GTM_QUERY_TYPE_AAAA"},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPoolV2_GetListByType(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetListByType([]global_lb.GTMQueryType{
		global_lb.GtmQueryTypeA,
		global_lb.GtmQueryTypeCname,
		global_lb.GtmQueryTypeAAAA,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPoolV2_GetList(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPoolV2_GetTTL(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTTL([]PoolID{
		{PoolName: "/Common/pool1", PoolType: "GTM_QUERY_TYPE_A"},
		{PoolName: "/Common/mx", PoolType: "GTM_QUERY_TYPE_MX"},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPoolV2_GetEnabledState(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetEnabledState([]PoolID{
		{PoolName: "/Common/pool1", PoolType: "GTM_QUERY_TYPE_A"},
		{PoolName: "/Common/mx", PoolType: "GTM_QUERY_TYPE_MX"},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPoolV2_GetObjectStatus(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetObjectStatus([]PoolID{
		{PoolName: "/Common/pool1", PoolType: "GTM_QUERY_TYPE_A"},
		{PoolName: "/Common/mx", PoolType: "GTM_QUERY_TYPE_MX"},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
