package gsmmap

import (
	"encoding/hex"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseSriSm(t *testing.T) {
	// Test cases
	tests := []struct {
		name                string
		hexString           string
		expectError         bool
		matchMarshaledBytes bool
	}{
		{
			name:                "Valid SRI SM",
			hexString:           "301380069122608538188101ff8206912260909899",
			expectError:         false,
			matchMarshaledBytes: true,
		},
		{
			name:                "Valid SRI SM - nonDER",
			hexString:           "3019800a915282051447720982f9810101820891328490001015f8",
			expectError:         false,
			matchMarshaledBytes: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Decode hex string to bytes
			originalBytes, err := hex.DecodeString(tc.hexString)
			if err != nil {
				t.Fatalf("Failed to decode hex string: %v", err)
			}

			// Parse bytes to SriSm struct
			sriSm, _, err := ParseSriSm(originalBytes)
			if err != nil {
				t.Fatalf("Failed to parse SriSm: %v", err)
			}

			// Marshal SriSm struct back to bytes
			marshaledBytes, err := sriSm.Marshal()
			if (err != nil) != tc.expectError {
				t.Fatalf("Unexpected error status: got %v, expected error: %v", err, tc.expectError)
			}

			if err == nil && tc.matchMarshaledBytes {
				// Compare original and marshaled bytes
				if diff := cmp.Diff(originalBytes, marshaledBytes); diff != "" {
					t.Errorf("Marshaled bytes don't match original (-original +marshaled):\n%s", diff)
				}
			}
		})
	}
}

func TestParseSriSmResp(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		hexString   string
		expectError bool
	}{
		{
			name:        "Valid SRI SM Response",
			hexString:   "3015040882131068584836f3a0098107917394950862f6",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Decode hex string to bytes
			originalBytes, err := hex.DecodeString(tc.hexString)
			if err != nil {
				t.Fatalf("Failed to decode hex string: %v", err)
			}

			// Parse bytes to SriSmResp struct
			sriSmResp, _, err := ParseSriSmResp(originalBytes)
			if err != nil {
				t.Fatalf("Failed to parse SriSmResp: %v", err)
			}

			// Marshal SriSmResp struct back to bytes
			marshaledBytes, err := sriSmResp.Marshal()
			if (err != nil) != tc.expectError {
				t.Fatalf("Unexpected error status: got %v, expected error: %v", err, tc.expectError)
			}

			if err == nil {
				// Compare original and marshaled bytes
				if diff := cmp.Diff(originalBytes, marshaledBytes); diff != "" {
					t.Errorf("Marshaled bytes don't match original (-original +marshaled):\n%s", diff)
				}
			}
		})
	}
}

func TestParseMtFsm(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		hexString   string
		expectError bool
	}{
		{
			name:        "Valid MT FSM",
			hexString:   "3077800832140080803138f684069169318488880463040b916971101174f40000422182612464805bd2e2b1252d467ff6de6c47efd96eb6a1d056cb0d69b49a10269c098537586e96931965b260d15613da72c29b91261bde72c6a1ad2623d682b5996d58331271375a0d1733eee4bd98ec768bd966b41c0d",
			expectError: false,
		},
		{
			name:        "Valid MT FSM Concatenated (part 1)",
			hexString:   "3081b7800826610011829761f6840891328490000005f704819e4009d047f6dbfe06000042217251400000a00500035f020190e53c0b947fd741e8b0bd0c9abfdb6510bcec26a7dd67d09c5e86cf41693728ffaecb41f2f2393da7cbc3f4f4db0d82cbdfe3f27cee0241d9e5f0bc0c32bfd9ecf71d44479741ecb47b0da2bf41e3771bce2ed3cb203abadc0685dd64d09c1e96d341e4323b6d2fcbd3ee33888e96bfeb6734e8c87edbdf2190bc3c96d7d3f476d94d77d5e70500",
			expectError: false,
		},
		{
			name:        "Valid MT FSM Concatenated (part 2)",
			hexString:   "3042800826610011829761f6840891328490000005f7042c4409d047f6dbfe060000422172514000001d0500035f0202cae8ba5c9e2ecb5de377fb157ea9d1b0d93b1e06",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Decode hex string to bytes
			originalBytes, err := hex.DecodeString(tc.hexString)
			if err != nil {
				t.Fatalf("Failed to decode hex string: %v", err)
			}

			// Parse bytes to MtFsm struct
			mtFsm, _, err := ParseMtFsm(originalBytes)
			if err != nil {
				t.Fatalf("Failed to parse MtFsm: %v", err)
			}

			// Marshal MoFsm struct back to bytes
			marshaledBytes, err := mtFsm.Marshal()
			if (err != nil) != tc.expectError {
				t.Fatalf("Unexpected error status: got %v, expected error: %v", err, tc.expectError)
			}

			if err == nil {
				// Compare original and marshaled bytes
				if diff := cmp.Diff(originalBytes, marshaledBytes); diff != "" {
					t.Errorf("Marshaled bytes don't match original (-original +marshaled):\n%s", diff)
				}
			}
		})
	}
}

func TestParseMoFsm(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		hexString   string
		expectError bool
	}{
		{
			name:        "Valid MO FSM",
			hexString:   "302d84069122609098998206912260539128041b01510a912260716622000011d972180d4a82eee13928cc7ebbcb20",
			expectError: false,
		},
		{
			name:        "Valid MO FSM Concatenated (part 1)",
			hexString:   "3081ab84069122609098998206912260532023048198413f0a9122600650150000a0050003020201a8e8f41c949e83c220f6db7d06b5cbf379f85cd6819a61f93deca6a2d373507a0e0a83d86ff719d42ecfe7e17359076a86e5f7b09b8a4ecf41e939280c62bfdd6750bb3c9f87cf651da81996dfc36e2a3a3d07a5e7a03088fd769f41edf27c1e3e9775a066587e0fbba9e8f41c949e83c220f6db7d06b5cbf379f85cd6819a61f93deca6a2d3",
			expectError: false,
		},
		{
			name:        "Valid MO FSM Concatenated (part 2)",
			hexString:   "303c84069122609098998206912260532023042a41400a912260065015000022050003020202e6a0f41c1406b1dfee33a85d9ecfc3e7b20ed40ccbef6137",
			expectError: false,
		},
		{
			name:        "Invalid Packet for MO FSM 1",
			hexString:   "301380069122608538188101ff8206912260909899",
			expectError: true,
		},
		{
			name:        "Invalid Packet for MO FSM 2",
			hexString:   "3081b7800826610011829761f6840891328490000005f704819e4009d047f6dbfe06000042217251400000a00500035f020190e53c0b947fd741e8b0bd0c9abfdb6510bcec26a7dd67d09c5e86cf41693728ffaecb41f2f2393da7cbc3f4f4db0d82cbdfe3f27cee0241d9e5f0bc0c32bfd9ecf71d44479741ecb47b0da2bf41e3771bce2ed3cb203abadc0685dd64d09c1e96d341e4323b6d2fcbd3ee33888e96bfeb6734e8c87edbdf2190bc3c96d7d3f476d94d77d5e70500",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Decode hex string to bytes
			originalBytes, err := hex.DecodeString(tc.hexString)
			if err != nil {
				t.Fatalf("Failed to decode hex string: %v", err)
			}

			// Parse bytes to MoFsm struct
			moFsm, _, err := ParseMoFsm(originalBytes)
			if (err != nil) != tc.expectError {
				t.Fatalf("Unexpected error status during parsing: got %v, expected error: %v", err, tc.expectError)
			}

			// If we expect an error and got one, test passes
			if tc.expectError && err != nil {
				t.Logf("Expected error occurred in test case '%s': %v", tc.name, err)
				return
			}

			// Marshal MoFsm struct back to bytes
			marshaledBytes, err := moFsm.Marshal()
			if err != nil {
				t.Fatalf("Failed to marshal MoFsm: %v", err)
			}

			if err == nil {
				// Compare original and marshaled bytes
				if diff := cmp.Diff(originalBytes, marshaledBytes); diff != "" {
					t.Errorf("Marshaled bytes don't match original (-original +marshaled):\n%s", diff)
				}
			}
		})
	}
}
