package wide_ip_v2

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/WideIPV2"

// IWideIPV2
// Introduced : BIG-IP_v12.0.0
// he WideIPV2 interface enables you to work with typed wide IPs and their settings,
// along with their typed wide IP pools, wide IP aliases and wide IP rules.
// This interface allows for greater flexibility and more granular control over the responses that GTM sends.
// Typed wide IPs (like A, AAAA, CNAME, MX, SRV, and NAPTR) may contain pools only of the same type.
// There is one exception to this rule,
// where CNAME pools can actually be added to a wide IP of any type,
// not just CNAME type wide IPs.
// Typed wide IPs can also be used to back non-terminal pool members.
// While previously in the GlobalLB::Pool interface, members were only virtual servers,
// in the PoolV2 interface we also have non-terminal type members
// which must correspond (exactly or via wide IP wildcard match) to an existing wide IP of the appropriate type.
// For the specifics of which type of wide IP is allowed to back which type of non-terminal pool member,
// please see the GlobalLB::PoolV2 documentation.
type IWideIPV2 interface {
	GetList() ([]global_lb.WideIPID, error)
	GetListByType(types []global_lb.GTMQueryType) ([][]global_lb.WideIPID, error)
	GetWideIpPool(wideIPs []global_lb.WideIPID) ([][]global_lb.PoolID, error)
	GetLBMethod(wideIPs []global_lb.WideIPID) ([]global_lb.LBMethod, error)
	GetObjectStatus(wideIPs []global_lb.WideIPID) ([]common.ObjectStatus, error)
	GetEnabledState(wideIPs []global_lb.WideIPID) ([]common.EnabledState, error)
	GetWideIpPoolRatio([]global_lb.WideIPID, [][]global_lb.PoolID) ([][]int64, error)
}

var _ IWideIPV2 = (*WideIPV2)(nil)

type WideIPV2 struct {
	c *soap.Client
}

func New(c *soap.Client) *WideIPV2 {
	return &WideIPV2{c: c}
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
				Item []struct {
					WideipName struct {
						Text string `xml:",chardata"`
					} `xml:"wideip_name"`
					WideipType struct {
						Text global_lb.GTMQueryType `xml:",chardata"`
					} `xml:"wideip_type"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

func (w *WideIPV2) GetList() ([]global_lb.WideIPID, error) {

	bt, err := w.c.Call(context.Background(), getListReq{
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

	var res []global_lb.WideIPID
	for _, v := range resp.Body.GetListResponse.Return.Item {
		res = append(res, global_lb.WideIPID{
			WideIPName: v.WideipName.Text,
			WideIPType: v.WideipType.Text,
		})
	}

	return res, nil
}

type getListByTypeReq struct {
	soap.BaseEnvEnvelope
	Body getListByTypeBody `xml:"env:Body"`
}

type getListByTypeBody struct {
	GetListByType getListByType `xml:"tns:get_list_by_type"`
}

type getListByType struct {
	Types struct {
		Item []global_lb.GTMQueryType `xml:"item"`
	} `xml:"types"`
}

type getListByTypeResp struct {
	XMLName xml.Name `xml:"Envelope" json:"envelope,omitempty"`
	Body    struct {
		GetListByTypeResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						WideipName struct {
							Text string `xml:",chardata"`
						} `xml:"wideip_name"`
						WideipType struct {
							Text global_lb.GTMQueryType `xml:",chardata"`
						} `xml:"wideip_type"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_list_by_typeResponse"`
	} `xml:"Body"`
}

func (w *WideIPV2) GetListByType(types []global_lb.GTMQueryType) ([][]global_lb.WideIPID, error) {

	bt, err := w.c.Call(context.Background(), getListByTypeReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getListByTypeBody{GetListByType: getListByType{Types: struct {
			Item []global_lb.GTMQueryType `xml:"item"`
		}(struct{ Item []global_lb.GTMQueryType }{Item: types})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getListByTypeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]global_lb.WideIPID
	for _, v := range resp.Body.GetListByTypeResponse.Return.Item {
		var resItem []global_lb.WideIPID
		for _, v2 := range v.Item {
			resItem = append(resItem, global_lb.WideIPID{
				WideIPName: v2.WideipName.Text,
				WideIPType: v2.WideipType.Text,
			})
		}
		res = append(res, resItem)
	}

	return res, nil

}

type getWideIpPoolReq struct {
	soap.BaseEnvEnvelope
	Body getWideIpPoolBody `xml:"env:Body"`
}

type getWideIpPoolBody struct {
	GetWideIpPool getWideIpPool `xml:"tns:get_wide_ip_pool"`
}

type getWideIpPool struct {
	WideIps struct {
		Item []global_lb.WideIPID `xml:"item"`
	} `xml:"wide_ips"`
}

type getWideIpPoolResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetWideIpPoolResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						PoolName struct {
							Text string `xml:",chardata"`
						} `xml:"pool_name"`
						PoolType struct {
							Text global_lb.GTMQueryType `xml:",chardata"`
						} `xml:"pool_type"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_wide_ip_poolResponse"`
	} `xml:"Body"`
}

func (w *WideIPV2) GetWideIpPool(wideIPs []global_lb.WideIPID) ([][]global_lb.PoolID, error) {

	bt, err := w.c.Call(context.Background(), getWideIpPoolReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getWideIpPoolBody{GetWideIpPool: getWideIpPool{WideIps: struct {
			Item []global_lb.WideIPID `xml:"item"`
		}(struct{ Item []global_lb.WideIPID }{Item: wideIPs})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getWideIpPoolResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]global_lb.PoolID
	for _, v := range resp.Body.GetWideIpPoolResponse.Return.Item {
		var resItem []global_lb.PoolID
		for _, v2 := range v.Item {
			resItem = append(resItem, global_lb.PoolID{
				PoolName: v2.PoolName.Text,
				PoolType: v2.PoolType.Text,
			})
		}
		res = append(res, resItem)
	}

	return res, nil
}

type getLBMethodReq struct {
	soap.BaseEnvEnvelope
	Body GetLBMethodBody `xml:"env:Body"`
}

type GetLBMethodBody struct {
	GetLBMethod getLBMethod `xml:"tns:get_lb_method"`
}

type getLBMethod struct {
	WideIps struct {
		Item []global_lb.WideIPID `xml:"item"`
	} `xml:"wide_ips"`
}

type getLBMethodResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetLbMethodResponse struct {
			Return struct {
				Item []global_lb.LBMethod `xml:"item"`
			} `xml:"return"`
		} `xml:"get_lb_methodResponse"`
	} `xml:"Body"`
}

func (w *WideIPV2) GetLBMethod(wideIPs []global_lb.WideIPID) ([]global_lb.LBMethod, error) {

	bt, err := w.c.Call(context.Background(), getLBMethodReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetLBMethodBody{GetLBMethod: getLBMethod{WideIps: struct {
			Item []global_lb.WideIPID `xml:"item"`
		}(struct{ Item []global_lb.WideIPID }{Item: wideIPs})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getLBMethodResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetLbMethodResponse.Return.Item, nil
}

type getObjectStatusReq struct {
	soap.BaseEnvEnvelope
	Body getObjectStatusBody `xml:"env:Body"`
}

type getObjectStatusBody struct {
	GetObjectStatus getObjectStatus `xml:"tns:get_object_status"`
}

type getObjectStatus struct {
	WideIPs struct {
		Item []global_lb.WideIPID `xml:"item"`
	} `xml:"wide_ips"`
}

type getObjectStatusResp struct {
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

func (w *WideIPV2) GetObjectStatus(wideIPs []global_lb.WideIPID) ([]common.ObjectStatus, error) {

	bt, err := w.c.Call(context.Background(), getObjectStatusReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getObjectStatusBody{GetObjectStatus: getObjectStatus{struct {
			Item []global_lb.WideIPID `xml:"item"`
		}(struct{ Item []global_lb.WideIPID }{Item: wideIPs})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getObjectStatusResp
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

type getEnabledStateReq struct {
	soap.BaseEnvEnvelope
	Body getEnabledStateBody `xml:"env:Body"`
}

type getEnabledStateBody struct {
	GetEnabledState getEnabledState `xml:"tns:get_enabled_state"`
}

type getEnabledState struct {
	WideIPs struct {
		Item []global_lb.WideIPID `xml:"item"`
	} `xml:"wide_ips"`
}

type getEnabledStateResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetEnabledStateResponse struct {
			Return struct {
				Item []common.EnabledState `xml:"item"`
			} `xml:"return"`
		} `xml:"get_enabled_stateResponse"`
	} `xml:"Body"`
}

func (w *WideIPV2) GetEnabledState(wideIPs []global_lb.WideIPID) ([]common.EnabledState, error) {

	bt, err := w.c.Call(context.Background(), getEnabledStateReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getEnabledStateBody{GetEnabledState: getEnabledState{struct {
			Item []global_lb.WideIPID `xml:"item"`
		}(struct{ Item []global_lb.WideIPID }{Item: wideIPs})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getEnabledStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetEnabledStateResponse.Return.Item, nil
}

type getWideIpPoolRatioReq struct {
	soap.BaseEnvEnvelope
	Body getWideIpPoolRatioBody `xml:"env:Body"`
}

type getWideIpPoolRatioBody struct {
	GetWideIpPoolRatio getWideIpPoolRatio `xml:"tns:get_wide_ip_pool_ratio"`
}

type getWideIpPoolRatio struct {
	WideIPs struct {
		Item []global_lb.WideIPID `xml:"item"`
	} `xml:"wide_ips"`
	WideIPPools WideIPPools `xml:"wide_ip_pools"`
}

type WideIPPools struct {
	Item []WideIPPool `xml:"item"`
}

type WideIPPool struct {
	Item []global_lb.PoolID `xml:"item"`
}

type getWideIpPoolRatioResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetWideIpPoolRatioResponse struct {
			Return struct {
				Item []struct {
					Item []int64 `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_wide_ip_pool_ratioResponse"`
	} `xml:"Body"`
}

// GetWideIpPoolRatio Gets the ratio of the specified wide IP pools on the specified wide IPs.
// Introduced : BIG-IP_v12.0.0
func (w *WideIPV2) GetWideIpPoolRatio(wideIPs []global_lb.WideIPID, poolIDs [][]global_lb.PoolID) ([][]int64, error) {

	wideIPPools := make([]WideIPPool, 0)
	for _, poolID := range poolIDs {
		item := WideIPPool{Item: []global_lb.PoolID{}}
		item.Item = append(item.Item, poolID...)
		wideIPPools = append(wideIPPools, item)
	}

	bt, err := w.c.Call(context.Background(), getWideIpPoolRatioReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getWideIpPoolRatioBody{GetWideIpPoolRatio: getWideIpPoolRatio{
			// struct{ Item []global_lb.WideIPID }{Item: }
			WideIPs: struct {
				Item []global_lb.WideIPID `xml:"item"`
			}(struct{ Item []global_lb.WideIPID }{Item: wideIPs}),
			WideIPPools: WideIPPools{Item: wideIPPools},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getWideIpPoolRatioResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]int64
	for _, item := range resp.Body.GetWideIpPoolRatioResponse.Return.Item {
		data := make([]int64, 0)
		data = append(data, item.Item...)
		res = append(res, data)
	}

	return res, nil
}
