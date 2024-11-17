package tzsp_layer_test

import (
	"testing"

	"github.com/google/gopacket"

	tzsp "github.com/Yiivgeny/tzsp-layer"
)

var testPacket = []byte{
	0x01, 0x00, 0x00, 0x01, 0x01, 0x2c, 0xc8, 0x1b, 0x22, 0x85, 0x53, 0x42, 0x01, 0x1e, 0x96, 0x4f, 0x4d, 0x08, 0x00,
	0x45, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06, 0x06, 0xd4, 0xc0, 0xa8, 0x59, 0x1a, 0xc0, 0xa8, 0x59,
	0x01, 0xc7, 0xf4, 0x20, 0x63, 0xb3, 0x4c, 0xef, 0xe3, 0xa6, 0x39, 0xe3, 0x68, 0x80, 0x18, 0x0b, 0xb6, 0x86, 0x03,
	0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xcd, 0x5b, 0x5e, 0xc3, 0x27, 0xfc, 0x25, 0x3e, 0x82, 0x06, 0x00, 0x70, 0x71,
	0x94, 0x18, 0xf0, 0xb1, 0x40, 0x0b, 0x74, 0x4b, 0x47, 0xbe, 0x46, 0x50, 0x00, 0x78, 0xa1, 0x0b, 0x14, 0x55, 0x50,
	0x0e, 0x8d, 0x07, 0xf1, 0xf8, 0xd3, 0x5a, 0xa2, 0x1b, 0x69, 0x0b, 0xfc, 0x51, 0x96, 0x87, 0x25, 0x2e, 0x98, 0x37,
	0xab, 0x8c, 0x06, 0xe5, 0x35, 0xfb, 0xda, 0x8f, 0x7d, 0xc0, 0x1a, 0x2d, 0xd0, 0x46, 0x5d, 0x3c, 0x25, 0xfc, 0xbc,
	0x24, 0x60, 0xf7, 0xc3, 0xc0, 0x85, 0x9e, 0x0e, 0x47, 0x9d, 0x22, 0xeb, 0x17, 0x1f, 0x2b, 0xc8, 0xb9, 0xe9, 0x37,
	0x30, 0x32, 0x00, 0x67, 0xcb, 0xf6, 0xc5, 0xfd, 0x09, 0xfb, 0x63, 0x23, 0x63, 0x20, 0xe0, 0xd1, 0x52, 0x78, 0x51,
	0x28, 0x05, 0xe3, 0x2e, 0xee, 0x97, 0x23, 0x15, 0x1a, 0xfd, 0xf9, 0x79, 0x7a, 0x44, 0xdc, 0xf6, 0x0a, 0xf9, 0xa9,
	0xe1, 0x2b, 0x84, 0xea, 0x5a, 0x8b, 0xd0, 0x5e, 0x84, 0xdd, 0x5d, 0xa1, 0x01,
}

func Test_TZSP_Layer(t *testing.T) {
	// Create a packet
	p := gopacket.NewPacket(testPacket, tzsp.LayerTypeTZSP, gopacket.Default)

	// Extract the TZSP layer
	layer := p.Layer(tzsp.LayerTypeTZSP)
	if layer == nil {
		t.Fatalf("Expected layer of type %v, but got nil", tzsp.LayerTypeTZSP)
	}

	// Check layer type conversion
	layerTZSP, ok := layer.(*tzsp.TZSP)
	if !ok {
		t.Fatalf("Expected layer of type *tzsp.TZSP, but got %T", layer)
	}

	// Verify TZSP layer fields
	if layerTZSP.Version != 1 {
		t.Errorf("Expected Version=1, but got %d", layerTZSP.Version)
	}
	if layerTZSP.Type != 0 {
		t.Errorf("Expected Type=0, but got %d", layerTZSP.Type)
	}
	if layerTZSP.EncapProtocol != tzsp.ProtocolEthernet {
		t.Errorf("Expected EncapProtocol=%v, but got %v", tzsp.ProtocolEthernet, layerTZSP.EncapProtocol)
	}

	// Verify other layers
	if p.LinkLayer().LayerType().String() != "Ethernet" {
		t.Errorf("Expected LinkLayer type to be 'Ethernet', but got '%s'", p.LinkLayer().LayerType().String())
	}
	if p.NetworkLayer().LayerType().String() != "IPv4" {
		t.Errorf("Expected NetworkLayer type to be 'IPv4', but got '%s'", p.NetworkLayer().LayerType().String())
	}
	if p.TransportLayer().LayerType().String() != "TCP" {
		t.Errorf("Expected TransportLayer type to be 'TCP', but got '%s'", p.TransportLayer().LayerType().String())
	}

	// Verify application payload
	expectedPayload := testPacket[71:]
	actualPayload := p.ApplicationLayer().Payload()
	if string(expectedPayload) != string(actualPayload) {
		t.Errorf("Expected application payload to be %v, but got %v", expectedPayload, actualPayload)
	}
}