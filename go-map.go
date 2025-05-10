package gsmmap

import (
	"github.com/warthog618/sms/encoding/tpdu"
)

type SriSm struct {
	MSISDN               string
	SmRpPri              bool
	ServiceCentreAddress string
	// TODO: add optional fields if needed
}

type SriSmResp struct {
	IMSI                 string
	LocationInfoWithLMSI LocationInfoWithLMSI
	// TODO: add optional fields if needed
}

type LocationInfoWithLMSI struct {
	NetworkNodeNumber string
	// TODO: add optional fields if needed
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
