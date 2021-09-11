package wide_ip_v2

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.2.0.39/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
		//soap.WithDebug(),
	)
}

func TestWideIPV2_GetList(t *testing.T) {
	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(list)
}

func TestWideIPV2_GetListByType(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetListByType([]global_lb.GTMQueryType{
		global_lb.GtmQueryTypeA,
		global_lb.GtmQueryTypeAAAA,
		global_lb.GtmQueryTypeCname,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(list)

}

func TestWideIPV2_GetWideIpPool(t *testing.T) {

	p := New(newClient(t))

	//list, err := p.GetList()
	//if err != nil {
	//	t.Fatal(err)
	//}

	pools, err := p.GetWideIpPool([]global_lb.WideIPID{
		{
			"/Common/web01y.fatp.nb", "GTM_QUERY_TYPE_A",
		},
		{
			"/Common/repo.nbcb.com", "GTM_QUERY_TYPE_A",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pools)
}

func TestWideIPV2_GetLBMethod(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	pools, err := p.GetLBMethod(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pools)
}

func TestWideIPV2_GetObjectStatus(t *testing.T) {
	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	pools, err := p.GetObjectStatus(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pools)
}

func TestWideIPV2_GetEnabledState(t *testing.T) {
	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	pools, err := p.GetEnabledState(list)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pools)
}

func TestWideIPV2_GetWideIpPoolRatio(t *testing.T) {

	p := New(newClient(t))

	list, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	pools, err := p.GetWideIpPool(list)
	if err != nil {
		t.Fatal(err)
	}
	ratio, err := p.GetWideIpPoolRatio(list, pools)

	//ratio, err := p.GetWideIpPoolRatio([]global_lb.WideIPID{
	//	{
	//		"/Common/web01y.fatp.nb", "GTM_QUERY_TYPE_A",
	//	},
	//	{
	//		"/Common/repo.nbcb.com", "GTM_QUERY_TYPE_A",
	//	},
	//}, [][]global_lb.PoolID{
	//	{
	//		{
	//			"/Common/pool_web01y.fatp.nb", "GTM_QUERY_TYPE_A",
	//		},
	//	},
	//	{
	//		{
	//			"/Common/pool_repo.nbcb.com", "GTM_QUERY_TYPE_A",
	//		},
	//	},
	//})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ratio)
}
