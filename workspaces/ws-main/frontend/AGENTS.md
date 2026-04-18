# AGENTS (Frontend)

## 开始前必读顺序
1. `../knowledge/ai/rules.md`
2. `./.rules.md`
3. `./.anti-patterns.md`
4. `./spec/acceptance-criteria.md`
5. `./spec/pages.md`
6. `./spec/design-tokens.md`
7. `./spec/blog-ui.md`
8. `./docs/implementation-notes.md`

## 交付强约束
- 未更新相关 spec，不得开始 UI/交互实现。
- 页面职责变更必须先改 `spec/pages.md`。
- 视觉约束变更必须先改 `spec/design-tokens.md`。
- 每次改动至少补 1 条测试场景到 `spec/testing/test-cases.md`。

## 完成后必须更新
- `docs/implementation-notes.md`：记录关键决策与 trade-off。
- `./.anti-patterns.md`：新增踩坑要及时沉淀。
- `./README.md`：仅在目标或目录职责变化时更新。

## 禁止事项
- 禁止在无 spec 情况下直接堆页面。
- 禁止散落魔法样式值绕过 tokens。
- 禁止未验证就声明完成。
