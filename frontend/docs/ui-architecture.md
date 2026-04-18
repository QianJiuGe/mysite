# UI Architecture (v1)

## 分层
- `pages`: 路由页面。
- `components`: 可复用 UI 组件。
- `styles`: token 与全局样式。

## 约束
1. 页面仅负责组合，不堆积视觉细节。
2. 通用卡片、布局组件优先复用。
3. 样式变量集中在 `styles/tokens.css`。
