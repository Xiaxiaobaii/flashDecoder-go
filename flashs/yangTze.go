package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
	"strconv"
)

type YangTzeDecoder struct {
}

func (y YangTzeDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "YM")
}

func (y YangTzeDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Vendor = "YangTze"

	utils.RetShiftChars(&partNumber, 2)

	if utils.RetShiftChars(&partNumber, 1) == "N" {
		info.Type = flashinfo.NAND
	}

	dieSize := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), Yang_Die_Size, []string{"", ""})
	info.CellLevel = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"S": "SLC",
		"M": "MLC",
		"T": "TLC",
		"Q": "QLC",
	}, "")

	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
		"A": "Vcc=2.7~3.6V; VccQ=3.3V",
		"B": "Vcc=2.7~3.6V; VccQ=1.8V",
		"C": "Vcc=2.35~3.6V; VccQ=1.2V",
		"D": "Vcc=2.35~3.6V; VccQ=3.3/1.8V",
		"E": "Vcc=2.35~3.6V; VccQ=1.8/1.2V",
	}, "")

	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"1": 8,
		"2": 16,
	}, -1)

	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), map[string]string{
		"W0": "Wafer",
		"T1": "TSOP-48; 12 x 20 x 1.x",
		"B1": "BGA-132; 12 x 18 x 1.x",
		"B2": "BGA-152; 14 x 18 x 1.x",
		"B3": "BGA-272; 14 x 18 x 1.x",
		"L1": "LGA-52; 12 x 17 x 1.x",
	}, "")

	classIf := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), YangTzeClassIfIcation, []int{-1, -1, -1, -1})
	info.CE = classIf[1]
	info.CH = classIf[3]
	info.RB = classIf[2]
	info.Die = classIf[0]

	if info.Die > 0 {
		size, _ := strconv.Atoi(dieSize[0])
		info.Capacity = strconv.Itoa(size*info.Die) + dieSize[1]
	} else {
		info.Capacity = dieSize[0] + dieSize[1]
	}

	additionalInfo := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), map[string][]string{
		"C0": {"0°C ~ 70°C", "ONFI 3.2; Max Speed=400MB/s"},
		"C1": {"0°C ~ 70°C", "ONFI 3.2; Max Speed=533MB/s"},
		"C2": {"0°C ~ 70°C", "ONFI 4.0; Max Speed=667MB/s"},
		"C3": {"0°C ~ 70°C", "ONFI 4.0; Max Speed=800MB/s"},
		"C4": {"0°C ~ 70°C", "ONFI 4.0; Max Speed=1200MB/s (UNOFFICIAL)"},
		"C5": {"0°C ~ 70°C", "ONFI 4.0; Max Speed=1333MB/s (UNOFFICIAL)"},
		"C6": {"0°C ~ 70°C", "ONFI 4.0; Max Speed=1600MB/s (UNOFFICIAL)"},
	}, []string{"", ""})

	if additionalInfo[0] != "" {
		info.TempErature = additionalInfo[0]
		info.SpeedGrade = additionalInfo[1]
	}

	info.Toggle = false
	info.Async = true
	info.Sync = true

	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}, -1)

	return info
}

var Yang_Die_Size = map[string][]string{
	"06": {
		"64",
		"Gbits",
	},
	"07": {
		"128",
		"Gbits",
	},
	"08": {
		"256",
		"Gbits",
	},
	"09": {
		"512",
		"Gbits",
	},
	"10": {
		"1",
		"Tbits",
	},
}

var YangTzeClassIfIcation = map[string][]int{
	"A": {1, 1, 1, 1}, //Die, CE, Rb, Ch
	"B": {2, 1, 1, 1},
	"C": {2, 2, 2, 1},
	"D": {2, 2, 2, 2},
	"E": {4, 2, 2, 1},
	"F": {4, 4, 4, 1},
	"G": {4, 2, 2, 2},
	"H": {4, 4, 4, 2},
	"Q": {4, 4, 4, 4},
	"I": {8, 2, 2, 2},
	"J": {8, 4, 4, 2},
	"K": {8, 4, 4, 4},
	"L": {8, 8, 4, 4},
	"R": {8, 8, 4, 2},
	"M": {16, 4, 4, 2},
	"N": {16, 4, 4, 4},
	"O": {16, 8, 4, 2},
	"P": {16, 8, 4, 4},
}

func YangTzeDecoderDefault() YangTzeDecoder {
	return YangTzeDecoder{}
}
