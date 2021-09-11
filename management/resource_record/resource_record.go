package resource_record

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/management"
)

const tns = "urn:iControl:Management/ResourceRecord"

// IResourceRecord
// Introduced : BIG-IP_v9.0.3
// The ResourceRecord interface contains all the calls necessary for manipulating Resource Records:
// adding/deleting/updating This interface does not support transactions.
type IResourceRecord interface {
	GetRRS(viewZones []management.ViewZone) ([][]string, error)
	GetRRSDetailed(viewZones []management.ViewZone) ([]management.RRList, error)
}

var _ IResourceRecord = (*ResourceRecord)(nil)

type ResourceRecord struct {
	c *soap.Client
}

func New(c *soap.Client) IResourceRecord {
	return &ResourceRecord{c: c}
}

type getRRSReq struct {
	soap.BaseEnvEnvelope
	Body getRRSBody `xml:"env:Body"`
}

type getRRSBody struct {
	GetRRS getRRS `xml:"tns:get_rrs"`
}

type getRRS struct {
	ViewZones struct {
		Item []management.ViewZone `xml:"item"`
	} `xml:"view_zones"`
}

type getRRSResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetRrsResponse struct {
			Return struct {
				Item []struct {
					Item []string `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_rrsResponse"`
	} `xml:"Body"`
}

func (r *ResourceRecord) GetRRS(viewZones []management.ViewZone) ([][]string, error) {

	bt, err := r.c.Call(context.Background(), getRRSReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getRRSBody{GetRRS: getRRS{ViewZones: struct {
			Item []management.ViewZone `xml:"item"`
		}(struct{ Item []management.ViewZone }{Item: viewZones})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getRRSResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]string
	for _, v := range resp.Body.GetRrsResponse.Return.Item {
		res = append(res, v.Item)
	}

	return res, nil
}

type getRRSDetailedReq struct {
	soap.BaseEnvEnvelope
	Body getRRSDetailedBody `xml:"env:Body"`
}

type getRRSDetailedBody struct {
	GetRRSDetailed getRRSDetailed `xml:"tns:get_rrs_detailed"`
}

type getRRSDetailed struct {
	ViewZones struct {
		Item []management.ViewZone `xml:"item"`
	} `xml:"view_zones"`
}

type getRRSDetailedResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetRrsDetailedResponse struct {
			Return struct {
				Item []struct {
					AList struct {
						Item []management.ARecord `xml:"item"`
					} `xml:"a_list"`
					NsList struct {
						Item []management.NSRecord `xml:"item"`
					} `xml:"ns_list" json:"ns_list,omitempty"`
					CnameList struct {
						Item []management.CNAMERecord `xml:"item"`
					} `xml:"cname_list"`
					SoaList struct {
						Item []management.SOARecord `xml:"item"`
					} `xml:"soa_list"`
					PtrList struct {
						Item []management.PTRRecord `xml:"item"`
					} `xml:"ptr_list"`
					HInfoList struct {
						Item []management.HINFORecord `xml:"item"`
					} `xml:"hinfo_list"`
					MxList struct {
						Item []management.MXRecord `xml:"item"`
					} `xml:"mx_list"`
					TxtList struct {
						Item []management.TXTRecord `xml:"item"`
					} `xml:"txt_list"`
					SrvList struct {
						Item []management.SRVRecord `xml:"item"`
					} `xml:"srv_list"`
					KeyList struct {
						Item []management.KEYRecord `xml:"item"`
					} `xml:"key_list"`
					SigList struct {
						Item []management.SIGRecord `xml:"item"`
					} `xml:"sig_list"`
					NxtList struct {
						Item []management.NXTRecord `xml:"item"`
					} `xml:"nxt_list"`
					AAAAList struct {
						Item []management.AAAARecord `xml:"item"`
					} `xml:"aaaa_list"`
					A6List struct {
						Item []management.A6Record `xml:"item"`
					} `xml:"a6_list"`
					DnameList struct {
						Item []management.DNAMERecord `xml:"item"`
					} `xml:"dname_list"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_rrs_detailedResponse"`
	} `xml:"Body"`
}

func (r *ResourceRecord) GetRRSDetailed(viewZones []management.ViewZone) ([]management.RRList, error) {

	bt, err := r.c.Call(context.Background(), getRRSDetailedReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getRRSDetailedBody{GetRRSDetailed: getRRSDetailed{ViewZones: struct {
			Item []management.ViewZone `xml:"item"`
		}(struct{ Item []management.ViewZone }{Item: viewZones})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getRRSDetailedResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []management.RRList
	for _, v := range resp.Body.GetRrsDetailedResponse.Return.Item {
		res = append(res, management.RRList{
			AList:     v.AList.Item,
			NSList:    v.NsList.Item,
			CNAMEList: v.CnameList.Item,
			SOAList:   v.SoaList.Item,
			PTRList:   v.PtrList.Item,
			HInfoList: v.HInfoList.Item,
			MXList:    v.MxList.Item,
			TXTList:   v.TxtList.Item,
			SRVList:   v.SrvList.Item,
			KeyList:   v.KeyList.Item,
			SIGList:   v.SigList.Item,
			NXTList:   v.NxtList.Item,
			AAAAList:  v.AAAAList.Item,
			A6List:    v.A6List.Item,
			DNAMEList: v.DnameList.Item,
		})
	}

	return res, err
}
