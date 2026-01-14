package utils

import "strings"

// / 减去字符串前len位并返回被减去的字符串
func Substr(str string, len int) string {
	return str[len:]
}

// / 返回字符串起始字符是否和prifix匹配
func RetStartBAEq(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// 减去str前num位并返回减去的内容
func RetShiftChars(str *string, num int) string {
	if num > len(*str) {
		return ""
	}
	ret := (*str)[:num]
	*str = Substr(*str, num)
	return ret
}

// 从数组中批量与字符串起始位置进行匹配，匹配成功后从字符串中提取并返回
func StartMatchFromArray(str *string, info []string, def string) string {
	for _, i := range info {
		if RetStartBAEq(*str, i) {
			RetShiftChars(str, len(i))
			return i
		}
	}
	return def
}

// 从哈希表Key中批量与字符串起始位置进行匹配，匹配成功后返回匹配的哈希表值
func StartMatchFromMap[T comparable](str *string, info map[string]T, def T) T {
	for k, v := range info {
		if RetStartBAEq(*str, k) {
			RetShiftChars(str, len(k))
			return v
		}
	}
	return def
}

// 将str和哈希表Key值匹配，匹配返回Value，不匹配返回Def
func GetOrDefault[T any](str string, info map[string]T, Def T) T {
	v, ok := info[str]
	if ok {
		return v
	} else {
		return Def
	}
}

//返回字符串中从起始+offset开始的length字符串，无越界检查
func SubOffsetStr(str string, offset int, length int) string {
	ret := str[offset : offset+length]
	// *str = (*str)[:offset+1] + (*str)[offset+length:]
	return ret
}

func Inarray[T comparable](needstr T, array []T) bool {
	for _, v := range array {
		if v == needstr {
			return true
		}
	}
	return false
}

func InMapkey[T any](needstr string, Map map[string]T) bool {
	for k := range Map {
		if needstr == k {
			return true
		}
	}
	return false
}

func IntCellTString(cell int) string {
	if cell == 1 {
		return "SLC"
	} else if cell == 2 {
		return "MLC"
	} else if cell == 3 {
		return "TLC"
	} else if cell == 4 {
		return "QLC"
	}else {
		return ""
	}
}