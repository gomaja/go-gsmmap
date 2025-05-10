package main

import (
	"fmt"
	"time"

	"github.com/gomaja/go-gsmmap"
	"github.com/warthog618/sms/encoding/tpdu"
)

func main() {
	imsi := "234100080813836"
	serviceCentreAddressOA := "9613488888"

	TPOA := "96170111474"
	protocolID := uint8(0x00) // TP-PID
	dataCoding := uint8(0x00) // TP-DCS (0x00 ⇒ GSM7 default)
	tpduDeliver, _ := tpdu.NewDeliver()
	tpduDeliver.OA = tpdu.NewAddress(tpdu.FromNumber(TPOA))
	tpduDeliver.OA.SetNumberingPlan(tpdu.NpISDN)
	tpduDeliver.OA.SetTypeOfNumber(tpdu.TonInternational)

	tpduDeliver.PID = protocolID
	tpduDeliver.DCS = tpdu.DCS(dataCoding)
	tpduDeliver.SCTS = tpdu.Timestamp{Time: time.Now()}
	tpduDeliver.UD = []byte("Hello! This is a message")
	tpduDeliver.FirstOctet = tpdu.FoMMS // to indicate that no more messages are waiting (simple message)

	// Create a Forward Short Message request
	mtFsm := &gsmmap.MtFsm{
		IMSI:                   imsi,
		ServiceCentreAddressOA: serviceCentreAddressOA,
		TPDU:                   *tpduDeliver,
		MoreMessagesToSend:     false,
	}

	// Marshal to ASN.1 DER format
	data, err := mtFsm.Marshal()
	if err != nil {
		fmt.Printf("Error marshaling MT-ForwardSM: %v\n", err)
		return
	}

	fmt.Printf("%x\n", data)

	// Use the marshaled data in your TCAP/SCCP stack
	// ...
}
