package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBCDEncoding(t *testing.T) {
	cases := []struct {
		description string
		str         string
		bytes       []byte
	}{
		{
			"imsi",
			"123451234567890",
			[]byte{0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0},
		},
	}

	for _, c := range cases {
		t.Run("Str2Bytes/"+c.description, func(t *testing.T) {
			swapped, err := EncodeTBCDDigits(c.str)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(swapped, c.bytes); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Bytes2Str/"+c.description, func(t *testing.T) {
			str, err := DecodeTBCDDigits(c.bytes)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(str, c.str); diff != "" {
				t.Error(diff)
			}
		})
	}
}
