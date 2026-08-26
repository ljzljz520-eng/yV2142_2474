package clock

import "time"

type Sequence struct {
	Values []time.Time
	Index  int
}

func (s *Sequence) Now() time.Time {
	if len(s.Values) == 0 {
		return time.Unix(0, 0).UTC()
	}
	if s.Index >= len(s.Values) {
		return s.Values[len(s.Values)-1]
	}
	v := s.Values[s.Index]
	s.Index++
	return v
}
func At(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
