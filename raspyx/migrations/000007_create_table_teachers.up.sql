CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.teachers (
    uuid VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(256) NOT NULL,
    second_name VARCHAR(256) NOT NULL,
    middle_name VARCHAR(256)
);

ALTER TABLE raspyx.teachers OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.teachers IS 'Таблица с преподавателями';
GRANT ALL ON TABLE raspyx.teachers TO raspyxuser;