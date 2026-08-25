package api

import (
	"example.com/emblem/model"
	"net/http"
)

func ValidateTeamRequest(t model.Team) error            { return t.Validate() }
func MethodAllowed(r *http.Request, method string) bool { return r.Method == method }
func PathID(path, prefix string) string {
	if len(path) < len(prefix) {
		return ""
	}
	return path[len(prefix):]
}
