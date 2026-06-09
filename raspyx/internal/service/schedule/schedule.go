package schedule

import (
	"net/http"
	"raspyx2/config"
	"raspyx2/internal/repository"
	"sync"
	"time"
)

type ScheduleService struct {
	cfg              *config.Config
	repo             *repository.Repository
	client           *http.Client
	linkRefreshMu    sync.Mutex
	linkRefreshCache map[string]time.Time
}

func NewScheduleService(cfg *config.Config, repo *repository.Repository) *ScheduleService {
	requestTimeout := time.Duration(cfg.Parser.RequestTimeout) * time.Millisecond
	if requestTimeout <= 0 {
		requestTimeout = 5 * time.Second
	}

	return &ScheduleService{
		cfg:              cfg,
		repo:             repo,
		client:           &http.Client{Timeout: requestTimeout},
		linkRefreshCache: make(map[string]time.Time),
	}
}
