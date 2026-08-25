package idiom

import (
	"fmt"
	"idiomchain/internal/model"
	"strings"
)

type Result struct {
	Normalized string
	Tail       rune
	Reason     string
}

func Normalize(v string) string { return strings.Join(strings.Fields(strings.TrimSpace(v)), "") }
func ValidateLink(previous, candidate string) Result {
	n := Normalize(candidate)
	if n == "" {
		return Result{Reason: "empty idiom"}
	}
	if !model.IsHanText(n) {
		return Result{Reason: "idiom must contain two Han characters"}
	}
	if previous != "" && model.FirstRune(n) != model.LastRune(previous) {
		return Result{Normalized: n, Reason: fmt.Sprintf("must start with %c", model.LastRune(previous))}
	}
	return Result{Normalized: n, Tail: model.LastRune(n)}
}
func IsValid(previous, candidate string) bool { return ValidateLink(previous, candidate).Reason == "" }
func Explain(previous, candidate string) string {
	r := ValidateLink(previous, candidate)
	if r.Reason == "" {
		return "valid"
	}
	return r.Reason
}
