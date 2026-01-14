package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type SkHynixDecoder struct {
}

func (s SkHynixDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "H2")
}

func (s SkHynixDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Vendor = "SkHynix"

	if utils.Inarray(utils.RetShiftChars(&partNumber, 3), []string{"H2J", "H2D", "H26", "H23"}) {
		info.Type = flashinfo.CON
		info.Unsupported_Reason = "Skhynix_Unsupport"
		return info
	} else {
		info.Type = flashinfo.NAND
	}

	voltage := utils.RetShiftChars(&partNumber, 1)
	info.Voltage = utils.GetOrDefault(voltage, SkHynix_Voltage, "")
	info.Capacity = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), SkHynix_Capacity, "")
	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"8": 8,
		"6": 16,
		"L": 8,
		"I": 8,
		"D": 8,
	}, -1)

	classif := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_ClassIfication, []int{-1, -1, 0})

	info.CellLevel = utils.IntCellTString(classif[0])

	mode := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_Mode, []int{-1, -1, -1, -1})

	info.CE = mode[0]
	info.RB = mode[1]
	info.Die = classif[1]
	info.CH = mode[3]

	info.Toggle = false
	info.Async = true

	if utils.Inarray(voltage, []string{"Q", "T"}) {
		info.Sync = true
	} else {
		info.Sync = false
	}

	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Samsung_Generation, -1)
	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_Package, "")

	packageMaterial := utils.RetShiftChars(&partNumber, 1)

	if utils.Inarray(packageMaterial, []string{"R", "P"}) {
		info.Lead_Free = true
	} else if packageMaterial == "L" {
		info.Lead_Free = false
	}

	// Todo Halogen Free
	// Todo Wafer

	utils.RetShiftChars(&partNumber, 1)

	// Todo Bad Block
	// Todo Operation

	return info
}

var SkHynix_Voltage = map[string]string{
	"U": "Vcc: 3.3V, VccQ: 3.3V",
	"L": "2.7V",
	"S": "1.8V",
	"J": "2.7V~3.6V/1.2V",
	"Q": "Vcc: 2.7V~3.6V, VccQ: 1.7V~1.95V/2.7V~3.6V",
	"T": "Vcc: 3.3V, VccQ: 1.8V/3.3V",
}

var SkHynix_Capacity = map[string]string{
	"64": "64",
	"25": "256",
	"1G": "1Gbit",
	"4G": "4Gbit",
	"AG": "16Gbit",
	"CG": "64Gbit",
	"12": "128",
	"51": "512",
	"2G": "2Gbit",
	"8G": "8Gbit",
	"BG": "32Gbit",
	"DG": "128Gbit",
	"EG": "256Gbit",
	"FG": "512Gbit",
	"PG": "192Gbit",
	"RG": "384Gbit",
	"VG": "768Gbit",
	"1T": "1Tbit",
	"2T": "2Tbit",
	"4T": "4Tbit",
}

var Sk_Small_Block = 0
var Sk_Large_Block = 1

var SkHynix_ClassIfication = map[string][]int{
	//Type, Die, Block
	"S": {1, 1, Sk_Small_Block},
	"A": {1, 2, Sk_Small_Block},
	"B": {1, 4, Sk_Small_Block},
	"F": {1, 1, Sk_Large_Block},
	"G": {1, 2, Sk_Large_Block},
	"H": {1, 4, Sk_Large_Block},
	"J": {1, 8, Sk_Large_Block},
	"K": {1, 2, Sk_Large_Block}, //Double Stack Package
	"T": {2, 1, Sk_Large_Block},
	"U": {2, 2, Sk_Large_Block},
	"V": {2, 4, Sk_Large_Block},
	"W": {2, 2, Sk_Large_Block}, //Double Stack Package
	"Y": {2, 8, Sk_Large_Block},
	"R": {2, 6, Sk_Large_Block},
	"Z": {2, 12, Sk_Large_Block},
	"C": {2, 16, Sk_Large_Block},
	"M": {3, 1, Sk_Large_Block},
	"N": {3, 2, Sk_Large_Block},
	"P": {3, 4, Sk_Large_Block},
	"Q": {3, 8, Sk_Large_Block},
	"2": {2, 1, Sk_Large_Block},
	"4": {2, 2, Sk_Large_Block},
	"3": {2, 4, Sk_Large_Block},
	"5": {2, 8, Sk_Large_Block},
	"D": {2, 1, Sk_Large_Block},
	"L": {3, 16, Sk_Large_Block},
	//TODO: more
}

var SkHynix_Mode = map[string][]int{
	"1": {1, 1, 1, 1}, //CE, RB, Sync
	"2": {1, 1, 0, 1},
	"4": {2, 2, 1, 1},
	"5": {2, 2, 0, 1},
	"D": {2, 2, 0, 2}, //Dual Interface
	"F": {4, 4, 0, 2}, //Dual Interface
	"T": {5, 5, 0, 1},
	"U": {6, 6, 0, 1},
	"V": {8, 8, 0, 1},
	"M": {4, 1, 1, 2}, //Dual Interface
	"G": {4, 2, 1, 2}, //Dual Interface
	"W": {6, 6, 1, 2}, //Dual Interface
	"H": {8, 8, 1, 2}, //Dual Interface
	"E": {4, 4, 1, 4},
	"Q": {4, 4, 1, 4},
	"A": {4, 4, 1, 2}, //TODO: confirm
}

var SkHynix_Package = map[string]string{
	"T": "TSOP1",
	"V": "WSOP",
	"S": "USOP",
	"N": "LSOP1",
	"F": "FBGA",
	"X": "LGA",
	"M": "WLGA",
	"Y": "VLGA",
	"U": "ULGA",
	"W": "Wafer",
	"C": "PGD1 (chip)",
	"K": "KGD",
	"D": "PGD2 (wafer)",
	"I": "VFBGA-100",
	"J": "LFBGA-100",
	"A": "VLGA",
	"H": "XLGA",
	"8": "FBGA-152",
	"9": "FBGA-152",
	"2": "FBGA-316",
	"3": "FBGA-316",
	//TODO: confirm
	"6": "BGA-132",
	"0": "BGA-132",
	"5": "BGA-132",
	"L": "BGA-132",
	"4": "BGA-132",
}

func SkHynixDecoderDefault() SkHynixDecoder {
	return SkHynixDecoder{}
}
