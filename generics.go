package flashdecoder

import (
	flashinfo "flashDecoder/info"
	"flashDecoder/utils"
)

func GenItoString(info flashinfo.Flashinfo) string {
	if info.Generation > 0 {
		switch info.Vendor {
		case "Intel":
			return utils.IntelGenString[info.Generation-1]
		case "Micron":
			return utils.MicronGenString[info.Generation-1]
		default:
			return ""
		}
	} else {
		return ""
	}

}
