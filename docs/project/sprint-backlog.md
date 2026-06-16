# Sprint Backlog

## Sprint 0：專案初始化

- 建立 monorepo 結構。
- Backend 初始化：Go module、Gin、config、logger、health check。
- Frontend 初始化：Next.js、TypeScript、Tailwind、PWA 基礎設定。
- Docker Compose：MySQL、Redis、backend、frontend。
- 建立 `.env.example`。
- 建立 migration 工具。
- 建立 README 開發與部署說明。

## Sprint 1：使用者與家庭

- Auth：register/login/logout/me。
- 密碼雜湊與 JWT middleware。
- Families：create/get/update。
- Family members：owner 建立與 membership check。
- Children：CRUD。
- Frontend：註冊、登入、Dashboard、建立孩子。
- Tests：Auth/Family/Children API tests。

## Sprint 2：角色與地圖

- Characters：CRUD。
- Regions：seed 8 個預設區域與列表 API。
- Frontend：建立角色、角色列表、王國地圖。
- Tests：角色權限隔離、region seed。

## Sprint 3：故事生成

- AI provider interface 與 mock provider。
- Prompt assembly。
- Story generation job。
- JSON schema validation。
- Safety check。
- Stories/stories_memories/ai_usage_logs 寫入。
- Frontend：故事生成表單、loading、結果導頁。
- Tests：AI mock happy path、AI failed、safety blocked。

## Sprint 4：故事閱讀與時光書

- Stories list/detail/update/delete。
- Timebook grouped API。
- Frontend：故事閱讀頁、故事列表、時光書頁、篩選器。
- Tests：E2E happy path。

## Sprint 5：QA、修正、部署

- 補 API regression tests。
- Playwright E2E。
- k6 smoke load test。
- OWASP ZAP baseline。
- Docker deployment guide。
- MVP release checklist。
