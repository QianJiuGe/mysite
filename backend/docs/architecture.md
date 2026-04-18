# Backend Architecture (v1)

## 分层
- `server`: HTTP 路由与中间件。
- `service`: 应用服务编排。
- `biz`: 领域规则（注册/审批/登录约束）。
- `data`: MySQL + Redis 访问层。

## 设计约束
1. server 层不直接操作数据库。
2. 领域规则只在 biz 层定义。
3. 数据访问失败统一映射到错误模型。
