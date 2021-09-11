package virtual_server_v2

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
)

const tns = "urn:iControl:GlobalLB/VirtualServerV2"

// IVirtualServerV2
//Introduced : BIG-IP_v11.0.0
//The VirtualServer interface enables you to work with virtual servers associated with a server.
type IVirtualServerV2 interface {
	GetAddress(virtualServers []VirtualServerID) ([]common.IPPortDefinition, error)
}

var _ IVirtualServerV2 = (*VirtualServerV2)(nil)

type VirtualServerV2 struct {
	c *soap.Client
}

func New(c *soap.Client) *VirtualServerV2 {
	return &VirtualServerV2{c: c}
}

type VirtualServerID struct {
	Name   string `xml:"name"`
	Server string `xml:"server"`
}

type GetAddressBody struct {
	GetAddress GetAddress `xml:"tns:get_address"`
}

type GetAddress struct {
	VirtualServers VirtualServers `xml:"virtual_servers"`
}

type VirtualServers struct {
	Item []VirtualServerID `xml:"item"`
}

type AddressResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetAddressResponse struct {
			Return struct {
				Item []struct {
					Address struct {
						Text string `xml:",chardata"`
					} `xml:"address"`
					Port struct {
						Text int64 `xml:",chardata"`
					} `xml:"port"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_addressResponse"`
	} `xml:"Body"`
}

// GetAddress Gets the IP address and service associated with a set of virtual servers.
// Note: A set_address method is not supported.
func (v *VirtualServerV2) GetAddress(virtualServers []VirtualServerID) ([]common.IPPortDefinition, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetAddressBody `xml:"env:Body"`
	}

	bt, err := v.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetAddressBody{GetAddress: GetAddress{VirtualServers: VirtualServers{Item: virtualServers}}},
	})
	if err != nil {
		return nil, err
	}

	var resp AddressResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []common.IPPortDefinition
	for _, v := range resp.Body.GetAddressResponse.Return.Item {
		res = append(res, common.IPPortDefinition{
			Address: v.Address.Text,
			Port:    v.Port.Text,
		})
	}

	return res, nil
}
