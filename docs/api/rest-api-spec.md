# REST API 規格

## 共通規則

- Base URL：`/api/v1`
- Auth：`Authorization: Bearer <access_token>`
- Content-Type：`application/json`
- 時間格式：ISO 8601

## Error Response

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "欄位格式不正確",
    "details": {}
  }
}
```

## Auth

### POST /auth/register

Request:

```json
{
  "email": "parent@example.com",
  "password": "securePassword123",
  "display_name": "小雨爸爸"
}
```

Response 201:

```json
{
  "user": { "id": 1, "email": "parent@example.com", "display_name": "小雨爸爸" },
  "access_token": "jwt",
  "expires_in": 3600
}
```

### POST /auth/login

Request:

```json
{
  "email": "parent@example.com",
  "password": "securePassword123"
}
```

### POST /auth/logout

將目前 JWT 加入 Redis blacklist。

### GET /auth/me

回傳目前使用者與家庭摘要。

## Families

### POST /families

```json
{ "name": "小雨的童話城堡" }
```

### GET /families/me

回傳目前使用者所屬家庭。

### PATCH /families/{familyId}

```json
{ "name": "新的家庭名稱" }
```

## Children

### POST /children

```json
{
  "family_id": 1,
  "name": "小雨",
  "nickname": "雨雨",
  "birth_date": "2022-05-01",
  "gender_optional": null,
  "avatar_url": null
}
```

### GET /children

Query：`family_id`。

### GET /children/{childId}

### PATCH /children/{childId}

### DELETE /children/{childId}

軟刪除。

## Characters

### POST /characters

```json
{
  "family_id": 1,
  "child_id": 1,
  "real_name": "小雨",
  "story_name": "星光小魔女",
  "role_type": "月光魔法學徒",
  "personality_traits": ["好奇", "善良", "愛笑"],
  "likes": ["兔子", "草莓", "公主"],
  "fears": ["打雷", "黑暗"],
  "magic_power": "讓星星發出溫柔的光"
}
```

### GET /characters

Query：`family_id`, `child_id`。

### GET /characters/{characterId}

### PATCH /characters/{characterId}

### DELETE /characters/{characterId}

軟刪除。

## Regions

### GET /regions

回傳 8 個 MVP 預設區域：童話城堡、魔法森林、糖果村、星光湖、彩虹山谷、龍之谷、夢境花園、時光塔。

## Stories

### POST /stories/generate

Request:

```json
{
  "family_id": 1,
  "child_id": 1,
  "main_character_id": 1,
  "region_id": 2,
  "theme": "勇氣",
  "story_length": "5_min",
  "real_life_event_optional": "今天小雨第一次自己收玩具。",
  "tone": "睡前安撫",
  "language": "zh-TW"
}
```

Response 201/202:

```json
{
  "job_id": 1001,
  "status": "completed",
  "story": {
    "id": 501,
    "title": "星光小魔女的整理任務",
    "summary": "小雨在魔法森林學會把魔法石送回家。"
  }
}
```

### GET /stories

Query：`family_id`, `child_id`, `character_id`, `theme`, `region_id`, `year`, `page`, `page_size`。

### GET /stories/{storyId}

### PATCH /stories/{storyId}

MVP 允許更新 `title`、`summary`、`status`。若未來開放 content 編輯，需寫 audit log。

### DELETE /stories/{storyId}

軟刪除。

## Timebook

### GET /timebook

Query：`family_id`, `child_id`, `character_id`, `theme`, `region_id`, `year`。

Response:

```json
{
  "years": [
    {
      "year": 2026,
      "months": [
        {
          "month": 6,
          "stories": [
            { "id": 501, "title": "勇敢的星光小魔女", "created_at": "2026-06-13T20:00:00Z" }
          ]
        }
      ]
    }
  ]
}
```

### GET /timebook/{year}

回傳指定年份資料。


## MVP 實作狀態

目前 Sprint 1 以前的本地 MVP backend 已以 in-memory store 實作以下端點，方便前端串接與 E2E 流程驗證：

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/logout`
- `GET /auth/me`
- `POST /families`
- `GET /families/me`
- `POST /children`
- `GET /children`
- `POST /characters`
- `GET /characters`
- `GET /regions`
- `POST /stories/generate`
- `GET /stories`
- `GET /stories/{storyId}`
- `GET /timebook`

資料庫 migration 已同步定義核心資料表；後續 Sprint 可將 repository 從 in-memory 切換為 MySQL。

## 狀態碼規範

- 200：查詢或更新成功。
- 201：建立成功。
- 202：AI 任務已建立但尚未完成。
- 400：輸入格式錯誤。
- 401：未登入或 token 無效。
- 403：無權限。
- 404：資源不存在或不可見。
- 409：資源衝突，例如 email 重複。
- 429：超過 rate limit。
- 500：系統錯誤。
