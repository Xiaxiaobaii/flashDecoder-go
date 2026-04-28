package flashs

import "testing"

func TestSpecTekHalfPageAndSizeFlag(t *testing.T) {
	// Construct a minimal SpecTek PN that triggers config == 'M'.
	// Decoder flow (after removing first 3 chars "F??"):
	// [cellLevel 1][process 3][density 2][config 1]...
	// Use legacy density code "18" (1.8Gbit) so it won't be consumed by MicronCapacitys,
	// leaving config at the next character.
	pn := "FBX" + "L84C" + "18" + "M" + "D" + "A" + "A" + "E" + "AB"

	info := SpecTekDecoderDefault().Decode(pn)
	if info.Vendor != "SpecTek" {
		t.Fatalf("expected Vendor=SpecTek, got %q", info.Vendor)
	}
	if !info.HalfPageAndSize {
		t.Fatalf("expected HalfPageAndSize=true for config M")
	}
}
