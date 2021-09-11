package region

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/Region"

// IRegion
// Introduced : BIG-IP_v9.2.0
// The Region interface enables you to work with user-defined region definitions.
type IRegion interface {
	GetList() ([]RegionDefinition, error)
	GetRegionItem([]RegionDefinition) ([][]RegionItem, error)
}

var _ IRegion = (*Region)(nil)

// RegionDefinition
// Introduced : BIG-IP_v9.2.0
// A struct that describes a region definition.
type RegionDefinition struct {
	Name   string                 `xml:"name"`    // The region name.
	DBType global_lb.RegionDBType `xml:"db_type"` // The region’s database type.
}

// RegionItem
//Introduced : BIG-IP_v9.2.0
//A struct that describes a region item.
type RegionItem struct {
	Content string               // The region item’s content.
	Type    global_lb.RegionType // The region type.
	Negate  bool                 // The state indicating whether the region member to be interpreted as not equal to the region member options selected.
}

type Region struct {
	c *soap.Client
}

func New(c *soap.Client) *Region {
	return &Region{c: c}
}

type getListBody struct {
	GetList struct{} `xml:"tns:get_list"`
}

type getListReq struct {
	soap.BaseEnvEnvelope
	Body getListBody `xml:"env:Body"`
}

type getListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []struct {
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name"`
					DbType struct {
						Text global_lb.RegionDBType `xml:",chardata"`
					} `xml:"db_type"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

// GetList
//Introduced : BIG-IP_v9.2.0
//Gets a list of of region definitions.
func (r *Region) GetList() ([]RegionDefinition, error) {

	bt, err := r.c.Call(context.Background(), getListReq{
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

	var res []RegionDefinition
	for _, v := range resp.Body.GetListResponse.Return.Item {
		res = append(res, RegionDefinition{
			Name:   v.Name.Text,
			DBType: v.DbType.Text,
		})
	}

	return res, nil

}

type getRegionItemBodyReq struct {
	soap.BaseEnvEnvelope
	Body getRegionItemBody `xml:"env:Body"`
}

type getRegionItemBody struct {
	GetRegionItem getRegionItem `xml:"tns:get_region_item"`
}

type getRegionItem struct {
	Regions Regions `xml:"regions"`
}

type Regions struct {
	Item []RegionDefinition `xml:"item"`
}

type getRegionItemResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetRegionItemResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						Content struct {
							Text string `xml:",chardata"`
						} `xml:"content"`
						Type struct {
							Text global_lb.RegionType `xml:",chardata"`
						} `xml:"type"`
						Negate struct {
							Text bool `xml:",chardata"`
						} `xml:"negate"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_region_itemResponse"`
	} `xml:"Body"`
}

// GetRegionItem
// Introduced : BIG-IP_v9.2.0
// Gets the list of region items that define the specified regions.
func (r *Region) GetRegionItem(definitions []RegionDefinition) ([][]RegionItem, error) {

	bt, err := r.c.Call(context.Background(), getRegionItemBodyReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getRegionItemBody{GetRegionItem: getRegionItem{Regions: Regions{Item: definitions}}},
	})
	if err != nil {
		return nil, err
	}

	var resp getRegionItemResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]RegionItem
	for _, v := range resp.Body.GetRegionItemResponse.Return.Item {
		var regionItem []RegionItem
		for _, v2 := range v.Item {
			regionItem = append(regionItem, RegionItem{
				Content: v2.Content.Text,
				Type:    v2.Type.Text,
				Negate:  v2.Negate.Text,
			})
		}
		res = append(res, regionItem)
	}

	return res, nil
}
