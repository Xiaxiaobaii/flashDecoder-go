package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type SpecTekDecoder struct {
}

func (s SpecTekDecoder) Check(partNumber string) bool {
	code := utils.SubOffsetStr(partNumber, 0, 2)
	return utils.Inarray(code, []string{"FN", "FT", "FB", "FX", "CB"})
}

func (s SpecTekDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()

	info.Vendor = "SpecTek"

	info.Type = flashinfo.NAND

	if len(partNumber) == 13 {
		info.Unsupported_Reason = "SpecTek Old Number"
		return info
	}

	utils.RetShiftChars(&partNumber, 3)

	cellL := utils.RetShiftChars(&partNumber, 1)

	info.CellLevel = utils.GetOrDefault(cellL, map[string]string{
		"3": "SLC",
		"M": "SLC",
		"4": "MLC",
		"L": "MLC",
		"B": "TLC",
		"N": "QLC",
		"Q": "QLC",
	}, "")

	info.Process = cellL + utils.RetShiftChars(&partNumber, 3)

	info.Capacity = utils.StartMatchFromMap(&partNumber, MicronCapacitys, "")

	var grade string

	if info.Capacity == "" {
		den := utils.RetShiftChars(&partNumber, 2)
		info.Capacity = utils.GetOrDefault(den, Spec_Legacy_Density, "")
		if info.Capacity == "" {
			info.Capacity = utils.GetOrDefault(den, Spec_NEWER_DENSITY, "")
			if info.Capacity != "" {
				grade = den[1:]
			}
		}
	} else {
		grade = utils.RetShiftChars(&partNumber, 1)
	}

	if grade != "" {
		info.CapGrade = utils.GetOrDefault(grade, map[string]string{
			"1": "94-100%",
			"9": "90-100%",
			"6": "50-90%",
			"5": "40-60%",
			"0": "",
			"A": "",
		}, "")
	}

	config := utils.RetShiftChars(&partNumber, 1)

	if utils.Inarray(config, []string{"G", "P"}) {
		info.EccEnable = true
	}

	// Upstream FlashDetector exposes "halfPageAndSize" in extra info.
	// In this Go port we store it directly on Flashinfo for easier consumption.
	if config == "M" {
		info.HalfPageAndSize = true
	}

	info.DeviceWidth = utils.GetOrDefault(config, map[string]int{
		"G": 8,
		"L": 16,
		"H": 1,
		"M": 8,
		"J": 4,
		"P": 16,
		"K": 8,
		"N": 0,
	}, -1)

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"1": "Vcc: 1.8V, VccQ: 1.8V",
		"2": "Vcc: 2.7V",
		"3": "Vcc: 3.3V, VccQ:3.3V",
		"4": "Vcc: 5.0V",
		"D": "Vcc: 3.3V, VccQ: 1.8V, VssQ: 0V",
		"E": "Vcc: 3.3V, VccQ: 1.8V/3.3V, VssQ: 0V",
		"F": "Vcc: 3.3V, VccQ: 1.2V, VssQ: 0V",
		"J": "Vcc: 3.3V, VccQ: 1.8V/3.3V, VssQ: 0V",
		"L": "Vcc: 2.5V, VccQ: 1.2V, VssQ: 0V",
		"S": "Vcc: 3.3V, VccQ: 3.3V, VssQ: 0V",
		"T": "Vcc: 3.3V, VccQ: 1.8V/1.2V, VssQ: 0V",
	}, "")

	classif := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string][]int{
		"A": {1, 0, 0, 1}, //die, ce, rb, ch
		"B": {1, 1, 1, 1},
		"D": {2, 1, 1, 1},
		"E": {2, 2, 2, 2},
		"F": {2, 2, 2, 1},
		"G": {3, 3, 3, 3},
		"J": {4, 2, 2, 1},
		"K": {4, 2, 2, 2},
		"L": {4, 4, 4, 4},
		"M": {4, 4, 4, 2},
		"Q": {8, 4, 4, 4},
		"R": {8, 2, 2, 2},
		"T": {16, 8, 4, 2},
		"U": {8, 4, 4, 2},
		"V": {16, 8, 4, 4},
		//SpecTek, no R/nB documented
		"C": {3, 3, -1, 2},
		"H": {4, 1, -1, 1},
		"N": {6, 6, -1, 3},
		"P": {8, 8, -1, 2},
		"W": {16, 4, -1, 2},
		"X": {4, 4, -1, 2},
		"Y": {11, 7, -1, 4},
		"1": {16, 2, -1, 1},
		"2": {64, 8, -1, 2},
		"3": {8, 4, -1, 2},
		"4": {4, 4, -1, 1},
		"S": {16, 4, -1, 4},
	}, []int{0, 0, 0, 0})

	info.CE = classif[1]
	info.CH = classif[3]
	info.RB = classif[2]
	info.Die = classif[0]

	// Todo package Funcality Partial Type
	utils.RetShiftChars(&partNumber, 1)

	interFace := utils.RetShiftChars(&partNumber, 1)
	SetInterface(interFace, &info)

	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), MicronPackage, "")

	// Todo SpecTek Interface Info

	// Todo SppedGrade

	return info
}

var Spec_Legacy_Density = map[string]string{
	"1G": "1Gbit",
	"18": "1.8Gbit",
	"2G": "2Gbit",
	"38": "3.8Gbit",
	"4G": "4Gbit",
	"78": "7.8Gbit",
	"8G": "8Gbit",
	"F8": "15.8Gbit",
	"HG": "16Gbit",
	"31": "31Gbit",
	"32": "32Gbit",
	"64": "64Gbit",
	"NX": "128",
	"NY": "256",
	"NZ": "512",
}

var Spec_NEWER_DENSITY = map[string]string{
	"1": "1Gbit",
	"2": "4Gbit",
	"3": "8Gbit",
	"4": "16Gbit",
	"5": "32Gbit",
	"6": "64Gbit",
	"7": "128Gbit",
	"8": "256Gbit",
	"9": "512Gbit",
	"A": "1Tbit",
	"B": "2Tbit",
}

func SpecTekDecoderDefault() SpecTekDecoder {
	return SpecTekDecoder{}
}
