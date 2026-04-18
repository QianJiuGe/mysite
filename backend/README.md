# backend

## 定位
个人网站后端服务，承载 API 契约、业务规则、数据模型与后端实现。

## 当前状态
v1 最小闭环已完成：注册(pending) -> 管理员审批 -> 登录 -> 首页接口。

## 技术选型
- Go + Gin
- 分层架构（`server/service/biz/data`）
- MySQL + Redis（Docker）

## 快速启动（本地）
1. 启动依赖：`cd backend && docker compose up -d`
2. 启动服务：`go run ./cmd/server -config ./configs/app.example.yaml`
3. 健康检查：`GET http://localhost:8080/healthz`

## API（v1）
- `POST /v1/auth/register`
- `POST /v1/auth/login`
- `GET /v1/admin/users/pending`（需 admin token）
- `POST /v1/admin/users/{userId}/approve`（需 admin token）
- `GET /v1/home`（需 token）

## 默认管理员
- username: `admin` / password: `Admin@123456`
- 仅用于开发初始化。

## 目录说明
- `cmd/server/`: 启动入口。
- `internal/server/`: HTTP 路由与中间件。
- `internal/service/`: 应用服务编排与请求响应处理。
- `internal/biz/`: 登录注册审批核心规则。
- `internal/data/`: MySQL/Redis 存取。
- `internal/model/`: 数据模型。
- `internal/conf/`: 配置加载。
- `spec/`: API、领域与测试规格。
- `migrations/`: 数据库迁移脚本。

## 下一步
1. 增加鉴权中间件、审计日志、限流。
2. 增加接口自动化测试。
