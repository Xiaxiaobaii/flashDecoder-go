package flashdecoder

import (
	"errors"
	"flashDecoder/utils"
	"sort"
	"strings"
)

type FdbEntry struct {
	Vendor string
	Code   string
	Parts  []string
}

type FdbSummary struct {
	MicronCount  int
	SpectekCount int
}

func GetFdbSummary() FdbSummary {
	return FdbSummary{
		MicronCount:  len(utils.Mdb.Micron),
		SpectekCount: len(utils.Mdb.Spectek),
	}
}

func FindFdb(vendor string, code string) (FdbEntry, error) {
	vendorNorm := strings.ToLower(strings.TrimSpace(vendor))
	codeNorm := strings.ToUpper(strings.TrimSpace(code))
	if vendorNorm == "" || codeNorm == "" {
		return FdbEntry{}, errors.New("vendor and code are required")
	}
	switch vendorNorm {
	case "micron":
		part, ok := utils.FindMicronByCode(codeNorm)
		if !ok {
			return FdbEntry{}, errors.New("no fdb entry found")
		}
		return FdbEntry{Vendor: "Micron", Code: codeNorm, Parts: []string{part}}, nil
	case "spectek":
		parts, ok := utils.FindSpectekByCode(codeNorm)
		if !ok {
			return FdbEntry{}, errors.New("no fdb entry found")
		}
		return FdbEntry{Vendor: "SpecTek", Code: codeNorm, Parts: parts}, nil
	default:
		return FdbEntry{}, errors.New("unsupported vendor for current fdb")
	}
}

func SearchFdb(keyword string, limit int) []FdbEntry {
	kw := strings.ToUpper(strings.TrimSpace(keyword))
	if kw == "" {
		return nil
	}
	if limit <= 0 {
		limit = 20
	}
	results := make([]FdbEntry, 0, limit)
	for code, part := range utils.Mdb.Micron {
		if strings.Contains(code, kw) || strings.Contains(strings.ToUpper(part), kw) {
			results = append(results, FdbEntry{
				Vendor: "Micron",
				Code:   code,
				Parts:  []string{part},
			})
		}
	}
	for code, parts := range utils.Mdb.Spectek {
		joined := strings.ToUpper(strings.Join(parts, " "))
		if strings.Contains(code, kw) || strings.Contains(joined, kw) {
			copied := make([]string, len(parts))
			copy(copied, parts)
			results = append(results, FdbEntry{
				Vendor: "SpecTek",
				Code:   code,
				Parts:  copied,
			})
		}
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].Vendor != results[j].Vendor {
			return results[i].Vendor < results[j].Vendor
		}
		return results[i].Code < results[j].Code
	})
	if len(results) > limit {
		return results[:limit]
	}
	return results
}
