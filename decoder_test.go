package flashdecoder

import "testing"

func TestDecode_MicronAliasNW383(t *testing.T) {
	info, err := Decode("NW383")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if info.Vendor != "Micron" {
		t.Fatalf("expected Vendor=Micron, got %q", info.Vendor)
	}
	if info.Type != "Nand" {
		t.Fatalf("expected Type=Nand, got %q", info.Type)
	}
	if info.Capacity != "128Gbit" {
		t.Fatalf("expected Capacity=128Gbit, got %q", info.Capacity)
	}
	if info.PartNumber == "" {
		t.Fatalf("expected PartNumber to be non-empty")
	}
}

func TestDecode_SpecTekFpgaPF001(t *testing.T) {
	info, err := Decode("PF001")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if info.Vendor != "SpecTek" {
		t.Fatalf("expected Vendor=SpecTek, got %q", info.Vendor)
	}
	// Keep assertions to stable fields; avoid long text descriptions.
	if info.Process != "L84C" {
		t.Fatalf("expected Process=L84C, got %q", info.Process)
	}
	if info.DeviceWidth != 8 {
		t.Fatalf("expected DeviceWidth=8, got %d", info.DeviceWidth)
	}
}

func TestDecode_EmptyPartNumber(t *testing.T) {
	_, err := Decode("")
	if err == nil {
		t.Fatalf("expected error for empty PN")
	}
}
