package monitor

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
	)
}

func TestMonitor_GetTemplateType(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateType([]string{"/Common/HTTP-FOSFB", "/Common/HTTPS_APP_RTCP"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateList(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateList()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetParentTemplate(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetParentTemplate([]string{
		"/Common/HTTP_STATIC_HTML_OCTS-XNV",
		"/Common/HTTP_PIM-APP",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateState(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateState([]string{
		"/Common/HTTP_STATIC_HTML_OCTS-XNV",
		"/Common/HTTP_PIM-APP",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateDestination(t *testing.T) {

	p := New(newClient(t))

	arr, err := p.GetTemplateDestination([]string{
		"/Common/HTTP_STATIC_HTML_cftpaas4",
		"/Common/HTTP_STATIC_HTML_OCTS-XNV",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateStringProperty(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateStringProperty([]string{
		"/Common/HTTP_STATIC_HTML_cftpaas4",
		"/Common/HTTP_STATIC_HTML_cftpaas4",
		"/Common/HTTP_STATIC_HTML_cftpaas4",
		"/Common/HTTP_STATIC_HTML_cftpaas4",
		"/Common/HTTPS_APP_SAAS_1",
		"/Common/HTTPS_APP_SAAS_1",
		"/Common/tcp",
	}, []StrPropertyType{
		STYPE_SEND, STYPE_RECEIVE, STYPE_USERNAME, STYPE_PASSWORD, STYPE_CIPHER_LIST, STYPE_CLIENT_CERTIFICATE, STYPE_SEND,
	})
	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.Marshal(&arr)

	t.Logf("%s", data)
}

func TestMonitor_GetTemplateUserDefinedStringProperty(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateUserDefinedStringProperty([]string{
		"/Common/HTTP_STATIC_HTML_OCTS-XNV",
		"/Common/HTTP_PIM-APP",
	}, []string{
		"STYPE_UNSET",
		"STYPE_SEND",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateIntegerProperty(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateIntegerProperty(
		[]string{"/Common/HTTP_STATIC_HTML_cftpaas4", "/Common/HTTP_STATIC_HTML_cftpaas4", "/Common/gateway_icmp"},
		[]IntPropertyType{ITypeInterval, ITypeTimeOut, "ITYPE_PROBE_ATTEMPTS"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateReverseMode(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateReverseMode(
		[]string{"/Common/HTTP_STATIC_HTML_cftpaas4", "/Common/HTTP_STATIC_HTML_cftpaas4"},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetTemplateTransparentMode(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetTemplateTransparentMode(
		[]string{
			"/Common/HTTP_STATIC_HTML_cftpaas4", "/Common/HTTP_SSTS_YXP_80",
			"/Common/HTTP_OFST-APP-BANK", "/Common/HTTP_STATIC_HTML_GDMS_iaas",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}

func TestMonitor_GetIgnoreDownResponseState(t *testing.T) {
	p := New(newClient(t))

	arr, err := p.GetIgnoreDownResponseState(
		[]string{
			"/Common/HTTP_STATIC_HTML_cftpaas4", "/Common/HTTP_STATIC_HTML_OCTS-XNV",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", arr)
}
