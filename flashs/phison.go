package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type PhisonDecoder struct {
}

var Phison_REBRAND_VENDOR = map[string]string{
	"T": "Kioxia",
	"I": "Micron",
	"K": "Micron",
	"H": "Skhynix",
	"D": "Western Digital",
	"C": "Yangtze",
	"N": "Intel",
}

var Phison_PACKAGE = map[string]string{
	"A": "BGA132",
	"P": "BGA152",
	"C": "BGA272",
	"O": "LGA60-SAT",
	"K": "LGA60-SAT",
	"F": "TSOP48",
	"T": "TSOP48",
	"B": "BGA132 / LGA110",
	"Y": "BGA56-UFS",
}

var Phison_CLASSIFICATION = map[string][]int{
	"1": {1, 1}, //ce die
	"3": {1, 2},
	"5": {2, 2},
	"6": {2, 4},
	"7": {4, 4},
	"8": {4, 8},
	"A": {4, 16},
	"B": {8, 8},
	"C": {8, 16},
}

var Phison_DENSITY = map[string]string{
	"7G": "16Bits",
	"8G": "32GBits",
	"9G": "64GBits",
	"AG": "128GBits",
	"BG": "256GBits",
	"EG": "384GBits",
	"HG": "512GBits",
	"IG": "1TBits",
	"JG": "2TBits",
}

var Phison_PROCESS_NODE = map[string]map[string][]string{
	"Kioxia": {
		"H": {"24nm 2plane", "MLC"}, // process node, cell
		"P": {"A19nm 4plane", "MLC"},
		"R": {"15nm", "TLC"},
		"S": {"15nm 2plane", "MLC"},
		"U": {"15nm 4plane", "MLC"},
		"V": {"BiCS2", "TLC"},
		"I": {"BICS3", "TLC"},
		"W": {"BiCS4", "TLC"},
		"X": {"BiCS4.5", "TLC"},
		"Y": {"BiCS5", "TLC"},
	},
	"Intel": {
		"N": {"20nm", "MLC"},
		"P": {"16nm", "MLC"},
		"O": {"L06B/B16A/N18A 32L", ""},
		"V": {"B16A/B27A", "TLC"},
		"I": {"B27B", "TLC"},
		"X": {"B37R", "TLC"},
		"Y": {"B47R", "TLC"},
	},
	"Skhynix": {
		"P": {"16nm", "TLC"},
		"O": {"3DV4", "TLC"},
	},
	"Yangtze": {
		"O": {"JGS", "TLC"},
	},
}

func PhisonDecoderDefault() PhisonDecoder {
	return PhisonDecoder{}
}

func (p PhisonDecoder) Check(partNumber string) bool {
	if len(partNumber) == 10 && utils.InMapkey(partNumber[0:1], Phison_REBRAND_VENDOR) &&
		utils.InMapkey(partNumber[1:2], Phison_PACKAGE) {
		return true
	}
	return false
}

func (p PhisonDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Type = flashinfo.NAND

	info.Vendor = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Phison_REBRAND_VENDOR, "")

	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Phison_PACKAGE, "")

	clz := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Phison_CLASSIFICATION, []int{-1, -1})
	info.CE = clz[0]
	info.Die = clz[1]

	info.Capacity = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), Phison_DENSITY, "")

	utils.RetShiftChars(&partNumber, 3)

	vendor := info.Vendor
	switch vendor {
	case "Micron":
		vendor = "Intel"
	case "Western Digital":
		vendor = "Kioxia"
	}

	processs, ok := Phison_PROCESS_NODE[vendor]
	if ok {
		node := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), processs, []string{"", ""})
		info.CellLevel = node[1]
		info.Process = node[0]
	}

	info.DeviceWidth = 8

	return info
}
