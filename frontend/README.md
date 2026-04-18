# frontend

## 定位
个人网站前端，承载登录注册、个人主页、技术博客、项目展示与管理后台。

## 当前状态
v2 页面体系已落地：登录/注册/审批 + Home/Blog/Projects + Admin。

## 技术选型
- Vue 3 + TypeScript
- Vue Router 4
- Vite 5

## 快速启动（本地）
1. 安装依赖：`cd frontend && npm install`
2. 启动开发：`npm run dev`
3. 访问：`http://localhost:5173`

## 目录说明
- `src/pages/`: 页面入口。
- `src/components/`: 可复用组件（SiteShell 等）。
- `src/lib/`: API 客户端、认证、内容数据。
- `src/router/`: 路由定义与守卫。
- `src/styles/`: design tokens 和全局样式。
- `spec/`: 页面职责、设计 token、验收与测试。
- `docs/`: 实现笔记与 UI 架构。

## 下一步
1. 接入真实文章/项目后端 API（替换前端静态数据）。
2. 增加统一请求拦截器与错误码映射。
3. 补充目录锚点滚动与全文检索。
