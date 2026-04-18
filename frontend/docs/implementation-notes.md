# Implementation Notes (Frontend)

## 记录规则
- 仅记录关键决策与权衡，不写流水。

## 记录
- Date: 2026-04-09
- Change: 登录/注册/首页骨架（类 QQ 空间）
- Context: 需要快速建立可扩展的首版页面体系。
- Options Considered:
  - A: 直接做复杂首页模块（放弃，超范围）
  - B: 先做结构骨架 + 认证基础页（采用）
- Decision: 先固定结构与路由，再逐步注入业务内容。
- Trade-offs: UI 视觉简化，但演进弹性更高。
- Impact: 页面分层与 token 约束明确。
- Rollback Plan: 回退到仅登录/注册页面。
- Follow-up: 接入真实 API、完善状态与错误展示。

- Date: 2026-04-09
- Change: 前端接入真实后端 API（register/login/home）
- Context: 需要把页面从模拟跳转切换为真实链路。
- Options Considered:
  - A: 保持页面本地模拟（放弃）
  - B: 最小 API client + token 存储 + 路由守卫（采用）
- Decision: 使用 `src/lib/api.ts` + `src/lib/auth.ts` 统一请求和 token 管理。
- Trade-offs: 先采用 localStorage 简化实现，后续再升级更安全会话方案。
- Impact: 登录注册首页流程已与后端真实接口打通。
- Rollback Plan: 恢复页面模拟跳转逻辑。
- Follow-up: 增加统一请求拦截器、错误码映射和自动刷新 token。

- Date: 2026-04-09
- Change: 新增管理端前端（审批 UI）并重做登录页视觉
- Context: 需要在前端直接完成用户审批，并提升登录页可用性与视觉质量。
- Options Considered:
  - A: 继续使用命令行审批（放弃）
  - B: 增加 `/admin` 页面并接入审批接口（采用）
- Decision: 管理端新增用户审批主视图，内容管理/站点设置先留占位；登录页升级为品牌区+表单区布局。
- Trade-offs: 当前管理端功能聚焦审批，复杂内容运营能力后续逐步扩展。
- Impact: 管理员可在前端直接审批注册请求；登录页与首页视觉质量显著提升。
- Rollback Plan: 回退到旧版简单登录和无管理端路由。
- Follow-up: 增加管理员用户搜索、批量审批和内容管理真实接口。

- Date: 2026-04-09
- Change: 首页改版为个人站 IA（Home + Blog + Projects）并落地面包屑导航体系
- Context: 原首页仍是类 QQ 空间骨架，无法承载“个人主页 + 技术博客 + 项目展示”目标。
- Options Considered:
  - A: 在原骨架上局部修补（放弃）
  - B: 先更新 spec，再重建页面结构与视觉 token（采用）
- Decision: 新增 `SiteShell` 统一顶栏和面包屑，新增博客/项目列表与详情页，并以静态内容数据验证信息架构。
- Trade-offs: 当前内容源仍为前端静态数据，后续接入 CMS/后端前需要一次数据层替换。
- Impact: 用户端形成稳定的“首页-列表-详情”导航闭环，页面风格统一且可长期演进。
- Rollback Plan: 回退到 v1 QZone 风格首页和旧路由。
- Follow-up: 接入真实文章/项目 API，补充目录锚点滚动与全文检索。
