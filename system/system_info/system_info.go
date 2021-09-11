package system_info

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:System/SystemInfo"

// ISystemInfo
// Introduced : BIG-IP_v9.0
// The SystemInfo interface enables you to query identifying attributes of the system.
type ISystemInfo interface {
	GetVersion() (string, error)
	GetUpTime() (int64, error)
}

type SystemInfo struct {
	c *soap.Client
}

func New(c *soap.Client) ISystemInfo {
	return &SystemInfo{c: c}
}

type getVersionReq struct {
	soap.BaseEnvEnvelope
	Body getVersionBody `xml:"env:Body"`
}

type getVersionBody struct {
	GetVersion struct{} `xml:"tns:get_version"`
}

type getVersionResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetVersionResponse struct {
			Return struct {
				Text string `xml:",chardata"`
			} `xml:"return"`
		} `xml:"get_versionResponse"`
	} `xml:"Body"`
}

func (s *SystemInfo) GetVersion() (string, error) {

	bt, err := s.c.Call(context.Background(), getVersionReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getVersionBody{GetVersion: struct{}{}},
	})
	if err != nil {
		return "", err
	}

	var resp getVersionResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return "", err
	}

	return resp.Body.GetVersionResponse.Return.Text, err
}

type getUpTimeReq struct {
	soap.BaseEnvEnvelope
	Body getUpTimeBody `xml:"env:Body"`
}

type getUpTimeBody struct {
	GetUpTime struct{} `xml:"tns:get_uptime"`
}

type getUpTimeResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetUptimeResponse struct {
			Return struct {
				Text int64 `xml:",chardata"`
			} `xml:"return"`
		} `xml:"get_uptimeResponse"`
	} `xml:"Body"`
}

func (s *SystemInfo) GetUpTime() (int64, error) {

	bt, err := s.c.Call(context.Background(), getUpTimeReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getUpTimeBody{GetUpTime: struct{}{}},
	})
	if err != nil {
		return 0, err
	}

	var resp getUpTimeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return 0, err
	}

	return resp.Body.GetUptimeResponse.Return.Text, nil
}
