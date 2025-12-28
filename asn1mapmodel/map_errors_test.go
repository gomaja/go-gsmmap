package asn1mapmodel

import (
	"testing"
)

func TestErrorBusySubscriberExport(t *testing.T) {
	// Verify that ErrorBusySubscriber is exported and accessible
	var err Error = ErrorBusySubscriber

	// Check the string representation
	expectedString := "BusySubscriber"
	if str := GetErrorString(uint8(err)); str != expectedString {
		t.Errorf("Expected string for ErrorBusySubscriber to be %q, but got %q", expectedString, str)
	}
}
