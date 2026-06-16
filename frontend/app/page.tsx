const quickLinks = ['建立孩子', '建立角色', '王國地圖', '時光書'];

const mvpSteps = [
  '註冊 / 登入',
  '建立家庭',
  '建立孩子資料',
  '建立童話角色',
  '選擇場景與主題',
  '產生並保存故事',
];

export default function HomePage() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-castle-cream to-white px-6 py-8">
      <section className="mx-auto flex max-w-md flex-col gap-8">
        <header className="rounded-[2rem] bg-white/80 p-6 shadow-xl shadow-purple-100">
          <p className="text-sm font-semibold text-castle-purple">Fairy Castle MVP</p>
          <h1 className="mt-2 text-4xl font-bold tracking-tight text-castle-night">童話城堡</h1>
          <p className="mt-3 text-base leading-7 text-slate-600">
            把孩子、家人與日常回憶，轉化成會陪伴孩子長大的家庭童話王國。
          </p>
          <div className="mt-6 grid gap-3">
            <a className="block w-full rounded-full bg-castle-purple px-5 py-3 text-center text-base font-semibold text-white shadow-lg shadow-purple-200" href="./demo">
              開始本機 MVP 試用
            </a>
            <a className="block w-full rounded-full border border-purple-200 bg-white px-5 py-3 text-center text-base font-semibold text-castle-purple" href="#mvp-flow">
              查看 MVP 流程
            </a>
          </div>
        </header>

        <section aria-labelledby="today-story" className="rounded-3xl bg-castle-night p-5 text-white">
          <p id="today-story" className="text-sm font-medium text-purple-200">
            今日推薦故事
          </p>
          <h2 className="mt-2 text-2xl font-bold">星光湖的晚安魔法</h2>
          <p className="mt-2 text-sm leading-6 text-purple-100">適合睡前 5 分鐘朗讀，主題：睡前安撫。</p>
        </section>

        <nav className="grid grid-cols-2 gap-3" aria-label="主要功能">
          {quickLinks.map((item) => (
            <a key={item} className="rounded-3xl bg-white p-5 text-center font-semibold text-castle-night shadow-md shadow-purple-100" href="#mvp-flow">
              {item}
            </a>
          ))}
        </nav>

        <section id="mvp-flow" className="rounded-3xl bg-white p-5 shadow-md shadow-purple-100">
          <h2 className="text-xl font-bold text-castle-night">MVP 核心流程</h2>
          <ol className="mt-4 space-y-3">
            {mvpSteps.map((step, index) => (
              <li key={step} className="flex gap-3 text-sm text-slate-700">
                <span className="flex h-7 w-7 shrink-0 items-center justify-center rounded-full bg-purple-100 font-bold text-castle-purple">{index + 1}</span>
                <span className="pt-1">{step}</span>
              </li>
            ))}
          </ol>
        </section>
      </section>
    </main>
  );
}
