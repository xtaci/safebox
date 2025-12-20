# 🔐 SafeBox
[![GoDoc][1]][2] [![MIT licensed][3]][4]

[English](README.md) | [中文](README_zh.md)

[1]: https://godoc.org/github.com/xtaci/safebox?status.svg
[2]: https://pkg.go.dev/github.com/xtaci/safebox
[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

一个旨在让生活更轻松的统一密钥管理系统。

Safebox 的主要目标是简化密钥备份，通过单个主密钥派生其余所有密钥。您只需将**一个**约 **256KB** 的文件备份到可移动磁盘上，例如软盘、U盘、MO 磁盘（磁光盘）、DVD-RAM 等……

# 特性

1. **无限密钥派生**：单个主密钥可以派生无限数量的子密钥，但我们建议每个主密钥最多派生 16384 个子密钥。
2. **多源熵**：熵源来自按键输入、系统熵 (/dev/urandom)、启动时间、进程 PID 等……
3. **可插拔导出器**：适应不同场景，如区块链、SSH 等。

# Safebox 支持派生以下密钥：

1. SSH
2. Ethereum (以太坊)
3. Bitcoin (比特币)
4. Atom
5. Band
6. Persistence
7. Kava
8. Akash
9. Filecoin
10. NEM
11. Tron (波场)

更多可插拔导出插件持续更新中……

# 建议

1. 安装在 OpenBSD 上。
2. 将主密钥文件备份在**可移动且可靠的存储介质**上，例如 [DVD-RAM](https://en.wikipedia.org/wiki/DVD-RAM)、[MO](https://en.wikipedia.org/wiki/Magneto-optical_drive)。

常见存储介质寿命表：

| 存储介质 | 寿命 | 冷存储 |
|------|------|------|
|磁光盘 (Magneto Optical Disk)| 预计 50 年 | 是 |
|DVD-RAM|预计 30 年| 是 |
|软盘 (Floppy disk) | 10-20 年 | 是 |
|闪存 (SD卡/U盘/SSD)|5-10 年或更久 (取决于写入周期)| 是 |
|硬盘 (Hard Disk) | 3-5 年| **否** |
|可刻录 CD/DVD| 2-5 年 | 是|

# 文本用户界面 (TUI)

Safebox 采用了**复古风格**的文本用户界面 (TUI) 设计，因此像树莓派这样的设备可以作为离线存储的密钥管理盒。并且可以通过基于文本的 **二维码** 获取密钥。

![image](https://github.com/user-attachments/assets/cdfb6c13-30f6-4892-ac38-53ae090614d6)
![image](https://github.com/user-attachments/assets/431d773d-b597-4dd0-a7aa-5b7c45a738cb)
![1731123361910](https://github.com/user-attachments/assets/d9cf031d-8ec0-4708-b91b-04bcb4ff9a2f)
![image](https://github.com/user-attachments/assets/be031eb0-2fa7-4b0e-8558-d21667a58f9e)

![image](https://user-images.githubusercontent.com/2346725/116670086-e595f080-a9d1-11eb-92b1-b5724b5e764e.png)

# 状态

Beta (测试版)
