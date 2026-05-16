# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目定位

这是个人网站的 monorepo，根目录负责统一协作入口：

- `knowledge/`: 跨前后端共享的规则、产品规划、系统架构、约定和 playbooks。
- `backend/`: Go + Gin API 服务，承载 API 契约、业务规则、数据模型和后端实现。
- `frontend/`: Vue 3 + TypeScript + Vite 前端。
- `myresource/blog/`: 博客 Markdown 文件存储。

开始修改前按 `WORKSPACE.md` 的阅读优先级处理：先读 `knowledge/ai/rules.md`，再读目标目录的 `AGENTS.md`、`.rules.md`、`.anti-patterns.md`、相关 `spec/` 和 `docs/implementation-notes.md`。

## 常用命令

### 一键本地启动

```bash
./start-dev.sh
```

该脚本会启动 `backend/docker-compose.yml` 中的 MySQL/Redis，运行后端 `go run ./cmd/server -config ./configs/app.example.yaml`，并以 `0.0.0.0:5173` 启动前端。日志与 PID 写入 `.runtime/`。

### 后端

```bash
cd backend && docker compose up -d
cd backend && go run ./cmd/server -config ./configs/app.example.yaml
cd backend && go build ./cmd/server
cd backend && go test ./...
cd backend && go test ./internal/biz -run TestName
```

如果系统默认 `go` 不可用，本仓库的 Cursor 规则使用：

```bash
export GOROOT="$HOME/.local/go" && export PATH="$GOROOT/bin:$PATH"
```

后端健康检查：

```bash
curl --noproxy '*' -fsS http://localhost:8080/healthz
```

### 前端

```bash
cd frontend && npm install
cd frontend && npm run dev
cd frontend && npm run build
cd frontend && npm run preview
```

`frontend/package.json` 当前只定义了 `dev`、`build`、`preview`，没有 lint 或测试脚本；若添加测试框架，同步补充对应脚本和单测运行方式。

### 修改后的验证流程

`.cursor/rules/dev-workflow.mdc` 要求每次代码修改后同时验证前后端：

```bash
cd backend && go build ./cmd/server
cd frontend && npm run build
```

涉及运行时行为时，重启服务后做 smoke check（例如 `/healthz` 或关键 API）。后端 API 路径、参数或返回值变化时，必须同步检查 `frontend/src/lib/api.ts` 和相关页面；前端调用变化也要反查后端契约。

## 后端架构

后端入口是 `backend/cmd/server/main.go`：加载 `configs/app.example.yaml`，初始化数据存储，创建 auth/blog/memo usecase，组装 `service.Service`，再由 `internal/server/http.go` 注册 Gin 路由。

分层约束来自 `backend/.rules.md` 和 `backend/docs/architecture.md`：

```text
server -> service -> biz -> data
```

- `server`: HTTP 路由、CORS、鉴权与 admin 中间件。
- `service`: Gin handler，负责请求绑定、响应格式和错误码映射。
- `biz`: 领域规则；注册审批、登录、博客 slug 校验、备忘录归属校验都在这里。
- `data`: MySQL/Redis 访问层；业务层通过 store 读写用户、会话和备忘录。

API 契约的单一事实源是 `backend/spec/api/openapi.yaml`。修改 API 前先更新契约；涉及领域规则时更新 `backend/spec/domain/business-rules.md`；每次行为变更至少补充 `backend/spec/testing/test-cases.md` 中的验收场景。

主要业务边界：

- 注册用户初始状态为 `pending`，只能由 `admin` 审批为 `approved`。
- 默认开发管理员来自配置：`admin / Admin@123456`。
- 会话通过登录后返回的 bearer token 访问受保护接口；中间件将 session 写入 Gin context。
- 博客公开读，admin 可写；文章是 `myresource/blog/*.md`，frontmatter 提供标题、摘要、标签和日期。
- 备忘录需要登录；普通用户只能操作自己的 memo，admin 列表可见全部。

## 前端架构

前端入口是 `frontend/src/main.ts`，挂载 Vue app、Vue Router 和全局样式。路由定义在 `frontend/src/router/index.ts`，通过 route meta 和 `frontend/src/lib/auth.ts` 中的 token/role 做登录与 admin 守卫。

API 调用集中在 `frontend/src/lib/api.ts`：

- `VITE_API_BASE` 可覆盖后端地址。
- 默认 API base 是当前页面协议和 hostname 加 `:8080`，用于 LAN 访问。
- 需要登录的请求由 `request(..., true)` 自动附加 `Authorization: Bearer <token>`。

UI 分层来自 `frontend/docs/ui-architecture.md`：

- `pages`: 路由页面，负责组合页面逻辑。
- `components`: 可复用 UI 组件；当前主布局是 `components/AppLayout.vue`。
- `styles`: 设计 token 与全局样式；颜色、间距、字号等视觉约束集中在 `frontend/src/styles/tokens.css`，不要在新组件里绕过 token 散落魔法值。

页面职责和视觉约束分别以 `frontend/spec/pages.md`、`frontend/spec/design-tokens.md`、`frontend/spec/blog-ui.md` 为准。修改 UI/交互前先更新相关 spec，并在 `frontend/spec/testing/test-cases.md` 补充可验收场景。

## 文档与规格约束

- 全局 AI 规则在 `knowledge/ai/rules.md`：先规格后实现、先验收场景后实现、未验证不得声称完成。
- 跨端 API 约定在 `knowledge/conventions/api.md`：错误码必须可枚举、可解释、可测试。
- 关键决策与 trade-off 记录到目标目录的 `docs/implementation-notes.md`。
- 只有目录职责或阶段目标变化时才更新对应 README；普通实现细节不要膨胀 README。
