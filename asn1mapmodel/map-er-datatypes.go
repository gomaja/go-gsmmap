package asn1mapmodel

type UnknownSubscriberDiagnostic int
type AbsentSubscriberDiagnosticSM int
type AbsentSubscriberReason int
type RoamingNotAllowedCause int
type AdditionalRoamingNotAllowedCause int
type CallBarringCause int
type CUGRejectCause int
type PWRegistrationFailureCause int
type SMEnumeratedDeliveryFailureCause int
type FailureCauseParam int
type PositionMethodFailureDiagnostic int
type UnauthorizedLCSClientDiagnostic int

const (
	ImsiUnknown UnknownSubscriberDiagnostic = iota
	GprsEpsSubscriptionUnknown
	NpdbMismatch
)

// Maps for UnknownSubscriberDiagnostic
var unknownSubscriberDiagnosticMap = map[UnknownSubscriberDiagnostic]string{
	ImsiUnknown:                "IMSI Unknown",
	GprsEpsSubscriptionUnknown: "GPRS/EPS Subscription Unknown",
	NpdbMismatch:               "NPDB Mismatch",
}

const (
	NoPagingResponseViaTheMSC AbsentSubscriberDiagnosticSM = iota
	IMSIDetached
	RoamingRestriction
	DeregisteredInTheHLRForNonGPRS
	MSPurgedForNonGPRS
	NoPagingResponseViaTheSGSN
	GPRSDetached
	DeregisteredInTheHLRForGPRS
	MSPurgedForGPRS
	UnidentifiedSubscriberViaTheMSC
	UnidentifiedSubscriberViaTheSGSN
	DeregisteredInTheHSSHLRForIMS
	NoResponseViaTheIPSMGW

	reserved13to255 = 255
)

// Maps for AbsentSubscriberDiagnosticSM
var absentSubscriberDiagnosticSMMap = map[AbsentSubscriberDiagnosticSM]string{
	NoPagingResponseViaTheMSC:        "No Paging Response via the MSC",
	IMSIDetached:                     "IMSI Detached",
	RoamingRestriction:               "Roaming Restriction",
	DeregisteredInTheHLRForNonGPRS:   "Deregistered in the HLR for Non-GPRS",
	MSPurgedForNonGPRS:               "MS Purged for Non-GPRS",
	NoPagingResponseViaTheSGSN:       "No Paging Response via the SGSN",
	GPRSDetached:                     "GPRS Detached",
	DeregisteredInTheHLRForGPRS:      "Deregistered in the HLR for GPRS",
	MSPurgedForGPRS:                  "MS Purged for GPRS",
	UnidentifiedSubscriberViaTheMSC:  "Unidentified Subscriber via the MSC",
	UnidentifiedSubscriberViaTheSGSN: "Unidentified Subscriber via the SGSN",
	DeregisteredInTheHSSHLRForIMS:    "Deregistered in the HSS/HLR for IMS",
	NoResponseViaTheIPSMGW:           "No Response via the IPSMGW",
}

const (
	ImsiDetach AbsentSubscriberReason = iota
	RestrictedArea
	NoPageResponse
	PurgedMS
	MtRoamingRetry
	BusySubscriber
)

// Maps for AbsentSubscriberReason
var absentSubscriberReasonMap = map[AbsentSubscriberReason]string{
	ImsiDetach:     "IMSI Detach",
	RestrictedArea: "Restricted Area",
	NoPageResponse: "No Page Response",
	PurgedMS:       "Purged MS",
	MtRoamingRetry: "MT Roaming Retry",
	BusySubscriber: "Busy Subscriber",
}

const (
	SupportedRATTypesNotAllowed AdditionalRoamingNotAllowedCause = iota
)

// Maps for AdditionalRoamingNotAllowedCause
var additionalRoamingNotAllowedCauseMap = map[AdditionalRoamingNotAllowedCause]string{
	SupportedRATTypesNotAllowed: "Supported RAT Types Not Allowed",
}

const (
	PlmnRoamingNotAllowed RoamingNotAllowedCause = iota
	_
	_
	OperatorDeterminedBarring
)

// Maps for RoamingNotAllowedCause
var roamingNotAllowedCauseMap = map[RoamingNotAllowedCause]string{
	PlmnRoamingNotAllowed:     "PLMN Roaming Not Allowed",
	OperatorDeterminedBarring: "Operator Determined Barring",
}

const (
	BarringServiceActive CallBarringCause = iota
	OperatorBarring
)

// Maps for CallBarringCause
var callBarringCauseMap = map[CallBarringCause]string{
	BarringServiceActive: "Barring Service Active",
	OperatorBarring:      "Operator Barring",
}

const (
	IncomingCallsBarredWithinCUG CUGRejectCause = iota
	SubscriberNotMemberOfCUG
	_
	_
	_
	RequestedBasicServiceViolatesCUGConstraints
	_
	CalledPartySSInteractionViolation
)

// Maps for CUGRejectCause
var cugRejectCauseMap = map[CUGRejectCause]string{
	IncomingCallsBarredWithinCUG:                "Incoming Calls Barred Within CUG",
	SubscriberNotMemberOfCUG:                    "Subscriber Not Member of CUG",
	RequestedBasicServiceViolatesCUGConstraints: "Requested Basic Service Violates CUG Constraints",
	CalledPartySSInteractionViolation:           "Called Party SS Interaction Violation",
}

const (
	Undetermined PWRegistrationFailureCause = iota
	UnvalidFormat
	NewPasswordsMismatch
)

// Maps for PWRegistrationFailureCause
var pwRegistrationFailureCauseMap = map[PWRegistrationFailureCause]string{
	Undetermined:         "Undetermined",
	UnvalidFormat:        "Invalid Format",
	NewPasswordsMismatch: "New Passwords Mismatch",
}

const (
	MemoryCapacityExceeded SMEnumeratedDeliveryFailureCause = iota
	EquipmentProtocolError
	EquipmentNotSMEquipped
	UnknownServiceCentre
	ScCongestion
	InvalidSMEAddress
	SubscriberNotSCSubscriber
)

// Maps for SMEnumeratedDeliveryFailureCause
var smEnumeratedDeliveryFailureCauseMap = map[SMEnumeratedDeliveryFailureCause]string{
	MemoryCapacityExceeded:    "Memory Capacity Exceeded",
	EquipmentProtocolError:    "Equipment Protocol Error",
	EquipmentNotSMEquipped:    "Equipment Not SM-Equipped",
	UnknownServiceCentre:      "Unknown Service Centre",
	ScCongestion:              "SC Congestion",
	InvalidSMEAddress:         "Invalid SME Address",
	SubscriberNotSCSubscriber: "Subscriber Not SC Subscriber",
}

const (
	LimitReachedOnNumberOfConcurrentLocationRequests FailureCauseParam = iota
)

// Maps for FailureCauseParam
var failureCauseParamMap = map[FailureCauseParam]string{
	LimitReachedOnNumberOfConcurrentLocationRequests: "Limit Reached on Number of Concurrent Location Requests",
}

const (
	NoAdditionalInformation UnauthorizedLCSClientDiagnostic = iota
	ClientNotInMSPrivacyExceptionList
	CallToClientNotSetup
	PrivacyOverrideNotApplicable
	DisallowedByLocalRegulatoryRequirements
	UnauthorizedPrivacyClass
	UnauthorizedCallSessionUnrelatedExternalClient
	UnauthorizedCallSessionRelatedExternalClient
)

// Maps for UnauthorizedLCSClientDiagnostic
var unauthorizedLCSClientDiagnosticMap = map[UnauthorizedLCSClientDiagnostic]string{
	NoAdditionalInformation:                        "No Additional Information",
	ClientNotInMSPrivacyExceptionList:              "Client Not in MS Privacy Exception List",
	CallToClientNotSetup:                           "Call to Client Not Setup",
	PrivacyOverrideNotApplicable:                   "Privacy Override Not Applicable",
	DisallowedByLocalRegulatoryRequirements:        "Disallowed by Local Regulatory Requirements",
	UnauthorizedPrivacyClass:                       "Unauthorized Privacy Class",
	UnauthorizedCallSessionUnrelatedExternalClient: "Unauthorized Call Session Unrelated External Client",
	UnauthorizedCallSessionRelatedExternalClient:   "Unauthorized Call Session Related External Client",
}

const (
	Congestion PositionMethodFailureDiagnostic = iota
	InsufficientResources
	InsufficientMeasurementData
	InconsistentMeasurementData
	LocationProcedureNotCompleted
	LocationProcedureNotSupportedByTargetMS
	QoSNotAttainable
	PositionMethodNotAvailableInNetwork
	PositionMethodNotAvailableInLocationArea
)

// Maps for PositionMethodFailureDiagnostic
var positionMethodFailureDiagnosticMap = map[PositionMethodFailureDiagnostic]string{
	Congestion:                               "Congestion",
	InsufficientResources:                    "Insufficient Resources",
	InsufficientMeasurementData:              "Insufficient Measurement Data",
	InconsistentMeasurementData:              "Inconsistent Measurement Data",
	LocationProcedureNotCompleted:            "Location Procedure Not Completed",
	LocationProcedureNotSupportedByTargetMS:  "Location Procedure Not Supported by Target MS",
	QoSNotAttainable:                         "QoS Not Attainable",
	PositionMethodNotAvailableInNetwork:      "Position Method Not Available in Network",
	PositionMethodNotAvailableInLocationArea: "Position Method Not Available in Location Area",
}
