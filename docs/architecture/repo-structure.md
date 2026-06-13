# Repo 結構建議

## 建議：Monorepo

MVP 需要快速協調前端、後端、OpenAPI、Docker 與文件，因此建議使用 monorepo。

```text
/
  README.md
  docs/
    00-index.md
    product/
    architecture/
    api/
    ai/
    qa/
    project/
  backend/
    cmd/api/main.go
    internal/
      config/
      domain/
      application/services/
      infrastructure/db/
      infrastructure/redis/
      infrastructure/ai/
      infrastructure/storage/
      interfaces/http/handlers/
      interfaces/http/middlewares/
      interfaces/http/routes/
    pkg/logger/
    pkg/errors/
    pkg/validator/
    migrations/
    docs/
    Dockerfile
    go.mod
  frontend/
    app/
    components/
    features/
    lib/api/
    lib/auth/
    public/
    tests/
    package.json
  deployments/
    docker-compose.yml
    nginx/
  scripts/
  .github/workflows/
```

## Backend 原則

- `internal/domain` 不依賴 framework。
- `internal/application` 放 use case orchestration。
- `internal/infrastructure` 封裝 MySQL、Redis、AI、S3。
- `internal/interfaces/http` 放 Gin handlers 與 middleware。

## Frontend 原則

- `app` 放 Next.js routes。
- `features` 依領域切分：auth、children、characters、stories、timebook。
- `lib/api` 放 typed API client。
- `components` 放可重用 UI component。

## 文件原則

- `docs/api/rest-api-spec.md` 作為人讀 API 文件。
- 後續新增 `docs/api/openapi.yaml` 作為機器可讀規格。
- ERD、Prompt、QA 文件需隨需求變更同步更新。
