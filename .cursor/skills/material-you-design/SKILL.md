---
name: material-you-design
description: >-
  Material You (M3) 设计系统规范，用于前端 UI 开发。包含完整的色彩、排版、圆角、阴影、
  按钮、图标按钮、卡片、Chip、表单、布局等规范。当创建或修改前端页面、组件、样式时使用。
---

# Material You (M3) 设计系统

本项目前端遵循 Material Design 3 (Material You) 设计规范。

## 字体

- 主字体：`Google Sans`（备选 `system-ui`）
- 等宽字体：`Roboto Mono`
- 图标：Material Symbols Outlined

## 色彩系统

### Primary
| Token | 值 |
|-------|----|
| Primary | `#6750A4` |
| On-Primary | `#FFFFFF` |
| Primary-Container | `#EADDFF` |
| On-Primary-Container | `#21005D` |

### Secondary
| Token | 值 |
|-------|----|
| Secondary | `#625B71` |
| On-Secondary | `#FFFFFF` |
| Secondary-Container | `#E8DEF8` |
| On-Secondary-Container | `#1D192B` |

### Tertiary
| Token | 值 |
|-------|----|
| Tertiary | `#7D5260` |
| Tertiary-Container | `#FFD8E4` |
| On-Tertiary-Container | `#31111D` |

### Error
| Token | 值 |
|-------|----|
| Error | `#B3261E` |
| Error-Container | `#F9DEDC` |
| On-Error-Container | `#410E0B` |

### Surface 层级
| Token | 值 |
|-------|----|
| Surface | `#FFFBFE` |
| Surface-Container-Lowest | `#FFFFFF` |
| Surface-Container-Low | `#F7F2FA` |
| Surface-Container | `#F3EDF7` |
| Surface-Container-High | `#ECE6F0` |
| Surface-Container-Highest | `#E6E0E9` |
| Surface-Variant | `#E7E0EC` |

### 文本 / 边框
| Token | 值 |
|-------|----|
| On-Surface | `#1C1B1F` |
| On-Surface-Variant | `#49454F` |
| Outline | `#79747E` |
| Outline-Variant | `#CAC4D0` |

## 圆角

| 用途 | 值 |
|------|----|
| 按钮 / Chip | `9999px` |
| 卡片 / FAB | `16px` |
| Extended FAB | `28px` |
| 输入框 | `4px` |
| 图标按钮 | `50%` |

## 阴影 (Elevation)

| 级别 | 值 |
|------|----|
| 0 | 无 |
| 1 | `0 1px 2px rgba(0,0,0,0.3), 0 1px 3px 1px rgba(0,0,0,0.15)` |
| 2 | `0 1px 2px rgba(0,0,0,0.3), 0 2px 6px 2px rgba(0,0,0,0.15)` |
| 3 | `0 4px 8px 3px rgba(0,0,0,0.15), 0 1px 3px rgba(0,0,0,0.3)` |
| 4 | `0 6px 10px 4px rgba(0,0,0,0.15), 0 2px 3px rgba(0,0,0,0.3)` |

卡片 hover 时 Elevation 1 → 2。

## 按钮系统（5 种变体）

| 变体 | CSS 类 | 背景 | 文字 | 高度 |
|------|--------|------|------|------|
| Filled | `.btn-filled` | Primary | On-Primary | 40px |
| Tonal | `.btn-tonal` | Secondary-Container | On-Secondary-Container | 40px |
| Outlined | `.btn-outlined` | transparent + 1px Outline | Primary | 40px |
| Text | `.btn-text` | transparent | Primary | 40px |
| Elevated | `.btn-elevated` | Surface-Container-Low | Primary | 40px |

所有按钮圆角 `9999px`，图标 `18px`，hover 加 Elevation 1。

## 图标按钮（4 种变体）

尺寸 40×40px，图标 24px，圆形 `border-radius: 50%`。

| 变体 | CSS 类 | 背景 | 图标色 |
|------|--------|------|--------|
| Standard | `.icon-btn` | transparent | On-Surface-Variant |
| Filled | `.icon-btn .icon-btn-filled` | Primary | On-Primary |
| Tonal | `.icon-btn .icon-btn-tonal` | Secondary-Container | On-Secondary-Container |
| Outlined | `.icon-btn .icon-btn-outlined` | transparent + 1px Outline | On-Surface-Variant |

用 `<span class="material-symbols-outlined">icon_name</span>` 作为图标。

## FAB

| 类型 | CSS 类 | 尺寸 | 圆角 |
|------|--------|------|------|
| Standard | `.fab` | 56×56px | 16px |
| Extended | `.fab-extended` | auto × 56px | 28px |

背景 Primary-Container，Elevation 3，hover → Elevation 4。

## 卡片（3 种变体）

| 变体 | CSS 类 | 背景 | 边框 |
|------|--------|------|------|
| Elevated | `.card-elevated` | Surface | Elevation 1 |
| Filled | `.card-filled` | Surface-Container-Highest | 无 |
| Outlined | `.card-outlined` | Surface | 1px Outline-Variant |

圆角 `16px`，内边距 `24px`。

## Chip 系统

CSS 类：`.chip` + 变体类。高度 32px，圆角 `9999px`。

| 变体 | 额外类 | 背景 | 文字色 |
|------|--------|------|--------|
| Success | `.chip-success` | `rgba(52,199,89,0.12)` | `#1B873B` |
| Warning | `.chip-warning` | `rgba(255,149,0,0.12)` | `#9E6A03` |
| Error | `.chip-error` | Error-Container | On-Error-Container |
| Tonal | `.chip-tonal` | Secondary-Container | On-Secondary-Container |
| Outlined | `.chip-outlined` | transparent + 1px Outline | On-Surface-Variant |

## 排版

| 样式 | CSS 类 | 字号 / 字重 / 行高 |
|------|--------|-------------------|
| Display Small | `.display-small` | 36px / 500 / 44px |
| Headline Medium | `.headline-medium` | 28px / 400 / 36px |
| Title Large | `.title-large` | 22px / 500 / 28px |
| Title Small | `.title-small` | 14px / 500 / 20px |
| Body Large | `.body-large` | 16px / 400 / 24px |
| Body Medium | `.body-medium` | 14px / 400 / 20px |
| Body Small | `.body-small` | 12px / 400 / 16px |
| Label Large | `.label-large` | 14px / 500 / 20px |
| Label Medium | `.label-medium` | 12px / 500 / 16px |

## 布局

1. **导航抽屉**（280px，可折叠为 80px）：Logo + 菜单 + 用户卡片
2. **顶部应用栏**（64px）：菜单按钮 + 标题 + 操作按钮
3. **内容区**：max-width 960px，padding 24px 32px

## 动效

| 缓动 | 值 |
|------|----|
| Standard | `cubic-bezier(0.2, 0, 0, 1)` |
| Decelerate | `cubic-bezier(0.05, 0.7, 0.1, 1)` |
| Accelerate | `cubic-bezier(0.3, 0, 0.8, 0.15)` |

时长：Short 200ms / Medium 400ms / Long 600ms。

## 交互

- 按钮 hover 加阴影
- 卡片 hover 阴影加深
- 表格行 hover 背景变 Surface-Container
- 图标按钮 hover 8% 透明背景
- 禁用态 opacity 0.38

## 响应式断点

| 宽度 | 变化 |
|------|------|
| ≤1024px | 导航折叠为 80px 图标模式 |
| ≤768px | 导航隐藏，内容全宽 |

## 深色主题变量

```css
.dark-theme {
  --md-surface: #1C1B1F;
  --md-surface-container-lowest: #0F0D13;
  --md-surface-container-low: #1D1B20;
  --md-surface-container: #211F26;
  --md-surface-container-high: #2B2930;
  --md-surface-container-highest: #36343B;
  --md-on-surface: #E6E1E5;
  --md-on-surface-variant: #CAC4D0;
  --md-outline: #938F99;
  --md-outline-variant: #49454F;
}
```

## 参考资源

- [Theme Builder](https://m3.material.io/theme-builder)
- [Material Symbols](https://fonts.google.com/icons)
- [M3 官方文档](https://m3.material.io)
