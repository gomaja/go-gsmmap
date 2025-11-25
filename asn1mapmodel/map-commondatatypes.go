package asn1mapmodel

// ExtensionNo Constants representing the extension indicator
const (
	ExtensionNo = 0b10000000 // bit 8 set to 1, indicating no extension
)

// Constants representing the nature of address indicator (bits 7, 6, 5)
const (
	AddressNatureUnknown           = 0b000 << 4
	AddressNatureInternational     = 0b001 << 4
	AddressNatureNational          = 0b010 << 4
	AddressNatureNetworkSpecific   = 0b011 << 4
	AddressNatureSubscriber        = 0b100 << 4
	AddressNatureReserved          = 0b101 << 4
	AddressNatureAbbreviated       = 0b110 << 4
	AddressNatureReservedExtension = 0b111 << 4
)

// Constants representing the numbering plan indicator (bits 4, 3, 2, 1)
const (
	NumberingPlanUnknown           = 0b0000
	NumberingPlanISDN              = 0b0001
	NumberingPlanSpare1            = 0b0010
	NumberingPlanData              = 0b0011
	NumberingPlanTelex             = 0b0100
	NumberingPlanSpare2            = 0b0101
	NumberingPlanLandMobile        = 0b0110
	NumberingPlanSpare3            = 0b0111
	NumberingPlanNational          = 0b1000
	NumberingPlanPrivate           = 0b1001
	NumberingPlanReservedExtension = 0b1111
)

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
