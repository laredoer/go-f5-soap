package view

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/management"
)

const tns = "urn:iControl:Management/View"

// IView
// Introduced : BIG-IP_v9.0.3
// The View interface contains all calls necessary to manipulate views This interface does not support transactions.
type IView interface {
	GetList() ([]management.ViewInfo, error)
	GetView(viewNames []string) ([]management.ViewInfo, error)
}

var _ IView = (*View)(nil)

type View struct {
	c *soap.Client
}

func New(c *soap.Client) IView {
	return &View{c: c}
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
					ViewName struct {
						Text string `xml:",chardata"`
					} `xml:"view_name"`
					ViewOrder struct {
						Text int64 `xml:",chardata"`
					} `xml:"view_order"`
					OptionSeq struct {
						Item []string `xml:"item"`
					} `xml:"option_seq"`
					ZoneNames struct {
						Item []string `xml:"item"`
					} `xml:"zone_names"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

// GetList
// Introduced : BIG-IP_v9.0.3
// Get a sequence of ViewInfo structs from the server
func (v *View) GetList() ([]management.ViewInfo, error) {

	bt, err := v.c.Call(context.Background(), getListReq{
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

	var res []management.ViewInfo
	for _, v := range resp.Body.GetListResponse.Return.Item {
		res = append(res, management.ViewInfo{
			ViewName:  v.ViewName.Text,
			ViewOrder: v.ViewOrder.Text,
			OptionSeq: v.OptionSeq.Item,
			ZoneNames: v.ZoneNames.Item,
		})
	}

	return res, nil
}

type getViewReq struct {
	soap.BaseEnvEnvelope
	Body GetViewBody `xml:"env:Body"`
}

type GetViewBody struct {
	GetView getView `xml:"tns:get_view"`
}

type getView struct {
	ViewNames struct {
		Item []string `xml:"item"`
	} `xml:"view_names"`
}

type getViewResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetViewResponse struct {
			Return struct {
				Item []struct {
					ViewName  string `xml:"view_name"`
					ViewOrder int64  `xml:"view_order"`
					OptionSeq struct {
						Item []string `xml:"item"`
					} `xml:"option_seq"`
					ZoneNames struct {
						Item []string `xml:"item"`
					} `xml:"zone_names"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_viewResponse"`
	} `xml:"Body"`
}

func (v *View) GetView(viewNames []string) ([]management.ViewInfo, error) {

	bt, err := v.c.Call(context.Background(), getViewReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetViewBody{GetView: getView{struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: viewNames})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getViewResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []management.ViewInfo
	for _, v := range resp.Body.GetViewResponse.Return.Item {
		res = append(res, management.ViewInfo{
			ViewName:  v.ViewName,
			ViewOrder: v.ViewOrder,
			OptionSeq: v.OptionSeq.Item,
			ZoneNames: v.ZoneNames.Item,
		})
	}

	return res, nil
}
