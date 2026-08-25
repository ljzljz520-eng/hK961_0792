package model

import "strings"

func NormalizeSlug(v string) string   { return strings.ToLower(strings.TrimSpace(v)) }
func NormalizeSlogan(v string) string { return strings.Join(strings.Fields(v), " ") }
func IsPublished(a Announcement) bool { return a.Published && len(a.Body) > 0 }
func EventName(v string) string       { return strings.ToUpper(strings.TrimSpace(v)) }
func ActiveTeam(t Team) bool          { return t.Active }
