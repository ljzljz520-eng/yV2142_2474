package model

import (
	"strings"
	"unicode/utf8"
)

func CanonicalName(v string) string { return strings.Join(strings.Fields(strings.TrimSpace(v)), " ") }
func IsHanText(v string) bool {
	if utf8.RuneCountInString(v) < 2 {
		return false
	}
	for _, r := range v {
		if r < '一' || r > '龥' {
			return false
		}
	}
	return true
}
func FirstRune(v string) rune {
	for _, r := range v {
		return r
	}
	return 0
}
func LastRune(v string) rune {
	var last rune
	for _, r := range v {
		last = r
	}
	return last
}
func SafeLimit(v, fallback int) int {
	if v <= 0 {
		return fallback
	}
	if v > 1000 {
		return 1000
	}
	return v
}
