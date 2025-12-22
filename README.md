# gorogio

## 项目目的

gorogio 是一个 **Go + Gio 的最小实验场**，用于训练“可推理 GUI 世界”的能力。

- **核心目标**：学习即时模式 GUI（Immediate Mode UI）思维
- **适合人群**：Go 开发者、想理解 GUI 事件/状态循环的人
- **实验性质**：不追求控件丰富度或商业级产品，可随意重构和扩展

## Gio 哲学定位

- **名字**：Gio = Go + I/O 或 Go Immediate Mode
- **思维哲学**：每帧独立计算、状态集中、事件透明、函数化
- **设计哲学**：小核心库、函数化布局、可组合控件、实验/认知优先
- **定位**：GUI 认知训练工具，而非生产级控件库

## 快速运行

```bash
git clone git@github.com:rocwg/gorogio.git
cd gorogio
go mod tidy
go run .
