package utils

import (
	"encoding/hex"
	"fmt"
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
