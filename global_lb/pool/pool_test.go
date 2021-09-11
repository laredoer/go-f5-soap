package pool

import (
	"crypto/tls"
	"testing"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

func newClient(t *testing.T) *soap.Client {
	return soap.NewClient("https://10.2.0.44/iControl/iControlPortal.cgi",
		soap.WithBasicAuth("admin", "admin"),
		soap.WithTLS(&tls.Config{InsecureSkipVerify: true}),
	)
}

func TestPool_GetAlternateLbMethod(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetAlternateLBMethod([]string{"/Common/pool1", "/Common/pool2"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetAlternateLBMethod(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetPreferredLBMethod([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetTTL(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetTTL([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetVerifyMemberAvailabilityState(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetVerifyMemberAvailabilityState([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetAnswersToReturn(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetAnswersToReturn([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetObjectStatus(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetObjectStatus([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPool_GetEnabledState(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetEnabledState([]string{"/Common/pool1", "/Common/pool2", "/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPool_GetList(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetList()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPool_GetMemberV2(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetMemberV2([]string{"/Common/boom-ac-29-vip_pool", "/Common/boom-ac-28-vip_pool"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPool_GetMonitorAssociation(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetMonitorAssociation([]string{"/Common/A1_UMSP-APP-BAT-NO_pool"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestPool_GetMemberRatio(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetMemberRatio(
		[]string{"/Common/boom-ac-29-vip_pool", "/Common/boom-ac-28-vip_pool"},
		[][]global_lb.VirtualServerID{
			{
				{Name: "vs_76_147_241_195_3306", Server: "/Common/boom-ac-29-vip_76.147.241.195"},
				{Name: "vs_76_147_241_196_3306", Server: "/Common/boom-ac-29-vip_76.147.241.196"},
			},
			{
				{Name: "vs_76_147_225_195_3306", Server: "/Common/boom-ac-28-vip_76.147.225.195"},
				{Name: "vs_76_147_232_195_3306", Server: "//Common/boom-ac-28-vip_76.147.232.195"},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
