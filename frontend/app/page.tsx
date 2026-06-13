const quickLinks = ['我的孩子', '我的角色', '王國地圖', '時光書'];

export default function HomePage() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-castle-cream to-white px-6 py-8">
      <section className="mx-auto flex max-w-md flex-col gap-8">
        <header className="rounded-[2rem] bg-white/80 p-6 shadow-xl shadow-purple-100">
          <p className="text-sm font-semibold text-castle-purple">Fairy Castle</p>
          <h1 className="mt-2 text-4xl font-bold tracking-tight text-castle-night">童話城堡</h1>
          <p className="mt-3 text-base leading-7 text-slate-600">
            把孩子、家人與日常回憶，轉化成會陪伴孩子長大的家庭童話王國。
          </p>
          <button className="mt-6 w-full rounded-full bg-castle-purple px-5 py-3 text-base font-semibold text-white shadow-lg shadow-purple-200">
            產生新故事
          </button>
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
            <a key={item} className="rounded-3xl bg-white p-5 text-center font-semibold text-castle-night shadow-md shadow-purple-100" href="#">
              {item}
            </a>
          ))}
        </nav>
      </section>
    </main>
  );
}
