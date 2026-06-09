CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.subject_types (
    uuid VARCHAR(36) PRIMARY KEY,
    type VARCHAR(256) NOT NULL UNIQUE
);

ALTER TABLE raspyx.subject_types OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.subject_types IS 'Таблица с типами дисциплин';
GRANT ALL ON TABLE raspyx.subject_types TO raspyxuser;