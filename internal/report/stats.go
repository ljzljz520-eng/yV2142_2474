package report

import "idiomchain/internal/model"

type Stats struct{ Accepted, Rejected int }

func Compute(entries []model.IdiomEntry, rejections []model.Rejection) Stats {
	return Stats{Accepted: len(entries), Rejected: len(rejections)}
}
func Empty(s Stats) bool { return s.Accepted == 0 && s.Rejected == 0 }
func Ratio(s Stats) float64 {
	if s.Accepted+s.Rejected == 0 {
		return 0
	}
	return float64(s.Accepted) / float64(s.Accepted+s.Rejected)
}
