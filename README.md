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
- CI backend check：完成。
