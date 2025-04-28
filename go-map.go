package gsmmap

import "github.com/fkgi/sms"

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

type Fsm struct {
	IMSI                   string
	ServiceCentreAddressOA string
	TPDU                   sms.Deliver

	MoreMessagesToSend bool
}
