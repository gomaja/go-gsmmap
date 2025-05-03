package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/gomaja/go-gsmmap/asn1mapmodel"
	"github.com/gomaja/go-gsmmap/utils"
)

func (sriSm *SriSm) Marshal() ([]byte, error) {
	// Create RoutingInfoForSMArg
	routingInfo := convertSriSmToRoutingInfoForSMArg(sriSm)

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(routingInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	// return complete IE, tag length and value
	return dataIE, nil
}

func convertSriSmToRoutingInfoForSMArg(sriSm *SriSm) asn1mapmodel.RoutingInfoForSMArg {
	var routingInfo asn1mapmodel.RoutingInfoForSMArg

	msisdnTBCDbytes, _ := utils.EncodeTBCDDigits(sriSm.MSISDN)

	serviceCentreAddressTBCDbytes, _ := utils.EncodeTBCDDigits(sriSm.ServiceCentreAddress)

	// prepare msisdn
	encodedMsisdn := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, msisdnTBCDbytes)

	// prepare serviceCenterAddress
	encodedServiceCentreAddress := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, serviceCentreAddressTBCDbytes)

	// fill the fields
	routingInfo.MSISDN = asn1mapmodel.ISDNAddressString(encodedMsisdn)
	routingInfo.SmRpPri = sriSm.SmRpPri
	routingInfo.ServiceCentreAddress = asn1mapmodel.AddressString(encodedServiceCentreAddress)

	return routingInfo
}

func (sriSmResp *SriSmResp) Marshal() ([]byte, error) {
	// Create RoutingInfoForSMRes
	routingInfoResp := convertSriSmRespToRoutingInfoForSMRes(sriSmResp)

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(routingInfoResp)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	// return complete IE, tag length and value
	return dataIE, nil
}

func convertSriSmRespToRoutingInfoForSMRes(sriSm *SriSmResp) asn1mapmodel.RoutingInfoForSMRes {
	var routingInfoResp asn1mapmodel.RoutingInfoForSMRes

	imsiTBCDbytes, _ := utils.EncodeTBCDDigits(sriSm.IMSI)

	mscTBCDbytes, _ := utils.EncodeTBCDDigits(sriSm.LocationInfoWithLMSI.NetworkNodeNumber)

	// prepare msc-Number to embed in locationInfoWithLMSI
	encodedMsc := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, mscTBCDbytes)

	// fill the fields
	routingInfoResp.IMSI = asn1mapmodel.IMSI(imsiTBCDbytes)
	routingInfoResp.LocationInfoWithLMSI = asn1mapmodel.LocationInfoWithLMSI{
		NetworkNodeNumber: asn1mapmodel.ISDNAddressString(encodedMsc),
	}

	return routingInfoResp
}

func (mtFsm *MtFsm) Marshal() ([]byte, error) {
	// Create MTForwardSMArg
	mtFsmArg := convertMtFsmToMTForwardSMArg(mtFsm)

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(mtFsmArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 CreateForwardSM: %v", err)
	}

	// return complete IE, tag length and value
	return dataIE, nil
}

func convertMtFsmToMTForwardSMArg(mtFsm *MtFsm) asn1mapmodel.MTForwardSMArg {
	var mtFsmArg asn1mapmodel.MTForwardSMArg

	imsiTBCDbytes, _ := utils.EncodeTBCDDigits(mtFsm.IMSI)

	serviceCentreAddressOATBCDbytes, _ := utils.EncodeTBCDDigits(mtFsm.ServiceCentreAddressOA)

	// prepare serviceCenterAddress
	encodedServiceCentreAddressOA := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, serviceCentreAddressOATBCDbytes)

	// fill the fields
	mtFsmArg.IMSI = imsiTBCDbytes
	mtFsmArg.ServiceCentreAddressOA = encodedServiceCentreAddressOA
	mtFsmArg.SmRPUI = mtFsm.TPDU.MarshalTP()

	if mtFsm.MoreMessagesToSend {
		mtFsmArg.MoreMessagesToSend = asn1.RawValue{Class: asn1.ClassUniversal, Tag: asn1.TagNull, Bytes: []byte{}}
	}

	return mtFsmArg
}

func (moFsm *MoFsm) Marshal() ([]byte, error) {
	// Create MOForwardSMArg
	moFsmArg := convertMoFsmToMOForwardSMArg(moFsm)

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(moFsmArg)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 CreateMoForwardSM: %v", err)
	}

	// return complete IE, tag length and value
	return dataIE, nil
}

func convertMoFsmToMOForwardSMArg(moFsm *MoFsm) asn1mapmodel.MOForwardSMArg {
	var moFsmArg asn1mapmodel.MOForwardSMArg

	serviceCentreAddressDATBCDbytes, _ := utils.EncodeTBCDDigits(moFsm.ServiceCentreAddressDA)

	// prepare serviceCenterAddress
	encodedServiceCentreAddressDA := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, serviceCentreAddressDATBCDbytes)

	msisdnBytes, _ := utils.EncodeTBCDDigits(moFsm.MSISDN)

	// prepare MSISDN
	msisdn := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, msisdnBytes)

	// fill the fields
	moFsmArg.ServiceCentreAddressDA = encodedServiceCentreAddressDA
	moFsmArg.MSISDN = msisdn
	moFsmArg.SmRPUI = moFsm.TPDU.MarshalTP()

	return moFsmArg
}
