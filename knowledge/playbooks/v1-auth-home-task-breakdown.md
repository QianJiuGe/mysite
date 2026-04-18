# Task Breakdown: v1 Auth + Home

> 状态：全部完成 (2026-04-09)

## Task 1: Backend API Contract [DONE]
- 输入: PRD + 业务规则
- 约束: 注册默认 pending，审批前禁止登录
- 完成定义: `spec/api/openapi.yaml` 覆盖 register/login/approve/home

## Task 2: Backend Data & Runtime Baseline [DONE]
- 输入: 数据模型与部署约束
- 约束: MySQL + Redis 本地 docker 可启动
- 完成定义: `migrations/001_init.sql` + `docker-compose.yml` 落地

## Task 3: Frontend Page Skeleton [DONE]
- 输入: 页面职责定义
- 约束: 个人站风格，不做 tab
- 完成定义: `/login` `/register` `/pending-approval` `/home` 可路由访问

## Task 4: Flow Verification [DONE]
- 输入: 前后端 spec 与页面流程
- 约束: 先测试场景后实现
- 完成定义: `spec/testing/test-cases.md` 覆盖关键路径
