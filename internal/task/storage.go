package task

import (
	"errors"
	"strconv"
	"sync"
)

type inMemoryStorage struct {
	mu      sync.RWMutex
	counter int
	tasks   map[string]*Task
}

var store = &inMemoryStorage{
	tasks: make(map[string]*Task),
}

// generateID возвращает уникальный числовой ID как строку.
func (s *inMemoryStorage) generateID() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	return strconv.Itoa(s.counter)
}

func (s *inMemoryStorage) save(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
}

func (s *inMemoryStorage) get(id string) (*Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *inMemoryStorage) delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}

func (s *inMemoryStorage) update(id string, updateFn func(*Task)) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	updateFn(task)
	return nil
}
