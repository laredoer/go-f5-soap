package bigip

import (
	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/global_lb/data_center"
	"github.com/wule61/go-f5-soap/global_lb/monitor"
	"github.com/wule61/go-f5-soap/global_lb/pool"
	"github.com/wule61/go-f5-soap/global_lb/pool_member"
	"github.com/wule61/go-f5-soap/global_lb/pool_v2"
	"github.com/wule61/go-f5-soap/global_lb/virtual_server"
	"github.com/wule61/go-f5-soap/global_lb/virtual_server_v2"
	"github.com/wule61/go-f5-soap/management/resource_record"
	"github.com/wule61/go-f5-soap/management/view"
	"github.com/wule61/go-f5-soap/management/zone"
	"github.com/wule61/go-f5-soap/system/system_info"
)

// GlobalLB The GlobalLB module contains the Global Load Balancing interfaces
// that enable you to work with the components of a global load balancer system,
// such as data centers, servers, virtual servers, wide IPs, pools ….
// You can also use the interfaces in this module to work with topology attributes and global variables.
type GlobalLB struct {
	Pool            pool.IPool
	PoolMember      pool_member.IPoolMember
	PoolV2          pool_v2.IPoolV2
	Monitor         monitor.IMonitor
	VirtualServer   virtual_server.IVirtualServer
	VirtualServerV2 virtual_server_v2.IVirtualServerV2
	DataCenter      data_center.IDataCenter
}

// Management
// Introduced : BIG-IP_v9.0
// The Management module contains all the interfaces necessary to manage the system.
type Management struct {
	Zone           zone.IZone
	View           view.IView
	ResourceRecord resource_record.IResourceRecord
}

type System struct {
	SystemInfo system_info.ISystemInfo
}

type BigIP struct {
	GlobalLB   *GlobalLB
	Management *Management
	System     *System
}

func New(c *soap.Client) *BigIP {

	return &BigIP{
		GlobalLB: &GlobalLB{
			Pool:            pool.New(c),
			PoolMember:      pool_member.New(c),
			PoolV2:          pool_v2.New(c),
			Monitor:         monitor.New(c),
			VirtualServer:   virtual_server.New(c),
			VirtualServerV2: virtual_server_v2.New(c),
			DataCenter:      data_center.New(c),
		},
		Management: &Management{
			Zone:           zone.New(c),
			View:           view.New(c),
			ResourceRecord: resource_record.New(c),
		},
		System: &System{
			SystemInfo: system_info.New(c),
		},
	}

}
