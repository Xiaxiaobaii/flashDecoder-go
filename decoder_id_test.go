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

func TestDecodeID_ParityVectors(t *testing.T) {
	cases := []struct {
		id     string
		vendor string
	}{
		{id: "2C123456", vendor: "Micron"},
		{id: "98A1B2C3", vendor: "Kioxia"},
		{id: "89FFFF00", vendor: "Intel"},
		{id: "ADF0F1F2", vendor: "SKHynix"},
		{id: "9B112233", vendor: "Yangtze"},
	}

	for _, tc := range cases {
		info, err := DecodeID(tc.id)
		if err != nil {
			t.Fatalf("%s: expected nil error, got %v", tc.id, err)
		}
		if info.Vendor != tc.vendor {
			t.Fatalf("%s: expected Vendor=%s, got %q", tc.id, tc.vendor, info.Vendor)
		}
	}
}
