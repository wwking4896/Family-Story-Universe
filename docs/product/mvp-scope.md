# MVP 功能清單

## MVP 原則

第一版只驗證「家庭角色 + AI 故事 + 時光書」核心閉環，不做大型遊戲化或複雜商務流程。

## P0 必做

| 模組 | 功能 | 說明 |
|---|---|---|
| Auth | 註冊 / 登入 / 登出 / Me | JWT 驗證，密碼安全儲存 |
| Family | 建立家庭 / 我的家庭 / 更新家庭 | 一位使用者 MVP 預設一個主要家庭 |
| Children | 孩子 CRUD | 性別 optional，支援 avatar_url |
| Characters | 童話角色 CRUD | 包含喜好、害怕、魔法能力、level、exp |
| Regions | 預設地圖區域 | 8 個 MVP 區域，以 seed data 建立 |
| Story Generation | 建立任務、呼叫 AI、保存故事 | 需處理成功、失敗、blocked 狀態 |
| Story Reader | 故事閱讀頁 | 手機友善、段落清楚 |
| Timebook | 年月分組與篩選 | 依孩子、角色、主題、區域、年份 |
| Security | 家庭資料隔離 | 所有家庭資源查詢必須檢查 membership |
| AI Safety | JSON schema + safety check | 不安全故事不可發布 |

## P1 可做但不阻塞 MVP

- 訂閱方案頁靜態展示。
- Admin 基礎 AI 用量檢視。
- 故事標題/摘要編輯。
- 基本 dashboard 統計。

## 明確不做

- AI 語音與聲音複製。
- 實體書下單流程。
- PDF 匯出。
- 多人即時共創。
- 社群分享。
- 原生 iOS/Android App。
- 複雜成就與 3D 城堡。
- Vector database。

## MVP 驗收總標準

1. 使用者可註冊登入。
2. 使用者可建立家庭、孩子、童話角色。
3. 使用者可選角色、場景、主題、長度與語氣生成故事。
4. 故事必須保存並可重新閱讀。
5. 時光書可依年月顯示。
6. 不同家庭資料完全隔離。
7. 系統可用 Docker 啟動。
8. 有基礎 API / E2E / AI 品質測試計畫。
