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

type MtFsm struct {
	IMSI                   string
	ServiceCentreAddressOA string
	TPDU                   sms.Deliver

	MoreMessagesToSend bool
}

type MoFsm struct {
	ServiceCentreAddressDA string
	MSISDN                 string
	TPDU                   sms.Submit
}
