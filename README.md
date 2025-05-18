# go-gsmmap

[![Go CI](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml/badge.svg)](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gomaja/go-gsmmap.svg)](https://pkg.go.dev/github.com/gomaja/go-gsmmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/gomaja/go-gsmmap)](https://goreportcard.com/report/github.com/gomaja/go-gsmmap)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, lightweight implementation of the MAP (Mobile Application Part) protocol in Go.

## Overview

The `go-gsmmap` package provides simple and painless handling of MAP in the mobile networks, implemented in the Go Programming Language. It's designed to be straightforward to integrate into existing Go applications that need to interact with mobile networks.

Though MAP is an ASN.1-based protocol, this implementation does not use any ASN.1 files or ASN.1 parsers. The MAP structures in this library are directly defined based on the ASN.1 definition, making them lightweight and efficient.

## Installation

```bash
go get github.com/gomaja/go-gsmmap
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

- [github.com/warthog618/sms](https://github.com/warthog618/sms): For SMS message handling

## Author

Marwan Jadid

## License

This project is licensed under the MIT License—see the [LICENSE](https://github.com/gomaja/go-gsmmap/blob/main/LICENSE) file for details.
