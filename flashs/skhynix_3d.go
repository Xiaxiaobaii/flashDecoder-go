package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type SkHynix3DDecoder struct{}

func (s SkHynix3DDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "H25")
}

func (s SkHynix3DDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Vendor = "SkHynix"
	info.Type = flashinfo.NAND
	info.Toggle = true
	info.ToggleN = "2.0"

	utils.RetShiftChars(&partNumber, 3)
	info.ToggleN = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"B": "4.0",
		"Q": "2.0",
	}, "")
	info.Voltage = "Vcc: 2.7V~3.6V, VccQ: 1.7V~1.95V/1.14V~1.26V"

	utils.RetShiftChars(&partNumber, 1)
	level := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"M": 2,
		"T": 3,
	}, 0)
	info.CellLevel = utils.IntCellTString(level)
	info.DeviceWidth = 8

	utils.RetShiftChars(&partNumber, 1)
	densityCode := utils.RetShiftChars(&partNumber, 1)
	if level == 2 {
		info.Capacity = utils.GetOrDefault(densityCode, map[string]string{
			"A": "256Gbit",
		}, "")
	} else {
		info.Capacity = utils.GetOrDefault(densityCode, map[string]string{
			"A": "512Gbit",
			"B": "1Tbit",
			"D": "2Tbit",
			"F": "4Tbit",
			"G": "8Tbit",
		}, "")
	}

	utils.RetShiftChars(&partNumber, 1)
	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Samsung_Generation, 0)

	return info
}

func SkHynix3DDecoderDefault() SkHynix3DDecoder {
	return SkHynix3DDecoder{}
}
