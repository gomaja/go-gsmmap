package gsmmap

import (
	"encoding/asn1"
	"fmt"

	"github.com/fkgi/sms"
	"github.com/gomaja/go-map/asn1mapmodel"
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

// ParseFsm take a complete bytes IE
func ParseFsm(dataIE []byte) (*Fsm, []byte, error) {
	var mtFsm asn1mapmodel.MTForwardSMArg

	rest, err := asn1.Unmarshal(dataIE, &mtFsm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode ASN.1 CreateForwardSM: %v", err)
	}

	var fsm Fsm
	fsm.IMSI = mtFsm.GetImsiString()
	fsm.ServiceCentreAddressOA = mtFsm.GetServiceCentreAddressOAString()
	fsm.TPDU, _ = sms.UnmarshalDeliver(mtFsm.SmRPUI)

	if mtFsm.MoreMessagesToSend.Tag == asn1.TagNull {
		fsm.MoreMessagesToSend = true
	}

	return &fsm, rest, nil
}
