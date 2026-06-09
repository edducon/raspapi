package schedule

import (
	"context"
	"fmt"
	"raspyx2/internal/models"
	"raspyx2/internal/repository/constRepository"
)

func (r *ScheduleRepository) UpdateScheduleLinkByLesson(filters *models.UpdateScheduleLinkByLessonRequest) (int64, error) {
	query := fmt.Sprintf(`
UPDATE %s.%s s
SET link = $1
FROM %s.%s sbj, %s.%s st
WHERE s.subject_uuid = sbj.uuid
	AND s.subject_type_uuid = st.uuid
	AND s.group_uuid = $2
	AND s.start_time = $3
	AND s.start_date = $4
	AND s.end_date = $5
	AND s.weekday = $6
	AND s.is_session = $7
	AND sbj.name = $8
	AND st.type = $9
	AND COALESCE(s.link, '') <> $1`,
		constRepository.RASPYX_SCHEMA,
		constRepository.SCHEDULE_TABLE,
		constRepository.RASPYX_SCHEMA,
		constRepository.SUBJECTS_TABLE,
		constRepository.RASPYX_SCHEMA,
		constRepository.SUBJECT_TYPES_TABLE,
	)

	result, err := r.Pool.Exec(context.Background(), query,
		filters.Link,
		filters.GroupUUID,
		filters.StartTime,
		filters.StartDate,
		filters.EndDate,
		filters.Weekday,
		filters.IsSession,
		filters.SubjectName,
		filters.SubjectType,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
