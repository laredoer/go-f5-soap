package prober_pool

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:GlobalLB/ProberPool"

// IProberPool
// Introduced : BIG-IP_v11.0.0
// The ProberPool interface enables you to create and maintain a prober pool,
// which contains a set of members which will be used to monitor the Server&aposs
// resources (typically its virtual servers). The members named in the prober pool should be BIG-IP systems,
// set up as Servers of type bigip-standalone or bigip-redundant.
// If a Server doesn&apost name a prober pool,
// the Server will use the prober pool attached to the server&aposs data center,
// if that has been assigned. The probing members named in the prober pool will be chosen according to
// the load balancing method selected for the prober pool (e.g., round robin or global availability).
type IProberPool interface {
	GetList() ([]string, error)
	GetMember(pools []string) ([][]string, error)
	GetMemberOrder(pools []string, members [][]string) ([][]int64, error)
}

var _ IProberPool = (*ProberPool)(nil)

type ProberPool struct {
	c *soap.Client
}

func New(c *soap.Client) *ProberPool {
	return &ProberPool{c: c}
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
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

func (p *ProberPool) GetList() ([]string, error) {

	bt, err := p.c.Call(context.Background(), getListReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getListBody{},
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

type getMemberReq struct {
	soap.BaseEnvEnvelope
	Body getMemberBody `xml:"env:Body"`
}

type getMemberBody struct {
	GetMember getMember `xml:"tns:get_member"`
}

type getMember struct {
	Pools struct {
		Item []string `xml:"item"`
	} `xml:"pools"`
}

type getMemberResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMemberResponse struct {
			Return struct {
				Item []struct {
					Item []string `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_memberResponse"`
	} `xml:"Body"`
}

func (p *ProberPool) GetMember(pools []string) ([][]string, error) {

	bt, err := p.c.Call(context.Background(), getMemberReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getMemberBody{GetMember: getMember{Pools: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: pools})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getMemberResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]string
	for _, v := range resp.Body.GetMemberResponse.Return.Item {
		res = append(res, v.Item)
	}

	return res, nil
}

type getMemberOrderReq struct {
	soap.BaseEnvEnvelope
	Body getMemberOrderBody `xml:"env:Body"`
}

type getMemberOrderBody struct {
	GetMemberOrder getMemberOrder `xml:"tns:get_member_order"`
}

type getMemberOrder struct {
	Pools struct {
		Item []string `xml:"item"`
	} `xml:"pools"`
	Members Members `xml:"members"`
}

type Members struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Item []string `xml:"item"`
}

type getMemberOrderResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMemberOrderResponse struct {
			Return struct {
				Item []struct {
					Item []int64 `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_member_orderResponse"`
	} `xml:"Body"`
}

func (p *ProberPool) GetMemberOrder(pools []string, members [][]string) ([][]int64, error) {

	var memberItem []Item
	for _, v := range members {
		item := Item{Item: []string{}}
		item.Item = append(item.Item, v...)
		memberItem = append(memberItem, item)
	}

	bt, err := p.c.Call(context.Background(), getMemberOrderReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getMemberOrderBody{GetMemberOrder: getMemberOrder{
			Pools: struct {
				Item []string `xml:"item"`
			}{Item: pools},
			Members: Members{Item: memberItem},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getMemberOrderResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]int64
	for _, v := range resp.Body.GetMemberOrderResponse.Return.Item {
		res = append(res, v.Item)
	}

	return res, nil
}
