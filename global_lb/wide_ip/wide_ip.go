package wide_ip

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/WideIP"

// IWideIP
// Introduced : BIG-IP_v9.2.0
// The WideIP interface enables you to work with wide IPs,
// as well as with the pools and the virtual servers that make them up.
// For example, use the WideIP interface
// to get a list of wide IPs, to add a wide IP, or to remove a wide IP.
type IWideIP interface {
	GetList() (wideIPs []string, err error)
	GetWideIpPool(wideIPs []string) ([][]WideIPPool, error)
	GetLBMethod(wideIPs []string) ([]global_lb.LBMethod, error)
	GetObjectStatus(wideIPs []string) ([]common.ObjectStatus, error)
	GetEnabledState(wideIPs []string) ([]common.EnabledState, error)
}

// WideIPPool
// Introduced : BIG-IP_v9.2.0
// A struct that describes a wide IP&aposs pool.
type WideIPPool struct {
	PoolName string // The pool name.
	Order    int64  // The order given to the specified pool.
	Ratio    int64  // The ratio given to the specified pool.
}

var _ IWideIP = (*WideIP)(nil)

type WideIP struct {
	c *soap.Client
}

func New(c *soap.Client) IWideIP {
	return &WideIP{c: c}
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

func (w *WideIP) GetList() ([]string, error) {

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

	return resp.Body.GetListResponse.Return.Item, nil
}

type getWideIpPoolReq struct {
	soap.BaseEnvEnvelope
	Body getWideIpPoolBody `xml:"env:Body"`
}

type getWideIpPoolBody struct {
	GetWideipPool getWideipPool `xml:"tns:get_wideip_pool"`
}

type getWideipPool struct {
	WideIPs struct {
		Item []string `xml:"item"`
	} `xml:"wide_ips"`
}

type getWideIpPoolResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetWideipPoolResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						PoolName struct {
							Text string `xml:",chardata"`
						} `xml:"pool_name"`
						Order struct {
							Text int64 `xml:",chardata"`
						} `xml:"order"`
						Ratio struct {
							Text int64 `xml:",chardata"`
						} `xml:"ratio"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_wideip_poolResponse"`
	} `xml:"Body"`
}

func (w *WideIP) GetWideIpPool(wideIPs []string) ([][]WideIPPool, error) {

	bt, err := w.c.Call(context.Background(), getWideIpPoolReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getWideIpPoolBody{GetWideipPool: getWideipPool{WideIPs: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: wideIPs})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getWideIpPoolResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]WideIPPool
	for _, v := range resp.Body.GetWideipPoolResponse.Return.Item {
		var resItem []WideIPPool
		for _, v2 := range v.Item {
			resItem = append(resItem, WideIPPool{
				PoolName: v2.PoolName.Text,
				Order:    v2.Order.Text,
				Ratio:    v2.Ratio.Text,
			})
		}
		res = append(res, resItem)
	}

	return res, nil
}

type getLBMethodReq struct {
	soap.BaseEnvEnvelope
	Body getLBMethodBody `xml:"env:Body"`
}

type getLBMethodBody struct {
	GetLBMethod getLBMethod `xml:"tns:get_lb_method"`
}

type getLBMethod struct {
	WideIPs struct {
		Item []string `xml:"item"`
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

func (w *WideIP) GetLBMethod(wideIPs []string) ([]global_lb.LBMethod, error) {

	bt, err := w.c.Call(context.Background(), getLBMethodReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getLBMethodBody{GetLBMethod: getLBMethod{WideIPs: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: wideIPs})}},
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
		Item []string `xml:"item"`
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

func (w *WideIP) GetObjectStatus(wideIPs []string) ([]common.ObjectStatus, error) {

	bt, err := w.c.Call(context.Background(), getObjectStatusReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getObjectStatusBody{GetObjectStatus: getObjectStatus{WideIPs: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: wideIPs})}},
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
		Item []string `xml:"item"`
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

func (w *WideIP) GetEnabledState(wideIPs []string) ([]common.EnabledState, error) {

	bt, err := w.c.Call(context.Background(), getEnabledStateReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getEnabledStateBody{GetEnabledState: getEnabledState{WideIPs: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: wideIPs})}},
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
