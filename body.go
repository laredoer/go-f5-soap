package go_f5_soap

import "encoding/xml"

type BaseEnvEnvelope struct {
	XMLName xml.Name `xml:"env:Envelope"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Tns     string   `xml:"xmlns:tns,attr"`
	Env     string   `xml:"xmlns:env,attr"`
	Ins0    string   `xml:"xmlns:ins0,attr"`
}

func NewBaseEnvEnvelope(tns string) BaseEnvEnvelope {
	return BaseEnvEnvelope{
		XMLName: xml.Name{},
		Xsd:     "http://www.w3.org/2001/XMLSchema",
		Xsi:     "http://www.w3.org/2001/XMLSchema-instance",
		Tns:     tns,
		Env:     "http://schemas.xmlsoap.org/soap/envelope/",
		Ins0:    "urn:iControl",
	}
}
