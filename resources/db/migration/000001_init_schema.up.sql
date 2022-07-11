CREATE TABLE DATA(
    KEY VARCHAR PRIMARY KEY,
    VALUE VARCHAR
);

COMMENT ON TABLE DATA IS 'Таблица для хранения key=value';

CREATE TABLE SESSION(
    SESSION_ID VARCHAR PRIMARY KEY,
    TELEGRAM_CHAT_ID NUMERIC,
    EXPIRED_AT TIMESTAMP
);

CREATE INDEX SESSION_TELEGRAM_CHAT_ID ON SESSION(TELEGRAM_CHAT_ID);

COMMENT ON TABLE SESSION IS 'Таблица для сохранения сессии';