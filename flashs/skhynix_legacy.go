package flashs

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

type SkHynixLegacyDecoder struct{}

func (s SkHynixLegacyDecoder) Check(partNumber string) bool {
	return utils.RetStartBAEq(partNumber, "HY27")
}

func (s SkHynixLegacyDecoder) Decode(partNumber string) flashinfo.Flashinfo {
	info := flashinfo.Default()
	info.Vendor = "SkHynix"
	info.Type = flashinfo.NAND

	utils.RetShiftChars(&partNumber, 4)
	info.Voltage = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_Voltage, "")

	classif := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_ClassIfication, []int{-1, -1, Sk_Small_Block})
	info.CellLevel = utils.IntCellTString(classif[0])
	info.Die = classif[1]

	info.DeviceWidth = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), map[string]int{
		"08": 8,
		"16": 16,
		"32": 32,
	}, -1)

	info.Capacity = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 2), SkHynix_Capacity, "")

	mode := utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_Mode, []int{-1, -1, -1, -1})
	info.CE = mode[0]
	info.RB = mode[1]
	info.CH = mode[3]

	info.Sync = true
	info.Async = true
	info.Toggle = false

	info.Generation = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), Samsung_Generation, 0)

	utils.RetShiftChars(&partNumber, 1)
	info.Package = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), SkHynix_Package, "")

	if partNumber != "" {
		packageMaterial := utils.RetShiftChars(&partNumber, 1)
		info.Lead_Free = utils.Inarray(packageMaterial, []string{"P", "R"})
		info.Halogen_Free = utils.Inarray(packageMaterial, []string{"H", "R"})
	}
	if partNumber != "" {
		info.TempErature = utils.GetOrDefault(utils.RetShiftChars(&partNumber, 1), map[string]string{
			"L": "-40°C ~ +85°C",
			"I": "-40°C ~ +85°C",
			"J": "-25°C ~ +85°C",
			"K": "0°C ~ +70°C",
			"U": "-40°C ~ +105°C",
		}, "")
	}

	return info
}

func SkHynixLegacyDecoderDefault() SkHynixLegacyDecoder {
	return SkHynixLegacyDecoder{}
}
