// Package management
// Introduced : BIG-IP_v9.0
//The Management module contains all the interfaces necessary to manage the system.
package management

// ViewInfo
// Introduced : BIG-IP_v9.0.3
// a struct that describes a view
type ViewInfo struct {
	ViewName  string   //	The name of the view
	ViewOrder int64    //	The order of the view within the named.conf file 0 = first in zone 0xffffffff on a change means to move the view to last any other number will move the view to that position, and bump up any view(s) 1 if necessary
	OptionSeq []string // a sequence of options for the view
	ZoneNames []string // a sequence of zones in this view
}

// ViewZone
//Introduced : BIG-IP_v9.0.3
//A struct that describes a view/zone
type ViewZone struct {
	ViewName string `xml:"view_name"` // The view name.
	ZoneName string `xml:"zone_name"` // The zone name.
}

// ZoneInfo
// Introduced : BIG-IP_v9.0.3
// a struct that describes a zone
type ZoneInfo struct {
	ViewName  string   // The name of the view
	ZoneName  string   // The name of the zone
	ZoneType  ZoneType // one of the types of ZoneType enum
	ZoneFile  string   // The name of the file for the zone data
	OptionSeq []string // A sequence of options for the zone
}

type ZoneType string

const (
	// UNSET not yet initialized
	UNSET ZoneType = "UNSET"

	// MASTER a master zone
	MASTER ZoneType = "MASTER"

	// SLAVE a slave zone
	SLAVE ZoneType = "SLAVE"

	// STUB a stub zone
	STUB ZoneType = "STUB"

	// FORWARD a forward zone
	FORWARD ZoneType = "FORWARD"

	// HINT a hint zone, “.”
	HINT ZoneType = "HINT"
)

// RRList
// Introduced : BIG-IP_v9.0.3
// struct that contains sequences for all possible RRtypes in a zone
type RRList struct {
	AList     []ARecord     // contains all A records
	NSList    []NSRecord    // contains all NS records
	CNAMEList []CNAMERecord // contains all CNAME records
	SOAList   []SOARecord   // contains all SOA records
	PTRList   []PTRRecord   // contains all PTR records
	HInfoList []HINFORecord // contains all HINFO records
	MXList    []MXRecord    // contains all MX records
	TXTList   []TXTRecord   // contains all TXT records
	SRVList   []SRVRecord   // contains all SRV records
	KeyList   []KEYRecord   // contains all KEY records
	SIGList   []SIGRecord   // contains all SIG records
	NXTList   []NXTRecord   // contains all NXT records
	AAAAList  []AAAARecord  // contains all AAAA records
	A6List    []A6Record    // contains all A6 records
	DNAMEList []DNAMERecord // contains all DNAME records
}

type ARecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	IPAddress  string `xml:"ip_address"`  // The ip address of the record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type NSRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	HostName   string `xml:"host_name"`   // The hostname of the Name Server
	TTL        int64  `xml:"ttl"`         // The TTL of the record
}

type CNAMERecord struct {
	DomainName string `xml:"domain_name"` //The domain name of the record
	Cname      string `xml:"cname"`       //The cname of the record
	TTL        int64  `xml:"ttl"`         //The TTL for this record
}

type SOARecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the zone
	Primary    string `xml:"primary"`     // The primary server of the zone
	Email      string `xml:"email"`       // The email address of the person responsible for the zone
	Serial     int64  `xml:"serial"`      // The serial number to start with for this zone
	Refresh    int64  `xml:"refresh"`     // The refresh interval(secs) for the zone
	Retry      int64  `xml:"retry"`       // The interval(secs) between retries for the zone
	Expire     int64  `xml:"expire"`      // The upper limit(secs) before a zone expires
	NegTTL     int64  `xml:"neg_ttl"`     // The Negative TTL for any RR from this zone
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type PTRRecord struct {
	IPAddress string `xml:"ip_address"` // The ip address of the record
	Dname     string `xml:"dname"`      // The DNAME for this record
	TTL       int64  `xml:"ttl"`        // The TTL for this record
}

type HINFORecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	Hardware   string `xml:"hardware"`    // The hardware info for this record
	OS         string `xml:"os"`          // The OS info for the record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type MXRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	Preference int64  `xml:"preference"`  // The preference to use for this record
	Mail       string `xml:"mail"`        // The mail-exchanger for this record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type TXTRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	Text       string `xml:"text"`        // The text entry for the record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type SRVRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	Priority   int64  `xml:"priority"`    // The priority to use for this record
	Weight     int64  `xml:"weight"`      // The weight to use for this record
	Port       int64  `xml:"port"`        // The port for this service
	Target     string `xml:"target"`      // The target to use for this record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type KEYRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	Flags      int    `xml:"flags"`       // 16bit flag for this key
	Protocol   int    `xml:"protocol"`    // 8bit protocol indicator
	Algorithm  int    `xml:"algorithm"`   // 8bit algorithm
	PublicKey  string `xml:"public_key"`  // a string containing the public key
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type SIGRecord struct {
	DomainName    string `xml:"domain_name"`    // The domain name of the record
	TypeCovered   int    `xml:"type_covered"`   // type of RR covered by this sig( NXT etc)
	Algorithm     int    `xml:"algorithm"`      // algorithm number used
	Labels        int    `xml:"labels"`         // how many labels in the original sig RR owner name
	OrigTTL       int64  `xml:"orig_ttl"`       // original ttl
	SigExpiration string `xml:"sig_expiration"` // expiration     date for sig.(secs since Jan 1….)
	SigInception  string `xml:"sig_inception"`  // start date for sig.(secs since Jan 1….)
	KeyTag        int    `xml:"key_tag"`        // used to select between multiple keys
	SignerName    string `xml:"signer_name"`    // domain name of the signer that generates the sig
	Signature     string `xml:"signature"`      // actual signature portion
	TTL           int64  `xml:"ttl"`            // The TTL for this record
}

type NXTRecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	NxtDomain  string `xml:"nxt_domain"`  // The next domain
	Types      string `xml:"types"`       // a string containing all resource record types
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type AAAARecord struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	IPAddress  string `xml:"ip_address"`  // The ip address of the record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type A6Record struct {
	DomainName string `xml:"domain_name"` // The domain name of the record
	PrefixBits int    `xml:"prefix_bits"` // Number of bits contained in prefix
	IPAddress  string `xml:"ip_address"`  // The ip address of the record
	PrefixName string `xml:"prefix_name"` // Name to lookup to get prefix of address
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}

type DNAMERecord struct {
	Label      string `xml:"label"`       // The label of the record
	DomainName string `xml:"domain_name"` // domain name for this dname record
	TTL        int64  `xml:"ttl"`         // The TTL for this record
}
