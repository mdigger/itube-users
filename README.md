# Сервис авторизации и работы с пользователями

Доступны следующие сервисы:

- **Identity** — отвечает за регистрацию и авторизацию пользователя по логину
и паролю;
- **OpenID** — поддерживает авторизацию с помощью внешних провайдеров по 
протоколу [OpenID Connect](https://openid.net/connect/).
- **tokens** — генерация и проверка токенов для сброса пароля или подтверждения 
почтового авдреса пользователя

Описание gRPC-протокола находится в каталоге [`api/protobuf-spec/`](api/protobuf-spec/).

### Параметры для запуска:

- Для работы необходимо указать `DSN` для доступа PostgreSQL. 

- Для поддержки авторизации пользователей через Google необходимо задать
`GOOGLE_CLIENT_ID` и `GOOGLE_SECRET`. Получить эти значения можно 
зарегистрировав сервер в [консоле Google](https://console.developers.google.com/).

- Настройка доступа к почтовому серверу задается в виде URL `SMTP` 
(`smtp://login:password@host[:port]`). Поддерживается TLS при соединении.

- Чтобы указать путь к шаблонам почтовых сообщений используйте `TEMPLATES`.
Пример файла с шаблоном можно посмотреть в файле 
[`email_templates.yaml`](email_templates.yaml).

- Порт, используемый для сервиса gRPC задается как `PORT`. По умолчанию
используется `50051`.

Все параметры можно задать как через переменные окружения, там и в виде
параметров запуска.

## Шаблоны писем

Для каждого поддерживаемого домена необходимо определение шаблона для проверки 
email-адреса и для сброса пароля пользователя. Желательно задавать как 
текстовый, так и HTML варианты писем.

Программа [`email_template-gen`](cmd/email-templates-gen/) позволяет быстро создать 
[шабоны для писем](email_templates.yaml) по [описанию](email_config.yaml). 
Для добавления новых доменов просто продублируйте описание в том же файле и 
поправьте поля.