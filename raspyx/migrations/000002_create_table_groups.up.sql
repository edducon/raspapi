CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.groups (
    uuid VARCHAR(36) PRIMARY KEY,
    number VARCHAR(256) NOT NULL UNIQUE
);

ALTER TABLE raspyx.groups OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.groups IS 'Таблица с группами';
GRANT ALL ON TABLE raspyx.groups TO raspyxuser;