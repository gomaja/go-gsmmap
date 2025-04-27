package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/gomaja/go-map/asn1mapmodel"
	"github.com/gomaja/go-map/utils"
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

func (fsm *Fsm) Marshal() ([]byte, error) {
	// Create MTForwardSMArg
	mtFsm := convertFsmToMTForwardSMArg(fsm)

	// Encode to ASN.1 DER format
	dataIE, err := asn1.Marshal(mtFsm)
	if err != nil {
		return nil, fmt.Errorf("failed to encode ASN.1 CreateForwardSM: %v", err)
	}

	// return complete IE, tag length and value
	return dataIE, nil
}

func convertFsmToMTForwardSMArg(fsm *Fsm) asn1mapmodel.MTForwardSMArg {
	var mtFsm asn1mapmodel.MTForwardSMArg

	imsiTBCDbytes, _ := utils.EncodeTBCDDigits(fsm.IMSI)

	serviceCentreAddressOATBCDbytes, _ := utils.EncodeTBCDDigits(fsm.ServiceCentreAddressOA)

	// prepare serviceCenterAddress
	encodedServiceCentreAddressOA := asn1mapmodel.EncodeAddressString(asn1mapmodel.ExtensionNo, asn1mapmodel.AddressNatureInternational, asn1mapmodel.NumberingPlanISDN, serviceCentreAddressOATBCDbytes)

	// fill the fields
	mtFsm.IMSI = imsiTBCDbytes
	mtFsm.ServiceCentreAddressOA = encodedServiceCentreAddressOA
	mtFsm.SmRPUI = fsm.TPDU.MarshalTP()

	if fsm.MoreMessagesToSend {
		mtFsm.MoreMessagesToSend = asn1.RawValue{Class: asn1.ClassUniversal, Tag: asn1.TagNull, Bytes: []byte{}}
	}

	return mtFsm
}
