package flashdecoder

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

func GenItoString(info flashinfo.Flashinfo) string {
	if info.Generation > 0 {
		if info.Vendor == "Intel" {
			return utils.IntelGenString[info.Generation-1]
		} else if info.Vendor == "Micron" {
			return utils.MicronGenString[info.Generation-1]
		} else {
			return ""
		}
	} else {
		return ""
	}

}
