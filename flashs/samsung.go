package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type SamsungDecoder struct {
}

func (s SamsungDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "K9")
}

func (s SamsungDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()

	info.Vendor = "Samsung"
	info.Type = flashinfo.NAND

	utils.RetShiftChars(&partNumber, 2)

	classIfication := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), ClassIfication, []int{-1, -1})

	info.CellLevel = utils.IntCellTString(classIfication[0])

	info.Die = classIfication[1]

	info.Capacity = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), Samsung_Capacitys, "")

	technology := utils.RetShiftChars(&partNumber, 1)

	info.ToggleN = utils.GetOrDefault(technology, map[string]string{
		"D": "1.0",
		"Y": "2.0",
		"B": "3.0",
	}, "")

	info.Toggle = utils.Inarray(technology, []string{"D", "Y", "B"})

	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"0": 0,
		"8": 8,
		"6": 16,
	}, -1)

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "1.65V~3.6V",
		"B": "2.7V (2.5V~2.9V)",
		"C": "5.0V (4.5V~5.5V)",
		"D": "2.65V (2.4V ~ 2.9V)",
		"E": "2.3V~3.6V",
		"R": "1.8V (1.65V~1.95V)",
		"Q": "1.8V (1.7V ~ 1.95V)",
		"T": "2.4V~3.0V",
		"S": "Vcc: 3.3V (3V~3.6V), VccQ: 1.8V (1.65V~1.95V)",
		"U": "2.7V~3.6V",
		"V": "3.3V (3.0V~3.6V)",
		"W": "2.7V~5.5V, 3.0V~5.5V",
		"0": "None",
		"H": "Vcc: 3.3V, VccQ: 1.8V (UNOFFICIAL)", //TODO: Confirm
		// Upstream FlashDetector also leaves this as TODO; keep a non-empty placeholder
		// so CLI output is stable and can be improved later once confirmed.
		"J": "Vcc: 3.3V, VccQ: 1.8V (UNOFFICIAL)",
	}, "")

	mode := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string][]int{
		"0": {1, 1}, //CE, R/nB
		"1": {2, 2},
		"3": {3, 3},
		"4": {4, 1},
		"5": {4, 4},
		"6": {6, 2},
		"7": {8, 4},
		"8": {8, 2},
		"9": {-1, -1}, //1st block OTP
		"A": {-1, -1}, //Mask Option 1
		"L": {-1, -1}, //Low grade
		"C": {16, 2},
		"J": {16, 4},
	}, []int{-1, -1})

	info.CE = mode[0]
	info.RB = mode[1]

	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Samsung_Generation, 0)

	utils.RetShiftChars(&partNumber, 1)

	pack := utils.RetShiftChars(&partNumber, 1)
	info.Package = utils.GetOrDefault(pack, map[string]string{
		"8": "TSOP1",
		"9": "56-TSOP1",
		"A": "COB",
		"B": "FBGA",
		"C": "BGA316",
		"D": "TBGA63 or 316",
		"E": "ISM",
		"F": "WSOP",
		"G": "FBGA",
		"H": "BGA132 or 136",
		"I": "ULGA (12*17)",
		"J": "FBGA",
		"K": "ULGA (12*17)",
		"L": "ULGA (14*18)",
		"M": "ULGA52 (13*18)",
		"P": "TSOP1",
		"Q": "TSOP2",
		"R": "56-TSOP1",
		"S": "TSOP1",
		"T": "WSOP",
		"U": "COB (MMC)",
		"V": "WSOP",
		"W": "Wafer",
		"Y": "TSOP1",
		"Z": "WELP",
		"X": "BGA108",
		"1": "BGA108",
	}, "")

	info.Lead_Free = utils.Inarray(pack, []string{"8", "9", "B", "E", "F", "I", "J", "K", "L", "M", "P", "Q", "R", "S", "T", "Z"})

	return info
}

var ClassIfication = map[string][]int{
	//CellLevel, Die
	"3": {4, 1},
	"9": {4, 8},
	"A": {3, 1},
	"B": {3, 2},
	"C": {3, 4},
	"D": {3, 16},
	"F": {1, 1},
	"G": {2, 1},
	"H": {2, 4},
	"K": {1, 2},
	"L": {2, 2},
	"M": {2, 2},
	"N": {1, 2},
	"O": {3, 8},
	"P": {2, 8},
	"Q": {1, 8},
	"R": {2, 12},
	"S": {2, 6},
	"T": {1, 1},
	"U": {2, 16},
	"V": {1, 16},
	"W": {1, 4},
	"X": {4, -1}, //QLC ?die
}

var Samsung_Generation = map[string]int{
	"M": 1,
	"A": 2,
	"B": 3,
	"C": 4,
	"D": 5,
	"E": 6,
	"F": 7,
	"G": 8,
	"H": 9,
	"Y": 25,
	"Z": 26,
}

var Samsung_Capacitys = map[string]string{
	"12": "512",
	"16": "16",
	"28": "128",
	"32": "32",
	"40": "4",
	"56": "256",
	"64": "64",
	"80": "8",
	"1G": "1Gbit",
	"2G": "2Gbit",
	"4G": "4Gbit",
	"8G": "8Gbit",
	"AG": "16Gbit",
	"BG": "32Gbit",
	"CG": "64Gbit",
	"DG": "128Gbit",
	"EG": "256Gbit",
	"FG": "256Gbit",
	"GG": "384Gbit",
	"HG": "512Gbit",
	"LG": "24Gbit",
	"NG": "96Gbit",
	"ZG": "48Gbit",
	"PG": "171Gbit",
	"QG": "341Gbit",
	"RG": "683Gbit",
	"SG": "1365Gbit",
	"KG": "1Tbit",
	"MG": "2Tbit",
	"UG": "4Tbit",
	"VG": "8Tbit",
	"00": "0Tbit",
}

func SamsungDecoderDefault() SamsungDecoder {
	return SamsungDecoder{}
}
