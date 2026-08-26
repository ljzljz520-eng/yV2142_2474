package idiom

import "idiomchain/internal/model"

func Last(entries []model.IdiomEntry) string {
	if len(entries) == 0 {
		return ""
	}
	return entries[len(entries)-1].Text
}
func Contains(entries []model.IdiomEntry, text string) bool {
	for _, e := range entries {
		if e.Text == text {
			return true
		}
	}
	return false
}
func Filter(entries []model.IdiomEntry, prefix string) []model.IdiomEntry {
	out := make([]model.IdiomEntry, 0)
	for _, e := range entries {
		if prefix == "" || len(e.Text) > 0 && string([]rune(e.Text)[0]) == prefix {
			out = append(out, e)
		}
	}
	return out
}
