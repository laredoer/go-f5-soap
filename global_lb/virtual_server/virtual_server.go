package virtual_server

import (
	"context"
	"encoding/xml"
	"fmt"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/VirtualServer"

type MonitorAssociation struct {
	VirtualServer global_lb.VirtualServerDefinition
	MonitorRule   global_lb.MonitorRule
}

type VirtualServers struct {
	Item []global_lb.VirtualServerDefinition `xml:"item"`
}

type IVirtualServer interface {
	GetList() ([]global_lb.VirtualServerDefinition, error)
	GetMonitorAssociation([]global_lb.VirtualServerDefinition) ([]MonitorAssociation, error)
	GetServer([]global_lb.VirtualServerDefinition) ([]string, error)
}

var _ IVirtualServer = (*VirtualServer)(nil)

type VirtualServer struct {
	c *soap.Client
}

func New(c *soap.Client) *VirtualServer {
	return &VirtualServer{c: c}
}

type GetListBody struct {
	GetList struct{} `xml:"tns:get_list"`
}

type getListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetListResponse struct {
			Return struct {
				Item []struct {
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name" json:"name,omitempty"`
					Address struct {
						Text string `xml:",chardata"`
					} `xml:"address"`
					Port struct {
						Text int64 `xml:",chardata"`
					} `xml:"port"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_listResponse"`
	} `xml:"Body"`
}

func (v *VirtualServer) GetList() ([]global_lb.VirtualServerDefinition, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetListBody `xml:"env:Body"`
	}

	bt, err := v.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            GetListBody{GetList: struct{}{}},
	})
	if err != nil {
		return nil, err
	}

	var resp getListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []global_lb.VirtualServerDefinition
	for _, v := range resp.Body.GetListResponse.Return.Item {
		res = append(res, global_lb.VirtualServerDefinition{
			Name:    v.Name.Text,
			Address: v.Address.Text,
			Port:    v.Port.Text,
		})
	}

	return res, nil
}

type GetServerBody struct {
	GetServer GetServer `xml:"tns:get_server"`
}

type GetServer struct {
	VirtualServers VirtualServers `xml:"virtual_servers"`
}

type getServersResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetServersResponse struct {
			Return struct {
				Item []string `xml:"item"`
			} `xml:"return"`
		} `xml:"get_serverResponse"`
	} `xml:"Body"`
}

func (v *VirtualServer) GetServer(virtualServers []global_lb.VirtualServerDefinition) ([]string, error) {

	type Req struct {
		soap.BaseEnvEnvelope
		Body GetServerBody `xml:"env:Body"`
	}

	bt, err := v.c.Call(context.Background(), Req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetServerBody{GetServer: GetServer{
			VirtualServers: VirtualServers{
				Item: virtualServers,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getServersResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}
	fmt.Println(resp.Body.GetServersResponse.Return.Item)
	var res []string
	for _, v := range resp.Body.GetServersResponse.Return.Item {

		res = append(res, v)
	}

	return res, nil

}

type GetMonitorAssociationBody struct {
	GetMonitorAssociation GetMonitorAssociation `xml:"tns:get_monitor_association"`
}

type GetMonitorAssociation struct {
	VirtualServers VirtualServers `xml:"virtual_servers"`
}

type getMonitorAssociationResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetMonitorAssociationResponse struct {
			Return struct {
				Item []struct {
					VirtualServer struct {
						Name struct {
							Text string `xml:",chardata"`
						} `xml:"name"`
						Address struct {
							Text string `xml:",chardata"`
						} `xml:"address"`
						Port struct {
							Text int64 `xml:",chardata"`
						} `xml:"port"`
					} `xml:"virtual_server"`
					MonitorRule struct {
						Type struct {
							Text global_lb.MonitorRuleType `xml:",chardata"`
						} `xml:"type"`
						Quorum struct {
							Text int64 `xml:",chardata"`
						} `xml:"quorum"`
						MonitorTemplates struct {
							Item []string `xml:"item"`
						} `xml:"monitor_templates" `
					} `xml:"monitor_rule"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_monitor_associationResponse"`
	} `xml:"Body"`
}

func (v *VirtualServer) GetMonitorAssociation(virtualServers []global_lb.VirtualServerDefinition) ([]MonitorAssociation, error) {

	type Req struct {
		soap.BaseEnvEnvelope
		Body GetMonitorAssociationBody `xml:"env:Body"`
	}

	bt, err := v.c.Call(context.Background(), Req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetMonitorAssociationBody{GetMonitorAssociation: GetMonitorAssociation{
			VirtualServers: VirtualServers{
				Item: virtualServers,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getMonitorAssociationResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []MonitorAssociation
	for _, v := range resp.Body.GetMonitorAssociationResponse.Return.Item {

		res = append(res, MonitorAssociation{
			VirtualServer: global_lb.VirtualServerDefinition{
				Name:    v.VirtualServer.Name.Text,
				Address: v.VirtualServer.Address.Text,
				Port:    v.VirtualServer.Port.Text,
			},
			MonitorRule: global_lb.MonitorRule{
				Type:             v.MonitorRule.Type.Text,
				Quorum:           v.MonitorRule.Quorum.Text,
				MonitorTemplates: v.MonitorRule.MonitorTemplates.Item,
			},
		})
	}

	return res, nil

}
