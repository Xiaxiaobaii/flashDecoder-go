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
	var flashID string
	var searchKeyword string
	var searchLimit int
	var summaryMode bool
	var jsonOutput bool

	flag.StringVar(&partNumber, "part", "", "Flash part number to decode (e.g. NW383)")
	flag.StringVar(&partNumber, "pn", "", "Alias of --part")
	flag.StringVar(&flashID, "id", "", "Flash ID to decode (hex, separators allowed)")
	flag.StringVar(&searchKeyword, "search", "", "Search FDB by code/part keyword")
	flag.IntVar(&searchLimit, "limit", 20, "Search result limit for --search mode")
	flag.BoolVar(&summaryMode, "summary", false, "Show FDB summary counts")
	flag.BoolVar(&jsonOutput, "json", false, "Output JSON instead of human-readable text")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s --part <PN> [--json]\n  %s --id <HEX_ID> [--json]\n  %s --search <KW> [--limit 20] [--json]\n  %s --summary [--json]\n\nFlags:\n", os.Args[0], os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintf(os.Stderr, "  %s --part NW383\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --part NW383 --json\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --id EC-D7-94-7E\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --search NW101 --limit 5\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --summary\n", os.Args[0])
	}
	flag.Parse()

	if partNumber == "" && flashID == "" && searchKeyword == "" && !summaryMode {
		flag.Usage()
		os.Exit(2)
	}
	switch {
	case summaryMode:
		summary := flashdecoder.GetFdbSummary()
		printResult(jsonOutput, nil, map[string]any{"summary": summary})
	case searchKeyword != "":
		results := flashdecoder.SearchFdb(searchKeyword, searchLimit)
		printResult(jsonOutput, nil, map[string]any{"keyword": searchKeyword, "results": results})
	case flashID != "":
		info, err := flashdecoder.DecodeID(flashID)
		printResult(jsonOutput, err, info)
		if err != nil {
			os.Exit(1)
		}
	default:
		info, err := flashdecoder.Decode(partNumber)
		printResult(jsonOutput, err, info)
		if err != nil {
			os.Exit(1)
		}
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

func printResult(jsonOutput bool, err error, data any) {
	if jsonOutput {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(map[string]any{
			"ok":    err == nil,
			"error": errString(err),
			"data":  data,
		})
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decode error: %v\n\n", err)
	}
	switch typed := data.(type) {
	case flashinfo.Flashinfo:
		fmt.Print(formatText(typed))
	default:
		out, marshalErr := json.MarshalIndent(typed, "", "  ")
		if marshalErr == nil {
			fmt.Println(string(out))
			return
		}
		fmt.Printf("%v\n", typed)
	}
}
