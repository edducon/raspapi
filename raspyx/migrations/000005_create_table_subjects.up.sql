CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.subjects (
    uuid VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE
);

ALTER TABLE raspyx.subjects OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.subjects IS 'Таблица с дисциплинами';
GRANT ALL ON TABLE raspyx.subjects TO raspyxuser;