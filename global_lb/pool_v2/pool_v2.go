package pool_v2

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/PoolV2"

// The IPoolV2 interface enables you to work with typed pools and their attributes.
// Typed pools (like A, AAAA, CNAME, MX, SRV, and NAPTR) may contain members of the same type.
// This allows for greater flexibility and more granular control over the responses that GTM sends.
// Previously, in the GlobalLB::Pool interface, members were only virtual servers.
// Now, in the PoolV2 interface, members can be virtual servers (for type A or AAAA type pools)
// or non-terminal type members (for type CNAME, MX, SRV, or NAPTR type pools).
// Non-terminal members are specified with a required dname and some optional settings
// that depend upon the type of non-terminal member.
// Except in the case of a CNAME member with the static-target setting enabled,
// all non-terminal members are backed by a corresponding wide IP.
// The dname of the non-terminal member is non-folderized and must be a fully-qualified domain name.
// The dname of the non-terminal member must match (exactly or via wide IP wildcard match)
// the name of a corresponding wide IP (without the folder name),
// except in the case mentioned above where the non-terminal is a CNAME type member with static-target enabled.
type IPoolV2 interface {
	GetMember(pools []PoolID) ([][]Member, error)
	GetList() ([]PoolID, error)
	GetListByType(gtmQueryType []global_lb.GTMQueryType) ([][]PoolID, error)
	GetTTL(pools []PoolID) ([]int64, error)
	GetEnabledState(pools []PoolID) ([]common.EnabledState, error)
	GetObjectStatus(pools []PoolID) ([]common.ObjectStatus, error)
}

var _ IPoolV2 = (*PoolV2)(nil)

type PoolV2 struct {
	c *soap.Client
}

func New(c *soap.Client) *PoolV2 {
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
		soap.BaseEnvEnvelope
		Body GetMemberBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
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
	Item []global_lb.GTMQueryType `xml:"item"`
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

type GetListBody struct {
	GetList GetList `xml:"tns:get_list"`
}

type GetList struct {
}
type ListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []struct {
					PoolName struct {
						Text string `xml:",chardata" `
					} `xml:"pool_name" `
					PoolType struct {
						Text string `xml:",chardata" `
					} `xml:"pool_type" `
				} `xml:"item" `
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body" `
}

func (p *PoolV2) GetList() ([]PoolID, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetListBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetListBody{GetList: GetList{}},
	})
	if err != nil {
		return nil, err
	}

	var resp ListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []PoolID
	for _, v := range resp.Body.GetListResponse.Return.Item {
		res = append(res, PoolID{
			PoolName: v.PoolName.Text,
			PoolType: v.PoolType.Text,
		})
	}

	return res, nil

}

func (p *PoolV2) GetListByType(gtmQueryType []global_lb.GTMQueryType) ([][]PoolID, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetListByTypeBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
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

type GetTTLBody struct {
	GetTTL GetTTL `xml:"tns:get_ttl"`
}

type GetTTL struct {
	Pools Pools `xml:"pools"`
}

type TTLResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTtlResponse struct {
			Return struct {
				Item []int64 `xml:"item"`
			} `xml:"return"`
		} `xml:"get_ttlResponse"`
	} `xml:"Body"`
}

func (p *PoolV2) GetTTL(pools []PoolID) ([]int64, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetTTLBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTTLBody{GetTTL: GetTTL{
			Pools: Pools{Item: pools},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp TTLResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTtlResponse.Return.Item, nil
}

type GetEnabledStateBody struct {
	GetEnabledState GetEnabledState `xml:"tns:get_enabled_state"`
}

type GetEnabledState struct {
	Pools Pools `xml:"pools"`
}

type EnableStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetEnabledStateResponse struct {
			Return struct {
				Item []common.EnabledState `xml:"item"`
			} `xml:"return"`
		} `xml:"get_enabled_stateResponse"`
	} `xml:"Body"`
}

func (p *PoolV2) GetEnabledState(pools []PoolID) ([]common.EnabledState, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetEnabledStateBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetEnabledStateBody{GetEnabledState: GetEnabledState{Pools: Pools{
			Item: pools,
		}}},
	})
	if err != nil {
		return nil, err
	}

	var resp EnableStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetEnabledStateResponse.Return.Item, nil

}

type GetObjectStatusBody struct {
	GetObjectStatus GetObjectStatus `xml:"tns:get_object_status"`
}

type GetObjectStatus struct {
	Pools Pools `xml:"pools"`
}

type ObjectStatusResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetObjectStatusResponse struct {
			Return struct {
				Item []struct {
					Text               string `xml:",chardata" `
					AvailabilityStatus struct {
						Text common.AvailabilityStatus `xml:",chardata"`
					} `xml:"availability_status"`
					EnabledStatus struct {
						Text common.EnabledStatus `xml:",chardata" `
					} `xml:"enabled_status" `
					StatusDescription struct {
						Text string `xml:",chardata"`
					} `xml:"status_description"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_object_statusResponse"`
	} `xml:"Body"`
}

func (p *PoolV2) GetObjectStatus(pools []PoolID) ([]common.ObjectStatus, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetObjectStatusBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetObjectStatusBody{GetObjectStatus: GetObjectStatus{Pools: Pools{Item: pools}}},
	})
	if err != nil {
		return nil, err
	}

	var resp ObjectStatusResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []common.ObjectStatus
	for _, v := range resp.Body.GetObjectStatusResponse.Return.Item {
		res = append(res, common.ObjectStatus{
			AvailabilityStatus: v.AvailabilityStatus.Text,
			EnabledStatus:      v.EnabledStatus.Text,
			StatusDescription:  v.StatusDescription.Text,
		})
	}

	return res, nil
}
