package utils

import (
	"encoding/hex"
	"fmt"
	"net"
)

// EncodeTBCDDigits creates a TBCD-encoded byte slice according to "ETSI TS 129 002 V15.5.0 (2019-07)" page 459.
// TBCD (Telephony Binary Coded Decimal) swaps the nibbles of each byte.
// For odd-length strings, a filler 'F' is appended.
// Input must consist of valid hexadecimal digits (0-9, a-f, A-F).
func EncodeTBCDDigits(s string) ([]byte, error) {
	// Check if the input contains only valid hex digits
	for _, r := range s {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
			return nil, fmt.Errorf("invalid character in input: %c", r)
		}
	}

	var hexString string
	if len(s)%2 == 0 {
		hexString = s
	} else {
		hexString = s + "f" // If odd length, pad with 'f'
	}

	raw, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %w", err)
	}

	// Swap nibbles for each byte
	return swapNibbles(raw), nil
}

// DecodeTBCDDigits decodes a TBCD-encoded byte slice into a hexadecimal string.
// It removes trailing 'f' or 'F' padding if present.
func DecodeTBCDDigits(raw []byte) (string, error) {
	if raw == nil {
		return "", fmt.Errorf("input is nil")
	}

	swapped := swapNibbles(raw)
	s := hex.EncodeToString(swapped)

	// Check if the last digit was a filler 'f' or 'F'
	if len(s) > 0 && (s[len(s)-1] == 'f' || s[len(s)-1] == 'F') {
		s = s[:len(s)-1]
	}

	return s, nil
}

// swapNibbles swaps the high and low nibbles in each byte.
// For example, 0x12 becomes 0x21, 0xAB becomes 0xBA.
func swapNibbles(data []byte) []byte {
	swapped := make([]byte, len(data))
	for i, b := range data {
		// Swap the high and low nibbles
		swapped[i] = ((b & 0x0F) << 4) | ((b & 0xF0) >> 4)
	}
	return swapped
}

// BuildGSNAddress encodes an IPv4/IPv6 address into the GSN Address format.
//
// Octets are coded according to TS 3GPP TS 23.003 [17]
//
// Format:
//
//	-------------------------------------------------------
//
// | Address Type (2 bits) | Address Length (6 bits) | ...address bytes...
//
//	-------------------------------------------------------
func BuildGSNAddress(ipStr string) ([]byte, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP: %s", ipStr)
	}

	var addrType uint8
	var addrBytes []byte

	if ipv4 := ip.To4(); ipv4 != nil {
		// IPv4
		addrType = 0
		addrBytes = ipv4
	} else {
		// IPv6
		addrType = 1
		addrBytes = ip.To16()
	}

	addrLen := uint8(len(addrBytes)) // 4 or 16

	// Pack type and length:
	//   type = 2 MSB bits
	//   length = 6 LSB bits
	header := (addrType << 6) | (addrLen & 0x3F)

	// Final result: 1-byte header + raw bytes
	result := make([]byte, 1+len(addrBytes))
	result[0] = header
	copy(result[1:], addrBytes)

	return result, nil
}

// ParseGSNAddress decodes a GSN Address format into an IPv4/IPv6 string.
//
// Octets are coded according to TS 3GPP TS 23.003 [17]
//
// Format:
//
//	-------------------------------------------------------
//
// | Address Type (2 bits) | Address Length (6 bits) | ...address bytes...
//
//	-------------------------------------------------------
func ParseGSNAddress(data []byte) (string, error) {
	if len(data) < 1 {
		return "", fmt.Errorf("GSN address too short")
	}

	header := data[0]
	addrType := (header >> 6) & 0x03
	addrLen := header & 0x3F

	if len(data) < 1+int(addrLen) {
		return "", fmt.Errorf("GSN address data too short: expected %d bytes, got %d", 1+int(addrLen), len(data))
	}

	addrBytes := data[1 : 1+addrLen]

	switch addrType {
	case 0: // IPv4
		if addrLen != 4 {
			return "", fmt.Errorf("invalid IPv4 address length: %d", addrLen)
		}
		ip := net.IP(addrBytes)
		return ip.String(), nil
	case 1: // IPv6
		if addrLen != 16 {
			return "", fmt.Errorf("invalid IPv6 address length: %d", addrLen)
		}
		ip := net.IP(addrBytes)
		return ip.String(), nil
	default:
		return "", fmt.Errorf("unknown address type: %d", addrType)
	}
}
