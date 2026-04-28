# flashDecoder-go
The Go re-implementation of `iTXTech/FlashDetector` (work in progress).

Chinese version: `README.zh-CN.md`

This repository provides:
- A Go **library** for decoding flash part numbers (`flashdecoder.Decode(partNumber)`).
- A small CLI entry in `bin/` as a demo/utility.

## Build & Run

### CLI (demo)

Run directly:

```bash
go run ./bin --part NW383
go run ./bin --part NW383 --json
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
