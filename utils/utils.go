package utils

import (
	"encoding/hex"
)

// EncodeTBCDDigits creates a TBCD-encoded digits according to "ETSI TS 129 002 V15.5.0 (2019-07)" page 459
// filler string should be always "f"
func EncodeTBCDDigits(s string) ([]byte, error) {
	var raw []byte
	var err error
	if len(s)%2 == 0 {
		raw, err = hex.DecodeString(s)
	} else {
		raw, err = hex.DecodeString(s + "f") // If odd, pad with 'f'
	}
	if err != nil {
		return nil, err
	}

	// Each pair of digits is swapped, e.g., "1234" -> 0x21, 0x43
	return swap(raw), nil
}

func DecodeTBCDDigits(raw []byte) string {
	s := hex.EncodeToString(swap(raw))

	// Check if the last digit was a filler 'f'
	if s[len(s)-1] == 'f' {
		s = s[:len(s)-1]
	}

	return s
}

// Each pair of digits is swapped, e.g., "1234" -> 0x21, 0x43
func swap(raw []byte) []byte {
	swapped := make([]byte, len(raw))
	for n := range raw {
		t := ((raw[n] >> 4) & 0xf) + ((raw[n] << 4) & 0xf0)
		swapped[n] = t
	}
	return swapped
}
