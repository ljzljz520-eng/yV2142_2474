package service

import (
	"idiomchain/internal/model"
	"sort"
)

func SortEntries(v []model.IdiomEntry) []model.IdiomEntry {
	out := append([]model.IdiomEntry{}, v...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func Tail(v []model.IdiomEntry) string {
	if len(v) == 0 {
		return ""
	}
	return v[len(v)-1].Tail
}
func AcceptedCount(v []model.IdiomEntry) int { return len(v) }
