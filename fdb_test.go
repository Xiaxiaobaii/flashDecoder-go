package flashdecoder

import "testing"

func TestFindFdbMicron(t *testing.T) {
	entry, err := FindFdb("micron", "NW101")
	if err != nil {
		t.Fatalf("FindFdb returned error: %v", err)
	}
	if entry.Vendor != "Micron" {
		t.Fatalf("expected Micron vendor, got %q", entry.Vendor)
	}
	if len(entry.Parts) != 1 || entry.Parts[0] == "" {
		t.Fatalf("expected one non-empty part, got %#v", entry.Parts)
	}
}

func TestFindFdbUnknownVendor(t *testing.T) {
	_, err := FindFdb("intel", "ABCD")
	if err == nil {
		t.Fatalf("expected error for unsupported vendor")
	}
}

func TestSearchFdb(t *testing.T) {
	results := SearchFdb("NW101", 10)
	if len(results) == 0 {
		t.Fatalf("expected non-empty search results")
	}
}
