package asn1mapmodel

import (
	"encoding/asn1"

	"github.com/gomaja/go-gsmmap/utils"
)

// ExtensionNo Constants representing the extension indicator
const (
	ExtensionNo = 0b10000000 // bit 8 set to 1, indicating no extension
)

// Constants representing the nature of address indicator (bits 7, 6, 5)
const (
	AddressNatureUnknown           = 0b000 << 4
	AddressNatureInternational     = 0b001 << 4
	AddressNatureNational          = 0b010 << 4
	AddressNatureNetworkSpecific   = 0b011 << 4
	AddressNatureSubscriber        = 0b100 << 4
	AddressNatureReserved          = 0b101 << 4
	AddressNatureAbbreviated       = 0b110 << 4
	AddressNatureReservedExtension = 0b111 << 4
)

// Constants representing the numbering plan indicator (bits 4, 3, 2, 1)
const (
	NumberingPlanUnknown           = 0b0000
	NumberingPlanISDN              = 0b0001
	NumberingPlanSpare1            = 0b0010
	NumberingPlanData              = 0b0011
	NumberingPlanTelex             = 0b0100
	NumberingPlanSpare2            = 0b0101
	NumberingPlanLandMobile        = 0b0110
	NumberingPlanSpare3            = 0b0111
	NumberingPlanNational          = 0b1000
	NumberingPlanPrivate           = 0b1001
	NumberingPlanReservedExtension = 0b1111
)

// RoutingInfoForSMArg defines the ASN.1 SEQUENCE structure with optional fields
type RoutingInfoForSMArg struct {
	MSISDN               ISDNAddressString `asn1:"tag:0"` // Context Specific (tcap.ContextSpecific) Tag code [0] , primitive form (tcap.Primitive)
	SmRpPri              bool              `asn1:"tag:1"`
	ServiceCentreAddress AddressString     `asn1:"tag:2"`
	//ExtensionContainer        *ExtensionContainer     `asn1:"tag:6,optional"`
	//GPRSSupportIndicator      *asn1.RawValue          `asn1:"tag:7,optional"`
	//SM_RP_MTI                 *SM_RP_MTI              `asn1:"tag:8,optional"`
	//SM_RP_SMEA                *SM_RP_SMEA             `asn1:"tag:9,optional"`
	//SM_DeliveryNotIntended    *SM_DeliveryNotIntended `asn1:"tag:10,optional"`
	//IP_SM_GWGuidanceIndicator *asn1.RawValue          `asn1:"tag:11,optional"`
	//IMSI                      *IMSI                   `asn1:"tag:12,optional"`
}

// RoutingInfoForSMRes defines the ASN.1 SEQUENCE structure with optional fields
type RoutingInfoForSMRes struct {
	IMSI                 IMSI
	LocationInfoWithLMSI LocationInfoWithLMSI `asn1:"tag:0"`
	//ExtensionContainer   *ExtensionContainer  `asn1:"tag:4,optional"`
	//IPSMGWGuidance       *IPSMGWGuidance      `asn1:"tag:5,optional"`
}

// LocationInfoWithLMSI defines the ASN.1 SEQUENCE structure with optional fields
type LocationInfoWithLMSI struct {
	NetworkNodeNumber ISDNAddressString `asn1:"tag:1"`
	//LMSI                              *LMSI                       `asn1:"optional"`
	//ExtensionContainer                *ExtensionContainer         `asn1:"optional"`
	//GPRSNodeIndicator                 *asn1.RawValue              `asn1:"tag:5,optional"` // NULL
	//AdditionalNumber                  *AdditionalNumber           `asn1:"tag:6,optional"`
	//NetworkNodeDiameterAddress        *NetworkNodeDiameterAddress `asn1:"tag:7,optional"`
	//AdditionalNetworkNodeDiameterAddr *NetworkNodeDiameterAddress `asn1:"tag:8,optional"`
	//ThirdNumber                       *AdditionalNumber           `asn1:"tag:9,optional"`
	//ThirdNetworkNodeDiameterAddr      *NetworkNodeDiameterAddress `asn1:"tag:10,optional"`
	//IMSNodeIndicator                  *asn1.RawValue              `asn1:"tag:11,optional"` // NULL
}

func (riSmReq *RoutingInfoForSMArg) GetMsisdnString() string {
	_, _, _, Digits := DecodeAddressString(riSmReq.MSISDN)
	return utils.DecodeTBCDDigits(Digits)
}

func (riSmReq *RoutingInfoForSMArg) GetServiceCenterAddressString() string {
	_, _, _, Digits := DecodeAddressString(riSmReq.ServiceCentreAddress)
	return utils.DecodeTBCDDigits(Digits)
}

func (riSmRes *RoutingInfoForSMRes) GetImsiString() string {
	return utils.DecodeTBCDDigits(riSmRes.IMSI)
}

func (riSmRes *RoutingInfoForSMRes) GetNetworkNodeNumberString() string {
	_, _, _, Digits := DecodeAddressString(riSmRes.LocationInfoWithLMSI.NetworkNodeNumber)
	return utils.DecodeTBCDDigits(Digits)
}

// MTForwardSMArg defines the ASN.1 SEQUENCE structure with optional fields and Extensibility
type MTForwardSMArg struct {
	SMRPDA asn1.RawValue

	SMRPOA asn1.RawValue

	//
	SmRPUI SignalInfo // sm-RP-UI

	// In case present, the value must be: asn1.RawValue{Class: asn1.ClassUniversal, Tag: asn1.TagNull, Bytes: []byte{}}
	MoreMessagesToSend asn1.RawValue `asn1:"optional"` // moreMessagesToSend - optional // does not carry data, will hold the NULL if present

	//ExtensionContainer     *ExtensionContainer    `asn1:"optional"`       // extensionContainer - optional
	//SmDeliveryTimer        *SM_DeliveryTimerValue `asn1:"optional"`       // smDeliveryTimer - optional
	//SmDeliveryStartTime    *Time                  `asn1:"optional"`       // smDeliveryStartTime - optional
	//SmsOverIPOnlyIndicator *asn1.RawValue         `asn1:"tag:0,optional"` // smsOverIP-OnlyIndicator - optional, Context Specific
}

// MTForwardSMRes defines the ASN.1 SEQUENCE structure with optional fields
type MTForwardSMRes struct {
	//SmRPUI            *SignalInfo        `asn1:"optional,tag:0"`
	//ExtensionContainer *ExtensionContainer `asn1:"optional,tag:1"`
}

// MOForwardSMArg defines the ASN.1 SEQUENCE structure with optional fields
type MOForwardSMArg struct {
	SMRPDA asn1.RawValue

	SMRPOA asn1.RawValue

	//
	SmRPUI SignalInfo

	//ExtensionContainer *ExtensionContainer `asn1:"optional,tag:3"`
	//IMSI               *IMSI               `asn1:"optional,tag:4"`
}

// MOForwardSMRes defines the ASN.1 SEQUENCE structure with optional fields
type MOForwardSMRes struct {
	//SmRPUI             *SignalInfo         `asn1:"optional,tag:0"`
	//ExtensionContainer *ExtensionContainer `asn1:"optional,tag:1"`
}

type SMRPDA struct {
	// SMRPDA defines the ASN.1 CHOICE structure for SM-RP-DA
	IMSI                   IMSI          `asn1:"tag:0,optional"` // [0] IMSI
	LMSI                   LMSI          `asn1:"tag:1,optional"` // [1] LMSI
	ServiceCentreAddressDA AddressString `asn1:"tag:4,optional"` // [4] AddressString
	NoSMRPDA               asn1.RawValue `asn1:"tag:5,optional"` // [5] NULL
}

func (smRpDa *SMRPDA) GetImsiString() string {
	return utils.DecodeTBCDDigits(smRpDa.IMSI)
}

func (smRpDa *SMRPDA) GetServiceCentreAddressDAString() string {
	_, _, _, Digits := DecodeAddressString(smRpDa.ServiceCentreAddressDA)
	return utils.DecodeTBCDDigits(Digits)
}

type SMRPOA struct {
	// SMRPOA defines the ASN.1 CHOICE structure for SM-RP-OA
	MSISDN                 ISDNAddressString `asn1:"tag:2,optional"` // [2] ISDN-AddressString
	ServiceCentreAddressOA AddressString     `asn1:"tag:4,optional"` // [4] AddressString
	NoSMRPOA               asn1.RawValue     `asn1:"tag:5,optional"` // [5] NULL
}

func (smRpOa *SMRPOA) GetMsisdnString() string {
	_, _, _, Digits := DecodeAddressString(smRpOa.MSISDN)
	return utils.DecodeTBCDDigits(Digits)
}

func (smRpOa *SMRPOA) GetServiceCentreAddressOAString() string {
	_, _, _, Digits := DecodeAddressString(smRpOa.ServiceCentreAddressOA)
	return utils.DecodeTBCDDigits(Digits)
}
