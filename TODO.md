# flashDecoder-go 对齐 TODO（基于 iTXTech/FlashDetector）

更新时间：2026-04-29
对齐基线：`https://github.com/iTXTech/FlashDetector/tree/master/FlashDetector`

## 1) 当前对齐概览
- 已有 PN 解码器（Go）：`Intel/Kioxia/Micron/MicronFpga/Phison/Samsung/SpecTek/SkHynix/YangTze`
- 上游 PN 解码器（PHP）额外存在：`SKHynix3D`、`SKHynixLegacy`、`WesternDigital`、`WesternDigitalShortCode`
- Go 端仍缺主链路能力：`FlashId` 解码入口、`FDB` 合并检索接口、CLI 的 `--id` 与检索模式
- 已知残留 TODO（代码内）：`flashs/skhynix.go`、`flashs/samsung.go` 仍有待确认映射项

## 2) 差距清单（按优先级）
- P0 解码覆盖：补齐 `WesternDigital` 与 `WesternDigitalShortCode`，并明确 `SKHynix3D/Legacy` 的拆分策略（独立 decoder 或并入 `skhynix.go`）
- P0 解码入口：实现 `FlashId` 主链路（类似 `Decode(partNumber)` 的公开 API）
- P1 数据能力：补齐 `FDB` 合并、按 vendor/part 搜索、基础映射查询
- P1 CLI 能力：新增 `--id`，并提供 `search`/`summary` 输出模式
- P1 质量保障：建立与上游样例的 PN/ID 回归测试集，覆盖 WD、SKHynix、SpecTek、Samsung 边界样例

## 3) 已挂任务（Slock）
- task #12：补齐缺失解码器 Phase A（SKHynixLegacy/SKHynix3D/WD/WD Short）
- task #13：实现 FlashId 解码主链路与公开 API
- task #14：补齐 FDB 合并与搜索接口
- task #15：CLI 增加 `--id` 与 `search/summary` 模式
- task #16：增加对齐回归测试集（PN/ID）

## 4) 执行顺序建议
1. 先做 task #12（解码覆盖）
2. 再做 task #13（主链路）
3. 接着 task #14（数据接口）
4. 然后 task #15（CLI）
5. 最后 task #16（回归集扩充并收口）
