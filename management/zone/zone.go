package zone

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/management"
)

const tns = "urn:iControl:Management/Zone"

// IZone
// Introduced : BIG-IP_v9.0
// The Zone interface enables the user to perform “zone” operations on a dns database
// This interface does not support transactions.
type IZone interface {
	// Deprecated: Please use get_zone_v2.
	GetZone(viewZones []management.ViewZone) ([]management.ZoneInfo, error)
	GetZoneV2(viewZones []management.ViewZone) ([]management.ZoneInfo, error)
	GetZoneName(viewNames []string) ([]management.ViewZone, error)
}

var _ IZone = (*Zone)(nil)

type Zone struct {
	c *soap.Client
}

func New(c *soap.Client) IZone {
	return &Zone{c: c}
}

type getZoneNameReq struct {
	soap.BaseEnvEnvelope
	Body getZoneNameBody `xml:"env:Body"`
}

type getZoneNameBody struct {
	GetZoneName getZoneName `xml:"tns:get_zone_name"`
}

type getZoneName struct {
	ViewNames struct {
		Item []string `xml:"item"`
	} `xml:"view_names"`
}

type getZoneNameResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetZoneNameResponse struct {
			Return struct {
				Item []struct {
					ViewName struct {
						Text string `xml:",chardata"`
					} `xml:"view_name"`
					ZoneName struct {
						Text string `xml:",chardata"`
					} `xml:"zone_name"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_zone_nameResponse"`
	} `xml:"Body"`
}

// GetZoneName
// Introduced : BIG-IP_v9.0
// Gets the list of zone names for the specified views.
func (z *Zone) GetZoneName(viewNames []string) ([]management.ViewZone, error) {

	bt, err := z.c.Call(context.Background(), getZoneNameReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getZoneNameBody{GetZoneName: getZoneName{ViewNames: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: viewNames})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getZoneNameResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []management.ViewZone
	for _, v := range resp.Body.GetZoneNameResponse.Return.Item {
		res = append(res, management.ViewZone{
			ViewName: v.ViewName.Text,
			ZoneName: v.ZoneName.Text,
		})
	}

	return res, nil
}

type getZoneReq struct {
	soap.BaseEnvEnvelope
	Body getZoneBody `xml:"env:Body"`
}

type getZoneBody struct {
	GetZone getZone `xml:"tns:get_zone"`
}

type getZone struct {
	ViewZones struct {
		Item []management.ViewZone `xml:"item"`
	} `xml:"view_zones"`
}

type getZoneResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetZoneResponse struct {
			Return struct {
				Item []struct {
					ViewName struct {
						Text string `xml:",chardata"`
					} `xml:"view_name"`
					ZoneName struct {
						Text string `xml:",chardata"`
					} `xml:"zone_name"`
					ZoneType struct {
						Text management.ZoneType `xml:",chardata"`
					} `xml:"zone_type"`
					ZoneFile struct {
						Text string `xml:",chardata"`
					} `xml:"zone_file"`
					OptionSeq struct {
						Item []string `xml:"item"`
					} `xml:"option_seq"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_zoneResponse"`
	} `xml:"Body"`
}

// GetZone
// Introduced : BIG-IP_v9.0
// Gets the ZoneInfo structs for the specified zones in the specified views.
// This method has been deprecated due to an inconsistency in the format of the options_seq field.
// Deprecated: Please use get_zone_v2. Gets the ZoneInfo structs for the specified zones in the specified views.
func (z *Zone) GetZone(viewZones []management.ViewZone) ([]management.ZoneInfo, error) {

	bt, err := z.c.Call(context.Background(), getZoneReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getZoneBody{GetZone: getZone(struct {
			ViewZones struct {
				Item []management.ViewZone `xml:"item"`
			}
		}{ViewZones: struct {
			Item []management.ViewZone `xml:"item"`
		}(struct{ Item []management.ViewZone }{Item: viewZones})})},
	})
	if err != nil {
		return nil, err
	}

	var resp getZoneResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []management.ZoneInfo
	for _, v := range resp.Body.GetZoneResponse.Return.Item {
		res = append(res, management.ZoneInfo{
			ViewName:  v.ViewName.Text,
			ZoneName:  v.ZoneName.Text,
			ZoneType:  v.ZoneType.Text,
			ZoneFile:  v.ZoneFile.Text,
			OptionSeq: v.OptionSeq.Item,
		})
	}

	return res, nil
}

type getZoneV2Req struct {
	soap.BaseEnvEnvelope
	Body getZoneV2Body `xml:"env:Body"`
}

type getZoneV2Body struct {
	GetZoneV2 getZoneV2 `xml:"tns:get_zone_v2"`
}

type getZoneV2 struct {
	ViewZones struct {
		Item []management.ViewZone `xml:"item"`
	} `xml:"view_zones"`
}

type getZoneV2Resp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetZoneV2Response struct {
			Return struct {
				Item []struct {
					ViewName struct {
						Text string `xml:",chardata"`
					} `xml:"view_name"`
					ZoneName struct {
						Text string `xml:",chardata"`
					} `xml:"zone_name"`
					ZoneType struct {
						Text management.ZoneType `xml:",chardata"`
					} `xml:"zone_type"`
					ZoneFile struct {
						Text string `xml:",chardata"`
					} `xml:"zone_file"`
					OptionSeq struct {
						Item []string `xml:"item"`
					} `xml:"option_seq"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_zone_v2Response"`
	} `xml:"Body"`
}

// GetZoneV2
// Introduced : BIG-IP_v12.0.0
// Gets the ZoneInfo structs for the specified zones in the specified views.
func (z *Zone) GetZoneV2(viewZones []management.ViewZone) ([]management.ZoneInfo, error) {

	bt, err := z.c.Call(context.Background(), getZoneV2Req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getZoneV2Body{GetZoneV2: getZoneV2(struct {
			ViewZones struct {
				Item []management.ViewZone `xml:"item"`
			}
		}{ViewZones: struct {
			Item []management.ViewZone `xml:"item"`
		}(struct{ Item []management.ViewZone }{Item: viewZones})})},
	})
	if err != nil {
		return nil, err
	}

	var resp getZoneV2Resp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []management.ZoneInfo
	for _, v := range resp.Body.GetZoneV2Response.Return.Item {
		res = append(res, management.ZoneInfo{
			ViewName:  v.ViewName.Text,
			ZoneName:  v.ZoneName.Text,
			ZoneType:  v.ZoneType.Text,
			ZoneFile:  v.ZoneFile.Text,
			OptionSeq: v.OptionSeq.Item,
		})
	}

	return res, nil
}
