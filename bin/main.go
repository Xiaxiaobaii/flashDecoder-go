package main

import (
	flashdecoder "flashDecoder"
	"fmt"
)

func main() {
	Data, _ := flashdecoder.Decode("NW383")
	fmt.Printf("Flash Search: \n闪存型号: %v\n闪存品牌: %v\n闪存类型: %v   闪存容量: %v   >闪存制程: %v\nCE数量: %v   Die数量: %v   R/B: %v   闪存位宽: %v\n工作电压: %v\n 闪存ID: %v\n闪存代次: %v\n",
Data.PartNumber, Data.Vendor, Data.Type, Data.Capacity, Data.Process, Data.CE, Data.Die, Data.RB, Data.DeviceWidth, Data.Voltage, Data.IDs, Data.Generation)
}