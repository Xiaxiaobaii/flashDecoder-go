package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type MicronFpgaDecoder struct {
}

func MicronFpgaDecoderDefault() MicronFpgaDecoder {
	return MicronFpgaDecoder{}
}

func (decoder MicronFpgaDecoder) Check(partNumber string) bool {
	for _, v := range []string{"NW", "NX", "NQ", "PF", "NY", "NC"} {
		if utils.RetStartBAEq(partNumber, v) || (len(partNumber) == 10 && utils.SubOffsetStr(partNumber, 5, 2) == v) {
			return true
		}
	}
	return false
}

func (decoder MicronFpgaDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	//var i string
	if len(partNumber) == 10 {
		utils.RetShiftChars(&partNumber, 5)
	}

	info.PartNumber = partNumber
	pn := utils.MFpga2Pn(partNumber)
	if pn != "" {
		info.Retry = true
		info.PartNumber = pn
	}

	return info
}
