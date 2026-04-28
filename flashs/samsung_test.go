package flashs

import "testing"

func TestSamsungVoltageCodeJ(t *testing.T) {
	// Construct a minimal Samsung PN that hits the voltage code = 'J'.
	// Format in decoder:
	// K9 + [classification 1] + [capacity 2] + [technology 1] + [width 1] + [voltage 1] + [mode 1] + [generation 1] + [-] + [package 1] ...
	pn := "K931GD8J0M-B"

	info := SamsungDecoderDefault().Decode(pn)
	if info.Vendor != "Samsung" {
		t.Fatalf("expected Vendor=Samsung, got %q", info.Vendor)
	}
	if info.Voltage == "" {
		t.Fatalf("expected Voltage to be non-empty for code J")
	}
	if info.Voltage != "Vcc: 3.3V, VccQ: 1.8V (UNOFFICIAL)" {
		t.Fatalf("unexpected Voltage: %q", info.Voltage)
	}
}

