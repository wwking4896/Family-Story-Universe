# AI Story Engine 規格

## 目標

AI Story Engine 需將孩子資料、童話角色、世界區域、主題、真實事件與過去記憶，轉換成安全、適齡、繁體中文且可保存的童話故事。

## Pipeline

1. Request Validation：驗證使用者、家庭、孩子、角色、區域、主題、長度、quota。
2. Memory Retrieval：從 `story_memories` 取得最近且相關的 5～10 筆記憶。
3. Prompt Assembly：組合 system prompt、user prompt 與 JSON schema 約束。
4. AI Generation：呼叫 AI Provider。
5. JSON Parse：解析回應，失敗則 retry 或標記 failed。
6. Schema Validation：確認必要欄位存在且型別正確。
7. Safety Check：檢查死亡、血腥、成人內容、歧視、恐怖與危險行為。
8. Persist：保存 stories、story_memories、ai_usage_logs、story_generation_jobs。
9. Return：回傳故事或任務狀態。

## AIService Interface

```text
GenerateStory(ctx, input) -> output
GenerateSummary(ctx, input) -> output
ExtractMemoryTags(ctx, input) -> []memory_tag
SafetyCheck(ctx, input) -> safety_result
GenerateImage(ctx, input) -> output // V2
```

## Story JSON Schema

```json
{
  "title": "故事標題",
  "summary": "一句話摘要",
  "age_range": "3-6",
  "theme": "勇氣",
  "region": "魔法森林",
  "main_character": "星光小魔女",
  "content": "完整故事內容",
  "moral": "故事寓意",
  "memory_tags": ["第一次收玩具", "勇氣", "責任"],
  "safety_check": {
    "violence": false,
    "death": false,
    "adult_content": false,
    "scary_content": false,
    "discrimination": false,
    "unsafe_behavior": false
  }
}
```

## 內容安全規則

- 不得包含死亡、血腥、成人內容、仇恨、歧視或不當暗示。
- 不得鼓勵孩子模仿危險行為。
- 不使用羞辱式管教語言。
- 可描述小挑戰，但必須可被理解、被照顧者支持且低焦慮。
- 睡前安撫故事需節奏平穩、結尾安心。
- 真實事件需被溫柔轉化，不直接暴露敏感家庭資訊。

## Memory Retrieval MVP

- 同 child_id 優先。
- 同 theme、region 加權。
- 最近 90 天優先。
- 最多 10 筆。
- 每筆 memory tag 限制長度，避免 prompt injection。

## Failure Handling

| 情境 | 處理 |
|---|---|
| AI timeout | job.status = failed，提示稍後再試 |
| JSON parse failed | 最多重試一次，仍失敗則 failed |
| safety failed | job.status = blocked，不發布故事 |
| quota exceeded | 回傳 429 或商業錯誤碼 |
| provider unavailable | 使用 mock/fallback 僅限 dev，不在 production 假成功 |


## MVP Rule-based Guardrails

目前 backend in-memory MVP 在呼叫 mock story generation 前會先做規則式驗證：

- 主題必須屬於 MVP 支援清單：勇氣、分享、禮貌、責任、同理心、情緒管理、睡前放鬆、親子陪伴。
- 故事長度必須是 `3_min`、`5_min` 或 `10_min`。
- 語氣若有提供，必須是溫柔、奇幻、搞笑、睡前安撫或冒險。
- 語言若有提供，必須是 `zh-TW`。
- `real_life_event_optional`、主題與語氣會做基礎兒童安全與 prompt injection 關鍵字檢查，例如死亡、血腥、成人內容、仇恨、歧視、恐怖，以及要求忽略系統提示的輸入。

這不是最終內容安全方案；正式 AI provider 串接後仍需保留 schema validation、模型安全檢查與人工可稽核 logs。
