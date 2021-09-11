package topology

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/Topology"

// ITopology The Topology interface enables you to work with topology attributes.
// For example, you can create and delete a topology.
// You can also use the Topology interface to add virtual server entries to,
// or remove virtual server entries from, a topology.
type ITopology interface {
	GetList() ([]TopologyRecord, error)
	GetOrder(records []TopologyRecord) ([]int64, error)
}

var _ ITopology = (*Topology)(nil)

type Topology struct {
	c *soap.Client
}

func New(c *soap.Client) *Topology {
	return &Topology{c: c}
}

type TopologyRecord struct {
	Server TopologyEndpoint
	LDns   TopologyEndpoint
}

type TopologyEndpoint struct {
	Type    global_lb.RegionType
	Content string
	Negate  bool
}

type GetListBody struct {
	GetList struct{} `xml:"tns:get_list"`
}

type ListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []struct {
					Server struct {
						Text string `xml:",chardata"`
					} `xml:"address"`
					Ldns struct {
						Text int64 `xml:",chardata"`
					} `xml:"port"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

// GetList Gets a list of of topology records.
func (t *Topology) GetList() ([]TopologyRecord, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetListBody `xml:"env:Body"`
	}

	bt, err := t.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetListBody{GetList: struct{}{}},
	})
	if err != nil {
		return nil, err
	}

	var resp ListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return nil, err
}

type GetOrderBody struct {
	GetOrder GetOrder `xml:"tns:get_order"`
}

type GetOrder struct {
	Records Records `xml:"records"`
}

type Records struct {
	Item []TopologyRecord `xml:"item"`
}

type OrderResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetOrderResponse struct {
			Return struct {
				Item []int64 `xml:"item"`
			} `xml:"return"`
		} `xml:"get_orderResponse"`
	} `xml:"Body"`
}

// GetOrder Gets the sort orders for the specified topology records.
func (t *Topology) GetOrder(records []TopologyRecord) ([]int64, error) {

	type Req struct {
		soap.BaseEnvEnvelope
		Body GetOrderBody `xml:"env:Body"`
	}

	bt, err := t.c.Call(context.Background(), Req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetOrderBody{GetOrder: GetOrder{Records: Records{Item: records}}},
	})
	if err != nil {
		return nil, err
	}

	var resp OrderResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetOrderResponse.Return.Item, nil
}
