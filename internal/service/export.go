package service

import (
	"idiomchain/internal/model"
	"strings"
)

func JoinTexts(v []model.IdiomEntry) string {
	a := make([]string, len(v))
	for i, e := range v {
		a[i] = e.Text
	}
	return strings.Join(a, " -> ")
}
func Reasons(v []model.Rejection) []string {
	a := make([]string, len(v))
	for i, e := range v {
		a[i] = e.Reason
	}
	return a
}
func SessionLabel(v model.ChainSession) string { return v.Name + " (" + v.ID + ")" }
