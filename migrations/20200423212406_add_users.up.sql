CREATE TABLE IF NOT EXISTS users (
  uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR(60),
  blocked BOOL NOT NULL DEFAULT FALSE,
  properties JSONB,
  created TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated TIMESTAMPTZ NOT NULL DEFAULT now(),
  logged TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT ON TABLE users IS 'Информация о зарегистрированных пользователях';
COMMENT ON COLUMN users.uid IS 'Уникальный идентификатор';
COMMENT ON COLUMN users.email IS 'Почтовый адрес';
COMMENT ON COLUMN users.password IS 'Хеш от пароля (bcrypt)';
COMMENT ON COLUMN users.blocked IS 'Флаг, что пользователь заблокирован';
COMMENT ON COLUMN users.properties IS 'Свойства пользователя';
COMMENT ON COLUMN users.created IS 'Дата и время регистрации';
COMMENT ON COLUMN users.updated IS 'Дата и время обновления информации или изменения пароля';
COMMENT ON COLUMN users.updated IS 'Дата и время последней успешной авторизации';