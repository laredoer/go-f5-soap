package data_center

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:GlobalLB/DataCenter"

// IDataCenter
// Introduced : BIG-IP_v9.2.0
// The DataCenter interface enables you to manipulate the data center attributes for a Global TM.
// For example, use the DataCenter interface to add or remove a data center,
// transfer server assignments from one data center to another,
// get and set data center attributes,
// remove a server from a data center, and so on.
type IDataCenter interface {
	GetList() ([]string, error)
	GetServer(dataCenters []string) ([]DataCenterServerDefinition, error)
}

// DataCenterServerDefinition
// Introduced : BIG-IP_v9.2.0
// A struct that contains definition for the data center and the associated servers.
type DataCenterServerDefinition struct {
	DataCenter string   // The name that identifies a data center.
	Servers    []string // The servers in the data center.
}

var _ IDataCenter = (*Client)(nil)

type Client struct {
	c *soap.Client
}

func New(c *soap.Client) *Client {
	return &Client{c: c}
}

type getListReq struct {
	soap.BaseEnvEnvelope
	Body getListBody `xml:"env:Body"`
}

type getListBody struct {
	GetList struct{} `xml:"tns:get_list"`
}

type getListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return" `
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

func (d *Client) GetList() ([]string, error) {

	bt, err := d.c.Call(context.Background(), getListReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getListBody{GetList: struct{}{}},
	})
	if err != nil {
		return nil, err
	}

	var resp getListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetListResponse.Return.Item, nil
}

type getServerReq struct {
	soap.BaseEnvEnvelope
	Body getServerBody `xml:"env:Body"`
}

type getServerBody struct {
	GetServer getServer `xml:"tns:get_server"`
}

type getServer struct {
	DataCenters DataCenters `xml:"data_centers"`
}

type DataCenters struct {
	Item []string `xml:"item"`
}

type getServerResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetServerResponse struct {
			Return struct {
				Item []struct {
					DataCenter struct {
						Text string `xml:",chardata"`
					} `xml:"data_center"`
					Servers struct {
						Item []string `xml:"item"`
					} `xml:"servers"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_serverResponse"`
	} `xml:"Body"`
}

// GetServer
// Introduced : BIG-IP_v9.2.0
// Gets a list of servers of the specified data centers.
func (d *Client) GetServer(dataCenters []string) ([]DataCenterServerDefinition, error) {

	bt, err := d.c.Call(context.Background(), getServerReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getServerBody{GetServer: getServer{DataCenters: DataCenters{dataCenters}}},
	})
	if err != nil {
		return nil, err
	}

	var resp getServerResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []DataCenterServerDefinition
	for _, v := range resp.Body.GetServerResponse.Return.Item {
		res = append(res, DataCenterServerDefinition{
			DataCenter: v.DataCenter.Text,
			Servers:    v.Servers.Item,
		})
	}

	return res, nil
}
