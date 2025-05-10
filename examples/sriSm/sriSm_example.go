package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gomaja/go-gsmmap"
)

func main() {
	// Create a new SRI-for-SM request
	sriSm := &gsmmap.SriSm{
		MSISDN:               "123456789",
		SmRpPri:              true,
		ServiceCentreAddress: "987654321",
	}

	// Marshal to ASN.1 DER format
	data, err := sriSm.Marshal()
	if err != nil {
		fmt.Printf("Error marshaling SRI-for-SM: %v\n", err)
		return
	}

	fmt.Printf("SRI-for-SM: %x\n", data)

	// Use the marshaled data in your TCAP/SCCP stack
	// ...

	responseData := "3015040882131068584836f3a0098107917394950862f6"
	// Decode the hex string to bytes. The hex string represents an encoded SRI-for-SM response,
	// which needs to be converted to a byte array for parsing by the ParseSriSmResp function.
	// This step ensures the response data is in the correct format for further processing.
	originalBytes, err := hex.DecodeString(responseData)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	// Parse an incoming SRI-for-SM response
	sriSmResp, _, err := gsmmap.ParseSriSmResp(originalBytes)
	if err != nil {
		fmt.Printf("Error parsing SRI-for-SM response: %v\n", err)
		return
	}

	fmt.Printf("IMSI: %s\n", sriSmResp.IMSI)
	fmt.Printf("MSC: %s\n", sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber)
}
