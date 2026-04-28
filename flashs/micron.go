package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

var MicronCapacitys = map[string]string{
	"1G":   "1Gbit",
	"2G":   "2Gbit",
	"4G":   "4Gbit",
	"8G":   "8Gbit",
	"16G":  "16Gbit",
	"21G":  "21Gbit",
	"32G":  "32Gbit",
	"42G":  "42Gbit",
	"64G":  "64Gbit",
	"84G":  "84Gbit",
	"128G": "128Gbit",
	"168G": "168Gbit",
	"192G": "192Gbit",
	"256G": "256Gbit",
	"336G": "336Gbit",
	"384G": "384Gbit",
	"512G": "512Gbit",
	"768G": "768Gbit",
	"1T":   "1Tbit",
	"1T2":  "1.125Tbit",
	"1HT":  "1.5Tbit",
	"2T":   "2Tbit",
	"3T":   "3Tbit",
	"4T":   "4Tbit",
	"6T":   "6Tbit",
	"8T":   "8Tbit",
	"16T":  "16Tbit",
}

var MicronPackage = map[string]string{
	"WP": "48-pin TSOP I Center Package Leads (CPL) PB free",
	"WC": "48-pin TSOP I Off-center Package Leads (OCPL) PB free",
	"C5": "52-pad VLGA, 14 x 18 x 1.0 (SDP/DDP/QDP)",
	"G1": "272/352 ball VBGA, 14 x 18 x 1.0 (SDP, DDP, 3DP, QDP)",
	"G2": "272/352 ball TBGA, 14 x 18 x 1.3 (QDP, 8DP)",
	"G6": "272/352 ball LBGA, 14 x 18 x 1.5 (16DP)",
	"H1": "100/170 ball VBGA, 12 x 18 x 1.0",
	"H2": "100/170 ball TBGA, 12 x 18 x 1.2",
	"H3": "100/170 ball LBGA, 12 x 18 x 1.4 (8DP)",
	"H4": "63/120 ball VFBGA, 9 x 11 x 1.0",
	"HC": "63/120 ball VFBGA, 10.5 x 13 x 1.0",
	"H6": "152/221 ball VBGA, 14 x 18 x 1.0 (SDP, DDP)",
	"H7": "152/221 ball TBGA, 14 x 18 x 1.2 (QDP)",
	"H8": "152/221 ball LBGA, 14 x 18 x 1.4 (8DP)",
	"H9": "100-ball LBGA, 12 x 18 x 1.6 (16DP)",
	"J1": "132/187 ball VBGA, 12 x 18 x 1.0 (SDP, DDP)",
	"J2": "132/187 ball TBGA, 12 x 18 x 1.2 (QDP)",
	"J3": "132/187 ball LBGA, 12 x 18 x 1.4 (8DP)",
	"J4": "132/187 ball VBGA, 12 x 18 x 1.0 (SDP, DDP)",
	"J5": "132/187 all TBGA, 12 x 18 x 1.2 (QDP)",
	"J6": "132/187 ball LBGA, 12 x 18 x 1.4 (8DP)",
	"J7": "152/221 ball LBGA, 14 x 18 x 1.5 (16DP)",
	"J9": "132-ball LBGA, 12 x 18 x 1.5 (16DP)",
	//SpecTek
	"C3": "52-pad ULGA, 12 x 17 x 0.65",
	"C4": "52-pad VLGA, 12 x 17 x 1.0",
	"C6": "52-pad LLGA, 14 x 18 x 1.47",
	"C7": "48-pad LLGA, 12 x 20 x 1.47",
	"C8": "52-pad WLGA, 14 x 18 x 0.75",
	"D1": "52-pad VLGA, 11 x 14 x 0.9",
	"D4": "154/195 ball VFBGA, 13.5 x 11.5 x 1.0",
	"D5": "154/195 ball LFBGA, 13.5 x 11.5 x 1.3",
	"D6": "154/195 ball LFBGA, 13.5 x 11.5 x 1.5",
	"D7": "154/195 ball LFBGA, 13.5 x 11.5 x 1.6",
	"G4": "252/308 ball LFBGA, 12 x 18 x 1.5",
	"G5": "272/352 ball LFBGA, 14 x 18 x 1.4",
	"G7": "252/308 ball LFBGA, 12 x 18 x 1.0",
	"G8": "252/308 ball LFBGA, 12 x 18 x 1.3",
	"G9": "252/308 ball LFBGA, 12 x 18 x 1.4",
	"H5": "56/256 ball VFBGA, 12.8 x 9.5 x 1.0",
	"K3": "100/170 ball VLGA 12 x 18 x 0.9",
	"K4": "100/170 ball TLGA, 12 x 18 x 1.1",
	"K6": "152/221 ball LBGA, 14 x 18 x 1.3",
	"K7": "152/221 ball VLGA 14 x 18 x 0.9",
	"K8": "152/221 ball TLGA 14 x 18 x 1.1",
	"K9": "132/187 ball VLGA, 12 x 18 x 1.0",
	"L4": "308/368 ball VFBGA, 11.5x15.5x1.0",
	"MD": "130-ball VFBGA, 8 x 9 x 1.0",
	"M4": "132/187 ball TBGA, 12 x 18 x 1.3",
	"M5": "132/187 ball LBGA, 12 x 18 x 1.5",
	"M8": "55-ball VFBGA, 8 x 10 x 1.2", //should be M8Z
	"M9": "252/308 ball LFBGA 12 x 18 x 1.45",
}

var MicronClassIfication = map[string][]int{ // die, ce, rb, ch
	"A": {
		1, 0, 0, 1,
	},
	"B": {
		1, 1, 1, 1,
	},
	"D": {
		2, 1, 1, 1,
	},
	"E": {
		2, 2, 2, 2,
	},
	"F": {
		2, 2, 2, 1,
	},
	"G": {
		3, 3, 3, 3,
	},
	"J": {
		4, 2, 2, 1,
	},
	"K": {
		4, 2, 2, 2,
	},
	"L": {
		4, 4, 4, 4,
	},
	"M": {
		4, 4, 4, 2,
	},
	"Q": {
		8, 4, 4, 4,
	},
	"R": {
		8, 2, 2, 2,
	},
	"T": {
		16, 8, 4, 2,
	},
	"U": {
		8, 4, 4, 2,
	},
	"V": {
		16, 8, 4, 4,
	},
	//SpecTek, No R/nB
	"C": {
		3, 3, -1, 2,
	},
	"H": {
		4, 1, -1, 1,
	},
	"N": {
		6, 6, -1, 3,
	},
	"P": {
		8, 8, -1, 2,
	},
	"W": {
		16, 4, -1, 2,
	},
	"X": {
		4, 4, -1, 2,
	},
	"Y": {
		11, 7, -1, 4,
	},
	"1": {
		16, 2, -1, 1,
	},
	"2": {
		26, 8, -1, 2,
	},
	"3": {
		8, 4, -1, 2,
	},
	"4": {
		4, 4, -1, 1,
	},
	"S": {
		16, 4, -1, 4,
	},
}

type MicronDecoder struct {
}

func MicronDecoderDefault() MicronDecoder {
	return MicronDecoder{}
}

func (decoder MicronDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "MT")
}

func (decoder MicronDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.PartNumber = partNumber
	info.Vendor = "Micron"
	if !utils.RetStartBAEq(partNumber, "29") {
		partNumber = utils.Substr(partNumber, 2) // remove MT
	}
	partNumber = utils.Substr(partNumber, 2) // remove 29

	info.EnterPrise = utils.RetShiftChars(&partNumber, 1) == "E"

	info.Type = flashinfo.NAND

	info.Capacity = utils.StartMatchFromMap(&partNumber, MicronCapacitys, "")
	// NOTE: avoid printing from library code; callers (CLI/tests) should control output.
	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), map[string]int{
		"01": 1,
		"08": 8,
		"16": 16,
	}, -1)

	info.CellLevel = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "SLC",
		"C": "MLC",
		"E": "TLC",
		"G": "QLC",
	}, "")

	classIfication := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), MicronClassIfication, []int{-1, -1, -1, -1})
	info.Die = classIfication[0]
	info.CE = classIfication[1]
	info.RB = classIfication[2]
	info.CH = classIfication[3]

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "Vcc: 3.3V (2.70-3.60V), VccQ: 3.3V (2.70-3.60V)",
		"B": "1.8V (1.70-1.95V)",
		"C": "Vcc: 3.3V (2.70-3.60V), VccQ: 1.8V (1.70-1.95V)",
		"E": "Vcc: 3.3V (2.70-3.60V), VccQ: 3.3V (2.70-3.60V) or 1.8V (1.70-1.95V)",
		"F": "Vcc: 3.3V (2.50-3.60V), VccQ: 1.2V (1.14-1.26V)",
		"G": "Vcc: 3.3V (2.60-3.60V), VccQ: 1.8V (1.70-1.95V)",
		"H": "Vcc: 3.3V (2.50-3.60V), VccQ: 1.2V (1.14-1.26) or 1.8V (1.70-1.95V)",
		"J": "Vcc: 3.3V (2.50-3.60V), VccQ: 1.8V (1.70-1.95V)",
		"K": "Vcc: 3.3V (2.60-3.60V), VccQ: 3.3V (2.60-3.60V)",
		"L": "Vcc: 2.5V (2.35-3.60V), VccQ: 1.2V (1.14-1.26V)",
	}, "")

	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
	}, -1)

	SetInterface(utils.RetShiftChars(&partNumber, 1), &info)

	info.Toggle = false

	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), MicronPackage, "")

	return info
}

func SetInterface(inter string, info *flashinfo.Flashinfo) {
	interFace := utils.GetOrDefault(inter, map[string][]bool{
		"A": {false, true, false}, //sync, async, spi
		"B": {true, true, false},
		"C": {true, false, false},
		"D": {false, false, true},
		//SpecTek
		"E": {true, true, false},
		"F": {true, true, false},
		"G": {true, true, false},
		"M": {false, false, false},
		"N": {true, true, false},
		"H": {true, true, false},
	}, []bool{false, false, false})

	info.Sync = interFace[0]
	info.Async = interFace[1]
	info.Spi = interFace[2]
}
