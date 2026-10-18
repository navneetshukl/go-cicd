package todo_test

import (
	"errors"
	"testing"

	"github.com/navneetshukl/go-cicd/internal/todo"
)

func TestMemoryStore(t *testing.T) {
	s := todo.NewMemoryStore()

	if _, err := s.Create(""); !errors.Is(err, todo.ErrEmptyTitle) {
		t.Fatalf("want ErrEmptyTitle, got %v", err)
	}
	a, _ := s.Create("a")
	b, _ := s.Create("b")
	if a.ID != 1 || b.ID != 2 {
		t.Fatalf("unexpected ids %d %d", a.ID, b.ID)
	}
	items, _ := s.List()
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d", len(items))
	}
}
