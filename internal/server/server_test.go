package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/navneetshukl/go-cicd-demo/internal/server"
	"github.com/navneetshukl/go-cicd-demo/internal/todo"
)

func do(t *testing.T, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	r := server.New(todo.NewMemoryStore())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestHealthz(t *testing.T) {
	if w := do(t, http.MethodGet, "/healthz", ""); w.Code != http.StatusOK {
		t.Fatalf("got %d, want 200", w.Code)
	}
}

func TestCreateTodo(t *testing.T) {
	w := do(t, http.MethodPost, "/todos", `{"title":"write tests"}`)
	if w.Code != http.StatusCreated {
		t.Fatalf("got %d, want 201", w.Code)
	}
	if !strings.Contains(w.Body.String(), "write tests") {
		t.Fatalf("body missing title: %s", w.Body.String())
	}
}

func TestCreateTodoEmptyTitle(t *testing.T) {
	if w := do(t, http.MethodPost, "/todos", `{"title":"  "}`); w.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want 400", w.Code)
	}
}