CREATE TABLE IF NOT EXISTS tokens(
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  -- uid UUID NOT NULL REFERENCES users ON DELETE CASCADE,
  domain VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  type SMALLINT NOT NULL,
  sended BOOL NOT NULL DEFAULT FALSE,
  created TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (domain, email, type)
);

COMMENT ON TABLE tokens IS 'Токены для проврки почтового адреса и сброса пароля пользователя';
COMMENT ON COLUMN tokens.id IS 'Токен';
COMMENT ON COLUMN tokens.email IS 'Почтовый адрес';
COMMENT ON COLUMN tokens.domain IS 'Домен';
COMMENT ON COLUMN tokens.type IS 'Тип токена: почта или сброс пароля';
COMMENT ON COLUMN tokens.sended IS 'Влаг, что письмо отправлено';
COMMENT ON COLUMN tokens.created IS 'Дата и время создания';
