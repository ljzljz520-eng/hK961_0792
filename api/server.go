package api

import (
	"encoding/json"
	"example.com/emblem/model"
	"example.com/emblem/query"
	"example.com/emblem/service"
	"net/http"
	"strings"
)

type Server struct {
	Query   *query.Query
	Service *service.Service
}

func New(s *service.Service) *Server { return &Server{query.New(s), s} }
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/launch/", s.handleLaunch)
	m.HandleFunc("/teams", s.handleTeam)
	return m
}
func (s *Server) handleLaunch(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/launch/")
	v, e := s.Query.Launch(id)
	if e != nil {
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func (s *Server) handleTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "method", 405)
		return
	}
	var t model.Team
	if json.NewDecoder(r.Body).Decode(&t) != nil {
		http.Error(w, "json", 400)
		return
	}
	if e := s.Service.UpdateTeam(t); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(t)
}
func WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
func DecodeTeam(r *http.Request) (model.Team, error) {
	var t model.Team
	e := json.NewDecoder(r.Body).Decode(&t)
	return t, e
}
func StatusForError(e error) int {
	if e == nil {
		return 200
	}
	if strings.Contains(e.Error(), "inactive") {
		return 404
	}
	return 400
}
