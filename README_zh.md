# 🔐 SafeBox
[![GoDoc][1]][2] [![MIT licensed][3]][4]

[English](README.md) | [中文](README_zh.md)

[1]: https://godoc.org/github.com/xtaci/safebox?status.svg
[2]: https://pkg.go.dev/github.com/xtaci/safebox
[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

一个致力于简化密钥管理的统一密钥管理系统，让您的数字资产管理更轻松。

Safebox 的核心理念是通过单个主密钥派生所有子密钥，从而大幅简化密钥备份流程。您只需将**一个**约 **256KB** 的文件备份到可移动存储介质上，例如软盘、U 盘、磁光盘 (MO)、DVD-RAM 等，即可完成所有密钥的备份。

# 核心特性

1. **无限密钥派生**：单个主密钥可以派生无限数量的子密钥，建议每个主密钥最多派生 16,384 个子密钥以确保最佳安全性。
2. **多源熵池**：熵源包括按键输入时序、系统随机数 (/dev/urandom)、启动时间戳、进程 PID 等多维度混合，确保密钥生成的高度随机性。
3. **可插拔导出器架构**：灵活适配不同应用场景，支持区块链钱包、SSH 密钥等多种密钥格式导出。

# 支持的密钥类型

Safebox 目前支持派生以下类型的密钥：

1. **SSH** - 安全外壳协议密钥
2. **Ethereum (以太坊)** - 以太坊区块链钱包
3. **Bitcoin (比特币)** - 比特币区块链钱包
4. **Atom** - Cosmos 生态系统
5. **Band** - Band Protocol
6. **Persistence** - Persistence 链
7. **Kava** - Kava 链
8. **Akash** - Akash Network
9. **Filecoin** - Filecoin 去中心化存储
10. **NEM** - NEM 区块链
11. **Tron (波场)** - 波场区块链

更多导出插件持续开发中，欢迎贡献……

# 安全建议

1. **推荐运行环境**：建议在 OpenBSD 系统上运行，以获得最高的安全性保障。
2. **备份介质选择**：务必将主密钥文件备份到**可移动且可靠的冷存储介质**上，推荐使用 [DVD-RAM](https://en.wikipedia.org/wiki/DVD-RAM) 或 [磁光盘 (MO)](https://en.wikipedia.org/wiki/Magneto-optical_drive) 等长寿命存储介质。

## 常见存储介质寿命对比

| 存储介质 | 预期寿命 | 适合冷存储 |
|------|------|------|
| 磁光盘 (Magneto Optical Disk) | 约 50 年 | ✅ 是 |
| DVD-RAM | 约 30 年 | ✅ 是 |
| 软盘 (Floppy Disk) | 10-20 年 | ✅ 是 |
| 闪存 (SD 卡/U 盘/SSD) | 5-10 年或更久<br/>(取决于写入周期) | ✅ 是 |
| 硬盘 (Hard Disk) | 3-5 年 | ❌ **否** |
| 可刻录 CD/DVD | 2-5 年 | ✅ 是 |

# 文本用户界面 (TUI)

Safebox 采用**复古风格**的文本用户界面 (TUI) 设计，具有以下优势：

- **轻量级部署**：可在树莓派等低功耗设备上运行，打造离线密钥管理终端
- **离线安全**：完全脱离网络环境，确保密钥不被远程窃取
- **文本二维码**：支持通过基于文本的二维码展示和传输密钥，无需图形界面

## 界面预览

![image](https://github.com/user-attachments/assets/cdfb6c13-30f6-4892-ac38-53ae090614d6)
![image](https://github.com/user-attachments/assets/431d773d-b597-4dd0-a7aa-5b7c45a738cb)
![1731123361910](https://github.com/user-attachments/assets/d9cf031d-8ec0-4708-b91b-04bcb4ff9a2f)
![image](https://github.com/user-attachments/assets/be031eb0-2fa7-4b0e-8558-d21667a58f9e)

![image](https://user-images.githubusercontent.com/2346725/116670086-e595f080-a9d1-11eb-92b1-b5724b5e764e.png)

# 项目状态

当前版本：**Beta (测试版)**

> ⚠️ **安全提示**：本项目仍在测试阶段，请在生产环境使用前进行充分测试。建议先使用小额资产进行验证。
