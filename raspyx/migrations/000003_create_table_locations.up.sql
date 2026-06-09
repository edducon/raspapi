CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.locations (
    uuid VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE
);

ALTER TABLE raspyx.locations OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.locations IS 'Таблица с локациями';
GRANT ALL ON TABLE raspyx.locations TO raspyxuser;