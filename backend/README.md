# backend

## 定位
个人网站后端服务，承载 API 契约、业务规则、数据模型与后端实现。

## 当前状态
v2：注册审批 + 博客（本地 md 文件）+ 备忘录（MySQL） + 鉴权中间件。

## 技术选型
- Go + Gin
- 分层架构（`server/service/biz/data`）
- MySQL + Redis（Docker）

## 快速启动（本地）
1. 启动依赖：`cd backend && docker compose up -d`
2. 启动服务：`go run ./cmd/server -config ./configs/app.example.yaml`
3. 健康检查：`GET http://localhost:8080/healthz`

## API（v1）

### 公开
- `POST /v1/auth/register`
- `POST /v1/auth/login`
- `GET /v1/blog/posts` — 博客列表
- `GET /v1/blog/posts/:slug` — 博客详情

### 需登录
- `GET /v1/home`
- `GET /v1/memos` — 当前用户备忘录列表
- `POST /v1/memos` — 创建备忘录
- `PUT /v1/memos/:id` — 更新备忘录
- `DELETE /v1/memos/:id` — 删除备忘录

### 需 admin
- `GET /v1/admin/users/pending`
- `POST /v1/admin/users/:userId/approve`
- `POST /v1/blog/posts` — 创建博客文章
- `PUT /v1/blog/posts/:slug` — 更新博客文章
- `DELETE /v1/blog/posts/:slug` — 删除博客文章

## 默认管理员
- username: `admin` / password: `Admin@123456`
- 仅用于开发初始化。

## 目录说明
- `cmd/server/`: 启动入口。
- `internal/server/`: HTTP 路由与鉴权中间件。
- `internal/service/`: 应用服务编排与请求响应处理。
- `internal/biz/`: 核心业务规则（认证、博客、备忘录）。
- `internal/data/`: MySQL/Redis 存取。
- `internal/model/`: 数据模型。
- `internal/conf/`: 配置加载。
- `migrations/`: 数据库迁移脚本。
- `../myresource/blog/`: 博客 markdown 文件存储。
