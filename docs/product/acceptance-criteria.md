# Acceptance Criteria

## Auth

- Given 使用者輸入有效 email、密碼、display_name，When 註冊，Then 系統建立帳號並回傳 JWT。
- Given email 已存在，When 註冊，Then 回傳 409。
- Given 密碼錯誤，When 登入，Then 回傳 401。
- Given 未帶 JWT，When 呼叫受保護 API，Then 回傳 401。
- Given JWT 過期，When 呼叫 API，Then 回傳 401。

## Family

- Given 已登入使用者，When 建立家庭，Then 系統建立 family 與 owner family_member。
- Given 非家庭成員，When 讀取家庭資料，Then 回傳 403 或 404。
- Given owner 更新家庭名稱，When PATCH family，Then 名稱更新且寫入 audit log。

## Children

- Given 已登入家庭成員，When 建立孩子，Then child.family_id 必須為該使用者所屬家庭。
- Given gender_optional 空值，When 建立孩子，Then 系統接受。
- Given A 家庭 child_id，When B 家庭讀取，Then 不可存取。
- Given 刪除孩子，When DELETE，Then 軟刪除且既有故事不消失。

## Characters

- Given child_id 不屬於目前家庭，When 建立角色，Then 回傳 400/403。
- Given 角色建立成功，Then level 預設 1、exp 預設 0。
- Given 角色被刪除，When 查詢列表，Then 預設不顯示軟刪除角色。

## Stories

- Given 合法請求，When 生成故事，Then 建立 story_generation_job。
- Given AI 回傳合法且安全 JSON，Then 建立 story 與 story_memories。
- Given AI 回傳非法 JSON，Then job.status = failed，且不建立 published story。
- Given safety_check 有任一高風險 true，Then job.status = blocked 或 failed。
- Given 生成成功，Then response 包含 story_id、title、summary。

## Timebook

- Given 家庭有多篇故事，When 查詢 timebook，Then 依年份與月份分組。
- Given 篩選 child_id，When 查詢 timebook，Then 只回傳該孩子故事。
- Given B 家庭使用 A 家庭 story_id，When 查詢，Then 不可存取。

## AI 品質

- 產出必須是繁體中文與台灣自然用語。
- 產出必須符合 Story JSON Schema。
- 不得包含死亡、血腥、成人內容、歧視、過度恐怖或危險模仿行為。
- 睡前安撫語氣需低刺激，結尾需安全溫柔。
