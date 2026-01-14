package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
	"strconv"
)

type IntelDecoder struct {
}

func (decoder IntelDecoder) Check(partNumber string) bool {
	code := utils.Substr(partNumber, 2)
	if utils.RetStartBAEq(partNumber, "PF29F") || utils.RetStartBAEq(partNumber, "PF29R") {
		return true
	}
	for _, v := range []string{"JS", "29", "X2", "BK", "CU"} {
		if code == v {
			return true
		}
	}

	return false
}

func (decoder IntelDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	//test: PF29F04T2ANCQJ1

	info := flashinfo.Default()
	info.PartNumber = partNumber
	info.Vendor = "Intel"

	if utils.RetStartBAEq(partNumber, "X") {
		utils.RetShiftChars(&partNumber, 1)
	} else {
		pack := utils.RetShiftChars(&partNumber, 2)
		for _, v := range []string{"JS", "PF", "BK", "CU"} {
			if pack == v {
				info.Package = utils.GetOrDefault(pack, map[string]string{
					"JS": "TSOP48",
					"BK": "LGA",
					"PF": "BGA",
					"CU": "LSOP",
				}, "")

			}
		}
	}

	utils.RetShiftChars(&partNumber, 2)

	typ := utils.RetShiftChars(&partNumber, 1)
	info.Type = utils.GetOrDefault(typ, map[string]flashinfo.NandType{
		"F": flashinfo.NAND,
		"P": flashinfo.Xpoint_3D,
	}, flashinfo.UNKWON_NAND)

	if typ == "P" {
		return info
	}

	cap := utils.RetShiftChars(&partNumber, 3)
	info.Capacity = utils.GetOrDefault(cap, map[string]string{
		"01G": "1Gbit",
		"02G": "2Gbit",
		"04G": "4Gbit",
		"08G": "8Gbit",
		"16G": "16Gbit",
		"32G": "32Gbit",
		"64G": "64Gbit",
		"16B": "128Gbit",
		"32B": "256Gbit",
		"48B": "384Gbit",
		"64B": "512Gbit",
		"96B": "768Gbit",
		"01T": "1Tbit",
		"02T": "2Tbit",
		"03T": "3Tbit",
		"04T": "4Tbit",
		"06T": "6Tbit",
		"08T": "8Tbit",
		"16T": "16Tbit",
	}, "")

	width := utils.RetShiftChars(&partNumber, 2)
	info.DeviceWidth = utils.GetOrDefault(width, map[string]int{
		"08": 8,
		"16": 16,
		"2A": 8,
		"4A": 8,
		"A8": 8,
	}, 0)


	if cap[2] >= '0' && cap[2] <= '9' {
		return MicronDecoderDefault().Decode(info.PartNumber)
	}

	class := utils.RetShiftChars(&partNumber, 1)
	classIfication := utils.GetOrDefault(class, map[string][]int{
		"A": {1, 1, 1}, //Die, CE, RB
		"B": {2, 1, 1},
		"C": {2, 2, 2},
		"D": {2, 2, 2},
		"E": {2, 2, 2},
		"F": {4, 2, 2},
		"G": {4, 2, 2},
		"H": {8, 4, 4},
		"J": {4, 4, 4},
		"K": {8, 4, 4},
		"L": {1, 1, 1},
		"M": {2, 2, 2},
		"N": {4, 4, 4},
		"O": {8, 8, 4},
		"P": {8, 8, 4}, //L74
		"Q": {8, 2, -1},
		"S": {16, 4, 4},
		"W": {16, 8, 4},
		"Y": {16, 4, 4},
	}, []int{-1, -1, -1})

	info.Die = classIfication[0]
	info.CE = classIfication[1]
	info.RB = classIfication[2]

	info.Sync = utils.GetOrDefault(class, map[string]bool{
		"A": true, //Die, CE, RB, I/O Common/Separate (Sync/Async only)
		"B": true,
		"C": true,
		"D": true,
		"E": false,
		"F": true,
		"G": false,
		"H": true,
		"J": true,
		"K": false,
		"L": true,
		"M": true,
		"N": true,
		"O": true,
		"P": true, //L74
		"Q": true,
		"S": true,
		"W": true,
		"Y": true,
	}, true)

	if width == "4A" {
		info.CH = 4
	} else if width == "2A" {
		info.CH = 2
	} else {
		info.CH = 1
	}

	info.Toggle = false

	info.Async = true

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "3.3V (2.70V-3.60V)",
		"B": "1.8V (1.70V-1.95V)",
		"C": "Vcc: 3.3V, VccQ: 1.8V/1.2V",
	}, "")

	info.CellLevel = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"N": "SLC",
		"M": "MLC",
		"T": "TLC",
		"Q": "QLC",
	}, "")

	lithhography := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "90 nm",
		"B": "72 nm",
		"C": "50 nm",
		"D": "34 nm",
		"E": "25 nm",
		"F": "20 nm",
		"G": "3D1 32L",
		"H": "3D2 64L",
		"J": "3D3 96L",
		"K": "3D4 144L",
	}, "")
	info.Process = lithhography
	
	info.Generation, _ = strconv.Atoi(utils.RetShiftChars(&partNumber, 1))

	if info.CellLevel == "TLC" && lithhography == "3D1 32L" && info.Capacity == "1Tbit" {
		info.Capacity = "1.5Tbit"
	}

	return info
}

func IntelDecoderDefault() IntelDecoder {
	return IntelDecoder{}
}
