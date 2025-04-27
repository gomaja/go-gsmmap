# go-map

[![Go Reference](https://pkg.go.dev/badge/github.com/gomaja/go-map.svg)](https://pkg.go.dev/github.com/gomaja/go-map)
[![Go Report Card](https://goreportcard.com/badge/github.com/gomaja/go-map)](https://goreportcard.com/report/github.com/gomaja/go-map)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, lightweight implementation of the MAP (Mobile Application Part) protocol in Go for SS7/SIGTRAN networks.

## Overview

The `go-map` package provides simple and painless handling of MAP in the SS7/SIGTRAN protocol stack, implemented in the Go Programming Language. It's designed to be easy to integrate into existing Go applications that need to interact with mobile networks.

Though MAP is an ASN.1-based protocol, this implementation does not use any ASN.1 files or ASN.1 parsers. The MAP structures in this library are directly defined based on the ASN.1 definition, making it lightweight and efficient.

## Installation

```bash
go get github.com/gomaja/go-map
```

## Usage

### Send Routing Info for Short Message (SRI-for-SM)

```go
package main

import (
    "fmt"
    "github.com/gomaja/go-map"
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

    // Use the marshaled data in your TCAP/SCCP stack
    // ...

    // Parse an incoming SRI-for-SM response
    sriSmResp, _, err := gsmmap.ParseSriSmResp(responseData)
    if err != nil {
        fmt.Printf("Error parsing SRI-for-SM response: %v\n", err)
        return
    }

    fmt.Printf("IMSI: %s\n", sriSmResp.IMSI)
    fmt.Printf("MSC: %s\n", sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber)
}
```

### Forward Short Message (MT-ForwardSM)

```go
package main

import (
    "fmt"
    "github.com/fkgi/sms"
    "github.com/gomaja/go-map"
)

func main() {
    // Create a new SMS Deliver message
    deliver := sms.Deliver{
        OA: sms.Address{
            TON:  sms.TONInternational,
            NPI:  sms.NPIISDNTelephone,
            Addr: "123456789",
        },
        UD: sms.UserData{
            DCS:  sms.DCS7BIT,
            Data: []byte("Hello, World!"),
        },
    }

    // Create a Forward Short Message request
    fsm := &gsmmap.Fsm{
        IMSI:                   "123456789012345",
        ServiceCentreAddressOA: "987654321",
        TPDU:                   deliver,
        MoreMessagesToSend:     false,
    }

    // Marshal to ASN.1 DER format
    data, err := fsm.Marshal()
    if err != nil {
        fmt.Printf("Error marshaling MT-ForwardSM: %v\n", err)
        return
    }

    // Use the marshaled data in your TCAP/SCCP stack
    // ...
}
```

## Supported Features

### MAP Messages

| MAP Message                                           | Abbreviation     | Reference                                | Supported |
|-------------------------------------------------------|------------------|------------------------------------------|-----------|
| Invoke Send Routing Info For Short Message             | SRI-for-SM-Req   | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅        |
| Return Result Last Send Routing Info For Short Message | SRI-SM-Resp      | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅        |
| Invoke Forward Short Message                           | MT-ForwardSM     | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅        |
| Return Result Last Invoke Forward Short Message        | ReturnResultLast | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅        |
| Begin otid (concatenated message preparation)          | Begin-otid       | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅        |

## API Documentation

### Main Structures

- **SriSm**: Structure for Send Routing Info for Short Message requests
- **SriSmResp**: Structure for Send Routing Info for Short Message responses
- **LocationInfoWithLMSI**: Structure containing location information
- **Fsm**: Structure for Forward Short Message requests

### Marshal/Parse Functions

- **SriSm.Marshal()**: Converts SriSm to ASN.1 DER format
- **ParseSriSm()**: Parses ASN.1 DER data into SriSm structure
- **SriSmResp.Marshal()**: Converts SriSmResp to ASN.1 DER format
- **ParseSriSmResp()**: Parses ASN.1 DER data into SriSmResp structure
- **Fsm.Marshal()**: Converts Fsm to ASN.1 DER format
- **ParseFsm()**: Parses ASN.1 DER data into Fsm structure

## Dependencies

- [github.com/fkgi/sms](https://github.com/fkgi/sms): For SMS message handling
- [github.com/fkgi/teldata](https://github.com/fkgi/teldata): For telecom data handling (indirect)

## Author

Marwan Jadid

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/gomaja/go-map/blob/main/LICENSE) file for details.
