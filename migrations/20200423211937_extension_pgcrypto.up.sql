CREATE EXTENSION IF NOT EXISTS pgcrypto;

COMMENT ON EXTENSION pgcrypto IS 'Используется функция gen_random_uuid() для генерации уникальных идентификаторов';