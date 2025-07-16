package asn1mapmodel

import "fmt"

// Error type declaration
type Error int

// Error constants
const (
	_ Error = iota
	UnknownSubscriber
	_
	UnknownMSC
	_
	UnidentifiedSubscriber
	AbsentSubscriberSM
	UnknownEquipment
	RoamingNotAllowed
	IllegalSubscriber
	BearerServiceNotProvisioned
	TeleserviceNotProvisioned
	IllegalEquipment
	CallBarred
	ForwardingViolation
	CugReject
	IllegalSSOperation
	SsErrorStatus
	SsNotAvailable
	SsSubscriptionViolation
	SsIncompatibility
	FacilityNotSupported
	OngoingGroupCall
	_
	_
	NoHandoverNumberAvailable
	SubsequentHandoverFailure
	AbsentSubscriber
	IncompatibleTerminal
	ShortTermDenial
	LongTermDenial
	SubscriberBusyForMTSMS
	SmDeliveryFailure
	MessageWaitingListFull
	SystemFailure
	DataMissing
	UnexpectedDataValue
	PwRegistrationFailure
	NegativePWCheck
	NoRoamingNumberAvailable
	_
	_
	TargetCellOutsideGroupCallArea
	NumberOfPWAttemptsViolation
	NumberChanged
	busySubscriber // TODO: export with solving duplicate issue
	NoSubscriberReply
	ForwardingFailed
	OrNotAllowed
	AtiNotAllowed
	NoGroupCallNumberAvailable
	ResourceLimitation
	UnauthorizedRequestingNetwork
	UnauthorizedLCSClient
	PositionMethodFailure
	_
	_
	_
	UnknownOrUnreachableLCSClient
	MmEventNotSupported
	AtsiNotAllowed
	AtmNotAllowed
	InformationNotAvailable
	_
	_
	_
	_
	_
	_
	_
	_
	UnknownAlphabet
	UssdBusy
)

// Error to string mapping
var errorStrings = map[Error]string{
	UnknownSubscriber:              "UnknownSubscriber",
	UnknownMSC:                     "UnknownMSC",
	UnidentifiedSubscriber:         "UnidentifiedSubscriber",
	AbsentSubscriberSM:             "AbsentSubscriberSM",
	UnknownEquipment:               "UnknownEquipment",
	RoamingNotAllowed:              "RoamingNotAllowed",
	IllegalSubscriber:              "IllegalSubscriber",
	BearerServiceNotProvisioned:    "BearerServiceNotProvisioned",
	TeleserviceNotProvisioned:      "TeleserviceNotProvisioned",
	IllegalEquipment:               "IllegalEquipment",
	CallBarred:                     "CallBarred",
	ForwardingViolation:            "ForwardingViolation",
	CugReject:                      "CugReject",
	IllegalSSOperation:             "IllegalSSOperation",
	SsErrorStatus:                  "SsErrorStatus",
	SsNotAvailable:                 "SsNotAvailable",
	SsSubscriptionViolation:        "SsSubscriptionViolation",
	SsIncompatibility:              "SsIncompatibility",
	FacilityNotSupported:           "FacilityNotSupported",
	OngoingGroupCall:               "OngoingGroupCall",
	NoHandoverNumberAvailable:      "NoHandoverNumberAvailable",
	AbsentSubscriber:               "AbsentSubscriber",
	SubscriberBusyForMTSMS:         "SubscriberBusyForMTSMS",
	SmDeliveryFailure:              "SmDeliveryFailure",
	MessageWaitingListFull:         "MessageWaitingListFull",
	SystemFailure:                  "SystemFailure",
	DataMissing:                    "DataMissing",
	UnexpectedDataValue:            "UnexpectedDataValue",
	PwRegistrationFailure:          "PwRegistrationFailure",
	NegativePWCheck:                "NegativePWCheck",
	NoRoamingNumberAvailable:       "NoRoamingNumberAvailable",
	busySubscriber:                 "BusySubscriber",
	NoSubscriberReply:              "NoSubscriberReply",
	ForwardingFailed:               "ForwardingFailed",
	OrNotAllowed:                   "OrNotAllowed",
	AtiNotAllowed:                  "AtiNotAllowed",
	NoGroupCallNumberAvailable:     "NoGroupCallNumberAvailable",
	ResourceLimitation:             "ResourceLimitation",
	UnauthorizedRequestingNetwork:  "UnauthorizedRequestingNetwork",
	UnauthorizedLCSClient:          "UnauthorizedLCSClient",
	PositionMethodFailure:          "PositionMethodFailure",
	UnknownAlphabet:                "UnknownAlphabet",
	UssdBusy:                       "UssdBusy",
	InformationNotAvailable:        "InformationNotAvailable",
	AtmNotAllowed:                  "AtmNotAllowed",
	MmEventNotSupported:            "MmEventNotSupported",
	AtsiNotAllowed:                 "AtsiNotAllowed",
	NumberOfPWAttemptsViolation:    "NumberOfPWAttemptsViolation",
	NumberChanged:                  "NumberChanged",
	TargetCellOutsideGroupCallArea: "TargetCellOutsideGroupCallArea",
	UnknownOrUnreachableLCSClient:  "UnknownOrUnreachableLCSClient",
	SubsequentHandoverFailure:      "SubsequentHandoverFailure",
	LongTermDenial:                 "LongTermDenial",
	ShortTermDenial:                "ShortTermDenial",
	IncompatibleTerminal:           "IncompatibleTerminal",
}

// GetErrorString converts an error value to its string representation.
func GetErrorString(errCode uint8) string {
	// Look up the error string
	if errStr, exists := errorStrings[Error(errCode)]; exists {
		return errStr
	}

	return fmt.Sprintf("Error code not define in ETSI TS 129 002 V15.5.0 / 3GPP TS 29.002 version 15.5.0 Release 15: %d", errCode)
}
