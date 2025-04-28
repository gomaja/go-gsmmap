package asn1mapmodel

// -- general data types

// AddressString is OCTET STRING encoded as TBCD-String
type AddressString []byte

// ISDNAddressString inherits AddressString
type ISDNAddressString AddressString

// SignalInfo defines the ASN.1 OCTET STRING with size constraint (1..maxSignalInfoLength)
type SignalInfo []byte

// -- data types for numbering and identification

// IMSI digits of MCC, MNC, MSIN are concatenated in this order
type IMSI []byte

// EncodeAddressString encodes the AddressString into an OCTET STRING
func EncodeAddressString(extension, natureOfAddress, numberingPlan uint8, digits []byte) []byte {
	firstOctet := extension | natureOfAddress | (numberingPlan & 0x0F)
	return append([]byte{firstOctet}, digits...)
}

// DecodeAddressString decodes an OCTET STRING into its components: extension, natureOfAddress, numberingPlan, and digits
func DecodeAddressString(encoded []byte) (extensionIndicator, natureOfAddress, numberingPlan uint8, digits []byte) {
	if len(encoded) == 0 {
		return 0, 0, 0, nil // Handle empty input
	}

	// Extract the first octet to get extension, natureOfAddress, and numberingPlan
	firstOctet := encoded[0]

	extensionIndicator = firstOctet & 0b10000000 // Extension: bit 8 (MSB)
	natureOfAddress = firstOctet & 0b01110000    // Extracts bits 7, 6, 5
	numberingPlan = firstOctet & 0x0F            // Extracts bits 4, 3, 2, 1

	// Remaining bytes are the digits
	if len(encoded) > 1 {
		digits = encoded[1:]
	} else {
		digits = []byte{}
	}

	return extensionIndicator, natureOfAddress, numberingPlan, digits
}

// LMSI represents an ASN.1 OCTET STRING with a size constraint of 4 bytes.
type LMSI []byte
