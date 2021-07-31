package pool_v2

import (
	"context"
	"encoding/xml"
	go_f5_soap "github.com/wule61/go-f5-soap"
)

const tns = "urn:iControl:GlobalLB/PoolV2"

type GTMQueryType string

const (
	GTM_QUERY_TYPE_UNKNOWN GTMQueryType = "GTM_QUERY_TYPE_UNKNOWN"
	GTM_QUERY_TYPE_A       GTMQueryType = "GTM_QUERY_TYPE_A"
	GTM_QUERY_TYPE_CNAME   GTMQueryType = "GTM_QUERY_TYPE_CNAME"
	GTM_QUERY_TYPE_AAAA    GTMQueryType = "GTM_QUERY_TYPE_AAAA"
	GTM_QUERY_TYPE_SRV     GTMQueryType = "GTM_QUERY_TYPE_SRV"
	GTM_QUERY_TYPE_NAPTR   GTMQueryType = "GTM_QUERY_TYPE_NAPTR"
)

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
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMemberResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						Name struct {
							Text string `xml:",chardata"`
						} `xml:"name"`
						Server struct {
							Text string `xml:",chardata"`
						} `xml:"server"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_memberResponse"`
	} `xml:"Body"`
}

type Member struct {
	Name   string
	Server string
}

func (p *PoolV2) GetMember(pools []PoolID) ([][]Member, error) {

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

	var res [][]Member
	for _, v := range resp.Body.GetMemberResponse.Return.Item {
		var poolMembers []Member
		for _, v2 := range v.Item {
			poolMembers = append(poolMembers, Member{
				Name:   v2.Name.Text,
				Server: v2.Server.Text,
			})
		}
		res = append(res, poolMembers)
	}

	return res, nil
}

type GetListByTypeBody struct {
	GetListByType GetListByType `xml:"tns:get_list_by_type"`
}

type GetListByType struct {
	Types Types `xml:"types"`
}

type Types struct {
	Item []GTMQueryType `xml:"item"`
}

type ListByTypeResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListByTypeResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						PoolName struct {
							Text string `xml:",chardata" `
						} `xml:"pool_name" `
						PoolType struct {
							Text string `xml:",chardata" `
						} `xml:"pool_type" `
					} `xml:"item" `
				} `xml:"item" `
			} `xml:"return"`
		} `xml:"get_list_by_typeResponse"`
	} `xml:"Body" `
}

func (p *PoolV2) GetListByType(gtmQueryType []GTMQueryType) ([][]PoolID, error) {

	type req struct {
		go_f5_soap.BaseEnvEnvelope
		Body GetListByTypeBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: go_f5_soap.NewBaseEnvEnvelope(tns),
		Body: GetListByTypeBody{GetListByType: GetListByType{
			Types: Types{
				Item: gtmQueryType,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp ListByTypeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]PoolID
	for _, v := range resp.Body.GetListByTypeResponse.Return.Item {
		var poolIDArr []PoolID
		for _, v2 := range v.Item {
			poolIDArr = append(poolIDArr, PoolID{
				PoolName: v2.PoolName.Text,
				PoolType: v2.PoolType.Text,
			})
		}
		res = append(res, poolIDArr)
	}

	return res, nil
}
