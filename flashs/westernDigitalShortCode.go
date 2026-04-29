package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
	"strings"
)

type WesternDigitalShortCodeDecoder struct{}

func WesternDigitalShortCodeDecoderDefault() WesternDigitalShortCodeDecoder {
	return WesternDigitalShortCodeDecoder{}
}

func (w WesternDigitalShortCodeDecoder) Check(partNumber string) bool {
	if len(partNumber) != 10 {
		return false
	}
	parts := strings.SplitN(partNumber, "-", 2)
	return len(parts) == 2 && len(parts[0]) == 5 && len(parts[1]) == 4
}

func (w WesternDigitalShortCodeDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.PartNumber = partNumber
	info.Vendor = "WesternDigital"

	parts := strings.SplitN(partNumber, "-", 2)
	if len(parts) != 2 {
		return info
	}
	size := parts[1]
	info.Capacity = utils.GetOrDefault(size, map[string]string{
		"1024": "1Gbit",
		"2048": "2Gbit",
		"4096": "4Gbit",
		"001G": "8Gbit",
		"002G": "16Gbit",
		"004G": "32Gbit",
		"008G": "64Gbit",
		"016G": "128Gbit",
		"032G": "256Gbit",
		"064G": "512Gbit",
		"128G": "1Tbit",
		"256G": "2Tbit",
		"512G": "4Tbit",
		"1T00": "8Tbit",
	}, "")
	return info
}
