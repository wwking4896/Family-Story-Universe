# 童話城堡 Fairy Castle 文件索引

本資料夾保存第一階段交付文件，目標是先完成產品與系統設計，不直接進入程式碼實作。

## 文件清單

1. [完整 PRD](product/prd.md)
2. [MVP 功能清單](product/mvp-scope.md)
3. [User Stories](product/user-stories.md)
4. [Acceptance Criteria](product/acceptance-criteria.md)
5. [系統架構設計](architecture/system-architecture.md)
6. [ERD 資料庫設計](architecture/erd.md)
7. [REST API 規格](api/rest-api-spec.md)
8. [AI Story Engine 規格](ai/story-engine.md)
9. [Prompt Template](ai/prompt-templates.md)
10. [QA Test Plan](qa/test-plan.md)
11. [Sprint Backlog](project/sprint-backlog.md)
12. [Repo 結構建議](architecture/repo-structure.md)

## 第一階段決策摘要

- MVP 先採 H5/PWA，不做原生 App。
- Backend 採 Golang，預設 Gin + Clean Architecture。
- Database 採 MySQL，Cache 採 Redis，檔案採 S3 Compatible Storage。
- AI 以 Provider Interface 抽象，不把特定供應商寫死在商業邏輯。
- MVP Story Memory Engine 先用 `story_memories` 標籤檢索，V2 再導入 embedding/vector database。
- 兒童資料、家庭隔離、內容安全與 prompt injection 防護列為 P0 非功能需求。
