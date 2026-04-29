package flashdecoder

import (
	"errors"
	"flashDecoder/flashids"
	"flashDecoder/flashs"
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
	"fmt"
	"regexp"
	"strings"
)

type Decoder interface {
	Check(partNumber string) bool

	Decode(partNumber string) flashinfo.Flashinfo
}

type IdDecoder interface {
	Check(id []byte) bool

	Decode(partNumber string) flashinfo.Flashinfo
}

var VendorIds = map[byte]IdDecoder{}

var flashdecoders []Decoder
var flashIdDecoders []IdDecoder

func init() {
	flashdecoders = append(flashdecoders, flashs.MicronDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.MicronFpgaDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.IntelDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.KioxiaDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.PhisonDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SamsungDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SpecTekDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.WesternDigitalShortCodeDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.WesternDigitalDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SkHynix3DDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SkHynixLegacyDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SkHynixDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.YangTzeDecoderDefault())
	flashIdDecoders = append(flashIdDecoders, flashids.BasicIDDecoderDefault())
	mdb, _ := utils.LoadMdb()
	if mdb.Micron == nil {
		mdb.Micron = utils.Micron
	}
	utils.Mdb = mdb
}

func Decode(partNumber string) (flashinfo.Flashinfo, error) {
	partNumber = strings.ToUpper(partNumber)
	if strings.TrimSpace(partNumber) == "" {
		return flashinfo.Flashinfo{}, errors.New("empty partNumber")
	}
	for _, v := range flashdecoders {
		if v.Check(partNumber) {
			info := v.Decode(partNumber)
			if info.Retry {
				info.Retry = false
				return Decode(info.PartNumber)
			} else {
				if info.Unsupported_Reason == "" {
					return info, nil
				} else {
					return info, errors.New(info.Unsupported_Reason)
				}

			}
		}
	}
	return flashinfo.Flashinfo{}, errors.New("no Support flash partNumber found")
}

func DecodeID(id string) (flashinfo.Flashinfo, error) {
	normalized, bytes, err := normalizeFlashID(id)
	if err != nil {
		return flashinfo.Flashinfo{}, err
	}
	for _, decoder := range flashIdDecoders {
		if decoder.Check(bytes) {
			info := decoder.Decode(normalized)
			if info.Unsupported_Reason == "" {
				return info, nil
			}
			return info, errors.New(info.Unsupported_Reason)
		}
	}
	return flashinfo.Flashinfo{}, fmt.Errorf("no flash id decoder matched: %s", normalized)
}

func normalizeFlashID(raw string) (string, []byte, error) {
	cleaned := strings.ToUpper(strings.TrimSpace(raw))
	if cleaned == "" {
		return "", nil, errors.New("empty flash id")
	}
	hexOnly := regexp.MustCompile(`[^0-9A-F]`).ReplaceAllString(cleaned, "")
	if len(hexOnly) < 2 || len(hexOnly)%2 != 0 {
		return "", nil, fmt.Errorf("invalid flash id: %s", raw)
	}
	bytes := make([]byte, 0, len(hexOnly)/2)
	for offset := 0; offset < len(hexOnly); offset += 2 {
		var value byte
		_, parseErr := fmt.Sscanf(hexOnly[offset:offset+2], "%02X", &value)
		if parseErr != nil {
			return "", nil, fmt.Errorf("invalid flash id: %s", raw)
		}
		bytes = append(bytes, value)
	}
	return hexOnly, bytes, nil
}
