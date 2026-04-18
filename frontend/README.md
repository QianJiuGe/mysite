# frontend

## 定位
个人网站前端，Vue 3 + TypeScript + Vite 构建。

## 当前状态
v2：左侧导航布局，博客（公开）+ 备忘录（需登录）+ 管理端（admin）。

## 技术选型
- Vue 3 + Vue Router 4 + TypeScript
- Vite 5
- marked（Markdown 渲染）

## 页面结构

| 路由 | 组件 | 权限 |
|------|------|------|
| `/` | 重定向到 `/blog` | 公开 |
| `/blog` | BlogListPage | 公开 |
| `/blog/:slug` | BlogPostPage | 公开 |
| `/memo` | MemoPage | 需登录 |
| `/login` | LoginPage | 公开 |
| `/register` | RegisterPage | 公开 |
| `/pending-approval` | PendingApprovalPage | 公开 |
| `/admin` | AdminPage | admin only |

## 布局
- `AppLayout.vue`：统一布局，顶栏 + 左侧导航 + 右侧内容区
- 配色参考 VibeVibe 文档站风格：白底、浅灰边框、紫色 accent

## 快速启动
```bash
cd frontend
npm install
npm run dev
```
