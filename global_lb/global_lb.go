package global_lb

import "github.com/wule61/go-f5-soap/common"

// RegionType
// Introduced : BIG-IP_v9.2.0
// A list of topology endpoint types.
type RegionType string

const (
	// RegionTypeCIDR The IP subnet topology.
	RegionTypeCIDR RegionType = "REGION_TYPE_CIDR"
	// RegionTypeRegion The region type.
	RegionTypeRegion RegionType = "REGION_TYPE_REGION"
	// RegionTypeContinent The continent type.
	RegionTypeContinent RegionType = "REGION_TYPE_CONTINENT"
	// RegionTypeCountry The country type.
	RegionTypeCountry RegionType = "REGION_TYPE_COUNTRY"
	// RegionTypeState The state type.
	RegionTypeState RegionType = "REGION_TYPE_STATE"
	// RegionTypePool The pool type
	RegionTypePool RegionType = "REGION_TYPE_POOL"
	// RegionTypeDataCenter The data center type
	RegionTypeDataCenter RegionType = "REGION_TYPE_DATA_CENTER"
	// RegionTypeISPRegion The ISP region type
	RegionTypeISPRegion RegionType = "REGION_TYPE_ISP_REGION"
	// RegionTypeGEOIPISP The GeoIP ISP region type
	RegionTypeGEOIPISP RegionType = "REGION_TYPE_GEOIP_ISP"
)

type VirtualServerDefinition struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
	Port    int64  `xml:"port"`
}

type VirtualServerID struct {
	Name   string `xml:"name"`
	Server string `xml:"server"`
}

type MonitorRuleType string

const (
	// MonitorRuleTypeUndefined This monitor rule is undefined, uninitialized state. This value is returned in queries only, and an exception will be raised if used when creating monitor associations.
	MonitorRuleTypeUndefined MonitorRuleType = "MONITOR_RULE_TYPE_UNDEFINED"

	// MonitorRuleTypeNone The object’s monitoring is disabled, i.e. no explicit nor default monitor in use. This value is returned in queries only, and an exception will be raised if used when creating monitor associations.
	MonitorRuleTypeNone MonitorRuleType = "MONITOR_RULE_TYPE_NONE"

	// MonitorRuleTypeSingle This monitor rule is based on a single monitor.
	MonitorRuleTypeSingle MonitorRuleType = "MONITOR_RULE_TYPE_SINGLE"

	// MonitorRuleTypeAndList This monitor rule is based on an ANDed list of monitors, i.e. all monitors must return successfully for the node/member to be considered UP.
	MonitorRuleTypeAndList MonitorRuleType = "MONITOR_RULE_TYPE_AND_LIST"

	// MonitorRuleTypeMOfN This monitor rule is based on a list of N monitors, but at least M of which must return successfully for the node/member to be considered UP.
	MonitorRuleTypeMOfN MonitorRuleType = "MONITOR_RULE_TYPE_M_OF_N"
)

type MonitorRule struct {
	Type             MonitorRuleType `xml:"type"`
	Quorum           int64           `xml:"quorum"`
	MonitorTemplates []string        `xml:"monitor_templates"`
}

// RegionDBType
// Introduced : BIG-IP_v9.2.0
// A list of region database types.
type RegionDBType string

const (
	// RegionDBTypeUserDefined The region database based on user-defined settings .
	RegionDBTypeUserDefined RegionDBType = "REGION_DB_TYPE_USER_DEFINED"
	// RegionDBTypeACL The region database based on ACL lists.
	RegionDBTypeACL RegionDBType = "REGION_DB_TYPE_ACL"
	// RegionDBTypeISP The region database based on ISPs like AOL….
	RegionDBTypeISP RegionDBType = "REGION_DB_TYPE_ISP"
)

// LBMethod
// Introduced : BIG-IP_v9.2.0
// A list of load balancing modes.
type LBMethod string

const (
	// LBMethodReturnToDNS Return to DNS.
	LBMethodReturnToDNS LBMethod = "LB_METHOD_RETURN_TO_DNS"

	// LBMethodNULL No load balancing mode defined.
	LBMethodNULL LBMethod = "LB_METHOD_NULL"

	// LBMethodRoundRobin Round Robin load balancing mode.
	LBMethodRoundRobin LBMethod = "LB_METHOD_ROUND_ROBIN"

	// LBMethodRatio Ratio load balancing mode.
	LBMethodRatio LBMethod = "LB_METHOD_RATIO"

	// LBMethodTopology Topology load balancing mode.
	LBMethodTopology LBMethod = "LB_METHOD_TOPOLOGY"

	// LBMethodStaticPersist Static persist load balancing mode.
	LBMethodStaticPersist LBMethod = "LB_METHOD_STATIC_PERSIST"

	// LBMethodGlobalAvailability Global Availability load balancing mode.
	LBMethodGlobalAvailability LBMethod = "LB_METHOD_GLOBAL_AVAILABILITY"

	// LBMethodVSCapacity Virtual Server (VS) Capacity load balancing mode.
	LBMethodVSCapacity LBMethod = "LB_METHOD_VS_CAPACITY"

	// LBMethodLeastConn Least Connections load balancing mode.
	LBMethodLeastConn LBMethod = "LB_METHOD_LEAST_CONN"

	// LBMethodLowestRTT Lowest Round Trip Times load balancing mode.
	LBMethodLowestRTT LBMethod = "LB_METHOD_LOWEST_RTT"

	// LBMethodLowestHops Lowest hop count load balancing mode.
	LBMethodLowestHops LBMethod = "LB_METHOD_LOWEST_HOPS"

	// LBMethodPacketRate Packet rate load balancing mode.
	LBMethodPacketRate LBMethod = "LB_METHOD_PACKET_RATE"

	// LBMethodCPU CPU usage load balancing mode.
	LBMethodCPU LBMethod = "LB_METHOD_CPU"

	// LBMethodHitRatio Hit ratio load balancing mode.
	LBMethodHitRatio LBMethod = "LB_METHOD_HIT_RATIO"

	// LBMethodQOS Quality of Service load balancing mode.
	LBMethodQOS LBMethod = "LB_METHOD_QOS"

	// LBMethodBPS Bits per second load balancing mode.
	LBMethodBPS LBMethod = "LB_METHOD_BPS"

	// LBMethodDropPacket Drop the request (don’t answer).
	LBMethodDropPacket LBMethod = "LB_METHOD_DROP_PACKET"

	// LBMethodExplicitIP Return an explicit IP address, specified by the user
	LBMethodExplicitIP LBMethod = "LB_METHOD_EXPLICIT_IP"

	// LBMethodConnectionRate This enum is deprecated.
	LBMethodConnectionRate LBMethod = "LB_METHOD_CONNECTION_RATE"

	// LBMethodVSScore Virtual Server (VS) Score load balancing mode.
	LBMethodVSScore LBMethod = "LB_METHOD_VS_SCORE"
)

// WideIPID
// Introduced : BIG-IP_v12.0.0
// A struct that uniquely identifies a wide IP.
type WideIPID struct {
	WideIPName string       `xml:"wideip_name"` // The name of the wide IP.
	WideIPType GTMQueryType `xml:"wideip_type"` // The type of wide IP.
}

// GTMQueryType
// Introduced : BIG-IP_v12.0.0
// An enumeration of GTM query types.
type GTMQueryType string

const (
	// GtmQueryTypeUnknown The GTM query type is unknown (or unsupported by iControl).
	GtmQueryTypeUnknown GTMQueryType = "GTM_QUERY_TYPE_UNKNOWN"

	// GtmQueryTypeA The GTM query type that corresponds a to DNS record type of A (IPv4 address record).
	GtmQueryTypeA GTMQueryType = "GTM_QUERY_TYPE_A"

	// GtmQueryTypeCname The GTM query type that corresponds to a DNS record type of CNAME (canonical name record).
	GtmQueryTypeCname GTMQueryType = "GTM_QUERY_TYPE_CNAME"

	// GtmQueryTypeMX The GTM query type that corresponds to a DNS record type of MX (mail exchange record).
	GtmQueryTypeMX GTMQueryType = "GTM_QUERY_TYPE_MX"

	// GtmQueryTypeAAAA The GTM query type that corresponds to a DNS record type of MX (mail exchange record).
	GtmQueryTypeAAAA GTMQueryType = "GTM_QUERY_TYPE_AAAA"

	// GtmQueryTypeSRV The GTM query type that corresponds to a DNS record type of SRV (service location record).
	GtmQueryTypeSRV GTMQueryType = "GTM_QUERY_TYPE_SRV"

	// GtmQueryTypeNAPTR The GTM query type that corresponds to a DNS record type of NAPTR (naming authority pointer).
	GtmQueryTypeNAPTR GTMQueryType = "GTM_QUERY_TYPE_NAPTR"
)

// PoolID
// Introduced : BIG-IP_v12.0.0
// A struct that uniquely identifies a GTM pool.
type PoolID struct {
	PoolName string       `xml:"pool_name"` // The name of the pool.
	PoolType GTMQueryType `xml:"pool_type"` // The type of pool.
}

type MonitorInstance struct {
	TemplateName       string        // The monitor template used by this instance.
	InstanceDefinition MonitorIPPort // The IP:port of this instance.
}

type MonitorIPPort struct {
	AddressType AddressType             // The address type of the IP:port specified in ipport.
	IPPort      common.IPPortDefinition // The IP:port definition.
}

type AddressType string

const (
	// ATypeUnset The address type is unknown.
	ATypeUnset AddressType = "ATYPE_UNSET"

	// ATypeStarAddressStarPort For example, “:”.
	ATypeStarAddressStarPort AddressType = "ATYPE_STAR_ADDRESS_STAR_PORT"

	// ATypeStarAddressExplicitPort For example, “*:80”.
	ATypeStarAddressExplicitPort AddressType = "ATYPE_STAR_ADDRESS_EXPLICIT_PORT"

	// ATypeExplicitAddressExplicitPort For example, “10.10.10.1:80”.
	ATypeExplicitAddressExplicitPort AddressType = "ATYPE_EXPLICIT_ADDRESS_EXPLICIT_PORT"

	// ATypeStarAddress For example, “:”.
	ATypeStarAddress AddressType = "ATYPE_STAR_ADDRESS"

	// ATypeExplicitAddress For example, “10.10.10.1:80”.
	ATypeExplicitAddress AddressType = "ATYPE_EXPLICIT_ADDRESS"
)

type MonitorInstanceState struct {
	Instance      MonitorInstance          // The monitor instance definition.
	InstanceState MonitorInstanceStateType // The state of the monitor instance.
	EnabledState  bool                     // The state indicating whether the instance is enabled/disabled.
}

// MonitorInstanceStateType
// Introduced : BIG-IP_v9.2.0
// A list of monitor instance states.
type MonitorInstanceStateType string

const (
	// InstanceStateUnchecked The instance state is unknown.
	InstanceStateUnchecked MonitorInstanceStateType = "INSTANCE_STATE_UNCHECKED"

	// InstanceStateChecking The instance state is CHECKING.
	InstanceStateChecking MonitorInstanceStateType = "INSTANCE_STATE_CHECKING"

	// InstanceStateUp The instance state is UP.
	InstanceStateUp MonitorInstanceStateType = "INSTANCE_STATE_UP"

	// InstanceStateDown The instance state is DOWN.
	InstanceStateDown MonitorInstanceStateType = "INSTANCE_STATE_DOWN"

	// InstanceStateForcedDown The instance state is FORCED_DOWN.
	InstanceStateForcedDown MonitorInstanceStateType = "INSTANCE_STATE_FORCED_DOWN"

	// InstanceStateDisabled The instance state is DISABLED.
	InstanceStateDisabled MonitorInstanceStateType = "INSTANCE_STATE_DISABLED"

	// InstanceStateDownByIRULE The instance state is DOWN, marked by an iRule.
	InstanceStateDownByIRULE MonitorInstanceStateType = "INSTANCE_STATE_DOWN_BY_IRULE"

	// InstanceStateDownWaitForManualResume The instance state is DOWN, and should only be marked up manually.
	InstanceStateDownWaitForManualResume MonitorInstanceStateType = "INSTANCE_STATE_DOWN_WAIT_FOR_MANUAL_RESUME"
)
