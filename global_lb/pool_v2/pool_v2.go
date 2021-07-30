package pool_v2

import (
	"context"
	"encoding/xml"
	go_f5_soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:GlobalLB/PoolV2"

type PoolV2 struct {
	c *go_f5_soap.Client
}

func NewPoolV2(c *go_f5_soap.Client) *PoolV2 {
	return &PoolV2{
		c: c,
	}
}

type PoolID struct {
	PoolName string `xml:"pool_name"`
	PoolType string `xml:"pool_type"`
}

type GetMemberBody struct {
	GetMember GetMember `xml:"tns:get_member"`
}

type GetMember struct {
	Pools Pools `xml:"pools"`
}

type Pools struct {
	Item []PoolID `xml:"item"`
}

type MemberResp struct {
	XMLName xml.Name `xml:"Envelope" json:"envelope,omitempty"`
	Body    struct {
		GetMemberResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"name" json:"name,omitempty"`
						Server struct {
							Text string `xml:",chardata" json:"text,omitempty"`
						} `xml:"server" json:"server,omitempty"`
					} `xml:"item" json:"item,omitempty"`
				} `xml:"item" json:"item,omitempty"`
			} `xml:"return" json:"return,omitempty"`
		} `xml:"get_memberResponse" json:"get_memberresponse,omitempty"`
	} `xml:"Body" json:"body,omitempty"`
}

type Member struct {
	Name   string
	Server string
}

func (p *PoolV2) GetMember(pools []PoolID) ([]Member, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetMemberBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body: GetMemberBody{GetMember: GetMember{Pools: Pools{
			Item: pools,
		}}},
	})
	if err != nil {
		return nil, err
	}

	var resp MemberResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []Member
	for _, v := range resp.Body.GetMemberResponse.Return.Item {
		for _, v2 := range v.Item {
			res = append(res, Member{
				Name:   v2.Name.Text,
				Server: v2.Server.Text,
			})
		}
	}

	return res, nil
}
