# Family Story Universe / 童話城堡 Fairy Castle

童話城堡是一座陪伴孩子長大的家庭童話王國。第一階段已完成產品與系統設計文件；目前已開始 Sprint 0 專案骨架。

## 文件入口

請從 [docs/00-index.md](docs/00-index.md) 開始閱讀。機器可讀 API 規格位於 [docs/api/openapi.yaml](docs/api/openapi.yaml)。

核心文件包含：完整 PRD、MVP 功能清單、User Stories、Acceptance Criteria、系統架構設計、ERD、REST API、OpenAPI 規格、AI Story Engine、Prompt Template、QA Test Plan、Sprint Backlog 與 Repo 結構建議。

## 專案結構

```text
backend/      Golang API skeleton
frontend/     Next.js H5/PWA skeleton
deployments/  Docker Compose and deployment assets
docs/         Product, architecture, API, AI, QA, and sprint docs
```

## 本機開發

### Backend

```bash
cd backend
go test ./...
go run ./cmd/api
```

Backend health check：

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/api/v1/healthz
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend 預設網址：`http://localhost:3000`。

### MVP 一鍵試用頁

啟動 backend 與 frontend 後，可開啟 `http://localhost:3000/demo` 進行本機 MVP smoke test。此頁會依序呼叫 in-memory API：註冊帳號、建立家庭、建立孩子、建立童話角色，最後產生一篇睡前故事。

因 frontend 與 backend 開發時會分別跑在 `localhost:3000` 與 `localhost:8080`，backend 已加入 MVP CORS middleware，允許 H5/PWA 在本機測試時帶 `Authorization` header 呼叫 API。

### Docker Compose

```bash
cd deployments
docker compose up --build
```

服務：

- Frontend：`http://localhost:3000`
- Backend：`http://localhost:8080`
- MySQL：`localhost:3306`
- Redis：`localhost:6379`

## GitHub Pages 前端預覽

可以先把 `frontend/` 以靜態網站方式部署到 GitHub Pages，快速確認首頁、版面與 H5/PWA 前端骨架。

限制：GitHub Pages 只能跑靜態前端，不能執行 Golang backend、MySQL、Redis 或 AI story generation worker。需要完整 API 流程時，請搭配本機 backend、Docker Compose，或另外部署 backend 到 Render、Fly.io、Railway、Cloud Run、VPS 等可執行容器 / Go 服務的平台。

啟用方式：

1. 到 GitHub repo 的 `Settings` → `Pages`。
2. 將 `Build and deployment` 的 `Source` 設為 `GitHub Actions`。
3. 合併到 `main` 後，或手動執行 `Deploy Frontend Preview to GitHub Pages` workflow。
4. 預覽網址會是 `https://wwking4896.github.io/Family-Story-Universe/`。

此 repo 已新增 `.github/workflows/pages.yml`，會在 `main` 分支的 `frontend/**` 或 workflow 變更後，自動 build Next.js static export 並部署 `frontend/out`。

若 `actions/deploy-pages` 出現 `HttpError: Not Found` 或 `Ensure GitHub Pages has been enabled`，通常代表 repo 尚未啟用 GitHub Pages。請先到 `Settings` → `Pages`，確認 Source 已選擇 `GitHub Actions`，儲存後再重新執行 workflow。若 repo 是 private，請同時確認目前 GitHub 方案支援 private repo Pages。

若要測完整產品流程：

- 本機最快：`cd deployments && docker compose up --build`。
- 只測 backend：`cd backend && go run ./cmd/api`，再用 curl / Postman 呼叫 API。
- 線上完整測試：GitHub Pages 放 frontend，backend 另外部署到支援長駐服務的平台，之後用 `NEXT_PUBLIC_API_BASE_URL` 指向該 backend API。

## Sprint 0 狀態

- Monorepo 目錄骨架：完成。
- Backend Go API skeleton：完成。
- Frontend Next.js/PWA skeleton：完成。
- Docker Compose：完成。
- 環境變數範本：完成。
- CI backend check：完成，GitHub Actions 會執行 backend test 與 build。

## MVP API 目前可用端點

本地 backend 目前提供 in-memory MVP 流程，適合前端串接與 E2E 驗證：

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `GET /api/v1/auth/me`
- `POST /api/v1/families` / `GET /api/v1/families/me` / `PATCH /api/v1/families/{familyId}`
- `POST /api/v1/children` / `GET /api/v1/children?family_id=1` / `GET|PATCH|DELETE /api/v1/children/{childId}`
- `POST /api/v1/characters` / `GET /api/v1/characters?family_id=1` / `GET|PATCH|DELETE /api/v1/characters/{characterId}`
- `GET /api/v1/regions`
- `POST /api/v1/stories/generate`
- `GET /api/v1/stories?family_id=1` / `GET|PATCH|DELETE /api/v1/stories/{storyId}`
- `GET /api/v1/timebook?family_id=1` / `GET /api/v1/timebook/{year}?family_id=1`

注意：目前資料暫存在記憶體中，服務重啟後會清空；MySQL schema 已先由 migrations 定義，後續 Sprint 可切換 repository。故事生成目前會先檢查 MVP 支援主題、故事長度、語氣、語言與基礎兒童安全 / prompt injection 關鍵字。


## CI 說明

GitHub Actions workflow 位於 `.github/workflows/ci.yml`，會在 pull request 與 push 到 `main` 時執行：

```bash
cd backend
go test ./...
go build ./cmd/api
```

目前 backend 僅使用 Go standard library，沒有外部 module，因此 CI 不依賴 `go.sum`。


## 衝突處理說明

若 PR 顯示與 `main` 發生衝突，請優先保留本文件目前的完整多行版本，因為它包含 Sprint 0 狀態、MVP API 端點與 CI 說明。若 GitHub conflict editor 顯示 README 衝突，保留本版本即可。

本 repo 也已加入 `.gitignore`，避免 `frontend/node_modules/`、Next.js build output、Go build binary、local `.env` 等開發產物被誤加入 PR。
