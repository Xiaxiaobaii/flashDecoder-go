# flashDecoder-go
The Go implementation of Flash Detector

Chinese version: `README.zh-CN.md`

This repository provides:
- A Go **library** for decoding flash part numbers (`flashdecoder.Decode(partNumber)`).
- A Go **library** for decoding flash IDs (`flashdecoder.DecodeID(idHex)`).
- FDB helper APIs (`flashdecoder.FindFdb`, `flashdecoder.SearchFdb`, `flashdecoder.GetFdbSummary`).
- A small CLI entry in `bin/` as a demo/utility.

## Build & Run

### CLI (demo)

Run directly:

```bash
go run ./bin --part NW383
go run ./bin --part NW383 --json
go run ./bin --id EC-D7-94-7E
go run ./bin --search NW101 --limit 5
go run ./bin --summary
```

Install the CLI:

```bash
go install ./bin
```

Then run (binary name follows the folder name, typically `bin`):

```bash
bin --part NW383
```

## Library Usage

```go
package main

import (
	"fmt"
	flashdecoder "flashDecoder"
)

func main() {
	info, err := flashdecoder.Decode("NW383")
	if err != nil {
		// Unsupported or not found; `info` may still contain partial fields.
		fmt.Println("decode error:", err)
	}

	fmt.Println("Vendor:", info.Vendor)
	fmt.Println("Type:", info.Type)
	fmt.Println("Capacity:", info.Capacity)
}
```

## Project Layout

- `decoder.go`: public API entry (`Decode`)
- `flashs/`: vendor-specific decoders
- `utils/`: helpers + `mdb.json` loader
- `bin/`: CLI demo entry
