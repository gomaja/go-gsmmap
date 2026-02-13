package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/gomaja/go-asn1utils"
	"github.com/gomaja/go-gsmmap/asn1mapmodel"
	"github.com/gomaja/go-gsmmap/utils"
	"github.com/warthog618/sms"
)

const errFailedToDecodeIMSI = "failed to decode IMSI: %w"

// ParseSriSm takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseSriSm(dataIE []byte) (*SriSm, []byte, error) {

	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseSriSmDER(derBytes)
}

// ParseSriSmDER takes a complete bytes IE with DER ASN1 encoding
func ParseSriSmDER(dataIE []byte) (*SriSm, []byte, error) {
	var routingInfo asn1mapmodel.RoutingInfoForSMArg

	rest, err := asn1.Unmarshal(dataIE, &routingInfo)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	var sriSm SriSm
	var msisdnErr, scaErr error

	sriSm.MSISDN, msisdnErr = routingInfo.GetMsisdnString()
	if msisdnErr != nil {
		return nil, nil, fmt.Errorf("failed to decode MSISDN: %w", msisdnErr)
	}

	sriSm.SmRpPri = routingInfo.SmRpPri
	sriSm.ServiceCentreAddress, scaErr = routingInfo.GetServiceCenterAddressString()
	if scaErr != nil {
		return nil, nil, fmt.Errorf("failed to decode ServiceCentreAddress: %w", scaErr)
	}

	return &sriSm, rest, nil
}

func ParseSriSmResp(dataIE []byte) (*SriSmResp, []byte, error) {
	var routingInfoResp asn1mapmodel.RoutingInfoForSMRes

	rest, err := asn1.Unmarshal(dataIE, &routingInfoResp)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	var sriSmResp SriSmResp
	var imsiErr, nnnErr error

	sriSmResp.IMSI, imsiErr = routingInfoResp.GetImsiString()
	if imsiErr != nil {
		return nil, nil, fmt.Errorf(errFailedToDecodeIMSI, imsiErr)
	}

	sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber, nnnErr = routingInfoResp.GetNetworkNodeNumberString()
	if nnnErr != nil {
		return nil, nil, fmt.Errorf("failed to decode NetworkNodeNumber: %w", nnnErr)
	}

	return &sriSmResp, rest, nil
}

// ParseMtFsm takes a complete bytes IE
func ParseMtFsm(dataIE []byte) (*MtFsm, []byte, error) {
	var mtFsmArg asn1mapmodel.MTForwardSMArg

	_, err := asn1.Unmarshal(dataIE, &mtFsmArg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateForwardSM: %v", err)
	}

	var smRpDa asn1mapmodel.SMRPDA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpDaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: mtFsmArg.SMRPDA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpDaBytes, err := asn1.Marshal(smRpDaByteString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal smRpDaByteString: %w", err)
	}

	_, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpDa: %v", err)
	}

	var smRpOa asn1mapmodel.SMRPOA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpOaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: mtFsmArg.SMRPOA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpOaBytes, err := asn1.Marshal(smRpOaByteString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal smRpOaByteString: %w", err)
	}

	rest, err := asn1.Unmarshal(smRpOaBytes, &smRpOa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpOa: %v", err)
	}

	var mtFsm MtFsm
	var imsiErr, scaErr error

	mtFsm.IMSI, imsiErr = smRpDa.GetImsiString()
	if imsiErr != nil {
		return nil, nil, fmt.Errorf(errFailedToDecodeIMSI, imsiErr)
	}

	mtFsm.ServiceCentreAddressOA, scaErr = smRpOa.GetServiceCentreAddressOAString()
	if scaErr != nil {
		return nil, nil, fmt.Errorf("failed to decode ServiceCentreAddressOA: %w", scaErr)
	}

	tpduDeliver, tpduErr := sms.Unmarshal(mtFsmArg.SmRPUI, sms.AsMT)
	if tpduErr != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU: %w", tpduErr)
	}

	if tpduDeliver == nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU, nil returned: %w", tpduErr)
	}

	mtFsm.TPDU = *tpduDeliver

	if mtFsmArg.MoreMessagesToSend.Tag == asn1.TagNull {
		mtFsm.MoreMessagesToSend = true
	}

	return &mtFsm, rest, nil
}

// ParseMoFsm takes a complete bytes IE
func ParseMoFsm(dataIE []byte) (*MoFsm, []byte, error) {
	var moFsmArg asn1mapmodel.MOForwardSMArg

	_, err := asn1.Unmarshal(dataIE, &moFsmArg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM: %v", err)
	}

	var smRpDa asn1mapmodel.SMRPDA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpDaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: moFsmArg.SMRPDA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpDaBytes, err := asn1.Marshal(smRpDaByteString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal smRpDaByteString: %w", err)
	}

	_, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpDa: %v", err)
	}

	var smRpOa asn1mapmodel.SMRPOA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpOaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: moFsmArg.SMRPOA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpOaBytes, err := asn1.Marshal(smRpOaByteString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal smRpOaByteString: %w", err)
	}

	rest, err := asn1.Unmarshal(smRpOaBytes, &smRpOa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpOa: %v", err)
	}

	var moFsm MoFsm
	var scaErr, msisdnErr error

	moFsm.ServiceCentreAddressDA, scaErr = smRpDa.GetServiceCentreAddressDAString()
	if scaErr != nil {
		return nil, nil, fmt.Errorf("failed to decode ServiceCentreAddressDA: %w", scaErr)
	}

	moFsm.MSISDN, msisdnErr = smRpOa.GetMsisdnString()
	if msisdnErr != nil {
		return nil, nil, fmt.Errorf("failed to decode MSISDN: %w", msisdnErr)
	}

	tpduSubmit, tpduErr := sms.Unmarshal(moFsmArg.SmRPUI, sms.AsMO)
	if tpduErr != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU: %w", tpduErr)
	}

	if tpduSubmit == nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU, nil returned: %w", tpduErr)
	}
	moFsm.TPDU = *tpduSubmit

	return &moFsm, rest, nil
}

// ParseUpdateLocation takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseUpdateLocation(dataIE []byte) (*UpdateLocation, []byte, error) {
	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseUpdateLocationDER(derBytes)
}

// ParseUpdateLocationDER takes a complete bytes IE with DER ASN1 encoding
func ParseUpdateLocationDER(dataIE []byte) (*UpdateLocation, []byte, error) {
	var updLocArg asn1mapmodel.UpdateLocationArg

	rest, err := asn1.Unmarshal(dataIE, &updLocArg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 UpdateLocationArg: %v", err)
	}

	var updLoc UpdateLocation
	var imsiErr, mscErr, vlrErr error

	updLoc.IMSI, imsiErr = updLocArg.GetImsiString()
	if imsiErr != nil {
		return nil, nil, fmt.Errorf(errFailedToDecodeIMSI, imsiErr)
	}

	updLoc.MSCNumber, mscErr = updLocArg.GetMSCNumberString()
	if mscErr != nil {
		return nil, nil, fmt.Errorf("failed to decode MSCNumber: %w", mscErr)
	}

	updLoc.VLRNumber, vlrErr = updLocArg.GetVLRNumberString()
	if vlrErr != nil {
		return nil, nil, fmt.Errorf("failed to decode VLRNumber: %w", vlrErr)
	}

	// Parse optional VlrCapability if present
	if updLocArg.VlrCapability.SupportedCamelPhases.BitLength > 0 ||
		updLocArg.VlrCapability.SupportedLCSCapabilitySets.BitLength > 0 {
		updLoc.VlrCapability = convertAsn1ToVlrCapability(&updLocArg.VlrCapability)
	}

	return &updLoc, rest, nil
}

func convertAsn1ToVlrCapability(asn1VlrCap *asn1mapmodel.VlrCapability) *VlrCapability {
	vlrCap := &VlrCapability{}

	// Convert SupportedCamelPhases from BitString
	if asn1VlrCap.SupportedCamelPhases.BitLength > 0 {
		vlrCap.SupportedCamelPhases = convertAsn1ToSupportedCamelPhases(asn1VlrCap.SupportedCamelPhases)
	}

	// Convert SupportedLCSCapabilitySets from BitString
	if asn1VlrCap.SupportedLCSCapabilitySets.BitLength > 0 {
		vlrCap.SupportedLCSCapabilitySets = convertAsn1ToSupportedLCSCapabilitySets(asn1VlrCap.SupportedLCSCapabilitySets)
	}

	return vlrCap
}

func convertAsn1ToSupportedCamelPhases(bits asn1.BitString) *SupportedCamelPhases {
	camelPhases := &SupportedCamelPhases{}

	if bits.BitLength > asn1mapmodel.CamelPhase1 {
		camelPhases.Phase1 = bits.At(asn1mapmodel.CamelPhase1) == 1
	}
	if bits.BitLength > asn1mapmodel.CamelPhase2 {
		camelPhases.Phase2 = bits.At(asn1mapmodel.CamelPhase2) == 1
	}
	if bits.BitLength > asn1mapmodel.CamelPhase3 {
		camelPhases.Phase3 = bits.At(asn1mapmodel.CamelPhase3) == 1
	}
	if bits.BitLength > asn1mapmodel.CamelPhase4 {
		camelPhases.Phase4 = bits.At(asn1mapmodel.CamelPhase4) == 1
	}

	return camelPhases
}

func convertAsn1ToSupportedLCSCapabilitySets(bits asn1.BitString) *SupportedLCSCapabilitySets {
	lcsCaps := &SupportedLCSCapabilitySets{}

	if bits.BitLength > asn1mapmodel.LCSCapabilitySet1 {
		lcsCaps.LcsCapabilitySet1 = bits.At(asn1mapmodel.LCSCapabilitySet1) == 1
	}
	if bits.BitLength > asn1mapmodel.LCSCapabilitySet2 {
		lcsCaps.LcsCapabilitySet2 = bits.At(asn1mapmodel.LCSCapabilitySet2) == 1
	}
	if bits.BitLength > asn1mapmodel.LCSCapabilitySet3 {
		lcsCaps.LcsCapabilitySet3 = bits.At(asn1mapmodel.LCSCapabilitySet3) == 1
	}
	if bits.BitLength > asn1mapmodel.LCSCapabilitySet4 {
		lcsCaps.LcsCapabilitySet4 = bits.At(asn1mapmodel.LCSCapabilitySet4) == 1
	}
	if bits.BitLength > asn1mapmodel.LCSCapabilitySet5 {
		lcsCaps.LcsCapabilitySet5 = bits.At(asn1mapmodel.LCSCapabilitySet5) == 1
	}

	return lcsCaps
}

// ParseUpdateGprsLocation takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseUpdateGprsLocation(dataIE []byte) (*UpdateGprsLocation, []byte, error) {
	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseUpdateGprsLocationDER(derBytes)
}

// ParseUpdateGprsLocationDER takes a complete bytes IE with DER ASN1 encoding
func ParseUpdateGprsLocationDER(dataIE []byte) (*UpdateGprsLocation, []byte, error) {
	var updGprsLocArg asn1mapmodel.UpdateGprsLocationArg

	rest, err := asn1.Unmarshal(dataIE, &updGprsLocArg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 UpdateGprsLocationArg: %v", err)
	}

	var updGprsLoc UpdateGprsLocation
	var imsiErr, sgsnNumErr, sgsnAddrErr error

	updGprsLoc.IMSI, imsiErr = updGprsLocArg.GetImsiString()
	if imsiErr != nil {
		return nil, nil, fmt.Errorf(errFailedToDecodeIMSI, imsiErr)
	}

	updGprsLoc.SGSNNumber, sgsnNumErr = updGprsLocArg.GetSGSNNumberString()
	if sgsnNumErr != nil {
		return nil, nil, fmt.Errorf("failed to decode SGSNNumber: %w", sgsnNumErr)
	}

	updGprsLoc.SGSNAddress, sgsnAddrErr = updGprsLocArg.GetSGSNAddressString()
	if sgsnAddrErr != nil {
		return nil, nil, fmt.Errorf("failed to decode SGSNAddress: %w", sgsnAddrErr)
	}

	// Parse optional SGSNCapability if present
	if updGprsLocArg.SGSNCapability.GprsEnhancementsSupportIndicator.Tag != 0 ||
		updGprsLocArg.SGSNCapability.SupportedLCSCapabilitySets.BitLength > 0 {
		updGprsLoc.SGSNCapability = convertAsn1ToSGSNCapability(&updGprsLocArg.SGSNCapability)
	}

	return &updGprsLoc, rest, nil
}

func convertAsn1ToSGSNCapability(asn1SGSNCap *asn1mapmodel.SGSNCapability) *SGSNCapability {
	sgsnCap := &SGSNCapability{}

	// Check GprsEnhancementsSupportIndicator (NULL type with tag 3)
	if asn1SGSNCap.GprsEnhancementsSupportIndicator.Tag == 3 {
		sgsnCap.GprsEnhancementsSupportIndicator = true
	}

	// Convert SupportedLCSCapabilitySets from BitString
	if asn1SGSNCap.SupportedLCSCapabilitySets.BitLength > 0 {
		lcsCaps := &SupportedLCSCapabilitySets{}
		bits := asn1SGSNCap.SupportedLCSCapabilitySets

		if bits.BitLength > asn1mapmodel.LCSCapabilitySet1 {
			lcsCaps.LcsCapabilitySet1 = bits.At(asn1mapmodel.LCSCapabilitySet1) == 1
		}
		if bits.BitLength > asn1mapmodel.LCSCapabilitySet2 {
			lcsCaps.LcsCapabilitySet2 = bits.At(asn1mapmodel.LCSCapabilitySet2) == 1
		}
		if bits.BitLength > asn1mapmodel.LCSCapabilitySet3 {
			lcsCaps.LcsCapabilitySet3 = bits.At(asn1mapmodel.LCSCapabilitySet3) == 1
		}
		if bits.BitLength > asn1mapmodel.LCSCapabilitySet4 {
			lcsCaps.LcsCapabilitySet4 = bits.At(asn1mapmodel.LCSCapabilitySet4) == 1
		}
		if bits.BitLength > asn1mapmodel.LCSCapabilitySet5 {
			lcsCaps.LcsCapabilitySet5 = bits.At(asn1mapmodel.LCSCapabilitySet5) == 1
		}

		sgsnCap.SupportedLCSCapabilitySets = lcsCaps
	}

	return sgsnCap
}

// ParseUpdateLocationRes takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseUpdateLocationRes(dataIE []byte) (*UpdateLocationRes, []byte, error) {
	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseUpdateLocationResDER(derBytes)
}

// ParseUpdateLocationResDER takes a complete bytes IE with DER ASN1 encoding
func ParseUpdateLocationResDER(dataIE []byte) (*UpdateLocationRes, []byte, error) {
	var updLocRes asn1mapmodel.UpdateLocationRes

	rest, err := asn1.Unmarshal(dataIE, &updLocRes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 UpdateLocationRes: %v", err)
	}

	var result UpdateLocationRes
	var hlrErr error

	result.HLRNumber, hlrErr = updLocRes.GetHLRNumberString()
	if hlrErr != nil {
		return nil, nil, fmt.Errorf("failed to decode HLRNumber: %w", hlrErr)
	}

	return &result, rest, nil
}

// ParseUpdateGprsLocationRes takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseUpdateGprsLocationRes(dataIE []byte) (*UpdateGprsLocationRes, []byte, error) {
	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseUpdateGprsLocationResDER(derBytes)
}

// ParseUpdateGprsLocationResDER takes a complete bytes IE with DER ASN1 encoding
func ParseUpdateGprsLocationResDER(dataIE []byte) (*UpdateGprsLocationRes, []byte, error) {
	var updGprsLocRes asn1mapmodel.UpdateGprsLocationRes

	rest, err := asn1.Unmarshal(dataIE, &updGprsLocRes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 UpdateGprsLocationRes: %v", err)
	}

	var result UpdateGprsLocationRes
	var hlrErr error

	result.HLRNumber, hlrErr = updGprsLocRes.GetHLRNumberString()
	if hlrErr != nil {
		return nil, nil, fmt.Errorf("failed to decode HLRNumber: %w", hlrErr)
	}

	return &result, rest, nil
}

// ParseAnyTimeInterrogation takes a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseAnyTimeInterrogation(dataIE []byte) (*AnyTimeInterrogation, []byte, error) {
	derBytes, err := asn1utils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseAnyTimeInterrogationDER(derBytes)
}

// ParseAnyTimeInterrogationDER takes a complete bytes IE with DER ASN1 encoding
func ParseAnyTimeInterrogationDER(dataIE []byte) (*AnyTimeInterrogation, []byte, error) {
	var atiArg asn1mapmodel.AnyTimeInterrogationArg

	rest, err := asn1.Unmarshal(dataIE, &atiArg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 AnyTimeInterrogationArg: %v", err)
	}

	var ati AnyTimeInterrogation

	// Extract SubscriberIdentity from RawValue (CHOICE type)
	// Wrap the inner bytes in a SEQUENCE so encoding/asn1 can unmarshal into the struct
	var subId asn1mapmodel.SubscriberIdentity
	subIdSeq := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: atiArg.SubscriberIdentity.Bytes}
	subIdBytes, err := asn1.Marshal(subIdSeq)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal SubscriberIdentity wrapper: %w", err)
	}

	_, err = asn1.Unmarshal(subIdBytes, &subId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode SubscriberIdentity: %w", err)
	}

	// Determine which CHOICE field is populated
	if len(subId.IMSI) > 0 {
		ati.SubscriberIdentity.IMSI, err = subId.GetImsiString()
		if err != nil {
			return nil, nil, fmt.Errorf(errFailedToDecodeIMSI, err)
		}
	} else if len(subId.MSISDN) > 0 {
		ati.SubscriberIdentity.MSISDN, err = subId.GetMsisdnString()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode MSISDN: %w", err)
		}
	}

	// Parse RequestedInfo flags
	ri := atiArg.RequestedInfo
	ati.RequestedInfo.LocationInformation = isRawValuePresent(ri.LocationInformation)
	ati.RequestedInfo.SubscriberState = isRawValuePresent(ri.SubscriberState)
	ati.RequestedInfo.CurrentLocation = isRawValuePresent(ri.CurrentLocation)
	ati.RequestedInfo.MsClassmark = isRawValuePresent(ri.MsClassmark)
	ati.RequestedInfo.IMEI = isRawValuePresent(ri.IMEI)
	ati.RequestedInfo.MnpRequestedInfo = isRawValuePresent(ri.MnpRequestedInfo)
	ati.RequestedInfo.TAdsData = isRawValuePresent(ri.TAdsData)
	ati.RequestedInfo.ServingNodeIndication = isRawValuePresent(ri.ServingNodeIndication)
	ati.RequestedInfo.LocationInformationEPSSupported = isRawValuePresent(ri.LocationInformationEPSSupported)
	ati.RequestedInfo.LocalTimeZoneRequest = isRawValuePresent(ri.LocalTimeZoneRequest)

	// Parse RequestedDomain (ENUMERATED, default:-1 means absent)
	if ri.RequestedDomain != -1 {
		domain := DomainType(ri.RequestedDomain)
		// Per spec: reception of values > 1 shall be mapped to 'cs-Domain'
		if domain > PsDomain {
			domain = CsDomain
		}
		ati.RequestedInfo.RequestedDomain = &domain
	}

	// Parse RequestedNodes (BIT STRING)
	if ri.RequestedNodes.BitLength > 0 {
		ati.RequestedInfo.RequestedNodes = convertAsn1ToRequestedNodes(ri.RequestedNodes)
	}

	// Decode gsmSCF-Address
	_, _, _, scfDigits := asn1mapmodel.DecodeAddressString(atiArg.GsmSCFAddress)
	ati.GsmSCFAddress, err = utils.DecodeTBCDDigits(scfDigits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode GsmSCFAddress: %w", err)
	}

	return &ati, rest, nil
}

func isRawValuePresent(rv asn1.RawValue) bool {
	return len(rv.FullBytes) > 0
}

func convertAsn1ToRequestedNodes(requestedNodes asn1mapmodel.RequestedNodes) *RequestedNodes {
	nodes := &RequestedNodes{}

	if requestedNodes.BitLength > asn1mapmodel.RequestedNodeMME {
		nodes.MME = requestedNodes.At(asn1mapmodel.RequestedNodeMME) == 1
	}
	if requestedNodes.BitLength > asn1mapmodel.RequestedNodeSGSN {
		nodes.SGSN = requestedNodes.At(asn1mapmodel.RequestedNodeSGSN) == 1
	}

	return nodes
}
