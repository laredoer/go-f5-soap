package common

// EnabledState
//Introduced : BIG-IP_v9.0
//A list of enabled states.
type EnabledState string

const (
	// StateDisabled The object is disabled.
	StateDisabled EnabledState = "STATE_DISABLED"

	// StateEnabled The object is enabled.
	StateEnabled EnabledState = "STATE_ENABLED"
)

// AvailabilityStatus
// Introduced : BIG-IP_v9.0
// A list of possible values for an object&aposs availability status.
type AvailabilityStatus string

const (
	// AvailabilityStatusNone  Error scenario.
	AvailabilityStatusNone AvailabilityStatus = "AVAILABILITY_STATUS_NONE"

	// AvailabilityStatusGreen The object is available in some capacity.
	AvailabilityStatusGreen AvailabilityStatus = "AVAILABILITY_STATUS_GREEN"

	// AvailabilityStatusYellow The object is not available at the current moment, but may become available again even without user intervention.
	AvailabilityStatusYellow AvailabilityStatus = "AVAILABILITY_STATUS_YELLOW"

	// AvailabilityStatusRed The object is not available, and will require user intervention to make this object available again.
	AvailabilityStatusRed AvailabilityStatus = "AVAILABILITY_STATUS_RED"

	// AvailabilityStatusBlue The object’s availability status is unknown.
	AvailabilityStatusBlue AvailabilityStatus = "AVAILABILITY_STATUS_BLUE"

	// AvailabilityStatusGray  The object’s is unlicensed.
	AvailabilityStatusGray AvailabilityStatus = "AVAILABILITY_STATUS_GRAY"
)

// EnabledStatus
// Introduced : BIG-IP_v9.0
// A list of possible values for enabled status.
type EnabledStatus string

const (
	// EnabledStatusNone Error scenario.
	EnabledStatusNone EnabledStatus = "ENABLED_STATUS_NONE"

	// EnabledStatusEnabled
	// The object is active when in Green availability status.
	// It may or may not be active when in Blue availability status.
	EnabledStatusEnabled EnabledStatus = "ENABLED_STATUS_ENABLED"

	// EnabledStatusDisabled  The object is inactive regardless of availability status.
	EnabledStatusDisabled EnabledStatus = "ENABLED_STATUS_DISABLED"

	// EnabledStatusDisabledByParent
	// The object is inactive regardless of availability status because its parent has been disabled,
	// but the object itself is still enabled.
	EnabledStatusDisabledByParent EnabledStatus = "ENABLED_STATUS_DISABLED_BY_PARENT"
)

// ObjectStatus
// Introduced : BIG-IP_v9.0
// An struct that specifies an object status.
type ObjectStatus struct {
	AvailabilityStatus AvailabilityStatus // The availability color status of the object.
	EnabledStatus      EnabledStatus      // The enabled status of the object.
	StatusDescription  string             // The textual description of the object’s status.
}

type IPPortDefinition struct {
	Address string `xml:"address"`
	Port    int64  `xml:"port"`
}

type MemberRatio struct {
	Member IPPortDefinition
	Ratio  int64
}

type MemberObjectStatus struct {
	Member IPPortDefinition
	Status ObjectStatus
}
