CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.rooms (
    uuid VARCHAR(36) PRIMARY KEY,
    number VARCHAR(256) NOT NULL UNIQUE
);

ALTER TABLE raspyx.rooms OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.rooms IS 'Таблица с аудиториями';
GRANT ALL ON TABLE raspyx.rooms TO raspyxuser;