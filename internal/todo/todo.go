package todo

import (
	"errors"
	"strings"
	"sync"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var ErrEmptyTitle = errors.New("title is required")

type Store interface {
	Create(title string) (Todo, error)
	List() ([]Todo, error)
}

type MemoryStore struct {
	mu    sync.Mutex
	next  int
	items []Todo
}

func NewMemoryStore() *MemoryStore { return &MemoryStore{next: 1} }

func (m *MemoryStore) Create(title string) (Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Todo{}, ErrEmptyTitle
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	t := Todo{ID: m.next, Title: title}
	m.next++
	m.items = append(m.items, t)
	return t, nil
}

func (m *MemoryStore) List() ([]Todo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]Todo, len(m.items))
	copy(out, m.items)
	return out, nil
}
