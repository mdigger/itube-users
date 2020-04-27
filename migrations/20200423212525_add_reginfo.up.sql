CREATE TABLE IF NOT EXISTS reginfo (
  id BIGSERIAL PRIMARY KEY,
  uid UUID NOT NULL REFERENCES users ON DELETE CASCADE,
  email VARCHAR NOT NULL,
  domain VARCHAR NOT NULL,
  referer TEXT,
  utm JSONB,
  provider VARCHAR,
  created TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT ON TABLE reginfo IS 'Журнал регистрации пользователей';
COMMENT ON COLUMN reginfo.id IS 'Счетчик';
COMMENT ON COLUMN reginfo.uid IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN reginfo.email IS 'Почтовый адрес';
COMMENT ON COLUMN reginfo.domain IS 'Домен';
COMMENT ON COLUMN reginfo.provider IS 'Идентификатор провайдера авторизации';
COMMENT ON COLUMN reginfo.referer IS 'Ссылка на источник регистрации';
COMMENT ON COLUMN reginfo.utm IS 'Флаги с маркетинговой ифнормацией';
COMMENT ON COLUMN reginfo.created IS 'Дата и время создания';
