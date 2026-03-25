package application

import "sync"

type TrackStep string

const (
	TrackStepIdle    TrackStep = ""
	TrackStepLink    TrackStep = "link"
	TrackStepTags    TrackStep = "tags"
	TrackStepFilters TrackStep = "filters"
)

type TrackDialog struct {
	Step    TrackStep
	Link    string
	Tags    []string
	Filters []string
}

type TrackStateStorage struct {
	mu   sync.RWMutex
	data map[int64]TrackDialog
}

func NewTrackStateStorage() *TrackStateStorage {
	return &TrackStateStorage{
		data: make(map[int64]TrackDialog),
	}
}

func (s *TrackStateStorage) Start(chatID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[chatID] = TrackDialog{
		Step: TrackStepLink,
	}
}

func (s *TrackStateStorage) Get(chatID int64) (TrackDialog, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	state, ok := s.data[chatID]
	return state, ok
}

func (s *TrackStateStorage) Update(chatID int64, state TrackDialog) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[chatID] = state
}

func (s *TrackStateStorage) Delete(chatID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, chatID)
}

func (s *TrackStateStorage) Exists(chatID int64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[chatID]
	return ok
}