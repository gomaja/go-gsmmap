package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/gomaja/go-gsmmap/asn1mapmodel"
	"github.com/gomaja/go-gsmmap/utils"
)

func (sriSm *SriSm) Marshal() ([]byte, error) {
	// Create RoutingInfoForSMArg structure from SriSm
	routingInfo, err := convertSriSmToRoutingInfoForSMArg(sriSm)
	if err != nil {
		return nil, fmt.Errorf("failed to convert SriSm to RoutingInfoForSMArg: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(routingInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 RoutingInfoForSM: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertSriSmToRoutingInfoForSMArg(sriSm *SriSm) (asn1mapmodel.RoutingInfoForSMArg, error) {
	var routingInfo asn1mapmodel.RoutingInfoForSMArg

	// Encode MSISDN from TBCD format
	msisdnTBCDbytes, err := utils.EncodeTBCDDigits(sriSm.MSISDN)
	if err != nil {
		return routingInfo, fmt.Errorf("failed to encode MSISDN: %w", err)
	}

	// Encode ServiceCentreAddress from TBCD format
	serviceCentreAddressTBCDbytes, err := utils.EncodeTBCDDigits(sriSm.ServiceCentreAddress)
	if err != nil {
		return routingInfo, fmt.Errorf("failed to encode ServiceCentreAddress: %w", err)
	}

	// Create an encoded MSISDN address string
	encodedMsisdn := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		msisdnTBCDbytes)

	// Create an encoded ServiceCentreAddress address string
	encodedServiceCentreAddress := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		serviceCentreAddressTBCDbytes)

	// Fill the fields in the return structure
	routingInfo.MSISDN = asn1mapmodel.ISDNAddressString(encodedMsisdn)
	routingInfo.SmRpPri = sriSm.SmRpPri
	routingInfo.ServiceCentreAddress = asn1mapmodel.AddressString(encodedServiceCentreAddress)

	//// Set optional fields if they exist
	//if sriSm.TeleserviceID != nil {
	//	routingInfo.TeleserviceID = sriSm.TeleserviceID
	//}
	//
	//if sriSm.IMSI != "" {
	//	imsiTBCDbytes, err := utils.EncodeTBCDDigits(sriSm.IMSI)
	//	if err != nil {
	//		return routingInfo, fmt.Errorf("failed to encode IMSI: %w", err)
	//	}
	//	routingInfo.IMSI = asn1mapmodel.IMSI(imsiTBCDbytes)
	//}

	return routingInfo, nil
}

func (sriSmResp *SriSmResp) Marshal() ([]byte, error) {
	// Create RoutingInfoForSMRes structure from SriSmResp
	routingInfoResp, err := convertSriSmRespToRoutingInfoForSMRes(sriSmResp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert SriSmResp to RoutingInfoForSMRes: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(routingInfoResp)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 RoutingInfoForSM: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertSriSmRespToRoutingInfoForSMRes(sriSm *SriSmResp) (asn1mapmodel.RoutingInfoForSMRes, error) {
	var routingInfoResp asn1mapmodel.RoutingInfoForSMRes

	// Encode IMSI from TBCD format
	imsiTBCDbytes, err := utils.EncodeTBCDDigits(sriSm.IMSI)
	if err != nil {
		return routingInfoResp, fmt.Errorf("failed to encode IMSI: %w", err)
	}

	// Encode MSC number from TBCD format
	mscTBCDbytes, err := utils.EncodeTBCDDigits(sriSm.LocationInfoWithLMSI.NetworkNodeNumber)
	if err != nil {
		return routingInfoResp, fmt.Errorf("failed to encode NetworkNodeNumber: %w", err)
	}

	// Create an encoded MSC address string to embed in locationInfoWithLMSI
	encodedMsc := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		mscTBCDbytes)

	// Set IMSI in response
	routingInfoResp.IMSI = asn1mapmodel.IMSI(imsiTBCDbytes)

	// Set LocationInfoWithLMSI in response
	routingInfoResp.LocationInfoWithLMSI = asn1mapmodel.LocationInfoWithLMSI{
		NetworkNodeNumber: asn1mapmodel.ISDNAddressString(encodedMsc),
	}

	//// Check if LMSI exists in the input and set it if present
	//if len(sriSm.LocationInfoWithLMSI.LMSI) > 0 {
	//	routingInfoResp.LocationInfoWithLMSI.LMSI = sriSm.LocationInfoWithLMSI.LMSI
	//}

	return routingInfoResp, nil
}

func (mtFsm *MtFsm) Marshal() ([]byte, error) {
	// Create MTForwardSMArg structure from MtFsm
	mtFsmArg, err := convertMtFsmToMTForwardSMArg(mtFsm)
	if err != nil {
		return nil, fmt.Errorf("failed to convert MtFsm to MTForwardSMArg: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(mtFsmArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 MTForwardSM: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertMtFsmToMTForwardSMArg(mtFsm *MtFsm) (asn1mapmodel.MTForwardSMArg, error) {
	var mtFsmArg asn1mapmodel.MTForwardSMArg

	// Encode IMSI
	imsiTBCDbytes, err := utils.EncodeTBCDDigits(mtFsm.IMSI)
	if err != nil {
		return mtFsmArg, fmt.Errorf("failed to encode IMSI: %w", err)
	}

	// Encode ServiceCenterAddressOA
	serviceCentreAddressOATBCDbytes, err := utils.EncodeTBCDDigits(mtFsm.ServiceCentreAddressOA)
	if err != nil {
		return mtFsmArg, fmt.Errorf("failed to encode ServiceCentreAddressOA: %w", err)
	}

	// Prepare serviceCenterAddress
	encodedServiceCentreAddressOA := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		serviceCentreAddressOATBCDbytes)

	// Prepare SMRPDA (SM-RP-DA)
	var smRpDa asn1mapmodel.SMRPDA
	// Note: Check if IMSI should be an AddressString or raw bytes
	// This depends on the definition of the SMRPDA struct
	smRpDa.IMSI = imsiTBCDbytes
	smRpDaBytes, err := stripTagAndLength(smRpDa)
	if err != nil {
		return mtFsmArg, fmt.Errorf("failed to process SMRPDA: %w", err)
	}
	mtFsmArg.SMRPDA.FullBytes = smRpDaBytes

	// Prepare SMRPOA (SM-RP-OA)
	var smRpOa asn1mapmodel.SMRPOA
	smRpOa.ServiceCentreAddressOA = encodedServiceCentreAddressOA
	smRpOaBytes, err := stripTagAndLength(smRpOa)
	if err != nil {
		return mtFsmArg, fmt.Errorf("failed to process SMRPOA: %w", err)
	}
	mtFsmArg.SMRPOA.FullBytes = smRpOaBytes

	// Set SM-RP-UI (Short Message)
	mtFsmArg.SmRPUI, err = mtFsm.TPDU.MarshalBinary()
	if err != nil {
		return mtFsmArg, fmt.Errorf("failed to marshal MtFsm TPDU: %w", err)
	}

	// Set MoreMessagesToSend flag if needed
	if mtFsm.MoreMessagesToSend {
		mtFsmArg.MoreMessagesToSend = asn1.RawValue{
			Class: asn1.ClassUniversal,
			Tag:   asn1.TagNull,
			Bytes: []byte{},
		}
	}

	return mtFsmArg, nil
}

func (moFsm *MoFsm) Marshal() ([]byte, error) {
	// Create MOForwardSMArg structure from MoFsm
	moFsmArg, err := convertMoFsmToMOForwardSMArg(moFsm)
	if err != nil {
		return nil, fmt.Errorf("failed to convert MoFsm to MOForwardSMArg: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(moFsmArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 MOForwardSM: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertMoFsmToMOForwardSMArg(moFsm *MoFsm) (asn1mapmodel.MOForwardSMArg, error) {
	var moFsmArg asn1mapmodel.MOForwardSMArg
	var err error

	// Encode ServiceCenterAddress
	serviceCentreAddressDATBCDbytes, err := utils.EncodeTBCDDigits(moFsm.ServiceCentreAddressDA)
	if err != nil {
		return moFsmArg, fmt.Errorf("failed to encode ServiceCentreAddressDA: %w", err)
	}
	encodedServiceCentreAddressDA := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		serviceCentreAddressDATBCDbytes)

	// Encode MSISDN
	msisdnBytes, err := utils.EncodeTBCDDigits(moFsm.MSISDN)
	if err != nil {
		return moFsmArg, fmt.Errorf("failed to encode MSISDN: %w", err)
	}
	msisdn := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		msisdnBytes)

	// Prepare SMRPDA (SM-RP-DA)
	var smRpDa asn1mapmodel.SMRPDA
	smRpDa.ServiceCentreAddressDA = encodedServiceCentreAddressDA
	smRpDaBytes, err := stripTagAndLength(smRpDa)
	if err != nil {
		return moFsmArg, fmt.Errorf("failed to process SMRPDA: %w", err)
	}
	moFsmArg.SMRPDA.FullBytes = smRpDaBytes

	// Prepare SMRPOA (SM-RP-OA)
	var smRpOa asn1mapmodel.SMRPOA
	smRpOa.MSISDN = msisdn
	smRpOaBytes, err := stripTagAndLength(smRpOa)
	if err != nil {
		return moFsmArg, fmt.Errorf("failed to process SMRPOA: %w", err)
	}
	moFsmArg.SMRPOA.FullBytes = smRpOaBytes

	// Set SM-RP-UI (Short Message)
	moFsmArg.SmRPUI, err = moFsm.TPDU.MarshalBinary()
	if err != nil {
		return moFsmArg, fmt.Errorf("failed to marshal MoFsm TPDU: %w", err)
	}

	return moFsmArg, nil
}

// Helper function to strip ASN.1 tag and length
func stripTagAndLength(value interface{}) ([]byte, error) {
	bytesWithTag, err := asn1.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	var rawValue asn1.RawValue
	_, err = asn1.Unmarshal(bytesWithTag, &rawValue)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return rawValue.Bytes, nil
}

func (updLoc *UpdateLocation) Marshal() ([]byte, error) {
	// Create UpdateLocationArg structure from UpdateLocation
	updLocArg, err := convertUpdateLocationToUpdateLocationArg(updLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UpdateLocation to UpdateLocationArg: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(updLocArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 UpdateLocation: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertUpdateLocationToUpdateLocationArg(updLoc *UpdateLocation) (asn1mapmodel.UpdateLocationArg, error) {
	var updLocArg asn1mapmodel.UpdateLocationArg

	// Encode IMSI from TBCD format
	imsiTBCDbytes, err := utils.EncodeTBCDDigits(updLoc.IMSI)
	if err != nil {
		return updLocArg, fmt.Errorf("failed to encode IMSI: %w", err)
	}

	// Encode MSCNumber from TBCD format
	mscTBCDbytes, err := utils.EncodeTBCDDigits(updLoc.MSCNumber)
	if err != nil {
		return updLocArg, fmt.Errorf("failed to encode MSCNumber: %w", err)
	}

	// Encode VLRNumber from TBCD format
	vlrTBCDbytes, err := utils.EncodeTBCDDigits(updLoc.VLRNumber)
	if err != nil {
		return updLocArg, fmt.Errorf("failed to encode VLRNumber: %w", err)
	}

	// Create encoded MSCNumber address string
	encodedMSCNumber := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		mscTBCDbytes)

	// Create encoded VLRNumber address string
	encodedVLRNumber := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		vlrTBCDbytes)

	// Fill the fields in the return structure
	updLocArg.IMSI = imsiTBCDbytes
	updLocArg.MSCNumber = encodedMSCNumber
	updLocArg.VLRNumber = encodedVLRNumber

	// Set optional VlrCapability if present
	if updLoc.VlrCapability != nil {
		updLocArg.VlrCapability = convertVlrCapabilityToAsn1(updLoc.VlrCapability)
	}

	return updLocArg, nil
}

func convertVlrCapabilityToAsn1(vlrCap *VlrCapability) asn1mapmodel.VlrCapability {
	var asn1VlrCap asn1mapmodel.VlrCapability

	// Convert SupportedCamelPhases to BitString
	// SupportedCamelPhases ::= BIT STRING { phase1(0), phase2(1), phase3(2), phase4(3) } (SIZE 1..16)
	if vlrCap.SupportedCamelPhases != nil {
		camelPhases := vlrCap.SupportedCamelPhases

		// Build the bit string from phase flags, tracking the highest bit position set
		var byteVal byte
		var bitLength int
		if camelPhases.Phase1 {
			byteVal |= 0x80 // bit 0 (MSB)
			bitLength = 1
		}
		if camelPhases.Phase2 {
			byteVal |= 0x40 // bit 1
			bitLength = 2
		}
		if camelPhases.Phase3 {
			byteVal |= 0x20 // bit 2
			bitLength = 3
		}
		if camelPhases.Phase4 {
			byteVal |= 0x10 // bit 3
			bitLength = 4
		}

		// If no phases set but struct exists, use minimum length of 1
		if bitLength == 0 {
			bitLength = 1
		}

		asn1VlrCap.SupportedCamelPhases = asn1.BitString{
			Bytes:     []byte{byteVal},
			BitLength: bitLength,
		}
	}

	// Convert SupportedLCSCapabilitySets to BitString
	// SupportedLCS-CapabilitySets ::= BIT STRING { lcsCapabilitySet1(0), ..., lcsCapabilitySet5(4) } (SIZE 2..16)
	if vlrCap.SupportedLCSCapabilitySets != nil {
		lcsCaps := vlrCap.SupportedLCSCapabilitySets

		// Build the bit string from LCS capability flags, tracking the highest bit position set
		var byteVal byte
		var bitLength int
		if lcsCaps.LcsCapabilitySet1 {
			byteVal |= 0x80 // bit 0 (MSB)
			bitLength = 1
		}
		if lcsCaps.LcsCapabilitySet2 {
			byteVal |= 0x40 // bit 1
			bitLength = 2
		}
		if lcsCaps.LcsCapabilitySet3 {
			byteVal |= 0x20 // bit 2
			bitLength = 3
		}
		if lcsCaps.LcsCapabilitySet4 {
			byteVal |= 0x10 // bit 3
			bitLength = 4
		}
		if lcsCaps.LcsCapabilitySet5 {
			byteVal |= 0x08 // bit 4
			bitLength = 5
		}

		// If no sets set but struct exists, use minimum length of 2 per spec
		if bitLength == 0 {
			bitLength = 2
		}

		asn1VlrCap.SupportedLCSCapabilitySets = asn1.BitString{
			Bytes:     []byte{byteVal},
			BitLength: bitLength,
		}
	}

	return asn1VlrCap
}

func (updGprsLoc *UpdateGprsLocation) Marshal() ([]byte, error) {
	// Create UpdateGprsLocationArg structure from UpdateGprsLocation
	updGprsLocArg, err := convertUpdateGprsLocationToUpdateGprsLocationArg(updGprsLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UpdateGprsLocation to UpdateGprsLocationArg: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(updGprsLocArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 UpdateGprsLocation: %w", err)
	}

	// Return complete Information Element (IE) with tag, length, and value
	return dataIE, nil
}

func convertUpdateGprsLocationToUpdateGprsLocationArg(updGprsLoc *UpdateGprsLocation) (asn1mapmodel.UpdateGprsLocationArg, error) {
	var updGprsLocArg asn1mapmodel.UpdateGprsLocationArg

	// Encode IMSI from TBCD format
	imsiTBCDbytes, err := utils.EncodeTBCDDigits(updGprsLoc.IMSI)
	if err != nil {
		return updGprsLocArg, fmt.Errorf("failed to encode IMSI: %w", err)
	}

	// Encode SGSNNumber from TBCD format
	sgsnNumberTBCDbytes, err := utils.EncodeTBCDDigits(updGprsLoc.SGSNNumber)
	if err != nil {
		return updGprsLocArg, fmt.Errorf("failed to encode SGSNNumber: %w", err)
	}

	// Encode SGSNAddress from IP string
	sgsnAddressBytes, err := utils.BuildGSNAddress(updGprsLoc.SGSNAddress)
	if err != nil {
		return updGprsLocArg, fmt.Errorf("failed to encode SGSNAddress: %w", err)
	}

	// Create encoded SGSNNumber address string
	encodedSGSNNumber := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		sgsnNumberTBCDbytes)

	// Fill the fields in the return structure
	updGprsLocArg.IMSI = imsiTBCDbytes
	updGprsLocArg.SGSNNumber = encodedSGSNNumber
	updGprsLocArg.SGSNAddress = sgsnAddressBytes

	// Set optional SGSNCapability if present
	if updGprsLoc.SGSNCapability != nil {
		updGprsLocArg.SGSNCapability = convertSGSNCapabilityToAsn1(updGprsLoc.SGSNCapability)
	}

	return updGprsLocArg, nil
}

func convertSGSNCapabilityToAsn1(sgsnCap *SGSNCapability) asn1mapmodel.SGSNCapability {
	var asn1SGSNCap asn1mapmodel.SGSNCapability

	// Set GprsEnhancementsSupportIndicator as NULL if true
	if sgsnCap.GprsEnhancementsSupportIndicator {
		asn1SGSNCap.GprsEnhancementsSupportIndicator = asn1.RawValue{
			Class: asn1.ClassContextSpecific,
			Tag:   3,
			Bytes: []byte{},
		}
	}

	// Convert SupportedLCSCapabilitySets to BitString
	if sgsnCap.SupportedLCSCapabilitySets != nil {
		lcsCaps := sgsnCap.SupportedLCSCapabilitySets

		var byteVal byte
		var bitLength int
		if lcsCaps.LcsCapabilitySet1 {
			byteVal |= 0x80
			bitLength = 1
		}
		if lcsCaps.LcsCapabilitySet2 {
			byteVal |= 0x40
			bitLength = 2
		}
		if lcsCaps.LcsCapabilitySet3 {
			byteVal |= 0x20
			bitLength = 3
		}
		if lcsCaps.LcsCapabilitySet4 {
			byteVal |= 0x10
			bitLength = 4
		}
		if lcsCaps.LcsCapabilitySet5 {
			byteVal |= 0x08
			bitLength = 5
		}

		if bitLength == 0 {
			bitLength = 2
		}

		asn1SGSNCap.SupportedLCSCapabilitySets = asn1.BitString{
			Bytes:     []byte{byteVal},
			BitLength: bitLength,
		}
	}

	return asn1SGSNCap
}

func (updLocRes *UpdateLocationRes) Marshal() ([]byte, error) {
	// Create UpdateLocationRes structure from UpdateLocationRes
	updLocResArg, err := convertUpdateLocationResToAsn1(updLocRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UpdateLocationRes to asn1: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(updLocResArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 UpdateLocationRes: %w", err)
	}

	return dataIE, nil
}

func convertUpdateLocationResToAsn1(updLocRes *UpdateLocationRes) (asn1mapmodel.UpdateLocationRes, error) {
	var result asn1mapmodel.UpdateLocationRes

	// Encode HLRNumber from TBCD format
	hlrTBCDbytes, err := utils.EncodeTBCDDigits(updLocRes.HLRNumber)
	if err != nil {
		return result, fmt.Errorf("failed to encode HLRNumber: %w", err)
	}

	// Create encoded HLRNumber address string
	encodedHLRNumber := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		hlrTBCDbytes)

	result.HLRNumber = encodedHLRNumber

	return result, nil
}

func (updGprsLocRes *UpdateGprsLocationRes) Marshal() ([]byte, error) {
	// Convert UpdateGprsLocationRes to ASN.1 structure
	updGprsLocResArg, err := convertUpdateGprsLocationResToAsn1(updGprsLocRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UpdateGprsLocationRes to asn1: %w", err)
	}

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(updGprsLocResArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 UpdateGprsLocationRes: %w", err)
	}

	return dataIE, nil
}

func convertUpdateGprsLocationResToAsn1(updGprsLocRes *UpdateGprsLocationRes) (asn1mapmodel.UpdateGprsLocationRes, error) {
	var result asn1mapmodel.UpdateGprsLocationRes

	// Encode HLRNumber from TBCD format
	hlrTBCDbytes, err := utils.EncodeTBCDDigits(updGprsLocRes.HLRNumber)
	if err != nil {
		return result, fmt.Errorf("failed to encode HLRNumber: %w", err)
	}

	// Create encoded HLRNumber address string
	encodedHLRNumber := asn1mapmodel.EncodeAddressString(
		asn1mapmodel.ExtensionNo,
		asn1mapmodel.AddressNatureInternational,
		asn1mapmodel.NumberingPlanISDN,
		hlrTBCDbytes)

	result.HLRNumber = encodedHLRNumber

	return result, nil
}
