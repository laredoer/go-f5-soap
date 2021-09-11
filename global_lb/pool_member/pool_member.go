package pool_member

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
)

const tns = "urn:iControl:GlobalLB/PoolMember"

// IPoolMember Due to changing virtual server keys,
// this full interface is deprecated.
// Its functionality has been moved to the GlobalLB::Pool interface.
// The PoolMember interface enables you to work with the pool members and their settings, and statistics.
type IPoolMember interface {
	GetRatio(poolNames []string, members [][]common.IPPortDefinition) ([][]common.MemberRatio, error)
	GetObjectStatus(poolNames []string, members [][]common.IPPortDefinition) ([][]common.MemberObjectStatus, error)
}

var _ IPoolMember = (*PoolMember)(nil)

type PoolMember struct {
	c *soap.Client
}

func New(c *soap.Client) *PoolMember {
	return &PoolMember{
		c: c,
	}
}

type GetRatioBody struct {
	GetRatio GetRatio `xml:"tns:get_ratio"`
}

type GetRatio struct {
	PoolNames PoolNames `xml:"pool_names"`
	Members   Members   `xml:"members"`
}

type PoolNames struct {
	Item []string `xml:"item"`
}

type Members struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Item []common.IPPortDefinition `xml:"item"`
}

type RatioResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetRatioResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						Member struct {
							Address struct {
								Text string `xml:",chardata"`
							} `xml:"address"`
							Port struct {
								Text int64 `xml:",chardata"`
							} `xml:"port"`
						} `xml:"member"`
						Ratio struct {
							Text int64 `xml:",chardata"`
						} `xml:"ratio"`
					} `xml:"item"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_ratioResponse"`
	} `xml:"Body"`
}

// GetRatio Gets the ratios for the specified members in the specified pools.
// 获取指定池中指定成员的比率。
func (p *PoolMember) GetRatio(poolNames []string, members [][]common.IPPortDefinition) ([][]common.MemberRatio, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetRatioBody `xml:"env:Body"`
	}

	var reqItem []Item
	for _, v := range members {
		item := Item{Item: []common.IPPortDefinition{}}
		item.Item = append(item.Item, v...)
		reqItem = append(reqItem, item)
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetRatioBody{GetRatio: GetRatio{
			PoolNames: PoolNames{Item: poolNames},
			Members:   Members{Item: reqItem},
		}},
	})
	if err != nil {
		return nil, err
	}
	var resp RatioResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]common.MemberRatio
	for _, v := range resp.Body.GetRatioResponse.Return.Item {
		var item []common.MemberRatio
		for _, v2 := range v.Item {
			item = append(item, common.MemberRatio{
				Member: common.IPPortDefinition{
					Address: v2.Member.Address.Text,
					Port:    v2.Member.Port.Text,
				},
				Ratio: v2.Ratio.Text,
			})
		}

		res = append(res, item)
	}

	return res, nil
}

type GetObjectStatusBody struct {
	GetObjectStatus GetObjectStatus `xml:"tns:get_object_status"`
}

type GetObjectStatus struct {
	PoolNames PoolNames `xml:"pool_names"`
	Members   Members   `xml:"members"`
}

type ObjectStatusResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetObjectStatusResponse struct {
			Return struct {
				Item []struct {
					Item []struct {
						Member struct {
							Address struct {
								Text string `xml:",chardata"`
							} `xml:"address" `
							Port struct {
								Text int64 `xml:",chardata" `
							} `xml:"port"`
						} `xml:"member" `
						Status struct {
							AvailabilityStatus struct {
								Text string `xml:",chardata" `
							} `xml:"availability_status" `
							EnabledStatus struct {
								Text string `xml:",chardata" `
							} `xml:"enabled_status" `
							StatusDescription struct {
								Text string `xml:",chardata"`
							} `xml:"status_description" `
						} `xml:"status" `
					} `xml:"item" `
				} `xml:"item"`
			} `xml:"return" `
		} `xml:"get_object_statusResponse" `
	} `xml:"Body" `
}

// GetObjectStatus Gets the statuses for the specified members in the specified pools.
// 获取指定池中指定成员的状态。
func (p *PoolMember) GetObjectStatus(poolNames []string, members [][]common.IPPortDefinition) ([][]common.MemberObjectStatus, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetObjectStatusBody `xml:"env:Body"`
	}

	var reqItem []Item
	for _, v := range members {
		item := Item{Item: []common.IPPortDefinition{}}
		item.Item = append(item.Item, v...)
		reqItem = append(reqItem, item)
	}

	bt, err := p.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetObjectStatusBody{GetObjectStatus: GetObjectStatus{
			PoolNames: PoolNames{Item: poolNames},
			Members:   Members{Item: reqItem},
		}},
	})
	if err != nil {
		return nil, err
	}
	var resp ObjectStatusResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res [][]common.MemberObjectStatus
	for _, v := range resp.Body.GetObjectStatusResponse.Return.Item {
		var item []common.MemberObjectStatus
		for _, v2 := range v.Item {
			item = append(item, common.MemberObjectStatus{
				Member: common.IPPortDefinition{
					Address: v2.Member.Address.Text,
					Port:    v2.Member.Port.Text,
				},
				Status: common.ObjectStatus{
					AvailabilityStatus: common.AvailabilityStatus(v2.Status.AvailabilityStatus.Text),
					EnabledStatus:      common.EnabledStatus(v2.Status.EnabledStatus.Text),
					StatusDescription:  v2.Status.StatusDescription.Text,
				},
			})
		}

		res = append(res, item)
	}

	return res, nil

}
