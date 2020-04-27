CREATE TABLE IF NOT EXISTS openid (
  provider VARCHAR NOT NULL,
  subject VARCHAR NOT NULL,
  uid UUID NOT NULL REFERENCES users ON DELETE CASCADE,
  created TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (provider, subject)
);

COMMENT ON TABLE openid IS 'Информация о регистрации пользователей через OpenID Connect';
COMMENT ON COLUMN openid.provider IS 'Идентификатор провайдера авторизации';
COMMENT ON COLUMN openid.subject IS 'Уникальный идентификатор пользователя у провайдера';
COMMENT ON COLUMN openid.uid IS 'Внутренний уникальный идентификатор пользователя';
COMMENT ON COLUMN openid.created IS 'Дата и время регистрации';
