package api

import (
	"example.com/emblem/service"
	"example.com/emblem/storage"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestRoutes(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "a"))
	defer s.Close()
	r := New(service.New(s)).Routes()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/launch/missing", nil))
	if w.Code != 404 {
		t.Fatal(w.Code)
	}
}
