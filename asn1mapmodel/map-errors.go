package asn1mapmodel

import "fmt"

// Error type declaration
type Error int

// Error constants
const (
	_ Error = iota
	ErrorUnknownSubscriber
	_
	ErrorUnknownMSC
	_
	ErrorUnidentifiedSubscriber
	ErrorAbsentSubscriberSM
	ErrorUnknownEquipment
	ErrorRoamingNotAllowed
	ErrorIllegalSubscriber
	ErrorBearerServiceNotProvisioned
	ErrorTeleserviceNotProvisioned
	ErrorIllegalEquipment
	ErrorCallBarred
	ErrorForwardingViolation
	ErrorCugReject
	ErrorIllegalSSOperation
	ErrorSsErrorStatus
	ErrorSsNotAvailable
	ErrorSsSubscriptionViolation
	ErrorSsIncompatibility
	ErrorFacilityNotSupported
	ErrorOngoingGroupCall
	_
	_
	ErrorNoHandoverNumberAvailable
	ErrorSubsequentHandoverFailure
	ErrorAbsentSubscriber
	ErrorIncompatibleTerminal
	ErrorShortTermDenial
	ErrorLongTermDenial
	ErrorSubscriberBusyForMTSMS
	ErrorSmDeliveryFailure
	ErrorMessageWaitingListFull
	ErrorSystemFailure
	ErrorDataMissing
	ErrorUnexpectedDataValue
	ErrorPwRegistrationFailure
	ErrorNegativePWCheck
	ErrorNoRoamingNumberAvailable
	_
	_
	ErrorTargetCellOutsideGroupCallArea
	ErrorNumberOfPWAttemptsViolation
	ErrorNumberChanged
	ErrorBusySubscriber
	ErrorNoSubscriberReply
	ErrorForwardingFailed
	ErrorOrNotAllowed
	ErrorAtiNotAllowed
	ErrorNoGroupCallNumberAvailable
	ErrorResourceLimitation
	ErrorUnauthorizedRequestingNetwork
	ErrorUnauthorizedLCSClient
	ErrorPositionMethodFailure
	_
	_
	_
	ErrorUnknownOrUnreachableLCSClient
	ErrorMmEventNotSupported
	ErrorAtsiNotAllowed
	ErrorAtmNotAllowed
	ErrorInformationNotAvailable
	_
	_
	_
	_
	_
	_
	_
	_
	ErrorUnknownAlphabet
	ErrorUssdBusy
)

// Error to string mapping
var errorStrings = map[Error]string{
	ErrorUnknownSubscriber:              "UnknownSubscriber",
	ErrorUnknownMSC:                     "UnknownMSC",
	ErrorUnidentifiedSubscriber:         "UnidentifiedSubscriber",
	ErrorAbsentSubscriberSM:             "AbsentSubscriberSM",
	ErrorUnknownEquipment:               "UnknownEquipment",
	ErrorRoamingNotAllowed:              "RoamingNotAllowed",
	ErrorIllegalSubscriber:              "IllegalSubscriber",
	ErrorBearerServiceNotProvisioned:    "BearerServiceNotProvisioned",
	ErrorTeleserviceNotProvisioned:      "TeleserviceNotProvisioned",
	ErrorIllegalEquipment:               "IllegalEquipment",
	ErrorCallBarred:                     "CallBarred",
	ErrorForwardingViolation:            "ForwardingViolation",
	ErrorCugReject:                      "CugReject",
	ErrorIllegalSSOperation:             "IllegalSSOperation",
	ErrorSsErrorStatus:                  "SsErrorStatus",
	ErrorSsNotAvailable:                 "SsNotAvailable",
	ErrorSsSubscriptionViolation:        "SsSubscriptionViolation",
	ErrorSsIncompatibility:              "SsIncompatibility",
	ErrorFacilityNotSupported:           "FacilityNotSupported",
	ErrorOngoingGroupCall:               "OngoingGroupCall",
	ErrorNoHandoverNumberAvailable:      "NoHandoverNumberAvailable",
	ErrorAbsentSubscriber:               "AbsentSubscriber",
	ErrorSubscriberBusyForMTSMS:         "SubscriberBusyForMTSMS",
	ErrorSmDeliveryFailure:              "SmDeliveryFailure",
	ErrorMessageWaitingListFull:         "MessageWaitingListFull",
	ErrorSystemFailure:                  "SystemFailure",
	ErrorDataMissing:                    "DataMissing",
	ErrorUnexpectedDataValue:            "UnexpectedDataValue",
	ErrorPwRegistrationFailure:          "PwRegistrationFailure",
	ErrorNegativePWCheck:                "NegativePWCheck",
	ErrorNoRoamingNumberAvailable:       "NoRoamingNumberAvailable",
	ErrorBusySubscriber:                 "BusySubscriber",
	ErrorNoSubscriberReply:              "NoSubscriberReply",
	ErrorForwardingFailed:               "ForwardingFailed",
	ErrorOrNotAllowed:                   "OrNotAllowed",
	ErrorAtiNotAllowed:                  "AtiNotAllowed",
	ErrorNoGroupCallNumberAvailable:     "NoGroupCallNumberAvailable",
	ErrorResourceLimitation:             "ResourceLimitation",
	ErrorUnauthorizedRequestingNetwork:  "UnauthorizedRequestingNetwork",
	ErrorUnauthorizedLCSClient:          "UnauthorizedLCSClient",
	ErrorPositionMethodFailure:          "PositionMethodFailure",
	ErrorUnknownAlphabet:                "UnknownAlphabet",
	ErrorUssdBusy:                       "UssdBusy",
	ErrorInformationNotAvailable:        "InformationNotAvailable",
	ErrorAtmNotAllowed:                  "AtmNotAllowed",
	ErrorMmEventNotSupported:            "MmEventNotSupported",
	ErrorAtsiNotAllowed:                 "AtsiNotAllowed",
	ErrorNumberOfPWAttemptsViolation:    "NumberOfPWAttemptsViolation",
	ErrorNumberChanged:                  "NumberChanged",
	ErrorTargetCellOutsideGroupCallArea: "TargetCellOutsideGroupCallArea",
	ErrorUnknownOrUnreachableLCSClient:  "UnknownOrUnreachableLCSClient",
	ErrorSubsequentHandoverFailure:      "SubsequentHandoverFailure",
	ErrorLongTermDenial:                 "LongTermDenial",
	ErrorShortTermDenial:                "ShortTermDenial",
	ErrorIncompatibleTerminal:           "IncompatibleTerminal",
}

// GetErrorString converts an error value to its string representation.
func GetErrorString(errCode uint8) string {
	// Look up the error string
	if errStr, exists := errorStrings[Error(errCode)]; exists {
		return errStr
	}

	return fmt.Sprintf("Error code not define in ETSI TS 129 002 V15.5.0 / 3GPP TS 29.002 version 15.5.0 Release 15: %d", errCode)
}
