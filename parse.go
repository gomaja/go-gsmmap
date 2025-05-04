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

	tpdu, tpduErr := sms.UnmarshalDeliver(mtFsmArg.SmRPUI)
	if tpduErr != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU: %w", tpduErr)
	}
	mtFsm.TPDU = tpdu

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

	tpdu, tpduErr := sms.UnmarshalSubmit(moFsmArg.SmRPUI)
	if tpduErr != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal TPDU: %w", tpduErr)
	}
	moFsm.TPDU = tpdu

	return &moFsm, rest, nil
}
