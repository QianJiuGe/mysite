# mysite-front

## 仓库定位
个人网站前端仓库，承载登录、注册、首页骨架与前端规范演进。

## 当前阶段目标
完成 v1 最小页面闭环：`/login`、`/register`、`/pending-approval`、`/home`。

## 技术选型
- Vue 3
- Vue Router
- Vite

## 快速启动（本地）
1. 安装依赖：`npm install`
2. 启动开发：`npm run dev`

## 目录说明
- `spec/`: 页面职责、设计 token、验收与测试场景。
- `src/pages`: 页面入口。
- `src/components`: 可复用组件。
- `src/styles`: token 和全局样式。

## 下一步建议
1. 对接后端真实登录注册审批接口。
2. 增加鉴权守卫与错误态提示。
3. 在首页骨架中逐步注入业务模块。
