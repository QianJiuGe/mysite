# Backend Test Cases (v1)

## 用例
1. 注册成功
- 输入：合法 username/password
- 期望：201，返回 `pending`

2. 待审批登录失败
- 前置：用户状态 `pending`
- 输入：正确用户名密码
- 期望：403，`FORBIDDEN_PENDING_APPROVAL`

3. 管理员审批成功
- 前置：存在 `pending` 用户；请求者为 admin
- 期望：200，用户状态 `approved`

4. 非管理员审批失败
- 前置：请求者为普通用户
- 期望：403，`FORBIDDEN_ADMIN_ONLY`

5. 管理员拉取待审批列表
- 前置：请求者为 admin
- 期望：200，返回 `users[]`

6. 审批后登录成功
- 前置：用户状态 `approved`
- 期望：200，返回 token

7. 首页接口鉴权
- 未登录：401
- 已登录 approved 用户：200
