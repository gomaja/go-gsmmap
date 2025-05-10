# go-gsmmap

[![Go CI](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml/badge.svg)](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gomaja/go-gsmmap.svg)](https://pkg.go.dev/github.com/gomaja/go-gsmmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/gomaja/go-gsmmap)](https://goreportcard.com/report/github.com/gomaja/go-gsmmap)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, lightweight implementation of the MAP (Mobile Application Part) protocol in Go for SS7/SIGTRAN networks.

## Overview

The `go-gsmmap` package provides simple and painless handling of MAP in the SS7/SIGTRAN protocol stack, implemented in the Go Programming Language. It's designed to be easy to integrate into existing Go applications that need to interact with mobile networks.

Though MAP is an ASN.1-based protocol, this implementation does not use any ASN.1 files or ASN.1 parsers. The MAP structures in this library are directly defined based on the ASN.1 definition, making it lightweight and efficient.

## Installation

```bash
go get github.com/gomaja/go-gsmmap
```

## Usage

### Send Routing Info for Short Message (SRI-for-SM)

```go
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
```

### Mt Forward Short Message (MT-ForwardSM)

```go
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
```

## Supported Features

### MAP Messages

| MAP Message                                            | Abbreviation     | Reference                                | Supported |
|--------------------------------------------------------|------------------|------------------------------------------|-----------|
| Invoke Send Routing Info For Short Message             | SRI-for-SM-Req   | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Return Result Last Send Routing Info For Short Message | SRI-SM-Resp      | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke Mt Forward Short Message                        | MT-ForwardSM     | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke MO Forward Short Message                        | MO-ForwardSM     | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Return Result Last Invoke Forward Short Message        | ReturnResultLast | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Begin otid (concatenated message preparation)          | Begin-otid       | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |

## API Documentation

### Main Structures

- **SriSm**: Structure for Send Routing Info for Short Message requests
- **SriSmResp**: Structure for Send Routing Info for Short Message responses
- **LocationInfoWithLMSI**: Structure containing location information
- **MtFsm**: Structure for Forward Short Message requests
- **MoFsm**: Structure for Forward Short Message requests

### Marshal/Parse Functions

- **SriSm.Marshal()**: Converts SriSm to ASN.1 DER format
- **ParseSriSm()**: Parses ASN.1 DER data into SriSm structure
- **SriSmResp.Marshal()**: Converts SriSmResp to ASN.1 DER format
- **ParseSriSmResp()**: Parses ASN.1 DER data into SriSmResp structure
- **MtFsm.Marshal()**: Converts MtFsm to ASN.1 DER format
- **ParseMtFsm()**: Parses ASN.1 DER data into MtFsm structure
- **MoFsm.Marshal()**: Converts MoFsm to ASN.1 DER format
- **ParseMoFsm()**: Parses ASN.1 DER data into MoFsm structure

## Dependencies

- [github.com/fkgi/sms](https://github.com/fkgi/sms): For SMS message handling
- [github.com/fkgi/teldata](https://github.com/fkgi/teldata): For telecom data handling (indirect)

## Author

Marwan Jadid

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/gomaja/go-gsmmap/blob/main/LICENSE) file for details.
