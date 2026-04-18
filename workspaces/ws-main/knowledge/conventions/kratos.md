# Kratos Conventions (Project Manual)

## 来源
- https://go-kratos.dev/docs/getting-started/start/
- https://go-kratos.dev/docs/getting-started/usage/
- https://github.com/go-kratos/kratos-layout

## 目标
统一本项目使用 Go Kratos 的初始化、目录习惯、代码生成与协作约束。

## 初始化规范
1. 使用官方 CLI：`go install github.com/go-kratos/kratos/cmd/kratos/v2@latest`
2. 新服务基于官方布局：`kratos new <service-name>`
3. 对 proto 变更，统一通过 kratos CLI 和 `go generate ./...` 生成。
4. 仅在 `api/` 定义对外契约，不在 handler 手写重复结构体。

## 目录规范（对齐 kratos-layout 思路）
- `cmd/server`: 程序入口与依赖装配。
- `internal/service`: transport 层服务实现（HTTP/gRPC 绑定）。
- `internal/biz`: 领域规则。
- `internal/data`: 数据访问（MySQL/Redis）。
- `configs`: 配置文件。
- `api`: proto 契约与生成入口。

## 编码规范
1. 业务规则写在 `biz`，禁止写进 `service`。
2. `service` 层只做协议转换与调用编排。
3. 错误码统一映射到 spec 中定义的错误模型。
4. 配置必须可替换，不允许硬编码连接串。
5. proto 为事实源，接口变更先改 proto/spec 再改实现。

## 生成规范
1. 新增接口：先加 proto，再 `kratos proto` 相关命令生成代码。
2. 结构变更后执行 `go generate ./...`，必要时 `wire` 重新注入。
3. 生成代码禁止手改，业务逻辑放在非生成文件中。

## 协作约束
1. AI/开发者在提交前必须同步更新 `spec/api/openapi.yaml` 或 proto。
2. 每次接口变更必须补至少一个测试场景。
3. `implementation-notes.md` 记录 Kratos 相关关键决策与权衡。
