package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/fkgi/sms"
	"github.com/gomaja/go-gsmmap/asn1mapmodel"
)

// ParseSriSm take a complete bytes IE
func ParseSriSm(dataIE []byte) (*SriSm, []byte, error) {
	var routingInfo asn1mapmodel.RoutingInfoForSMArg

	rest, err := asn1.Unmarshal(dataIE, &routingInfo)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	var sriSm SriSm
	sriSm.MSISDN = routingInfo.GetMsisdnString()
	sriSm.SmRpPri = routingInfo.SmRpPri
	sriSm.ServiceCentreAddress = routingInfo.GetServiceCenterAddressString()

	return &sriSm, rest, nil
}

func ParseSriSmResp(dataIE []byte) (*SriSmResp, []byte, error) {
	var routingInfoResp asn1mapmodel.RoutingInfoForSMRes

	rest, err := asn1.Unmarshal(dataIE, &routingInfoResp)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateInvokeSriForSm: %v", err)
	}

	var sriSmResp SriSmResp
	sriSmResp.IMSI = routingInfoResp.GetImsiString()
	sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber = routingInfoResp.GetNetworkNodeNumberString()

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
	smRpDaBytes, _ := asn1.Marshal(smRpDaByteString)

	rest, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpDa: %v", err)
	}

	var smRpOa asn1mapmodel.SMRPOA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpOaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: mtFsmArg.SMRPOA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpOaBytes, _ := asn1.Marshal(smRpOaByteString)
	rest, err = asn1.Unmarshal(smRpOaBytes, &smRpOa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpOa: %v", err)
	}

	var mtFsm MtFsm
	mtFsm.IMSI = smRpDa.GetImsiString()
	mtFsm.ServiceCentreAddressOA = smRpOa.GetServiceCentreAddressOAString()
	mtFsm.TPDU, _ = sms.UnmarshalDeliver(mtFsmArg.SmRPUI)

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
	smRpDaBytes, _ := asn1.Marshal(smRpDaByteString)

	rest, err = asn1.Unmarshal(smRpDaBytes, &smRpDa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpDa: %v", err)
	}

	var smRpOa asn1mapmodel.SMRPOA
	// encapsulating the input byte to the proper one that can be understood by "encoding/binary"
	smRpOaByteString := asn1.RawValue{Tag: asn1.TagSequence, IsCompound: true, Bytes: moFsmArg.SMRPOA.FullBytes} // Tag = 16 with Constructor = 0x30
	smRpOaBytes, _ := asn1.Marshal(smRpOaByteString)
	rest, err = asn1.Unmarshal(smRpOaBytes, &smRpOa)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateMoForwardSM smRpOa: %v", err)
	}

	var moFsm MoFsm
	moFsm.ServiceCentreAddressDA = smRpDa.GetServiceCentreAddressDAString()
	moFsm.MSISDN = smRpOa.GetMsisdnString()

	moFsm.TPDU, _ = sms.UnmarshalSubmit(moFsmArg.SmRPUI)

	return &moFsm, rest, nil
}
