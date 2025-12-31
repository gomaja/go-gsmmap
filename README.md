# go-gsmmap

[![Go CI](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml/badge.svg)](https://github.com/gomaja/go-gsmmap/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gomaja/go-gsmmap.svg)](https://pkg.go.dev/github.com/gomaja/go-gsmmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/gomaja/go-gsmmap)](https://goreportcard.com/report/github.com/gomaja/go-gsmmap)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/gomaja/go-gsmmap)](https://github.com/gomaja/go-gsmmap)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gomaja_go-gsmmap&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gomaja_go-gsmmap)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=gomaja_go-gsmmap&metric=coverage)](https://sonarcloud.io/summary/new_code?id=gomaja_go-gsmmap)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=gomaja_go-gsmmap&metric=bugs)](https://sonarcloud.io/summary/new_code?id=gomaja_go-gsmmap)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=gomaja_go-gsmmap&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=gomaja_go-gsmmap)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=gomaja_go-gsmmap&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=gomaja_go-gsmmap)

A robust, lightweight implementation of the MAP (Mobile Application Part) protocol in Go.

## Overview

The `go-gsmmap` package provides simple and painless handling of MAP in the mobile networks, implemented in the Go Programming Language. It's designed to be straightforward to integrate into existing Go applications that need to interact with mobile networks.

The GSM-MAP structures in this library are directly defined as go structs with ASN.1 tags, making them lightweight and efficient.

## Installation

```bash
go get github.com/gomaja/go-gsmmap
```

## Supported Features

### MAP Messages

| MAP Message                                            | Abbreviation        | Reference                                | Supported |
|--------------------------------------------------------|---------------------|------------------------------------------|-----------|
| Invoke Send Routing Info For Short Message             | SRI-for-SM-Req      | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Return Result Last Send Routing Info For Short Message | SRI-SM-Resp         | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke Mt Forward Short Message                        | MT-ForwardSM        | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke MO Forward Short Message                        | MO-ForwardSM        | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke Update Location                                 | UpdateLocation      | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Return Result Update Location                          | UpdateLocation-Res  | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Invoke Update GPRS Location                            | UpdateGprsLocation  | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |
| Return Result Update GPRS Location                     | UpdateGprsLoc-Res   | 3GPP TS 29.002 version 15.5.0 Release 15 | ✅         |

## Dependencies

- [github.com/warthog618/sms](https://github.com/warthog618/sms): For SMS message handling

## Author

Marwan Jadid

## License

This project is licensed under the MIT License—see the [LICENSE](https://github.com/gomaja/go-gsmmap/blob/main/LICENSE) file for details.
