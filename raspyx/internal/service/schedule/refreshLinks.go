package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"raspyx2/internal/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	linkRefreshBefore   = 20 * time.Minute
	linkRefreshAfter    = 10 * time.Minute
	linkRefreshThrottle = 30 * time.Second
)

func (s *ScheduleService) refreshUpcomingLinksByGroupNumber(groupNumber, groupUUID string, isSession bool, now time.Time) error {
	scheduleNow := toScheduleLocation(now)
	pairNums := upcomingPairNums(scheduleNow)
	if len(pairNums) == 0 {
		return nil
	}

	cacheKey := fmt.Sprintf("%s:%t", groupNumber, isSession)
	if !s.claimLinkRefresh(cacheKey, scheduleNow) {
		return nil
	}

	sourceSchedule, err := s.fetchGroupSchedule(groupNumber, isSession)
	if err != nil {
		return err
	}

	weekday := strconv.Itoa(int(scheduleNow.Weekday()))
	if weekday == "0" {
		return nil
	}

	date := scheduleNow.Format("2006-01-02")
	for gridKey, day := range sourceSchedule.Grid {
		lessonWeekday := gridKey
		if gridKey == date {
			lessonWeekday = weekday
		}
		if lessonWeekday != weekday {
			continue
		}

		for _, pairNum := range pairNums {
			lessons := day[pairNum]
			for i := range lessons {
				lesson := &lessons[i]
				link := lessonLink(lesson)
				if link == "" {
					continue
				}

				startTime, _ := pairNumToStartEnd(pairNum)
				affected, err := s.repo.ScheduleRepository.UpdateScheduleLinkByLesson(&models.UpdateScheduleLinkByLessonRequest{
					GroupUUID:   groupUUID,
					SubjectName: removeTrash(lesson.Sbj),
					SubjectType: removeTrash(lesson.Type),
					StartTime:   startTime,
					StartDate:   lesson.Df,
					EndDate:     lesson.Dt,
					Weekday:     mustAtoi(lessonWeekday),
					Link:        link,
					IsSession:   isSession,
				})
				if err != nil {
					return err
				}
				_ = affected
			}
		}
	}

	return nil
}

func (s *ScheduleService) fetchGroupSchedule(groupNumber string, isSession bool) (*models.ParseScheduleResponse, error) {
	session := 0
	if isSession {
		session = 1
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		"GET",
		fmt.Sprintf("https://rasp.dmami.ru/site/group?group=%v&session=%v", groupNumber, session),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "https://rasp.dmami.ru/")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed models.ParseScheduleResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("error unmarshling response %v: %w", parsed, err)
	}
	if parsed.Status != "ok" {
		return nil, fmt.Errorf("schedule source response error: %s", parsed.Message)
	}

	return &parsed, nil
}

func (s *ScheduleService) claimLinkRefresh(key string, now time.Time) bool {
	s.linkRefreshMu.Lock()
	defer s.linkRefreshMu.Unlock()

	lastRefresh, ok := s.linkRefreshCache[key]
	if ok && now.Sub(lastRefresh) < linkRefreshThrottle {
		return false
	}

	s.linkRefreshCache[key] = now
	return true
}

func upcomingPairNums(now time.Time) []string {
	pairNums := make([]string, 0, 1)
	for pairNum, start := range pairStartClock() {
		pairStart := time.Date(now.Year(), now.Month(), now.Day(), start.hour, start.minute, 0, 0, now.Location())
		if !now.Before(pairStart.Add(-linkRefreshBefore)) && !now.After(pairStart.Add(linkRefreshAfter)) {
			pairNums = append(pairNums, pairNum)
		}
	}

	return pairNums
}

func pairStartClock() map[string]struct {
	hour   int
	minute int
} {
	return map[string]struct {
		hour   int
		minute int
	}{
		"1": {hour: 9, minute: 0},
		"2": {hour: 10, minute: 40},
		"3": {hour: 12, minute: 20},
		"4": {hour: 14, minute: 30},
		"5": {hour: 16, minute: 10},
		"6": {hour: 17, minute: 50},
		"7": {hour: 19, minute: 30},
	}
}

func pairNumToStartEnd(pair string) (string, string) {
	switch pair {
	case "1":
		return "09:00:00", "10:30:00"
	case "2":
		return "10:40:00", "12:10:00"
	case "3":
		return "12:20:00", "13:50:00"
	case "4":
		return "14:30:00", "16:00:00"
	case "5":
		return "16:10:00", "17:40:00"
	case "6":
		return "17:50:00", "19:20:00"
	case "7":
		return "19:30:00", "21:00:00"
	default:
		return "", ""
	}
}

func toScheduleLocation(t time.Time) time.Time {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return t
	}

	return t.In(location)
}

func lessonLink(lesson *models.Lesson) string {
	for _, auditory := range lesson.Auditories {
		if link := linkFromHTML(auditory.Title); link != "" {
			return link
		}
	}

	return ""
}

func linkFromHTML(html string) string {
	urlRegex := regexp.MustCompile(`https?://[^\s"'<>]+`)
	return urlRegex.FindString(html)
}

func removeTrash(s string) string {
	return strings.TrimSpace(removeEmojis(removeHTML(s)))
}

func removeEmojis(text string) string {
	emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]`)
	return strings.TrimSpace(emojiRegex.ReplaceAllString(text, ""))
}

func removeHTML(text string) string {
	htmlRegex := regexp.MustCompile(`>.*<`)
	newText := htmlRegex.FindString(text)
	if newText == "" {
		return text
	}
	return strings.TrimSpace(newText[1 : len(newText)-1])
}

func mustAtoi(value string) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return parsed
}
