// Package monitor
// Introduced : BIG-IP_v9.2.0
// The Monitor interface enables you to manipulate a load balancer&aposs monitor templates and instances.
// For example, use the Monitor interface to enable or disable a monitor instance,
// or to create a monitor template, or to get and set different attributes of a monitor template.
package monitor

import (
	"context"
	"encoding/xml"

	soap "github.com/wule61/go-f5-soap"
	"github.com/wule61/go-f5-soap/common"
	"github.com/wule61/go-f5-soap/global_lb"
)

const tns = "urn:iControl:GlobalLB/Monitor"

type TemplateType string

const (

	// TTypeUnset The template type is unknown.
	TTypeUnset TemplateType = "TTYPE_UNSET"

	// TTypeICMP The ICMP template type.
	TTypeICMP = "TTYPE_ICMP"

	// TTypeTCP The TCP template type.
	TTypeTCP = "TTYPE_TCP"

	// TTypeTCPEcho The TCP_ECHO template type.
	TTypeTCPEcho = "TTYPE_TCP_ECHO"

	// TTypeExternal The EXTERNAL template type.
	TTypeExternal = "TTYPE_EXTERNAL"

	// TTypeHTTP The HTTP template type.
	TTypeHTTP = "TTYPE_HTTP"

	// TTypeHTTPS The HTTPS template type.
	TTypeHTTPS = "TTYPE_HTTPS"

	// TTypeNNTP The NNTP template type.
	TTypeNNTP = "TTYPE_NNTP"

	// TTypeFTP The FTP template type.
	TTypeFTP = "TTYPE_FTP"

	// TTypePOP3 The POP3 template type.
	TTypePOP3 = "TTYPE_POP3"

	// TTypeSMTP The SMTP template type.
	TTypeSMTP = "TTYPE_SMTP"

	// TTypeMSSQL The MSSQL template type.
	TTypeMSSQL = "TTYPE_MSSQL"

	// TTypeGateway The GATEWAY template type.
	TTypeGateway = "TTYPE_GATEWAY"

	// TTypeIMAP The IMAP template type.
	TTypeIMAP = "TTYPE_IMAP"

	// TTypeRadius The RADIUS template type.
	TTypeRadius = "TTYPE_RADIUS"

	// TTypeLDAP The LDAP template type.
	TTypeLDAP = "TTYPE_LDAP"

	// TTypeWMI The WMI template type.
	TTypeWMI = "TTYPE_WMI"

	// TTypeSnmpDca The SNMP_DCA template type.
	TTypeSnmpDca = "TTYPE_SNMP_DCA"

	// TTypeSnmpDcaBase The SNMP_DCA_BASE template type.
	TTypeSnmpDcaBase = "TTYPE_SNMP_DCA_BASE"

	// TTypeRealServer The REAL_SERVER template type.
	TTypeRealServer = "TTYPE_REAL_SERVER"

	// TTypeUDP The UDP template type.
	TTypeUDP = "TTYPE_UDP"

	// TTypeNone Not using any monitor template.
	TTypeNone = "TTYPE_NONE"

	// TTypeOracle The ORACLE template type.
	TTypeOracle = "TTYPE_ORACLE"

	// TTypeSoap The SOAP template type.
	TTypeSoap = "TTYPE_SOAP"

	// TTypeGatewayICMP The GATEWAY_ICMP template type.
	TTypeGatewayICMP = "TTYPE_GATEWAY_ICMP"

	// TTypeSIP The SIP template type.
	TTypeSIP = "TTYPE_SIP"

	// TTypeTCPHalfOpen The TCP_HALF_OPEN template type.
	TTypeTCPHalfOpen = "TTYPE_TCP_HALF_OPEN"

	// TTypeScripted The SCRIPTED template type.
	TTypeScripted = "TTYPE_SCRIPTED"

	// TTypeWAP The WAP template type.
	TTypeWAP = "TTYPE_WAP"

	// TTypeBIGIP The BIGIP template type.
	TTypeBIGIP = "TTYPE_BIGIP"

	// TTypeBIGIPLink The BIGIP_LINK template type.
	TTypeBIGIPLink = "TTYPE_BIGIP_LINK"

	// TTypeSnmpGtm The SNMP_GTM template type.
	TTypeSnmpGtm = "TTYPE_SNMP_GTM"

	// TTypeSnmpLink The SNMP_LINK template type.
	TTypeSnmpLink = "TTYPE_SNMP_LINK"

	// TTypeFirePassGtm Template type for monitoring Firepass servers
	TTypeFirePassGtm = "TTYPE_FIREPASS_GTM"

	// TTypeRadiusAccounting The RADIUS ACCOUNTING template type.
	TTypeRadiusAccounting = "TTYPE_RADIUS_ACCOUNTING"

	// TTypeDiameter The Diameter authorization, authentication, and accounting server template type.
	TTypeDiameter = "TTYPE_DIAMETER"

	// TTypeMysql The MySQL monitor template type.This monitor verifies MySQL-based services.
	TTypeMysql = "TTYPE_MYSQL"

	// TTypePostgreSQL The PostgreSQL monitor template type.This monitor verifies PostgreSQL-based services.
	TTypePostgreSQL = "TTYPE_POSTGRESQL"
)

type IntPropertyType string

const (
	// ITypeUnset The integer property type is unknown.
	ITypeUnset IntPropertyType = "ITYPE_UNSET"

	// ITypeInterval The integer property type used to change the value of interval.
	ITypeInterval = "ITYPE_INTERVAL"

	// ITypeTimeOut The integer property type used to change the value of timeout.
	ITypeTimeOut = "ITYPE_TIMEOUT"

	// ITypeProbeInterval The integer property type used to change the value of the probing interval.
	ITypeProbeInterval = "ITYPE_PROBE_INTERVAL"

	// ITypeProbeTimeOut The integer property type used to change the value of the probing timeout.
	ITypeProbeTimeOut = "ITYPE_PROBE_TIMEOUT"

	// ITypeProbeNumProbes The integer property type used to change the number of probes.
	ITypeProbeNumProbes = "ITYPE_PROBE_NUM_PROBES"

	// ITypeProbeNumSuccesses The integer property type used to change the number of successful probes.
	ITypeProbeNumSuccesses = "ITYPE_PROBE_NUM_SUCCESSES"
)

type IntegerValue struct {
	Type  IntPropertyType `xml:"type"`  // The integer property type.
	Value int64           `xml:"value"` // The integer property value.
}

type StrPropertyType string

const (
	// STYPE_UNSET The string property type is unknown.
	STYPE_UNSET StrPropertyType = "STYPE_UNSET"

	// STYPE_SEND The string property type used to change a string value of a template (TCP, HTTP, HTTPS).
	STYPE_SEND = "STYPE_SEND"

	// STYPE_GET The string property type used to change a string value of a template (HTTP, HTTPS, FTP).
	STYPE_GET = "STYPE_GET"

	// STYPE_RECEIVE The string property type used to change a string value of a template (TCP, HTTP, HTTPS).
	STYPE_RECEIVE = "STYPE_RECEIVE"

	//The string property type used to change a string value of a template (HTTP, HTTPS, NNTP, FTP, POP3, SQL, IMAP, RADIUS, RADIUS_ACCOUNTING, LDAP, WMI).
	STYPE_USERNAME = "STYPE_USERNAME"

	//The string property type used to change a string value of a template (HTTP, HTTPS, NNTP, FTP, POP3, SQL, IMAP, RADIUS, LDAP, WMI).
	STYPE_PASSWORD = "STYPE_PASSWORD"

	//The string property specifying the name of the executeable file run in an external monitor. Monitor executable files are officially managed as external monitor file objects via the STYPE_RUN_V2 property and the System::ExternalMonitorFile interface. Thus this value has been deprecated.
	STYPE_RUN = "STYPE_RUN"

	//The string property type used to change a string value of a template (NNTP).
	STYPE_NEWSGROUP = "STYPE_NEWSGROUP"

	//The string property type used to change a string value of a template (SQL).
	STYPE_DATABASE = "STYPE_DATABASE"

	//The string property type used to change a string value of a template (SMTP).
	STYPE_DOMAIN = "STYPE_DOMAIN"

	//The string property type used to change a string value of a template (EXTERNAL).
	STYPE_ARGUMENTS = "STYPE_ARGUMENTS"

	//The string property type used to change a string value of a template (IMAP).
	STYPE_FOLDER = "STYPE_FOLDER"

	//	The string property type used to change a string value of a template (LDAP).
	STYPE_BASE = "STYPE_BASE"

	//The string property type used to change a string value of a template (LDAP).
	STYPE_FILTER = "STYPE_FILTER"

	//The string property type used to change a string value of a template (RADIUS, RADIUS_ACCOUNTING).
	STYPE_SECRET = "STYPE_SECRET"

	//	The string property type used to change a string value of a template (WMI, REAL_SERVER).
	STYPE_METHOD = "STYPE_METHOD"

	//The string property type used to change a string value of a template (WMI).
	STYPE_URL = "STYPE_URL"

	//The string property type used to change a string value of a template (WMI, REAL_SERVER).
	STYPE_COMMAND = "STYPE_COMMAND"

	//The string property type used to change a string value of a template (WMI, REAL_SERVER).
	STYPE_METRICS = "STYPE_METRICS"

	//The string property type used to change a string value of a template (WMI).
	STYPE_POST = "STYPE_POST"

	//The string property type used to change a string value of a template (WMI, REAL_SERVER).
	STYPE_USERAGENT = "STYPE_USERAGENT"

	//The string property type used to change a string value of a template (SNMP_DCA ).
	STYPE_AGENT_TYPE = "STYPE_AGENT_TYPE"

	//The string property type used to change a string value of a template (SNMP_DCA).
	STYPE_CPU_COEFFICIENT = "STYPE_CPU_COEFFICIENT"

	//The string property
	//type used to
	//change a string value of a template (SNMP_DCA).
	STYPE_CPU_THRESHOLD = "STYPE_CPU_THRESHOLD"

	//The string property
	//type used to
	//change a string value of a template (SNMP_DCA).
	STYPE_MEMORY_COEFFICIENT = "STYPE_MEMORY_COEFFICIENT"

	// The string property type used to  change a string value of a template (SNMP_DCA).
	STYPE_MEMORY_THRESHOLD = "STYPE_MEMORY_THRESHOLD"

	//The string property
	//type used to
	//change a string value of a template (SNMP_DCA).
	STYPE_DISK_COEFFICIENT = "STYPE_DISK_COEFFICIENT"

	//The string property
	//type used to
	//change a string value of a template (SNMP_DCA).
	STYPE_DISK_THRESHOLD = "STYPE_DISK_THRESHOLD"

	//The string property
	//type used to
	//change a string value of a template (SNMP_DCA, SNMP_DCA_BASE).
	STYPE_SNMP_VERSION = "STYPE_SNMP_VERSION"

	// The string property
	//type used to
	//change a string value of a template (SNMP_DCA, SNMP_DCA_BASE).
	STYPE_COMMUNITY = "STYPE_COMMUNITY"

	//This string property
	//type is no
	//longer effective and must be considered deprecated.
	STYPE_SEND_PACKETS = "STYPE_SEND_PACKETS"

	//This string property
	//type is no
	//longer effective and must be considered deprecated.
	STYPE_TIMEOUT_PACKETS = "STYPE_TIMEOUT_PACKETS"

	//The string property
	//type used to
	//disable new sessions upon a match (also known as receive disable).
	STYPE_RECEIVE_DRAIN = "STYPE_RECEIVE_DRAIN"

	// The string property
	//type used in
	//database template.
	STYPE_RECEIVE_ROW = "STYPE_RECEIVE_ROW"

	// The string property
	//type used in
	//database template.
	STYPE_RECEIVE_COLUMN = "STYPE_RECEIVE_COLUMN"

	//The string property
	//type used to
	//enable EAV logging.
	STYPE_DEBUG = "STYPE_DEBUG"

	// The string property
	//type used in
	//LDAP template.
	STYPE_SECURITY = "STYPE_SECURITY"

	// The string property
	//type used to
	//indicate UDPs passive mode or port.
	STYPE_MODE = "STYPE_MODE"

	// The string property
	//type used to
	//represent the HTTPS cipher list.
	STYPE_CIPHER_LIST = "STYPE_CIPHER_LIST"

	STYPE_NAMESPACE = "STYPE_NAMESPACE"

	STYPE_PARAMETER_NAME = "STYPE_PARAMETER_NAME"

	STYPE_PARAMETER_VALUE = "STYPE_PARAMETER_VALUE"

	STYPE_PARAMETER_TYPE = "STYPE_PARAMETER_TYPE"

	STYPE_RETURN_TYPE = "STYPE_RETURN_TYPE"

	STYPE_RETURN_VALUE = "STYPE_RETURN_VALUE"

	STYPE_SOAP_FAULT = "STYPE_SOAP_FAULT"

	STYPE_SSL_OPTIONS = "STYPE_SSL_OPTIONS"

	STYPE_CLIENT_CERTIFICATE = "STYPE_CLIENT_CERTIFICATE"

	STYPE_PROTOCOL = "STYPE_PROTOCOL"

	STYPE_MANDATORY_ATTRS = "STYPE_MANDATORY_ATTRS"

	STYPE_FILENAME = "STYPE_FILENAME"

	STYPE_ACCOUNTING_NODE = "STYPE_ACCOUNTING_NODE"

	STYPE_ACCOUNTING_PORT = "STYPE_ACCOUNTING_PORT"

	STYPE_SERVER_ID = "STYPE_SERVER_ID"

	STYPE_CALL_ID = "STYPE_CALL_ID"

	STYPE_SESSION_ID = "STYPE_SESSION_ID"

	STYPE_FRAMED_ADDRESS = "STYPE_FRAMED_ADDRESS"

	STYPE_SNMP_PORT = "STYPE_SNMP_PORT"

	STYPE_AGGREGATE_DYNAMIC_RATIOS = "STYPE_AGGREGATE_DYNAMIC_RATIOS"

	STYPE_DB_COUNT = "STYPE_DB_COUNT"

	STYPE_NAS_IP = "STYPE_NAS_IP"

	STYPE_CLIENT_KEY = "STYPE_CLIENT_KEY"

	STYPE_MAX_LOAD_AVERAGE = "STYPE_MAX_LOAD_AVERAGE"

	STYPE_CONCURRENCY_LIMIT = "STYPE_CONCURRENCY_LIMIT"

	STYPE_FILTER_NEG = "STYPE_FILTER_NEG"

	STYPE_REQUEST = "STYPE_REQUEST"

	STYPE_HEADERS = "STYPE_HEADERS"

	STYPE_DIAMETER_ACCT_APPLICATION_ID = "STYPE_DIAMETER_ACCT_APPLICATION_ID"

	STYPE_DIAMETER_AUTH_APPLICATION_ID = "STYPE_DIAMETER_AUTH_APPLICATION_ID"

	STYPE_DIAMETER_ORIGIN_HOST = "STYPE_DIAMETER_ORIGIN_HOST"

	STYPE_DIAMETER_ORIGIN_REALM = "STYPE_DIAMETER_ORIGIN_REALM"

	STYPE_DIAMETER_HOST_IP_ADDRESS = "STYPE_DIAMETER_HOST_IP_ADDRESS"

	STYPE_DIAMETER_VENDOR_ID = "STYPE_DIAMETER_VENDOR_ID"

	STYPE_DIAMETER_PRODUCT_NAME = "STYPE_DIAMETER_PRODUCT_NAME"

	STYPE_DIAMETER_VENDOR_SPECIFIC_VENDOR_ID = "STYPE_DIAMETER_VENDOR_SPECIFIC_VENDOR_ID"

	STYPE_DIAMETER_VENDOR_SPECIFIC_ACCT_APPLICATION_ID = "STYPE_DIAMETER_VENDOR_SPECIFIC_ACCT_APPLICATION_ID"

	STYPE_DIAMETER_VENDOR_SPECIFIC_AUTH_APPLICATION_ID = "STYPE_DIAMETER_VENDOR_SPECIFIC_AUTH_APPLICATION_ID"

	STYPE_RUN_V2 = "STYPE_RUN_V2"

	STYPE_CLIENT_CERTIFICATE_V2 = "STYPE_CLIENT_CERTIFICATE_V2"

	STYPE_CLIENT_KEY_V2 = "STYPE_CLIENT_KEY_V2"
)

type StringValue struct {
	Type  StrPropertyType `xml:"type"`
	Value string          `xml:"value"`
}

type UserDefinedStringValue struct {
	Name  string // The user-defined string property name.
	Value string // The user-defined string property value.
}
type IMonitor interface {
	GetTemplateList() ([]MonitorTemplate, error)
	GetTemplateType(templateNames []string) ([]TemplateType, error)
	GetParentTemplate(templateNames []string) ([]string, error)
	GetTemplateAddressType(templateNames []string) ([]global_lb.AddressType, error)
	GetTemplateDestination(templateNames []string) ([]global_lb.MonitorIPPort, error)
	GetTemplateIntegerProperty(templateNames []string, propertyTypes []IntPropertyType) ([]IntegerValue, error)
	GetTemplateState(templateNames []string) ([]common.EnabledState, error)
	GetTemplateStringProperty(templateNames []string, propertyTypes []StrPropertyType) ([]StringValue, error)
	GetTemplateUserDefinedStringProperty(templateNames []string, propertyNames []string) ([]UserDefinedStringValue, error)
	GetTemplateReverseMode(templateNames []string) ([]bool, error)
	GetTemplateTransparentMode(templateNames []string) ([]bool, error)
	GetIgnoreDownResponseState(templateNames []string) ([]common.EnabledState, error)
}

type MonitorTemplate struct {
	TemplateName string       // The template name.
	TemplateType TemplateType // The template type.

}

var _ IMonitor = (*Monitor)(nil)

type Monitor struct {
	c *soap.Client
}

type GetParentTemplateBody struct {
	GetParentTemplate GetParentTemplate `xml:"tns:get_parent_template"`
}

type GetParentTemplate struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getParentTemplateResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetParentTemplateResponse struct {
			Return struct {
				Item []string `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_parent_templateResponse" json:"get_parent_template_response"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetParentTemplate(templateNames []string) ([]string, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetParentTemplateBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetParentTemplateBody{GetParentTemplate: GetParentTemplate{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getParentTemplateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetParentTemplateResponse.Return.Item, nil
}

type getTemplateAddressTypeReq struct {
	soap.BaseEnvEnvelope
	Body getTemplateAddressTypeBody `xml:"env:Body"`
}

type getTemplateAddressTypeBody struct {
	GetTemplateAddressType getTemplateAddressType `xml:"tns:get_template_address_type"`
}

type getTemplateAddressType struct {
	TemplateNames struct {
		Item []string `xml:"item"`
	} `xml:"template_names"`
}

type getTemplateAddressTypeResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTemplateAddressTypeResponse struct {
			Return struct {
				Item []global_lb.AddressType `xml:"item"`
			} `xml:"return"`
		} `xml:"get_template_address_typeResponse"`
	} `xml:"Body"`
}

func (m *Monitor) GetTemplateAddressType(templateNames []string) ([]global_lb.AddressType, error) {

	bt, err := m.c.Call(context.Background(), getTemplateAddressTypeReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getTemplateAddressTypeBody{GetTemplateAddressType: getTemplateAddressType{TemplateNames: struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: templateNames})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateAddressTypeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateAddressTypeResponse.Return.Item, err
}

type getTemplateDestinationReq struct {
	soap.BaseEnvEnvelope
	Body getTemplateDestinationBody `xml:"env:Body"`
}

type getTemplateDestinationBody struct {
	GetTemplateDestination getTemplateDestination `xml:"tns:get_template_destination"`
}

type getTemplateDestination struct {
	TemplateNames struct {
		Item []string `xml:"item"`
	} `xml:"template_names"`
}

type getTemplateDestinationResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTemplateDestinationResponse struct {
			Return struct {
				Item []struct {
					AddressType global_lb.AddressType `xml:"address_type"`
					Ipport      struct {
						Address string `xml:"address"`
						Port    int64  `xml:"port"`
					} `xml:"ipport"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_template_destinationResponse"`
	} `xml:"Body"`
}

func (m *Monitor) GetTemplateDestination(templateNames []string) ([]global_lb.MonitorIPPort, error) {

	bt, err := m.c.Call(context.Background(), getTemplateDestinationReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getTemplateDestinationBody{GetTemplateDestination: getTemplateDestination{struct {
			Item []string `xml:"item"`
		}(struct{ Item []string }{Item: templateNames})}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateDestinationResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []global_lb.MonitorIPPort
	for _, v := range resp.Body.GetTemplateDestinationResponse.Return.Item {
		res = append(res, global_lb.MonitorIPPort{
			AddressType: v.AddressType,
			IPPort: common.IPPortDefinition{
				Address: v.Ipport.Address,
				Port:    v.Ipport.Port,
			},
		})
	}

	return res, nil
}

type getTemplateIntegerPropertyReq struct {
	soap.BaseEnvEnvelope
	Body getTemplateIntegerPropertyBody `xml:"env:Body"`
}

type getTemplateIntegerPropertyBody struct {
	GetTemplateIntegerProperty getTemplateIntegerProperty `xml:"tns:get_template_integer_property"`
}

type getTemplateIntegerProperty struct {
	TemplateNames struct {
		Item []string `xml:"item"`
	} `xml:"template_names"`
	PropertyTypes struct {
		Item []IntPropertyType `xml:"item"`
	} `xml:"property_types"`
}

type getTemplateIntegerPropertyResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTemplateIntegerPropertyResponse struct {
			Return struct {
				Item []IntegerValue `xml:"item"`
			} `xml:"return"`
		} `xml:"get_template_integer_propertyResponse"`
	} `xml:"Body"`
}

func (m *Monitor) GetTemplateIntegerProperty(templateNames []string, propertyTypes []IntPropertyType) ([]IntegerValue, error) {

	bt, err := m.c.Call(context.Background(), getTemplateIntegerPropertyReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: getTemplateIntegerPropertyBody{GetTemplateIntegerProperty: getTemplateIntegerProperty{
			TemplateNames: struct {
				Item []string `xml:"item"`
			}{Item: templateNames},
			PropertyTypes: struct {
				Item []IntPropertyType `xml:"item"`
			}{Item: propertyTypes},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateIntegerPropertyResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateIntegerPropertyResponse.Return.Item, nil
}

type GetTemplateStateBody struct {
	GetTemplateState GetTemplateState `xml:"tns:get_template_state"`
}

type GetTemplateState struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getTemplateStateResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetTemplateStateResponse struct {
			Return struct {
				Item []common.EnabledState `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_template_stateResponse" json:"get_template_state_response"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetTemplateState(templateNames []string) ([]common.EnabledState, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateStateBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateStateBody{GetTemplateState: GetTemplateState{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateStateResponse.Return.Item, nil
}

type GetTemplateStringPropertyBody struct {
	GetTemplateStringProperty GetTemplateStringProperty `xml:"tns:get_template_string_property"`
}

type PropertyTypes struct {
	Item []StrPropertyType `xml:"item"`
}

type GetTemplateStringProperty struct {
	TemplateNames TemplateNames `xml:"template_names"`
	PropertyTypes PropertyTypes `xml:"property_types"`
}

type getTemplateStringPropertyResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetTemplateStringPropertyResponse struct {
			Return struct {
				Item []StringValue `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_template_string_propertyResponse"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetTemplateStringProperty(templateNames []string, propertyTypes []StrPropertyType) ([]StringValue, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateStringPropertyBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateStringPropertyBody{GetTemplateStringProperty: GetTemplateStringProperty{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
			PropertyTypes: PropertyTypes{
				Item: propertyTypes,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateStringPropertyResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateStringPropertyResponse.Return.Item, nil
}

type GetTemplateUserDefinedStringPropertyBody struct {
	GetTemplateUserDefinedStringProperty GetTemplateUserDefinedStringProperty `xml:"tns:get_template_user_defined_string_property"`
}

type PropertyNames struct {
	Item []string `xml:"item"`
}

type GetTemplateUserDefinedStringProperty struct {
	TemplateNames TemplateNames `xml:"template_names"`
	PropertyNames PropertyNames `xml:"property_names"`
}

type getTemplateUserDefinedStringPropertyResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetTemplateUserDefinedStringPropertyResponse struct {
			Return struct {
				Item []UserDefinedStringValue `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_template_user_defined_string_propertyResponse"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetTemplateUserDefinedStringProperty(templateNames []string, propertyNames []string) ([]UserDefinedStringValue, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateUserDefinedStringPropertyBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateUserDefinedStringPropertyBody{GetTemplateUserDefinedStringProperty: GetTemplateUserDefinedStringProperty{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
			PropertyNames: PropertyNames{
				Item: propertyNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateUserDefinedStringPropertyResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateUserDefinedStringPropertyResponse.Return.Item, nil
}

func New(c *soap.Client) *Monitor {
	return &Monitor{
		c: c,
	}
}

type getTemplateListReq struct {
	soap.BaseEnvEnvelope
	Body getTemplateListBody `xml:"env:Body"`
}

type getTemplateListBody struct {
	GetTemplateList struct{} `xml:"tns:get_template_list"`
}

type getTemplateListResp struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetTemplateListResponse struct {
			Return struct {
				Item []struct {
					TemplateName struct {
						Text string `xml:",chardata"`
					} `xml:"template_name"`
					TemplateType struct {
						Text TemplateType `xml:",chardata"`
					} `xml:"template_type"`
				} `xml:"item"`
			} `xml:"return"`
		} `xml:"get_template_listResponse"`
	} `xml:"Body"`
}

func (m *Monitor) GetTemplateList() ([]MonitorTemplate, error) {

	bt, err := m.c.Call(context.Background(), getTemplateListReq{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body:            getTemplateListBody{GetTemplateList: struct{}{}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateListResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	var res []MonitorTemplate
	for _, v := range resp.Body.GetTemplateListResponse.Return.Item {
		res = append(res, MonitorTemplate{
			TemplateName: v.TemplateName.Text,
			TemplateType: v.TemplateType.Text,
		})
	}

	return res, err
}

type TemplateNames struct {
	Item []string `xml:"item"`
}

type GetTemplateTypeBody struct {
	GetTemplateType GetTemplateType `xml:"tns:get_template_type"`
}

type GetTemplateType struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getTemplateTypeResp struct {
	XMLName xml.Name `xml:"Envelope" json:"envelope,omitempty"`
	Body    struct {
		GetTemplateTypeResponse struct {
			Return struct {
				Item []TemplateType `xml:"item"`
			} `xml:"return" json:"return,omitempty"`
		} `xml:"get_template_typeResponse" json:"get_template_typeresponse,omitempty"`
	} `xml:"Body" json:"body,omitempty"`
}

func (m *Monitor) GetTemplateType(templateNames []string) ([]TemplateType, error) {

	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateTypeBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateTypeBody{GetTemplateType: GetTemplateType{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateTypeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateTypeResponse.Return.Item, nil

}

type GetTemplateReverseModeBody struct {
	GetTemplateReverseMode GetTemplateReverseMode `xml:"tns:get_template_reverse_mode"`
}

type GetTemplateReverseMode struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getTemplateReverseModeResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetTemplateReverseModeResponse struct {
			Return struct {
				Item []bool `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_template_reverse_modeResponse"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetTemplateReverseMode(templateNames []string) ([]bool, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateReverseModeBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateReverseModeBody{GetTemplateReverseMode: GetTemplateReverseMode{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateReverseModeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateReverseModeResponse.Return.Item, nil
}

type GetTemplateTransparentModeBody struct {
	GetTemplateTransparentMode GetTemplateTransparentMode `xml:"tns:get_template_transparent_mode"`
}

type GetTemplateTransparentMode struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getTemplateTransparentModeResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetTemplateTransparentModeResponse struct {
			Return struct {
				Item []bool `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_template_transparent_modeResponse"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetTemplateTransparentMode(templateNames []string) ([]bool, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetTemplateTransparentModeBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetTemplateTransparentModeBody{GetTemplateTransparentMode: GetTemplateTransparentMode{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getTemplateTransparentModeResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetTemplateTransparentModeResponse.Return.Item, nil
}

type GetIgnoreDownResponseStateBody struct {
	GetIgnoreDownResponseState GetIgnoreDownResponseState `xml:"tns:get_ignore_down_response_state"`
}

type GetIgnoreDownResponseState struct {
	TemplateNames TemplateNames `xml:"template_names"`
}

type getIgnoreDownResponseStateResp struct {
	XMLName xml.Name `xml:"Envelope" json:"xml_name"`
	Body    struct {
		GetIgnoreDownResponseStateResponse struct {
			Return struct {
				Item []common.EnabledState `xml:"item" json:"item"`
			} `xml:"return" json:"return"`
		} `xml:"get_ignore_down_response_stateResponse"`
	} `xml:"Body" json:"body"`
}

func (m *Monitor) GetIgnoreDownResponseState(templateNames []string) ([]common.EnabledState, error) {
	type req struct {
		soap.BaseEnvEnvelope
		Body GetIgnoreDownResponseStateBody `xml:"env:Body"`
	}

	bt, err := m.c.Call(context.Background(), req{
		BaseEnvEnvelope: soap.NewBaseEnvEnvelope(tns),
		Body: GetIgnoreDownResponseStateBody{GetIgnoreDownResponseState: GetIgnoreDownResponseState{
			TemplateNames: TemplateNames{
				Item: templateNames,
			},
		}},
	})
	if err != nil {
		return nil, err
	}

	var resp getIgnoreDownResponseStateResp
	if err := xml.Unmarshal(bt, &resp); err != nil {
		return nil, err
	}

	return resp.Body.GetIgnoreDownResponseStateResponse.Return.Item, nil
}
