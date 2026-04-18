# AGENTS (Backend)

## 开始前必读顺序
1. `../knowledge/ai/rules.md`
2. `../knowledge/conventions/kratos.md`
3. `./.rules.md`
4. `./.anti-patterns.md`
5. `./spec/acceptance-criteria.md`
6. `./spec/api/openapi.yaml`
7. `./spec/domain/business-rules.md`
8. `./spec/data/schema.md`
9. `./docs/implementation-notes.md`

## 交付强约束
- 未更新相关 spec，不得修改实现。
- 涉及 API 变更，必须先改 `openapi.yaml`。
- 涉及领域规则变更，必须改 `business-rules.md`。
- 涉及数据结构变更，必须改 `schema.md`。
- 每次变更至少补 1 条测试场景到 `spec/testing/test-cases.md`。

## 完成后必须更新
- `docs/implementation-notes.md`：记录关键决策与 trade-off。
- `./.anti-patterns.md`：若出现新坑，追加条目。
- `./README.md`：仅当目录职责或阶段目标变化时更新。

## 禁止事项
- 禁止先实现后补规格。
- 禁止修改无关模块。
- 禁止未验证就声明完成。
