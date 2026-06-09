CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.schedule (
    uuid VARCHAR(36) PRIMARY KEY,
    group_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.groups(uuid),
    subject_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.subjects(uuid),
    subject_type_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.subject_types(uuid),
    location_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.locations(uuid),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    weekday INT NOT NULL,
    link TEXT,
    is_session BOOL NOT NULL DEFAULT false
);

ALTER TABLE raspyx.schedule OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.schedule IS 'Таблица с расписанием';
GRANT ALL ON TABLE raspyx.schedule TO raspyxuser;