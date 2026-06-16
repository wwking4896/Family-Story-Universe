# QA Test Plan

## 測試目標

確保 MVP 能安全、穩定地完成「註冊登入 → 建立家庭/孩子/角色 → 生成故事 → 閱讀故事 → 時光書回顧」核心流程。

## 測試類型

| 類型 | 工具 | 範圍 |
|---|---|---|
| Unit Test | Go test / React Testing Library | domain、service、component |
| API Test | Go integration / Postman / Newman | REST API |
| E2E | Playwright | 使用者核心旅程 |
| AI Eval | JSON schema validator + rule checks | AI 輸出品質 |
| Load Test | k6 | 核心 API smoke load |
| Security Smoke | OWASP ZAP | 常見 web 風險 |

## 功能測試

### Auth

- 註冊成功。
- 重複 email 註冊失敗。
- 登入成功。
- 錯誤密碼登入失敗。
- 未登入不可呼叫受保護 API。
- 登出後 token blacklist 生效。

### Children

- 建立孩子成功。
- gender_optional 可為 null。
- 修改孩子成功。
- 軟刪除孩子成功。
- 不同家庭不可讀取孩子資料。

### Characters

- 建立童話角色成功。
- child_id 不屬於家庭時建立失敗。
- 修改角色成功。
- 軟刪除角色成功。

### Stories

- 合法輸入可建立 generation job。
- AI 成功時保存故事。
- AI 失敗時 job status 正確。
- 故事可重新讀取。
- 故事列表支援篩選。

### Timebook

- 依年份與月份分組。
- 支援 child、character、theme、region、year 篩選。
- 無資料時回傳空陣列而非錯誤。

## AI 測試

- 輸出合法 JSON。
- 使用繁體中文與台灣自然用語。
- 不包含死亡、血腥、成人內容、歧視、過度恐怖。
- 適合 2～10 歲。
- 結尾溫柔、有安全感。
- memory_tags 3～8 個。
- prompt injection 測試不可繞過安全規則。

## 權限測試

- A 家庭不可讀取 B 家庭 children。
- A 家庭不可讀取 B 家庭 characters。
- A 家庭不可讀取 B 家庭 stories。
- child_id 不屬於 family 時不可生成故事。
- main_character_id 不屬於 family 時不可生成故事。

## 安全測試

- SQL Injection。
- XSS。
- JWT 偽造。
- JWT 過期。
- Rate limit。
- Prompt injection。
- 檔案上傳副檔名與大小限制。

## UAT Checklist

- 手機上可完成註冊到第一篇故事。
- 故事生成過程有清楚 loading 狀態。
- 故事閱讀頁適合父母朗讀。
- 時光書可快速找到過去故事。
- 錯誤訊息對一般父母看得懂。
