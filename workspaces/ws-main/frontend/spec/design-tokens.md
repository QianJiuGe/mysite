# Design Tokens (v2 Personal Site)

## 目标
建立可长期演进的视觉系统，支撑“个人主页 + 技术博客 + 项目展示”，风格要求简洁、专业、耐看。

## Token 分类
- Color:
  - `color-bg-page`
  - `color-bg-elevated`
  - `color-surface`
  - `color-border`
  - `color-text-primary`
  - `color-text-secondary`
  - `color-accent`
  - `color-accent-soft`
- Typography:
  - `font-family-sans`
  - `font-family-serif`
  - `font-size-xs|sm|md|lg|xl|2xl|3xl`
  - `line-height-tight|base|relaxed`
- Spacing:
  - `space-4|8|12|16|20|24|32|40|56|72`
- Radius:
  - `radius-sm|md|lg|xl`
- Shadow:
  - `shadow-sm|md|lg`
- Layout:
  - `container-max`

## 组件级约束
1. 顶栏、卡片、正文容器统一使用 token，不允许散落魔法值。
2. 卡片默认采用“浅底 + 细边框 + 中小圆角 + 轻阴影”的稳定组合。
3. 标题层级严格控制：页面标题 > 分区标题 > 卡片标题 > 元信息。
4. 博客正文宽度与行高必须优先保证可读性，不追求满屏展示。
5. 面包屑在内容页固定位置，颜色与大小低于主标题，不与主信息竞争。
