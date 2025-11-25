package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/gomaja/go-gsmmap/asn1mapmodel"
	tcapUtils "github.com/gomaja/go-tcap/utils"
	"github.com/warthog618/sms"
)

// ParseSriSm take a complete bytes IE with any ASN1 encoding (DER and non-DER)
func ParseSriSm(dataIE []byte) (*SriSm, []byte, error) {

	derBytes, err := tcapUtils.MakeDER(dataIE)
	if err != nil {
		return nil, nil, err
	}

	return ParseSriSmDER(derBytes)
}

// ParseSriSmDER take a complete bytes IE with DER ASN1 encoding
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
		return nil, nil, fmt.Errorf("failed to decode IMSI: %w", imsiErr)
	}

	sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber, nnnErr = routingInfoResp.GetNetworkNodeNumberString()
	if nnnErr != nil {
		return nil, nil, fmt.Errorf("failed to decode NetworkNodeNumber: %w", nnnErr)
	}

	return &sriSmResp, rest, nil
}

// ParseMtFsm take a complete bytes IE
func ParseMtFsm(dataIE []byte) (*MtFsm, []byte, error) {
	var mtFsmArg asn1mapmodel.MTForwardSMArg

	rest, err := asn1.Unmarshal(dataIE, &mtFsmArg)
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

	rest, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
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

	rest, err = asn1.Unmarshal(smRpOaBytes, &smRpOa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpOa: %v", err)
	}

	var mtFsm MtFsm
	var imsiErr, scaErr error

	mtFsm.IMSI, imsiErr = smRpDa.GetImsiString()
	if imsiErr != nil {
		return nil, nil, fmt.Errorf("failed to decode IMSI: %w", imsiErr)
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

// ParseMoFsm take a complete bytes IE
func ParseMoFsm(dataIE []byte) (*MoFsm, []byte, error) {
	var moFsmArg asn1mapmodel.MOForwardSMArg

	rest, err := asn1.Unmarshal(dataIE, &moFsmArg)
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

	rest, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
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

	rest, err = asn1.Unmarshal(smRpOaBytes, &smRpOa)
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
	derBytes, err := tcapUtils.MakeDER(dataIE)
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
		return nil, nil, fmt.Errorf("failed to decode IMSI: %w", imsiErr)
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
		camelPhases := &SupportedCamelPhases{}
		bits := asn1VlrCap.SupportedCamelPhases

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

		vlrCap.SupportedCamelPhases = camelPhases
	}

	// Convert SupportedLCSCapabilitySets from BitString
	if asn1VlrCap.SupportedLCSCapabilitySets.BitLength > 0 {
		lcsCaps := &SupportedLCSCapabilitySets{}
		bits := asn1VlrCap.SupportedLCSCapabilitySets

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

		vlrCap.SupportedLCSCapabilitySets = lcsCaps
	}

	return vlrCap
}
