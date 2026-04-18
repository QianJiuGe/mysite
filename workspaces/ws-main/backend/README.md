# mysite-back

## 仓库定位
个人网站后端服务仓库，承载 API 契约、业务规则、数据模型与后端实现。

## 当前阶段目标
完成 v1 最小闭环：注册（pending）-> 管理员审批 -> 登录 -> 首页接口。

## 技术选型
- Go
- Kratos 规范分层（`server/service/biz/data`）
- MySQL + Redis（Docker）

## 快速启动（本地）
1. 启动依赖：`docker compose up -d`
2. 启动服务：`go run ./cmd/server -config ./configs/app.example.yaml`
3. 健康检查：`GET http://localhost:8080/healthz`

## API（v1）
- `POST /v1/auth/register`
- `POST /v1/auth/login`
- `POST /v1/admin/users/{userId}/approve`（需要 `Authorization: Bearer <token>` 且 admin）
- `GET /v1/home`（需要 `Authorization: Bearer <token>`）

## 默认管理员
- username: `admin`
- password: `Admin@123456`
- 仅用于开发初始化，后续需改密策略。

## 目录说明
- `cmd/server`: 启动入口。
- `internal/service`: HTTP 接口编排与请求响应处理。
- `internal/biz`: 登录注册审批核心规则。
- `internal/data`: MySQL/Redis 存取。
- `spec/`: API、领域、数据与测试规格。

## 下一步建议
1. 切换到 proto + Kratos transport（按 `knowledge/conventions/kratos.md`）。
2. 增加鉴权中间件、审计日志、限流。
3. 增加接口自动化测试。
