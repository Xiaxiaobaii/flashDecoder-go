package flashids

import (
	flashinfo "flashDecoder/info"
	"fmt"
	"strings"
)

type BasicIDDecoder struct{}

var vendorCodeMap = map[byte]string{
	0x2C: "Micron",
	0x98: "Kioxia",
	0x89: "Intel",
	0xEC: "Samsung",
	0xAD: "SKHynix",
	0x45: "Sandisk",
	0xC8: "XTX",
	0xC2: "Macronix",
	0xEF: "Winbond",
	0x9B: "Yangtze",
}

func (d BasicIDDecoder) Check(id []byte) bool {
	return len(id) > 0
}

func (d BasicIDDecoder) Decode(idHex string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Type = flashinfo.NAND
	info.IDs = splitHexBytes(idHex)
	if len(info.IDs) == 0 {
		info.Unsupported_Reason = "invalid flash id"
		return info
	}
	vendorCode := parseHexByte(info.IDs[0])
	info.Vendor = vendorCodeMap[vendorCode]
	if info.Vendor == "" {
		info.Vendor = fmt.Sprintf("Unknown(0x%02X)", vendorCode)
	}
	return info
}

func BasicIDDecoderDefault() BasicIDDecoder {
	return BasicIDDecoder{}
}

func splitHexBytes(idHex string) []string {
	if len(idHex) < 2 {
		return nil
	}
	res := make([]string, 0, len(idHex)/2)
	for offset := 0; offset+1 < len(idHex); offset += 2 {
		res = append(res, idHex[offset:offset+2])
	}
	return res
}

func parseHexByte(hexByte string) byte {
	var value byte
	_, err := fmt.Sscanf(strings.ToUpper(hexByte), "%02X", &value)
	if err != nil {
		return 0
	}
	return value
}
