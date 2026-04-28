# flashDecoder-go

`flashDecoder-go` 是对 `iTXTech/FlashDetector` 的 Go 语言重写（进行中）。

本仓库提供：
- 一个用于**按闪存型号（part number）解码**的 Go **库**：`flashdecoder.Decode(partNumber)`
- `bin/` 下的一个小型 **CLI** 入口（用于演示/工具化调用）

## 构建与运行

### CLI（演示）

直接运行：

```bash
go run ./bin --part NW383
go run ./bin --part NW383 --json
```

安装 CLI：

```bash
go install ./bin
```

然后运行（可执行文件名通常与文件夹名一致，一般是 `bin`）：

```bash
bin --part NW383
```

## 作为 Go 库使用

```go
package main

import (
	"fmt"
	flashdecoder "flashDecoder"
)

func main() {
	info, err := flashdecoder.Decode("NW383")
	if err != nil {
		// 型号不支持或未匹配到时会返回 error；此时 info 也可能包含部分字段。
		fmt.Println("decode error:", err)
	}

	fmt.Println("Vendor:", info.Vendor)
	fmt.Println("Type:", info.Type)
	fmt.Println("Capacity:", info.Capacity)
}
```

## 项目结构

- `decoder.go`：对外 API 入口（`Decode`）
- `flashs/`：各厂商/系列的型号解码器
- `utils/`：工具函数 + `mdb.json` 加载
- `bin/`：CLI 演示入口

