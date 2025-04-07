# Описание
Данный сервис должен стать платформой для тестирования, где можно создавать квизы, добавлять тестовые вопросы. Пользователи должны иметь возможность зарегистрироваться, получить доступ к тестам и решить тест. Также должна быть возможность получения результата.

# Охватываемый стек
- [ ] Go
- [ ] Реализация правильной архитектуры репозитория
- [ ] Написание тестов для usecase
- [ ] Использование testify и mock
- [ ] Использование асинхронных задач для подсчета результатов (использование брокера сообщений NATS?)
- [ ] Реализация CRON периодической задачи для рассылки результатов (куда-либо)
- [ ] Использование GORM и БД MySQL
- [ ] Оборачивание всего сервиса в docker
- [ ] Написание фронта на HTMX для SPA (возможно)
- [ ] Использование Fiber в качестве http-фреймворка
- [ ] Авторизация (?)
- [ ] Использовать Redis для сессий