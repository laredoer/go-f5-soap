package pool

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

func TestPool_GetAlternateLbMethodByPoolNames(t *testing.T) {


	p := NewPool(newClient(t))

	arr, err := p.GetAlternateLBMethodByPoolNames([]string{"/Common/pool1","/Common/pool2"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetAlternateLBMethodByPoolNames(t *testing.T) {
	p := NewPool(newClient(t))

	arr, err := p.GetPreferredLBMethodByPoolNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetTTLByPoolNames(t *testing.T) {

	p := NewPool(newClient(t))

	arr, err := p.GetTTLByPoolNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetVerifyMemberAvailabilityStateByPoolNames(t *testing.T) {
	p := NewPool(newClient(t))

	arr, err := p.GetVerifyMemberAvailabilityStateByPoolNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetAnswersToReturnByPoolNames(t *testing.T) {
	p := NewPool(newClient(t))

	arr, err := p.GetAnswersToReturnByPoolNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(arr)
}

func TestPool_GetObjectStatusByPoolNames(t *testing.T) {

	p := NewPool(newClient(t))

	arr, err := p.GetObjectStatusByPoolNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v",arr)
}

func TestPool_GetEnabledStateByNames(t *testing.T) {
	p := NewPool(newClient(t))

	arr, err := p.GetEnabledStateByNames([]string{"/Common/pool1","/Common/pool2","/Common/aaaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v",arr)
}