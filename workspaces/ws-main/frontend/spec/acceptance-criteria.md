# Frontend Acceptance Criteria (v2 Personal Site)

1. 提供登录页、注册页、待审批页。
2. 登录后提供 `/home`、`/blog`、`/blog/:slug`、`/projects`、`/projects/:id` 页面。
3. 用户端页面统一具备顶栏导航和面包屑结构。
4. 首页包含个人定位、最近文章、精选项目三个核心区块。
5. 博客列表页展示文章卡片；博客详情页具备高可读正文容器与上一篇/下一篇导航。
6. 项目列表页展示项目卡片；项目详情页包含背景、方案、结果三段式内容。
7. 页面样式遵循 design tokens，不出现明显魔法值堆叠。
8. 页面职责与 `spec/pages.md` 保持一致。
