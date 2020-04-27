CREATE TABLE IF NOT EXISTS emails (
  email VARCHAR PRIMARY KEY,
  updated TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT ON TABLE emails IS 'Информация о подтвержденных адресах пользователей';
COMMENT ON COLUMN emails.email IS 'Почтовый адрес';
COMMENT ON COLUMN emails.updated IS 'Дата и время последнего подтверждения';
