# Frontend Test Cases (v2)

1. 注册页提交后进入待审批提示流程。
2. 登录成功后跳转首页。
3. 待审批页文本提示准确。
4. 已登录用户访问 `/blog` 时可看到文章列表卡片与面包屑。
5. 已登录用户访问 `/blog/ship-fast-with-guardrails` 时可看到文章正文区与上一篇/下一篇入口。
6. 已登录用户访问 `/projects` 时可看到项目卡片列表。
7. 已登录用户访问 `/projects/edge-observability` 时可看到背景、方案、结果三个章节。
8. 未登录访问 `/home`、`/blog`、`/projects` 会被重定向到 `/login`。
