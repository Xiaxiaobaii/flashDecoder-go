package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type KioxiaDecoder struct {
}

func KioxiaDecoderDefault() KioxiaDecoder {
	return KioxiaDecoder{}
}

func (k KioxiaDecoder) Check(partNumber string) bool {
	if utils.RetStartBAEq(partNumber, "TH") || utils.RetStartBAEq(partNumber, "TC") {
		return true
	}
	return false
}

func (k KioxiaDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.PartNumber = partNumber
	info.Vendor = "Kioxia"

	utils.RetShiftChars(&partNumber, 2)

	utils.RetShiftChars(&partNumber, 2)

	typ := utils.RetShiftChars(&partNumber, 1)
	info.Type = utils.GetOrDefault(typ, map[string]flashinfo.NandType{
		"N": flashinfo.NAND,
		"D": flashinfo.NAND,
		"T": flashinfo.NAND,
		"L": flashinfo.NAND,
	}, flashinfo.UNKWON_NAND)

	for _, v := range []string{"T", "L"} {
		if v == typ {
			info.Toggle = true
			break
		} else {
			info.Toggle = false
		}
	}

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"V": "3.3V",
		"Y": "1.8V",
		"A": "Vcc: 3.3V, VccQ: 1.8V",
		"B": "Vcc: 3.3V, VccQ: 1.65V-3.6V",
		"D": "Vcc: 3.3V/1.8V, VccQ: 3.3V/1.8V",
		"E": "Vcc: 2.7V-3.6V, VccQ: 2.7V-3.6V/1.7V-1.95V",
		"F": "Vcc: 2.7V-3.6V, VccQ: 3.3V/1.8V (UNOFFICIAL)",
		"G": "Vcc: 2.7V~3.6V, VccQ: 1.8V  (UNOFFICIAL)",
		"J": "Vcc: 2.7V-3.6V, VccQ: 1.14V-1.26V/1.7V-1.95V",
		"K": "Vcc: 2.7V-3.6V, VccQ: 1.14V-1.26V/1.7V-1.95V",
	}, "")

	info.Capacity = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), map[string]string{
		"M8": "256Mbit",
		"M9": "512Mbit",
		"G0": "1Gbit",
		"G1": "2Gbit",
		"G2": "4Gbit",
		"G3": "8Gbit",
		"G4": "16Gbit",
		"GA": "24Gbit",
		"G5": "32Gbit",
		"GB": "48Gbit",
		"G6": "64Gbit",
		"GC": "96Gbit",
		"G7": "128Gbit",
		"GD": "192Gbit",
		"G8": "256Gbit",
		"GE": "384Gbit",
		"G9": "512Gbit",
		"GF": "768Gbit",
		"T0": "1Tbit",
		"T1": "2Tbit",
		"T2": "4Tbit",
		"T3": "8Tbit",
		"TG": "1.5Tbit",
	}, "")

	ep := utils.RetShiftChars(&partNumber, 1)
	info.CellLevel = utils.GetOrDefault(ep, map[string]string{
		"S": "SLC",
		"H": "eSLC", //eSLC
		"D": "MLC",
		"E": "eMLC", //eMLC
		"J": "MLC",
		"C": "MLC",
		"T": "TLC",
		"U": "eTLC", //eTLC
		"V": "TLC",
		"X": "TLC",
		"W": "TLC",
		"F": "QLC", //QLC
	}, "")

	for _, v := range []string{
		"H", "E", "U", "V",
	} {
		if v == ep {
			info.EnterPrise = true
			break
		}
	}

	width := utils.RetShiftChars(&partNumber, 1)
	if utils.Inarray(width, []string{"0", "1", "2", "3", "4", "A", "B", "C", "D", "F"}) {
		info.DeviceWidth = 8
	} else if utils.Inarray(width, []string{"5", "6", "7", "8", "9"}) {
		info.DeviceWidth = 16
	}

	size := utils.GetOrDefault(width, map[string][]string{
		"0": {"4KB", "256KB"},
		"1": {"4KB", "512KB"},
		"2": {">4KB", ">512KB"},
		"3": {"2KB", "128KB"},
		"4": {"2KB", "256KB"},
		"5": {"4KB", "256KB"},
		"6": {"4KB", "512KB"},
		"7": {">4KB", ">512KB"},
		"8": {"2KB", "128KB"},
		"9": {"2KB", "256KB"},
		"A": {"8KB/4KB", "2MB"},
		"B": {"16KB", "8MB"},
		"C": {"16KB 1pl", "4MB"},
		"D": {"16KB 2pl", "4MB"},
		"F": {"16KB 4pl", "4MB"},
	}, []string{"", ""})

	info.Page_size = size[0]
	info.Block_size = size[1]

	info.Process = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "130 nm",
		"B": "90 nm",
		"C": "70 nm",
		"D": "56 nm",
		"E": "43 nm",
		"F": "32 nm",
		"G": "24 nm A-type",
		"H": "24 nm B-type",
		"J": "19 nm/1x",
		"K": "A19 nm/1y",
		"L": "15 nm/1z",
		"2": "BiCS2",
		"3": "BiCS3",
		"4": "BiCS4",
		"M": "BiCS4.5",
		"5": "BiCS5",
	}, "")

	Pack := utils.RetShiftChars(&partNumber, 2)
	if utils.Inarray(Pack, []string{"FT", "TG", "TA"}) {
		info.Package = "TSTOP48"
	} else if utils.Inarray(Pack, []string{"XB", "XG", "BA", "BB"}) {
		info.Package = "BGA"
	} else if utils.Inarray(Pack, []string{"XL", "LA"}) {
		info.Package = "LGA"
	}

	classIfication := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string][]int{
		"0": {1, 1}, //Ch, nCE
		"I": {1, 1},
		"2": {1, 2},
		"K": {1, 2},
		"4": {2, 2},
		"M": {2, 2},
		"7": {1, 4},
		"R": {1, 4},
		"8": {0, 4},
		"S": {0, 4},
		"A": {0, 6},
		"U": {0, 6},
		"B": {0, 8},
		"V": {0, 8},
		"D": {4, 4},
		"E": {4, 8},
	}, []int{-1, -1})

	info.CH = classIfication[0]
	info.CE = classIfication[1]

	p := utils.RetShiftChars(&partNumber, 1)
	detailedPackage := map[string]map[string]string{
		"LGA": {
			"1": "LGA40 (12 x 18 x 0.7)",
			"2": "LGA40 (12 x 18 x 1.15)",
			"3": "LGA40 (12 x 17 x 0.65)",
			"4": "LGA40 (12 x 17 x 1.0)",
			"5": "LGA40 (12 x 17 x 1.04)",
			"6": "LGA40 (13 x 17 x 1.04)",
			"7": "LGA52 (14 x 18 x 1.4)",
			"8": "LGA52 (14 x 18 x 1.04)",
			"9": "LGA52 (14 x 18 x 1.0)",
			"A": "LGA52 (12 x 17 x 1.04/1.0)",
			"B": "LGA52 (12 x 17 x 1.4)",
			"C": "LGA52 (11 x 14 x 0.9)",
		},
		"BGA": {
			"1": "BGA224 (14 x 18 x 1.46)",
			"2": "BGA224 (14 x 18 x 1.46)",
			"3": "BGA60 (8.5 x 13)",
			"4": "BGA60 (9 x 11)",
			"5": "BGA60 (10 x 13)",
			"6": "BGA60 (8.5 x 13)",
			"7": "BGA60 (9 x 11)",
			"8": "BGA60 (10 x 13)",
			"9": "BGA132 (12 x 18 x 1.4)",
			"A": "BGA132 (12 x 18 x 1.85)",
			"B": "BGA224 (14 x 18 x 1.35)",
			"C": "BGA132",
			"D": "BGA132",
			"E": "BGA272",
			"F": "BGA272",
			"G": "BGA272",
			"H": "BGA132",
			"J": "BGA152",
			"K": "BGA152",
			"N": "BGA152",
			"P": "BGA132",
		},
	}

	v, ok := detailedPackage[Pack]
	if ok {
		info.Package = v[p]
	}

	return info

}
