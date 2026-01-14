package flashdecoder

import (
	"errors"
	"flashDecoder/flashs"
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
	"fmt"
	"strings"
)

type Decoder interface {
	Check(partNumber string) bool

	Decode(partNumber string) flashinfo.Flashinfo
}

type IdDecoder interface {
	Decode(partNumber string) flashinfo.Flashinfo
}

var VendorIds = map[string]IdDecoder {
	
}

var flashdecoders []Decoder

func init() {
	flashdecoders = append(flashdecoders, flashs.MicronDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.MicronFpgaDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.IntelDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.KioxiaDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.PhisonDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SamsungDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SpecTekDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.SkHynixDecoderDefault())
	flashdecoders = append(flashdecoders, flashs.YangTzeDecoderDefault())
	mdb, e := utils.LoadMdb()
	utils.Mdb = mdb
	if e != nil {
		fmt.Println("load mdb error: ", e.Error())
	}
}

func Decode(partNumber string) (flashinfo.Flashinfo, error) {
	partNumber = strings.ToUpper(partNumber)
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

// func IdDecode(id string) (flashinfo.Flashinfo, error) {
// 	for _, v := range flashIdDecoders {
// 		if v.Check(id) {
// 			info := v.Decode(id)
// 			if info.Retry {
// 				info.Retry = false
// 				return Decode(info.PartNumber)
// 			} else {
// 				if info.Unsupported_Reason == "" {
// 					return info, nil
// 				} else {
// 					return info, errors.New(info.Unsupported_Reason)
// 				}

// 			}
// 		}
// 	i
// 	return flashinfo.Flashinfo{}, errors.New("no Support flash partNumber found")
// }
