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

func TestDecode_SkHynixLegacyPrefix(t *testing.T) {
	info, err := Decode("HY27UF081G2M")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if info.Vendor != "SkHynix" {
		t.Fatalf("expected Vendor=SkHynix, got %q", info.Vendor)
	}
	if info.Type != "Nand" {
		t.Fatalf("expected Type=Nand, got %q", info.Type)
	}
}

func TestDecode_SkHynix3DPrefix(t *testing.T) {
	info, err := Decode("H25QTA82A")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if info.Vendor != "SkHynix" {
		t.Fatalf("expected Vendor=SkHynix, got %q", info.Vendor)
	}
	if info.Type != "Nand" {
		t.Fatalf("expected Type=Nand, got %q", info.Type)
	}
	if !info.Toggle {
		t.Fatalf("expected Toggle=true")
	}
}

func TestDecode_WesternDigitalShortCode(t *testing.T) {
	info, err := Decode("05055-032G")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if info.Vendor != "WesternDigital" {
		t.Fatalf("expected Vendor=WesternDigital, got %q", info.Vendor)
	}
	if info.Capacity != "256Gbit" {
		t.Fatalf("expected Capacity=256Gbit, got %q", info.Capacity)
	}
}

func TestDecode_ParityVectorsPN(t *testing.T) {
	cases := []struct {
		pn       string
		vendor   string
		nandType string
	}{
		{pn: "NW383", vendor: "Micron", nandType: "Nand"},
		{pn: "K9GAG08U0M", vendor: "Samsung", nandType: "Nand"},
		{pn: "TH58NVG7S0FTA20", vendor: "Kioxia", nandType: "Nand"},
		{pn: "YMN08TA1B1C0A", vendor: "YangTze", nandType: "Nand"},
		{pn: "SDTNQCDHE-032G", vendor: "WesternDigital", nandType: "Nand"},
	}

	for _, tc := range cases {
		info, err := Decode(tc.pn)
		if err != nil {
			t.Fatalf("%s: expected nil error, got %v", tc.pn, err)
		}
		if info.Vendor != tc.vendor {
			t.Fatalf("%s: expected Vendor=%s, got %q", tc.pn, tc.vendor, info.Vendor)
		}
		if string(info.Type) != tc.nandType {
			t.Fatalf("%s: expected Type=%s, got %q", tc.pn, tc.nandType, info.Type)
		}
	}
}
