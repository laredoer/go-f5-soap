package pool

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/Pool"

type GetAlternateLbMethodBody struct {
	GetAlternateLbMethod GetAlternateLbMethod `xml:"tns:get_alternate_lb_method"`
}

type GetAlternateLbMethod struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type PoolNames struct {
	Item []string `xml:"item"`
}

type MonitorAssociation struct {
	PoolName    string                `xml:"pool_name"`
	MonitorRule global_lb.MonitorRule `xml:"monitor_rule"`
}

// IPool The Pool interface enables you to work with pools and their attributes.
// Introduced : BIG-IP_v9.2.0
type IPool interface {
	GetList() (poolNames []string, err error)
	GetMemberV2(poolNames []string) (virtualServerIDs [][]global_lb.VirtualServerID, err error)
	GetMemberRatio(poolNames []string, members [][]global_lb.VirtualServerID) ([][]int64, error)
	GetMonitorAssociation(poolNames []string) ([]MonitorAssociation, error)
	GetAlternateLBMethod(poolNames []string) ([]string, error)
	GetPreferredLBMethod(poolNames []string) ([]string, error)
	GetTTL(poolNames []string) ([]int64, error)
	GetVerifyMemberAvailabilityState(poolNames []string) ([]string, error)
	GetAnswersToReturn(poolNames []string) ([]int64, error)
	GetObjectStatus(poolNames []string) ([]common.ObjectStatus, error)
	GetEnabledState(poolNames []string) ([]common.EnabledState, error)
}

var _ IPool = (*Client)(nil)

type Client struct {
	c *soap.Client
}

func New(c *soap.Client) *Client {
	return &Client{
		c: c,
	}
}

type GetMonitorAssociationBody struct {
	GetMonitorAssociation GetMonitorAssociation `xml:"tns:get_monitor_association"`
}

type GetMonitorAssociation struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type GetMonitorAssociationResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMonitorAssociationResponse struct {
			Return struct {
				Item []struct {
					PoolName struct {
						Text string `xml:",chardata"`
					} `xml:"pool_name"`
					MonitorRule struct {
						Type struct {
							Text global_lb.MonitorRuleType `xml:",chardata"`
						} `xml:"type"`
						Quorum struct {
							Text int64 `xml:",chardata"`
						} `xml:"quorum"`
						MonitorTemplates struct {
							Item []string `xml:"item"`
						} `xml:"monitor_templates"`
					} `xml:"monitor_rule"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_monitor_associationResponse"`
	} `xml:"Body"`
}

func (p *Client) GetMonitorAssociation(poolNames []string) ([]MonitorAssociation, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetMonitorAssociationBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetMonitorAssociationBody{GetMonitorAssociation: GetMonitorAssociation{
			PoolNames: PoolNames{
				Item: poolNames,
			},
		}},
	})

	if err != nil {
		return nil, err
	}

	var resp GetMonitorAssociationResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []MonitorAssociation
	for _, v := range resp.Body.GetMonitorAssociationResponse.Return.Item {
		res = append(res, MonitorAssociation{
			PoolName: v.PoolName.Text,
			MonitorRule: global_lb.MonitorRule{
				Type:             v.MonitorRule.Type.Text,
				Quorum:           v.MonitorRule.Quorum.Text,
				MonitorTemplates: v.MonitorRule.MonitorTemplates.Item,
			},
		})
	}

	return res, nil

}

type GetListBody struct {
	GetList struct{} `xml:"tns:get_list"`
}

type ListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

func (p *Client) GetList() (poolNames []string, err error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetListBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetListBody{struct{}{}},
	})

	if err != nil {
		return nil, err
	}

	var resp ListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetListResponse.Return.Item, nil
}

type GetMemberV2Body struct {
	GetMemberV2 GetMemberV2 `xml:"tns:get_member_v2"`
}

type GetMemberV2 struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type GetMemberV2Resp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMemberV2Response struct {
			Return struct {
				Item []struct {
					Item []struct {
						Name struct {
							Text string `xml:",chardata"`
						} `xml:"name" json:"name,omitempty"`
						Server struct {
							Text string `xml:",chardata"`
						} `xml:"server"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_member_v2Response"`
	} `xml:"Body"`
}

func (p *Client) GetMemberV2(poolNames []string) (virtualServerIDs [][]global_lb.VirtualServerID, err error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetMemberV2Body `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetMemberV2Body{GetMemberV2: GetMemberV2{
			PoolNames: PoolNames{
				Item: poolNames,
			},
		}},
	})

	if err != nil {
		return nil, err
	}

	var resp GetMemberV2Resp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	for _, v := range resp.Body.GetMemberV2Response.Return.Item {
		var vsItem []global_lb.VirtualServerID
		for _, v2 := range v.Item {
			vsItem = append(vsItem, global_lb.VirtualServerID{
				Name:   v2.Name.Text,
				Server: v2.Server.Text,
			})
		}
		virtualServerIDs = append(virtualServerIDs, vsItem)
	}

	return virtualServerIDs, nil
}

type AlternateLbMethodByPoolNamesResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetAlternateLbMethodResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_alternate_lb_methodResponse"`
	} `xml:"Body"`
}

func (p *Client) GetAlternateLBMethod(poolNames []string) ([]string, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetAlternateLbMethodBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetAlternateLbMethodBody{GetAlternateLbMethod{PoolNames{Item: poolNames}}},
	})

	if err != nil {
		return nil, err
	}

	var resp AlternateLbMethodByPoolNamesResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetAlternateLbMethodResponse.Return.Item, nil
}

type GetPreferredLBMethodBody struct {
	GetPreferredLBMethod GetPreferredLBMethod `xml:"tns:get_preferred_lb_method"`
}

type GetPreferredLBMethod struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type PreferredLBMethodByPoolNamesResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetPreferredLbMethodResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_preferred_lb_methodResponse"`
	} `xml:"Body"`
}

func (p *Client) GetPreferredLBMethod(poolNames []string) ([]string, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetPreferredLBMethodBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetPreferredLBMethodBody{GetPreferredLBMethod{PoolNames{Item: poolNames}}},
	})

	if err != nil {
		return nil, err
	}

	var resp PreferredLBMethodByPoolNamesResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetPreferredLbMethodResponse.Return.Item, nil
}

type GetTTLBody struct {
	GetTTL GetTTL `xml:"tns:get_ttl"`
}

type GetTTL struct {
	PoolNames PoolNames `xml:"pool_names"`
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

func (p *Client) GetTTL(poolNames []string) ([]int64, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetTTLBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetTTLBody{GetTTL{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp TTLResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTtlResponse.Return.Item, err
}

type GetVerifyMemberAvailabilityStateBody struct {
	GetVerifyMemberAvailabilityState GetVerifyMemberAvailabilityState `xml:"tns:get_verify_member_availability_state"`
}

type GetVerifyMemberAvailabilityState struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type VerifyMemberAvailabilityStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetVerifyMemberAvailabilityStateResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_verify_member_availability_stateResponse"`
	} `xml:"Body"`
}

func (p *Client) GetVerifyMemberAvailabilityState(poolNames []string) ([]string, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetVerifyMemberAvailabilityStateBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetVerifyMemberAvailabilityStateBody{GetVerifyMemberAvailabilityState{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp VerifyMemberAvailabilityStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetVerifyMemberAvailabilityStateResponse.Return.Item, err
}

type GetAnswersToReturnBody struct {
	GetAnswersToReturn GetAnswersToReturn `xml:"tns:get_answers_to_return"`
}

type GetAnswersToReturn struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type AnswersToReturnResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetAnswersToReturnResponse struct {
			Return struct {
				Item []int64 `xml:"item"`
			} `xml:"return"`
		} `xml:"get_answers_to_returnResponse"`
	} `xml:"Body"`
}

func (p *Client) GetAnswersToReturn(poolNames []string) ([]int64, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetAnswersToReturnBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetAnswersToReturnBody{GetAnswersToReturn{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp AnswersToReturnResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetAnswersToReturnResponse.Return.Item, err
}

type GetObjectStatusBody struct {
	GetObjectStatus GetObjectStatus `xml:"tns:get_object_status"`
}

type GetObjectStatus struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type ObjectStatusResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetObjectStatusResponse struct {
			Return struct {
				Item []struct {
					AvailabilityStatus struct {
						Text common.AvailabilityStatus `xml:",chardata"`
					} `xml:"availability_status"`
					EnabledStatus struct {
						Text common.EnabledStatus `xml:",chardata"`
					} `xml:"enabled_status"`
					StatusDescription struct {
						Text string `xml:",chardata"`
					} `xml:"status_description"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_object_statusResponse"`
	} `xml:"Body"`
}

func (p *Client) GetObjectStatus(poolNames []string) ([]common.ObjectStatus, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetObjectStatusBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetObjectStatusBody{GetObjectStatus{PoolNames{Item: poolNames}}},
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

type GetEnabledStateBody struct {
	GetEnabledState GetEnabledState `xml:"tns:get_enabled_state"`
}

type GetEnabledState struct {
	PoolNames PoolNames `xml:"pool_names"`
}

type EnabledStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetEnabledStateResponse struct {
			Return struct {
				Item []common.EnabledState `xml:"item"`
			} `xml:"return"`
		} `xml:"get_enabled_stateResponse"`
	} `xml:"Body"`
}

func (p *Client) GetEnabledState(poolNames []string) ([]common.EnabledState, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetEnabledStateBody `xml:"env:Body"`
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetEnabledStateBody{GetEnabledState{PoolNames{Item: poolNames}}},
	})
	if err != nil {
		return nil, err
	}

	var resp EnabledStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetEnabledStateResponse.Return.Item, nil
}

type getMemberRatioReq struct {
	soap.BaseEnvEnvelope
	Body getMemberRatioBody `xml:"env:Body"`
}

type getMemberRatioBody struct {
	GetMemberRatio getMemberRatio `xml:"tns:get_member_ratio"`
}

type getMemberRatio struct {
	PoolNames PoolNames `xml:"pool_names"`
	Members   Members   `xml:"members"`
}

type Members struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Item []global_lb.VirtualServerID `xml:"item"`
}

type getMemberRatioResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMemberRatioResponse struct {
			Return struct {
				Item []struct {
					Item []int64 `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_member_ratioResponse"`
	} `xml:"Body"`
}

// GetMemberRatio
// Introduced : BIG-IP_v11.0.0
// Gets the ratios for the specified members of the specified pools.
func (p *Client) GetMemberRatio(poolNames []string, members [][]global_lb.VirtualServerID) ([][]int64, error) {

	var memberItem []Item
	for _, v := range members {
		item := Item{Item: []global_lb.VirtualServerID{}}
		item.Item = append(item.Item, v...)
		memberItem = append(memberItem, item)
	}

	bt, err := p.c.Call(context.Background(), getMemberRatioReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getMemberRatioBody{GetMemberRatio: getMemberRatio{
			PoolNames: PoolNames{Item: poolNames},
			Members:   Members{Item: memberItem},
		},
		},
	})
	if err != nil {
		return nil, err
	}

	var resp getMemberRatioResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]int64
	for _, v := range resp.Body.GetMemberRatioResponse.Return.Item {
		res = append(res, v.Item)
	}

	return res, nil
}
