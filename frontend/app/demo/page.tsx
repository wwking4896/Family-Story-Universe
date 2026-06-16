'use client';

import { FormEvent, useMemo, useState } from 'react';

type DemoState = {
  token: string;
  familyId: number;
  childId: number;
  characterId: number;
  storyTitle: string;
  storyContent: string;
};

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080/api/v1';

export default function DemoPage() {
  const [email, setEmail] = useState(() => `parent-${Date.now()}@example.com`);
  const [password, setPassword] = useState('password123');
  const [displayName, setDisplayName] = useState('小雨爸爸');
  const [familyName, setFamilyName] = useState('小雨的童話城堡');
  const [childName, setChildName] = useState('小雨');
  const [characterName, setCharacterName] = useState('星光小魔女');
  const [realLifeEvent, setRealLifeEvent] = useState('今天小雨第一次自己收玩具，雖然一開始有點不想收，但後來還是完成了。');
  const [state, setState] = useState<DemoState>({ token: '', familyId: 0, childId: 0, characterId: 0, storyTitle: '', storyContent: '' });
  const [log, setLog] = useState<string[]>(['請先啟動 backend：cd backend && go run ./cmd/api']);
  const [isRunning, setIsRunning] = useState(false);

  const canGenerate = useMemo(() => Boolean(state.token && state.familyId && state.childId && state.characterId), [state]);

  async function runSetup(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsRunning(true);
    setLog(['開始建立 MVP 測試資料...']);
    try {
      const auth = await post<{ access_token: string }>('/auth/register', { email, password, display_name: displayName });
      appendLog('完成註冊 / 登入 token');

      const family = await post<{ id: number }>('/families', { name: familyName }, auth.access_token);
      appendLog(`建立家庭 #${family.id}`);

      const child = await post<{ id: number }>('/children', { family_id: family.id, name: childName, nickname: childName, birth_date: '2022-05-01' }, auth.access_token);
      appendLog(`建立孩子 #${child.id}`);

      const character = await post<{ id: number }>('/characters', {
        family_id: family.id,
        child_id: child.id,
        real_name: childName,
        story_name: characterName,
        role_type: '月光魔法學徒',
        personality_traits: ['好奇', '善良', '愛笑'],
        likes: ['兔子', '草莓', '公主'],
        fears: ['打雷', '黑暗'],
        magic_power: '讓星星發出溫柔的光',
      }, auth.access_token);
      appendLog(`建立童話角色 #${character.id}`);

      setState((current) => ({ ...current, token: auth.access_token, familyId: family.id, childId: child.id, characterId: character.id }));
      appendLog('測試資料完成，可以產生故事。');
    } catch (error) {
      appendLog(error instanceof Error ? error.message : '建立測試資料失敗');
    } finally {
      setIsRunning(false);
    }
  }

  async function generateStory() {
    setIsRunning(true);
    appendLog('開始產生故事...');
    try {
      const result = await post<{ story: { title: string; content: string } }>('/stories/generate', {
        family_id: state.familyId,
        child_id: state.childId,
        main_character_id: state.characterId,
        region_id: 2,
        theme: '勇氣',
        story_length: '5_min',
        tone: '睡前安撫',
        language: 'zh-TW',
        real_life_event_optional: realLifeEvent,
      }, state.token);
      setState((current) => ({ ...current, storyTitle: result.story.title, storyContent: result.story.content }));
      appendLog(`故事完成：${result.story.title}`);
    } catch (error) {
      appendLog(error instanceof Error ? error.message : '產生故事失敗');
    } finally {
      setIsRunning(false);
    }
  }

  function appendLog(message: string) {
    setLog((items) => [...items, message]);
  }

  return (
    <main className="min-h-screen bg-gradient-to-b from-castle-cream to-white px-6 py-8">
      <section className="mx-auto flex max-w-2xl flex-col gap-6">
        <a className="text-sm font-semibold text-castle-purple" href="../">← 回首頁</a>
        <header className="rounded-[2rem] bg-white/90 p-6 shadow-xl shadow-purple-100">
          <p className="text-sm font-semibold text-castle-purple">MVP Local Demo</p>
          <h1 className="mt-2 text-3xl font-bold text-castle-night">童話城堡一鍵試用流程</h1>
          <p className="mt-3 text-sm leading-6 text-slate-600">
            這個頁面會呼叫本機 backend，依序建立帳號、家庭、孩子、角色，最後產生一篇故事。請先確認 backend 正在 `http://localhost:8080` 執行。
          </p>
        </header>

        <form onSubmit={runSetup} className="grid gap-4 rounded-3xl bg-white p-5 shadow-md shadow-purple-100">
          <label className="grid gap-2 text-sm font-semibold text-castle-night">
            Email
            <input className="rounded-2xl border border-purple-100 px-4 py-3 font-normal" value={email} onChange={(event) => setEmail(event.target.value)} />
          </label>
          <label className="grid gap-2 text-sm font-semibold text-castle-night">
            密碼
            <input className="rounded-2xl border border-purple-100 px-4 py-3 font-normal" value={password} onChange={(event) => setPassword(event.target.value)} type="password" />
          </label>
          <label className="grid gap-2 text-sm font-semibold text-castle-night">
            顯示名稱
            <input className="rounded-2xl border border-purple-100 px-4 py-3 font-normal" value={displayName} onChange={(event) => setDisplayName(event.target.value)} />
          </label>
          <div className="grid gap-3 md:grid-cols-3">
            <input aria-label="家庭名稱" className="rounded-2xl border border-purple-100 px-4 py-3" value={familyName} onChange={(event) => setFamilyName(event.target.value)} />
            <input aria-label="孩子名稱" className="rounded-2xl border border-purple-100 px-4 py-3" value={childName} onChange={(event) => setChildName(event.target.value)} />
            <input aria-label="角色名稱" className="rounded-2xl border border-purple-100 px-4 py-3" value={characterName} onChange={(event) => setCharacterName(event.target.value)} />
          </div>
          <button className="rounded-full bg-castle-purple px-5 py-3 font-semibold text-white disabled:opacity-60" disabled={isRunning} type="submit">
            建立測試資料
          </button>
        </form>

        <section className="grid gap-4 rounded-3xl bg-white p-5 shadow-md shadow-purple-100">
          <label className="grid gap-2 text-sm font-semibold text-castle-night">
            今天發生的事
            <textarea className="min-h-28 rounded-2xl border border-purple-100 px-4 py-3 font-normal" value={realLifeEvent} onChange={(event) => setRealLifeEvent(event.target.value)} />
          </label>
          <button className="rounded-full bg-castle-night px-5 py-3 font-semibold text-white disabled:opacity-60" disabled={!canGenerate || isRunning} onClick={generateStory} type="button">
            產生睡前故事
          </button>
        </section>

        <section className="rounded-3xl bg-castle-night p-5 text-white">
          <h2 className="text-xl font-bold">執行紀錄</h2>
          <ol className="mt-3 list-decimal space-y-2 pl-5 text-sm text-purple-100">
            {log.map((item, index) => <li key={`${item}-${index}`}>{item}</li>)}
          </ol>
        </section>

        {state.storyTitle && (
          <article className="rounded-[2rem] bg-white p-6 shadow-xl shadow-purple-100">
            <p className="text-sm font-semibold text-castle-purple">故事閱讀頁預覽</p>
            <h2 className="mt-2 text-2xl font-bold text-castle-night">{state.storyTitle}</h2>
            <p className="mt-4 whitespace-pre-wrap text-base leading-8 text-slate-700">{state.storyContent}</p>
          </article>
        )}
      </section>
    </main>
  );
}

async function post<T>(path: string, payload: unknown, token?: string): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: JSON.stringify(payload),
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(`API ${path} 回傳 ${response.status}: ${errorText}`);
  }

  return response.json() as Promise<T>;
}
