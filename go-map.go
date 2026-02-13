package gsmmap

import (
	"github.com/warthog618/sms/encoding/tpdu"
)

type SriSm struct {
	MSISDN               string
	SmRpPri              bool
	ServiceCentreAddress string
	// Add remaining optional fields if needed
}

type SriSmResp struct {
	IMSI                 string
	LocationInfoWithLMSI LocationInfoWithLMSI
	// Add remaining optional fields if needed
}

type LocationInfoWithLMSI struct {
	NetworkNodeNumber string
	// Add remaining optional fields if needed
}

type MtFsm struct {
	IMSI                   string
	ServiceCentreAddressOA string
	TPDU                   tpdu.TPDU // tpdu.MT (Deliver type TPDU)

	MoreMessagesToSend bool
}

type MoFsm struct {
	ServiceCentreAddressDA string
	MSISDN                 string
	TPDU                   tpdu.TPDU // tpdu.MO (Submit type TPDU)
}

type UpdateLocation struct {
	IMSI      string
	MSCNumber string
	VLRNumber string

	VlrCapability *VlrCapability
}

type VlrCapability struct {
	SupportedCamelPhases       *SupportedCamelPhases
	SupportedLCSCapabilitySets *SupportedLCSCapabilitySets
}

type SupportedCamelPhases struct {
	Phase1 bool
	Phase2 bool
	Phase3 bool
	Phase4 bool
}

type SupportedLCSCapabilitySets struct {
	LcsCapabilitySet1 bool
	LcsCapabilitySet2 bool
	LcsCapabilitySet3 bool
	LcsCapabilitySet4 bool
	LcsCapabilitySet5 bool
}

type UpdateGprsLocation struct {
	IMSI        string
	SGSNNumber  string
	SGSNAddress string

	SGSNCapability *SGSNCapability
}

type SGSNCapability struct {
	GprsEnhancementsSupportIndicator bool
	SupportedLCSCapabilitySets       *SupportedLCSCapabilitySets
}

type UpdateLocationRes struct {
	HLRNumber string
}

type UpdateGprsLocationRes struct {
	HLRNumber string
}

// DomainType represents the requested domain
// DomainType ::= ENUMERATED { cs-Domain (0), ps-Domain (1), ...}
type DomainType int

const (
	CsDomain DomainType = 0
	PsDomain DomainType = 1
)

// RequestedNodes represents which nodes are requested
// RequestedNodes ::= BIT STRING { mme (0), sgsn (1)} (SIZE (1..8))
type RequestedNodes struct {
	MME  bool
	SGSN bool
}

// SubscriberIdentity represents the subscriber identity CHOICE — set exactly one
type SubscriberIdentity struct {
	IMSI   string // if identifying by IMSI
	MSISDN string // if identifying by MSISDN
}

// RequestedInfo represents the requested information flags for ATI
type RequestedInfo struct {
	LocationInformation             bool
	SubscriberState                 bool
	CurrentLocation                 bool
	RequestedDomain                 *DomainType // nil = absent
	IMEI                            bool
	MsClassmark                     bool
	MnpRequestedInfo                bool
	LocationInformationEPSSupported bool
	TAdsData                        bool
	RequestedNodes                  *RequestedNodes // nil = absent
	ServingNodeIndication           bool
	LocalTimeZoneRequest            bool
}

// AnyTimeInterrogation represents the public API for the ATI request
type AnyTimeInterrogation struct {
	SubscriberIdentity SubscriberIdentity
	RequestedInfo      RequestedInfo
	GsmSCFAddress      string // gsmSCF-Address (required)
}
