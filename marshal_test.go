package gsmmap

import (
	"testing"
)

// TestMtFsmMarshal is commented out because it requires proper initialization of sms.Deliver
// which depends on the github.com/fkgi/sms package API
/*
func TestMtFsmMarshal(t *testing.T) {
	// This test requires proper initialization of sms.Deliver
	// which depends on the github.com/fkgi/sms package API
}
*/

func TestSriSmMarshal(t *testing.T) {
	// Create a simple SriSm struct for testing
	sriSm := &SriSm{
		MSISDN:               "123456789",
		SmRpPri:              true,
		ServiceCentreAddress: "12345",
	}

	// Marshal the struct
	marshaledBytes, err := sriSm.Marshal()
	if err != nil {
		t.Fatalf("Failed to marshal SriSm: %v", err)
	}

	// Parse the marshaled bytes back to a struct
	parsedSriSm, _, err := ParseSriSm(marshaledBytes)
	if err != nil {
		t.Fatalf("Failed to parse marshaled bytes: %v", err)
	}

	// Compare the original and parsed structs
	if sriSm.MSISDN != parsedSriSm.MSISDN {
		t.Errorf("MSISDN mismatch: got %s, expected %s", parsedSriSm.MSISDN, sriSm.MSISDN)
	}
	if sriSm.SmRpPri != parsedSriSm.SmRpPri {
		t.Errorf("SmRpPri mismatch: got %v, expected %v", parsedSriSm.SmRpPri, sriSm.SmRpPri)
	}
	if sriSm.ServiceCentreAddress != parsedSriSm.ServiceCentreAddress {
		t.Errorf("ServiceCentreAddress mismatch: got %s, expected %s", parsedSriSm.ServiceCentreAddress, sriSm.ServiceCentreAddress)
	}
}

func TestSriSmRespMarshal(t *testing.T) {
	// Create a simple SriSmResp struct for testing
	sriSmResp := &SriSmResp{
		IMSI: "123456789012345",
		LocationInfoWithLMSI: LocationInfoWithLMSI{
			NetworkNodeNumber: "12345",
		},
	}

	// Marshal the struct
	marshaledBytes, err := sriSmResp.Marshal()
	if err != nil {
		t.Fatalf("Failed to marshal SriSmResp: %v", err)
	}

	// Parse the marshaled bytes back to a struct
	parsedSriSmResp, _, err := ParseSriSmResp(marshaledBytes)
	if err != nil {
		t.Fatalf("Failed to parse marshaled bytes: %v", err)
	}

	// Compare the original and parsed structs
	if sriSmResp.IMSI != parsedSriSmResp.IMSI {
		t.Errorf("IMSI mismatch: got %s, expected %s", parsedSriSmResp.IMSI, sriSmResp.IMSI)
	}
	if sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber != parsedSriSmResp.LocationInfoWithLMSI.NetworkNodeNumber {
		t.Errorf("NetworkNodeNumber mismatch: got %s, expected %s",
			parsedSriSmResp.LocationInfoWithLMSI.NetworkNodeNumber,
			sriSmResp.LocationInfoWithLMSI.NetworkNodeNumber)
	}
}

// Test edge cases
func TestMarshalEdgeCases(t *testing.T) {
	// Test with empty strings
	emptySriSm := &SriSm{
		MSISDN:               "",
		SmRpPri:              false,
		ServiceCentreAddress: "",
	}

	_, err := emptySriSm.Marshal()
	if err != nil {
		t.Errorf("Failed to marshal SriSm with empty strings: %v", err)
	}

	// Test with very long strings
	longString := "1234567890123456789012345678901234567890123456789012345678901234567890"
	longSriSm := &SriSm{
		MSISDN:               longString,
		SmRpPri:              true,
		ServiceCentreAddress: longString,
	}

	_, err = longSriSm.Marshal()
	if err != nil {
		t.Errorf("Failed to marshal SriSm with long strings: %v", err)
	}
}
