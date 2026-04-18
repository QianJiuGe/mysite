# ADR-0001: Auth + Home Bootstrap

- Status: Accepted
- Date: 2026-04-09

## Context
首轮要建立登录/注册/审批与首页最小闭环，且保持可扩展。

## Decision
1. 采用单管理员模型。
2. 新用户默认 pending，管理员审批后可登录。
3. 首页先交付类 QQ 空间结构骨架，不做 tab。
4. 后端使用 Gin 框架，按分层规范演进。

## Consequences
- 可快速启动开发并保持文档驱动。
- 需要后续补强：首登改密、审计、限流、真实鉴权中间件。
