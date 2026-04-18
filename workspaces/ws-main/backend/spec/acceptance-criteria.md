# Backend Acceptance Criteria (v1 Auth + Home)

1. 注册成功后用户状态为 `pending`。
2. `pending` 用户登录返回 `FORBIDDEN_PENDING_APPROVAL`。
3. 管理员可审批 `pending` 用户。
4. 审批后用户可登录并获取 token。
5. 首页接口对已登录用户返回基础信息。
6. API 语义、错误码与 `openapi.yaml` 一致。
