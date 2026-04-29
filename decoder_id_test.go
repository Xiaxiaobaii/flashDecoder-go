package flashdecoder

import "testing"

func TestDecodeID_NormalizesAndParsesVendor(t *testing.T) {
	info, err := DecodeID("ec d7 94 7e")
	if err != nil {
		t.Fatalf("DecodeID returned error: %v", err)
	}
	if info.Vendor != "Samsung" {
		t.Fatalf("expected vendor Samsung, got %q", info.Vendor)
	}
	if len(info.IDs) != 4 || info.IDs[0] != "EC" {
		t.Fatalf("unexpected ids: %#v", info.IDs)
	}
}

func TestDecodeID_UnknownVendor(t *testing.T) {
	info, err := DecodeID("AA1122")
	if err != nil {
		t.Fatalf("DecodeID returned error: %v", err)
	}
	if info.Vendor != "Unknown(0xAA)" {
		t.Fatalf("unexpected unknown vendor text: %q", info.Vendor)
	}
}

func TestDecodeID_InvalidInput(t *testing.T) {
	_, err := DecodeID("xyz")
	if err == nil {
		t.Fatalf("expected error for invalid input")
	}
}
