package schedule

import (
	"raspyx2/internal/models"
	"testing"
	"time"
)

func TestLessonLinkFallsBackToAuditoryHTML(t *testing.T) {
	lesson := &models.Lesson{
		Auditories: []struct {
			Title string `json:"title"`
			Color string `json:"color"`
		}{
			{Title: `<a href="https://online.example.test/lesson">online</a>`},
		},
	}

	if got := lessonLink(lesson); got != "https://online.example.test/lesson" {
		t.Fatalf("lessonLink() = %q", got)
	}
}

func TestUpcomingPairNums(t *testing.T) {
	location := time.FixedZone("MSK", 3*60*60)

	beforeFirstPair := time.Date(2026, 6, 9, 8, 41, 0, 0, location)
	if got := upcomingPairNums(beforeFirstPair); len(got) != 1 || got[0] != "1" {
		t.Fatalf("upcomingPairNums() before first pair = %#v", got)
	}

	tooEarly := time.Date(2026, 6, 9, 8, 39, 0, 0, location)
	if got := upcomingPairNums(tooEarly); len(got) != 0 {
		t.Fatalf("upcomingPairNums() too early = %#v", got)
	}
}
