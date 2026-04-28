package main

import (
	"encoding/json"
	"flag"
	flashdecoder "flashDecoder"
	flashinfo "flashDecoder/info"
	"fmt"
	"os"
)

func main() {
	var partNumber string
	var jsonOutput bool

	flag.StringVar(&partNumber, "part", "", "Flash part number to decode (e.g. NW383)")
	flag.StringVar(&partNumber, "pn", "", "Alias of --part")
	flag.BoolVar(&jsonOutput, "json", false, "Output JSON instead of human-readable text")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s --part <PN> [--json]\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintf(os.Stderr, "  %s --part NW383\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --part NW383 --json\n", os.Args[0])
	}
	flag.Parse()

	if partNumber == "" {
		flag.Usage()
		os.Exit(2)
	}

	info, err := flashdecoder.Decode(partNumber)
	if jsonOutput {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(map[string]any{
			"ok":    err == nil,
			"error": errString(err),
			"data":  info,
		})
	} else {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Decode error: %v\n\n", err)
		}
		fmt.Print(formatText(info))
	}

	if err != nil {
		os.Exit(1)
	}
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func formatText(info flashinfo.Flashinfo) string {
	return fmt.Sprintf(
		"Flash Search:\n"+
			"闪存型号: %v\n"+
			"闪存品牌: %v\n"+
			"闪存类型: %v\n"+
			"闪存容量: %v\n"+
			"闪存制程: %v\n"+
			"CE 数量: %v\n"+
			"Die 数量: %v\n"+
			"R/B: %v\n"+
			"闪存位宽: %v\n"+
			"工作电压: %v\n"+
			"闪存 ID: %v\n"+
			"闪存代次: %v\n",
		info.PartNumber,
		info.Vendor,
		info.Type,
		info.Capacity,
		info.Process,
		info.CE,
		info.Die,
		info.RB,
		info.DeviceWidth,
		info.Voltage,
		info.IDs,
		info.Generation,
	)
}
