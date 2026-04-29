package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type WesternDigitalDecoder struct{}

func WesternDigitalDecoderDefault() WesternDigitalDecoder {
	return WesternDigitalDecoder{}
}

func (w WesternDigitalDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "SD")
}

func (w WesternDigitalDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.PartNumber = partNumber
	info.Vendor = "WesternDigital"

	utils.RetShiftChars(&partNumber, 2)
	if utils.RetStartBAEq(partNumber, "IN") {
		info.Type = flashinfo.INAND
		info.Unsupported_Reason = "Sandisk_iNand_Not_Supported"
		return info
	}
	if utils.RetStartBAEq(partNumber, "IS") {
		info.Type = flashinfo.CON
		info.Unsupported_Reason = "Sandisk_iSSD_Not_Supported"
		return info
	}

	info.Type = flashinfo.NAND
	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"T": "TSOP",
		"Y": "BGA",
		"Z": "LGA",
		"Q": "BGA",
	}, "")

	if partNumber == "" {
		return info
	}
	if utils.RetStartBAEq(partNumber, "N") {
		utils.RetShiftChars(&partNumber, 1)
	}

	info.Process = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "BiCS2",
		"B": "BiCS3",
		"C": "BiCS4",
		"L": "56 nm",
		"M": "43 nm",
		"N": "32 nm",
		"P": "24 nm",
		"Q": "19 nm",
		"R": "1y nm",
		"S": "15 nm",
	}, "")

	info.CellLevel = utils.IntCellTString(utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"C": 3,
		"F": 2,
		"G": 2,
		"H": 2,
		"I": 3,
		"M": 2,
		"N": 3,
		"Q": 2,
	}, 0))

	info.CE = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"A": 1,
		"B": 2,
		"C": 4,
		"D": 8,
	}, -1)

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"H": "2.7V~3.6V",
	}, "")
	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"E": 8,
		"S": 16,
	}, -1)

	if utils.RetStartBAEq(partNumber, "M") {
		info.Lead_Free = true
		utils.RetShiftChars(&partNumber, 1)
	}

	if utils.RetStartBAEq(partNumber, "-") {
		utils.RetShiftChars(&partNumber, 1)
		info.Capacity = utils.StartMatchFromMap(&partNumber, map[string]string{
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
	}

	return info
}
