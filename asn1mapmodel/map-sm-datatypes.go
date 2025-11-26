package asn1mapmodel

import (
	"encoding/asn1"

	"github.com/gomaja/go-gsmmap/utils"
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

func (riSmReq *RoutingInfoForSMArg) GetMsisdnString() (string, error) {
	_, _, _, Digits := DecodeAddressString(riSmReq.MSISDN)
	return utils.DecodeTBCDDigits(Digits)
}

func (riSmReq *RoutingInfoForSMArg) GetServiceCenterAddressString() (string, error) {
	_, _, _, Digits := DecodeAddressString(riSmReq.ServiceCentreAddress)
	return utils.DecodeTBCDDigits(Digits)
}

func (riSmRes *RoutingInfoForSMRes) GetImsiString() (string, error) {
	return utils.DecodeTBCDDigits(riSmRes.IMSI)
}

func (riSmRes *RoutingInfoForSMRes) GetNetworkNodeNumberString() (string, error) {
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

func (smRpDa *SMRPDA) GetImsiString() (string, error) {
	return utils.DecodeTBCDDigits(smRpDa.IMSI)
}

func (smRpDa *SMRPDA) GetServiceCentreAddressDAString() (string, error) {
	_, _, _, Digits := DecodeAddressString(smRpDa.ServiceCentreAddressDA)
	return utils.DecodeTBCDDigits(Digits)
}

type SMRPOA struct {
	// SMRPOA defines the ASN.1 CHOICE structure for SM-RP-OA
	MSISDN                 ISDNAddressString `asn1:"tag:2,optional"` // [2] ISDN-AddressString
	ServiceCentreAddressOA AddressString     `asn1:"tag:4,optional"` // [4] AddressString
	NoSMRPOA               asn1.RawValue     `asn1:"tag:5,optional"` // [5] NULL
}

func (smRpOa *SMRPOA) GetMsisdnString() (string, error) {
	_, _, _, Digits := DecodeAddressString(smRpOa.MSISDN)
	return utils.DecodeTBCDDigits(Digits)
}

func (smRpOa *SMRPOA) GetServiceCentreAddressOAString() (string, error) {
	_, _, _, Digits := DecodeAddressString(smRpOa.ServiceCentreAddressOA)
	return utils.DecodeTBCDDigits(Digits)
}
