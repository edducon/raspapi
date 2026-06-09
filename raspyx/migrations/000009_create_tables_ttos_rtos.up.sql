CREATE SCHEMA IF NOT EXISTS raspyx;

ALTER SCHEMA raspyx OWNER TO raspyxuser;

COMMENT ON SCHEMA raspyx IS 'Схема для учебного расписания';

CREATE TABLE IF NOT EXISTS raspyx.teachers_to_schedule (
    teacher_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.teachers(uuid),
    schedule_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.schedule(uuid) ON DELETE CASCADE,
    PRIMARY KEY(teacher_uuid, schedule_uuid)
);

ALTER TABLE raspyx.teachers_to_schedule OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.teachers_to_schedule IS 'Таблица связи преподавателей с расписанием';
GRANT ALL ON TABLE raspyx.teachers_to_schedule TO raspyxuser;

CREATE TABLE IF NOT EXISTS raspyx.rooms_to_schedule (
    room_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.rooms(uuid),
    schedule_uuid VARCHAR(36) NOT NULL REFERENCES raspyx.schedule(uuid) ON DELETE CASCADE,
    PRIMARY KEY(room_uuid, schedule_uuid)
);

ALTER TABLE raspyx.rooms_to_schedule OWNER TO raspyxuser;
COMMENT ON TABLE raspyx.rooms_to_schedule IS 'Таблица связи аудиторий с расписанием';
GRANT ALL ON TABLE raspyx.rooms_to_schedule TO raspyxuser;