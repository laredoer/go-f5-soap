package go_f5_soap

import "encoding/xml"

type BaseEnvEnvelope struct {
	XMLName xml.Name `xml:"env:Envelope" json:"envelope,omitempty"`
	Xsd     string   `xml:"xmlns:xsd,attr" json:"xsd,omitempty"`
	Xsi     string   `xml:"xmlns:xsi,attr" json:"xsi,omitempty"`
	Tns     string   `xml:"xmlns:tns,attr" json:"tns,omitempty"`
	Env     string   `xml:"xmlns:env,attr" json:"env,omitempty"`
	Ins0    string   `xml:"xmlns:ins0,attr" json:"ins0,omitempty"`
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
