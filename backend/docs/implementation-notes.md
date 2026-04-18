# Implementation Notes (Backend)

## 记录规则
- 仅记录关键决策与权衡，不写流水。

## 记录
- Date: 2026-04-09
- Change: v1 登录注册 + 审批 + 首页骨架
- Context: 需快速建立用户端/管理端最小闭环。
- Options Considered:
  - A: 直接实现复杂 RBAC（放弃，超范围）
  - B: 单管理员 + 用户状态机（采用）
- Decision: 采用单管理员与 `pending->approved` 流转。
- Trade-offs: 功能简单但安全与扩展性需后续补强。
- Impact: API、数据模型、首页访问路径统一为最小闭环。
- Rollback Plan: 回退到仅 health 接口版本。
- Follow-up: 首登强制改密、审计日志、限流。

- Date: 2026-04-09
- Change: 接通真实 service/biz/data 链路（MySQL + Redis）
- Context: 需要把文档骨架推进到可运行后端实现。
- Options Considered:
  - A: 仅继续占位 handler（放弃）
  - B: 直接实现最小业务闭环（采用）
- Decision: 使用 Gin 承载接口，按分层规范组织代码（server/service/biz/data）。
- Trade-offs: Gin 轻量敏捷，满足当前阶段需求。
- Impact: 注册/登录/审批/首页逻辑具备真实数据读写与会话验证。
- Rollback Plan: 回退到骨架版 `cmd/server/main.go`。
- Follow-up: 统一错误中间件。

- Date: 2026-04-09
- Change: 新增待审批用户列表接口 `GET /v1/admin/users/pending`
- Context: 前端管理端审批页面需要先展示待处理用户。
- Options Considered:
  - A: 继续依赖人工输入 userId（放弃）
  - B: 后端提供 pending 列表查询（采用）
- Decision: 在 admin 能力内新增列表接口，仅管理员可访问。
- Trade-offs: 当前只返回最小字段，后续可扩展注册时间等信息。
- Impact: 管理端前端可完成完整审批闭环。
- Rollback Plan: 停用管理端列表功能并回退命令行审批。
- Follow-up: 增加分页、搜索与批量审批。

- Date: 2026-04-09
- Change: MySQL/Redis 数据目录改为宿主机绑定挂载（`~/mysite/data`）
- Context: 容器重启后命名 volume 与环境迁移不够直观，且需要明确可见的数据持久化路径。
- Options Considered:
  - A: 继续使用 Docker named volume（放弃）
  - B: 绑定到宿主机目录 `~/mysite/data`（采用）
- Decision: `docker-compose.yml` 改为 bind mount：MySQL -> `${HOME}/mysite/data/mysql`，Redis -> `${HOME}/mysite/data/redis`。
- Trade-offs: 本地路径更可控且便于备份，但需要注意目录权限与跨机器路径一致性。
- Impact: 重启容器后数据不会丢失，排查和备份流程更直接。
- Rollback Plan: 恢复 named volume 配置。
- Follow-up: 如需团队共享环境，后续补充统一数据目录约定与备份脚本。
