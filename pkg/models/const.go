package models

const (
	StartCommandString = `"Привет! Я чат-бот Розетта

									В боте доступны модели: gpt-4, gpt-3.5-turbo и Midjourney v 5.2
									
									Чатбот умеет:
									1. Создавать изображения
									2. Писать и редактировать тексты
									3. Переводить с любого языка на любой
									4. Писать и редактировать код
									
									Общайтесь с ботом, как с живым собеседником, задавая вопросы на любом языке. Обратите внимание, что иногда GPT придумывает факты и обладает ограниченными знаниями о событиях после 2021 года.
									
									✉️ Чтобы получить текстовый ответ, просто напишите в чат ваш вопрос (выбор модели GPT в разделе /settings).
									
									🌅 Чтобы создать изображение с помощью Midjourney, начните запрос с команды /imagine а затем добавьте описание (доступно в подписке /premium)."`

	SettingCommandString = `"В боте доступны 4 модели ChatGPT: 
									✔️ gpt-3.5-turbo — Это наиболее популярная и доступная модель GPT. Она отлично оптимизированна для чата и прекрасно справляется с генерацией текста. Лимит токенов: 4096.
									✔️ gpt-3.5-turbo-instruct — Новая модель, с улучшенной оптимизацией для ответов на вопросы и конкретных задач: переведи, сделай и тд. Лимит токенов: 4096.
									✔️ gpt-3.5-turbo-16k Имеет те же возможности, что и основная модель, но поддерживает в 4 раза больше текста. Лимит токенов: 16384.
									✔️ gpt-4 — Новейшая на сегодняшний день модель понимания и генерации языка, способная справляться со сложными задачами. Лимит токенов: 8192.
									
									Лимит токенов определяет макcимально возможную длину вашего запроса + сгенерированного ответа GPT. 1 токен равен примерно 4 символам английского языка или 1 символу на других языках.
									
									В бесплатной версии бота доступны модели gpt-3.5-turbo и -instruct. Пользователи с подпиской могут выбрать также модель 16k. Доступ к GPT-4 можно приобрести отдельно.`

	AccountCommandString = `"Тип подписки: стандартная ✔️
									Модель GPT: gpt-3.5-turbo /settings
									Вопросов GPT-3.5 за сегодня: 2/20
									Вопросов GPT-4: 0/0
									Картинок Midjourney: 0/0
									
									Нужно больше? Подключите подписку на месяц за 290 руб.
									
									Премиум-подписка включает:
									✅ 100 запросов к GPT-3.5 ежедневно;
									⭐️ 10 изображений /Midjourney
									✅ gpt-3.5 16K - тексты в 4 раза длиннее;
									✅ нет паузы между запросами;
									✅ высокая скорость работы, даже в период повышенной нагрузки.
									
									
									Чтобы подключить, перейдите в раздел /premium"`

	GenerateImageCommandString = `"🌅 Для создания уникальных изображений с помощью Midjourney введите сначала команду /imagine, а затем короткое описание (промпт) на любом языке.

									В боте по умолчанию работает последняя версия Midjourney 5.2. 

									Все пользователи бота с подпиской могут отправлять 10 запросов к Midjourney бесплатно.

									⏳ Генерация изображений занимает обычно 1-3 минуты.

									Вместе с картинками появляются кнопки:
									U - увеличивает выбранное изображение и добавляет детализацию;
									V - создает 4 новых вариации выбранного изображения;
									🔄 - генерирует изображения заново с тем же промптом.

									Соблюдайте Правила Midjourney:
									⛔️ не используйте в промпте агрессивную или оскорбительную лексику;
									⛔️ не используйте лексику 18+;"`

	HelpCommandString = `"📝 Генерация текстов
									Для генерации текстов при помощи GPT проcто напишите запрос в чат. Пользователи с подпиской /premium могут также отправлять голосовые сообщения. Бот не распознает изображения. Только текст и аудио.
									
									/settings - выбор модели GPT. Доступны: gpt-4, gpt-3.5-turbo и gpt-3.5-turbo-16k
									
									💬 Поддержка контекста
									По умолчанию бот запоминает контекст. Т.е. при подготовке ответа учитывает не только ваш текущий запрос, но и свой предыдущий ответ. Это позволяет вести диалог и задавать уточняющие вопросы. Чтобы начать новый диалог без учета контекста, используйте команду /deletecontext
									
									🌅 Генерация изображений
									Для генерации изображений в боте используется последняя версия Midjourney. Доступна пользователям с подпиской /premium или в рамках отдельного пакета.
									
									/imagine и краткое описание - генерация изображения
									
									⚙️ Другие команды
									/start - описание бота
									/account - ваш профиль и баланс
									/premium - подключение премиум-подписки GPT, Midjourney, Suno и Eightify
									/donations - чаевые для создателей бота
									/terms - пользовательское соглашение
									
									По всем вопросам также можно написать администраторам @webtimr|@Hironix_Rotifer"`

	UnknownCommandString = "You entered an unknown command"
)
