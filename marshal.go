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
