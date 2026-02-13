package asn1mapmodel

import (
	"encoding/asn1"

	"github.com/gomaja/go-gsmmap/utils"
)

// UpdateLocationArg represents the UpdateLocation operation argument
// UpdateLocationArg ::= SEQUENCE {
//
//	imsi                        IMSI,
//	msc-Number                  [1] ISDN-AddressString,
//	vlr-Number                  ISDN-AddressString,
//	lmsi                        [10] LMSI OPTIONAL,
//	extensionContainer          ExtensionContainer OPTIONAL,
//	...,
//	vlr-Capability              [6] VLR-Capability OPTIONAL,
//	informPreviousNetworkEntity [11] NULL OPTIONAL,
//	cs-LCS-NotSupportedByUE     [12] NULL OPTIONAL,
//	v-gmlc-Address              [2] GSN-Address OPTIONAL,
//	add-info                    [13] ADD-Info OPTIONAL,
//	pagingArea                  [14] PagingArea OPTIONAL,
//	skipSubscriberDataUpdate    [15] NULL OPTIONAL,
//	restorationIndicator        [16] NULL OPTIONAL,
//	eplmn-List                  [3] EPLMN-List OPTIONAL,
//	mme-DiameterAddress         [4] NetworkNodeDiameterAddress OPTIONAL }
type UpdateLocationArg struct {
	IMSI      IMSI              // Required
	MSCNumber ISDNAddressString `asn1:"tag:1"` // Required [1]
	VLRNumber ISDNAddressString // Required
	//LMSI                        LMSI                        `asn1:"tag:10,optional"`  // [10] OPTIONAL
	//ExtensionContainer          *ExtensionContainer         `asn1:"optional"`         // OPTIONAL
	VlrCapability VlrCapability `asn1:"tag:6,optional"` // [6] OPTIONAL
	//InformPreviousNetworkEntity asn1.RawValue               `asn1:"tag:11,optional"` // [11] NULL OPTIONAL
	//CSLCSNotSupportedByUE       asn1.RawValue               `asn1:"tag:12,optional"` // [12] NULL OPTIONAL
	//VGMLCAddress                GSNAddress                  `asn1:"tag:2,optional"`  // [2] OPTIONAL
	//AddInfo                     *ADDInfo                    `asn1:"tag:13,optional"` // [13] OPTIONAL
	//PagingArea                  *PagingArea                 `asn1:"tag:14,optional"` // [14] OPTIONAL
	//SkipSubscriberDataUpdate    asn1.RawValue               `asn1:"tag:15,optional"` // [15] NULL OPTIONAL
	//RestorationIndicator        asn1.RawValue               `asn1:"tag:16,optional"` // [16] NULL OPTIONAL
	//EPLMNList                   *EPLMNList                  `asn1:"tag:3,optional"`  // [3] OPTIONAL
	//MMEDiameterAddress          *NetworkNodeDiameterAddress `asn1:"tag:4,optional"`  // [4] OPTIONAL
}

// VlrCapability represents the capabilities of a VLR
// VLR-Capability ::= SEQUENCE{
//
//	supportedCamelPhases          [0] SupportedCamelPhases OPTIONAL,
//	extensionContainer            ExtensionContainer OPTIONAL,
//	solsaSupportIndicator         [2] NULL OPTIONAL,
//	istSupportIndicator           [1] IST-SupportIndicator OPTIONAL,
//	superChargerSupportedInServingNetworkEntity [3] SuperChargerInfo OPTIONAL,
//	longFTN-Supported             [4] NULL OPTIONAL,
//	supportedLCS-CapabilitySets   [5] SupportedLCS-CapabilitySets OPTIONAL,
//	offeredCamel4CSIs             [6] OfferedCamel4CSIs OPTIONAL,
//	supportedRAT-TypesIndicator   [7] SupportedRAT-Types OPTIONAL,
//	longGroupID-Supported         [8] NULL OPTIONAL,
//	mtRoamingForwardingSupported  [9] NULL OPTIONAL,
//	msisdn-lessOperation-Supported [10] NULL OPTIONAL }
//
// SupportedCamelPhases indicates which CAMEL phases are supported
// SupportedCamelPhases ::= BIT STRING {
//
//	phase1 (0),
//	phase2 (1),
//	phase3 (2),
//	phase4 (3)} (SIZE (1..16))
//
// SupportedLCSCapabilitySets indicates which LCS capability sets are supported
// SupportedLCS-CapabilitySets ::= BIT STRING {
//
//	lcsCapabilitySet1 (0),
//	lcsCapabilitySet2 (1),
//	lcsCapabilitySet3 (2),
//	lcsCapabilitySet4 (3),
//	lcsCapabilitySet5 (4) } (SIZE (2..16))
type VlrCapability struct {
	SupportedCamelPhases       asn1.BitString `asn1:"tag:0,optional"`
	SupportedLCSCapabilitySets asn1.BitString `asn1:"tag:5,optional"`
}

// CAMEL Phase bit positions
const (
	CamelPhase1 = 0
	CamelPhase2 = 1
	CamelPhase3 = 2
	CamelPhase4 = 3
)

// LCS Capability Set bit positions
const (
	LCSCapabilitySet1 = 0
	LCSCapabilitySet2 = 1
	LCSCapabilitySet3 = 2
	LCSCapabilitySet4 = 3
	LCSCapabilitySet5 = 4
)

func (updLoc *UpdateLocationArg) GetImsiString() (string, error) {
	return utils.DecodeTBCDDigits(updLoc.IMSI)
}

// UpdateGprsLocationArg represents the UpdateGprsLocation operation argument
// UpdateGprsLocationArg ::= SEQUENCE {
//
//	imsi                         IMSI,
//	sgsn-Number                  ISDN-AddressString,
//	sgsn-Address                 GSN-Address,
//	extensionContainer           ExtensionContainer OPTIONAL,
//	...,
//	sgsn-Capability              [0] SGSN-Capability OPTIONAL,
//	informPreviousNetworkEntity  [1] NULL OPTIONAL,
//	ps-LCS-NotSupportedByUE      [2] NULL OPTIONAL,
//	v-gmlc-Address               [3] GSN-Address OPTIONAL,
//	add-info                     [4] ADD-Info OPTIONAL,
//	eps-info                     [5] EPS-Info OPTIONAL,
//	servingNodeTypeIndicator     [6] NULL OPTIONAL,
//	skipSubscriberDataUpdate     [7] NULL OPTIONAL,
//	usedRAT-Type                 [8] Used-RAT-Type OPTIONAL,
//	gprsSubscriptionDataNotNeeded [9] NULL OPTIONAL,
//	nodeTypeIndicator            [10] NULL OPTIONAL,
//	areaRestricted               [11] NULL OPTIONAL,
//	ue-reachableIndicator        [12] NULL OPTIONAL,
//	epsSubscriptionDataNotNeeded [13] NULL OPTIONAL,
//	ue-srvcc-Capability          [14] UE-SRVCC-Capability OPTIONAL,
//	eplmn-List                   [15] EPLMN-List OPTIONAL,
//	mmeNumberforMTSMS            [16] ISDN-AddressString OPTIONAL,
//	smsRegisterRequest           [17] SMSRegisterRequest OPTIONAL,
//	sms-Only                     [18] NULL OPTIONAL }
type UpdateGprsLocationArg struct {
	IMSI           IMSI              // Required
	SGSNNumber     ISDNAddressString // Required
	SGSNAddress    GSNAddress        // Required
	SGSNCapability SGSNCapability    `asn1:"tag:0,optional"` // [0] OPTIONAL
}

// SGSNCapability represents the capabilities of an SGSN
// SGSN-Capability ::= SEQUENCE{
//
//	solsaSupportIndicator                        NULL OPTIONAL,
//	extensionContainer                           [1] ExtensionContainer OPTIONAL,
//	...,
//	superChargerSupportedInServingNetworkEntity  [2] SuperChargerInfo OPTIONAL,
//	gprsEnhancementsSupportIndicator             [3] NULL OPTIONAL,
//	supportedCamelPhases                         [4] SupportedCamelPhases OPTIONAL,
//	supportedLCS-CapabilitySets                  [5] SupportedLCS-CapabilitySets OPTIONAL,
//	offeredCamel4CSIs                            [6] OfferedCamel4CSIs OPTIONAL,
//	smsCallBarringSupportIndicator               [7] NULL OPTIONAL,
//	supportedRAT-TypesIndicator                  [8] SupportedRAT-Types OPTIONAL,
//	supportedFeatures                            [9] SupportedFeatures OPTIONAL,
//	t-adsDataRetrieval                           [10] NULL OPTIONAL,
//	homogeneousSupportOfIMSVoiceOverPSSessions   [11] BOOLEAN OPTIONAL,
//	cancellationTypeInitialAttach                [12] NULL OPTIONAL,
//	additionalMsisdnSupport                      [13] NULL OPTIONAL,
//	msisdn-lessOperation-Supported               [14] NULL OPTIONAL }
//
// SupportedLCS-CapabilitySets ::= BIT STRING {
//
//	lcsCapabilitySet1 (0),
//	lcsCapabilitySet2 (1),
//	lcsCapabilitySet3 (2),
//	lcsCapabilitySet4 (3),
//	lcsCapabilitySet5 (4) } (SIZE (2..16))
type SGSNCapability struct {

	// In a case observed, the value appeared: asn1.RawValue{Class: asn1.ClassContextSpecific, Tag: asn1.TagBitString}
	GprsEnhancementsSupportIndicator asn1.RawValue `asn1:"tag:3,optional"`

	SupportedLCSCapabilitySets asn1.BitString `asn1:"tag:5,optional"`
}

func (updGprsLoc *UpdateGprsLocationArg) GetImsiString() (string, error) {
	return utils.DecodeTBCDDigits(updGprsLoc.IMSI)
}

func (updGprsLoc *UpdateGprsLocationArg) GetSGSNNumberString() (string, error) {
	_, _, _, Digits := DecodeAddressString(updGprsLoc.SGSNNumber)
	return utils.DecodeTBCDDigits(Digits)
}

func (updLoc *UpdateLocationArg) GetMSCNumberString() (string, error) {
	_, _, _, Digits := DecodeAddressString(updLoc.MSCNumber)
	return utils.DecodeTBCDDigits(Digits)
}

func (updLoc *UpdateLocationArg) GetVLRNumberString() (string, error) {
	_, _, _, Digits := DecodeAddressString(updLoc.VLRNumber)
	return utils.DecodeTBCDDigits(Digits)
}

func (updGprsLoc *UpdateGprsLocationArg) GetSGSNAddressString() (string, error) {
	return utils.ParseGSNAddress(updGprsLoc.SGSNAddress)
}

// UpdateLocationRes represents the UpdateLocation operation result
// UpdateLocationRes ::= SEQUENCE {
//
//	hlr-Number             ISDN-AddressString,
//	extensionContainer     ExtensionContainer OPTIONAL,
//	...,
//	add-Capability         NULL OPTIONAL,
//	pagingArea-Capability  [0]NULL OPTIONAL }
type UpdateLocationRes struct {
	HLRNumber ISDNAddressString // Required
}

func (updLocRes *UpdateLocationRes) GetHLRNumberString() (string, error) {
	_, _, _, Digits := DecodeAddressString(updLocRes.HLRNumber)
	return utils.DecodeTBCDDigits(Digits)
}

// UpdateGprsLocationRes represents the UpdateGprsLocation operation result
// UpdateGprsLocationRes ::= SEQUENCE {
//
//	hlr-Number                    ISDN-AddressString,
//	extensionContainer            ExtensionContainer OPTIONAL,
//	...,
//	add-Capability                NULL OPTIONAL,
//	sgsn-mmeSeparationSupported   [0] NULL OPTIONAL,
//	mmeRegisteredforSMS           [1] NULL OPTIONAL }
type UpdateGprsLocationRes struct {
	HLRNumber ISDNAddressString // Required
}

func (updGprsLocRes *UpdateGprsLocationRes) GetHLRNumberString() (string, error) {
	_, _, _, Digits := DecodeAddressString(updGprsLocRes.HLRNumber)
	return utils.DecodeTBCDDigits(Digits)
}

// RequestedInfo represents the ASN.1 SEQUENCE of request flags
// RequestedInfo ::= SEQUENCE {
// locationInformation	[0] NULL		OPTIONAL,
// subscriberState	[1] NULL		OPTIONAL,
// extensionContainer	[2] ExtensionContainer	OPTIONAL,
// ...,
// currentLocation	[3] NULL		OPTIONAL,
// requestedDomain	[4] DomainType	OPTIONAL,
// imei			[6] NULL		OPTIONAL,
// ms-classmark	[5] NULL		OPTIONAL,
// mnpRequestedInfo	[7] NULL 		OPTIONAL,
// locationInformationEPS-Supported	[11] NULL 	OPTIONAL,
// t-adsData		[8] NULL		OPTIONAL,
// requestedNodes	[9] RequestedNodes	OPTIONAL,
// servingNodeIndication	[10] NULL		OPTIONAL,
// localTimeZoneRequest	[12] NULL		OPTIONAL
// }
//
// --	currentLocation and locationInformationEPS-Supported shall be absent if
// --	locationInformation is absent
// --	t-adsData shall be absent in messages sent to the VLR
// --	requestedNodes shall be absent if requestedDomain is "cs-Domain"
// --	servingNodeIndication shall be absent if locationInformation is absent;
// --	servingNodeIndication shall be absent if current location is present;
// --	servingNodeIndication indicates by its presence that only the serving node's
// --	address (MME-Name or SGSN-Number or VLR-Number) is requested.
type RequestedInfo struct {
	LocationInformation asn1.RawValue `asn1:"tag:0,optional"`
	SubscriberState     asn1.RawValue `asn1:"tag:1,optional"`
	//ExtensionContainer omitted (tag:2)
	CurrentLocation                 asn1.RawValue  `asn1:"tag:3,optional"`
	RequestedDomain                 DomainType     `asn1:"tag:4,optional,default:-1"`
	IMEI                            asn1.RawValue  `asn1:"tag:6,optional"`
	MsClassmark                     asn1.RawValue  `asn1:"tag:5,optional"`
	MnpRequestedInfo                asn1.RawValue  `asn1:"tag:7,optional"`
	LocationInformationEPSSupported asn1.RawValue  `asn1:"tag:11,optional"`
	TAdsData                        asn1.RawValue  `asn1:"tag:8,optional"`
	RequestedNodes                  RequestedNodes `asn1:"tag:9,optional"`
	ServingNodeIndication           asn1.RawValue  `asn1:"tag:10,optional"`
	LocalTimeZoneRequest            asn1.RawValue  `asn1:"tag:12,optional"`
}

// DomainType represents the domain type enumeration
// DomainType ::= ENUMERATED {
//
//	cs-Domain   (0),
//	ps-Domain   (1),
//	...}
//
// -- exception handling:
// -- reception of values > 1 shall be mapped to 'cs-Domain'
type DomainType = asn1.Enumerated

const (
	CsDomain DomainType = 0
	PsDomain DomainType = 1
)

// RequestedNodes bit positions
// RequestedNodes ::= BIT STRING {
//
//	mme   (0),
//	sgsn  (1)} (SIZE (1..8))
//
// -- Other bits than listed above shall be discarded.
type RequestedNodes = asn1.BitString

const (
	RequestedNodeMME  = 0
	RequestedNodeSGSN = 1
)

// AnyTimeInterrogationArg represents the ATI request
// AnyTimeInterrogationArg ::= SEQUENCE {
// subscriberIdentity	[0] SubscriberIdentity,
// requestedInfo	[1] RequestedInfo,
// gsmSCF-Address	[3] ISDN-AddressString,
// extensionContainer	[2] ExtensionContainer	OPTIONAL,
// ...}
type AnyTimeInterrogationArg struct {
	SubscriberIdentity asn1.RawValue     `asn1:"tag:0"` // CHOICE wrapped in context-specific tag
	RequestedInfo      RequestedInfo     `asn1:"tag:1"`
	GsmSCFAddress      ISDNAddressString `asn1:"tag:3"`
	//ExtensionContainer omitted (tag:2)
}
