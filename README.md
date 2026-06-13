# Family Story Universe / 童話城堡 Fairy Castle

童話城堡是一座陪伴孩子長大的家庭童話王國。第一階段已完成產品與系統設計文件；目前已開始 Sprint 0 專案骨架。

## 文件入口

請從 [docs/00-index.md](docs/00-index.md) 開始閱讀。

核心文件包含：完整 PRD、MVP 功能清單、User Stories、Acceptance Criteria、系統架構設計、ERD、REST API、AI Story Engine、Prompt Template、QA Test Plan、Sprint Backlog 與 Repo 結構建議。

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
- `POST /api/v1/families`
- `GET /api/v1/families/me`
- `POST /api/v1/children` / `GET /api/v1/children?family_id=1`
- `POST /api/v1/characters` / `GET /api/v1/characters?family_id=1`
- `GET /api/v1/regions`
- `POST /api/v1/stories/generate`
- `GET /api/v1/stories?family_id=1`
- `GET /api/v1/timebook?family_id=1`

注意：目前資料暫存在記憶體中，服務重啟後會清空；MySQL schema 已先由 migrations 定義，後續 Sprint 可切換 repository。


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
