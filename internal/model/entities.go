package model

import "time"

type IdiomEntry struct {
	ID, SessionID, Text, Tail string
	CreatedAt                 time.Time
}
type ChainSession struct {
	ID, Name  string
	CreatedAt time.Time
}
type Rejection struct {
	ID, SessionID, Input, Reason string
	CreatedAt                    time.Time
}
type Setting struct {
	Key, Value string
	UpdatedAt  time.Time
}

func NewEntry(id, session, text, tail string, at time.Time) IdiomEntry {
	return IdiomEntry{ID: id, SessionID: session, Text: text, Tail: tail, CreatedAt: at}
}
func NewSession(id, name string, at time.Time) ChainSession {
	return ChainSession{ID: id, Name: name, CreatedAt: at}
}
func NewRejection(id, session, input, reason string, at time.Time) Rejection {
	return Rejection{ID: id, SessionID: session, Input: input, Reason: reason, CreatedAt: at}
}
func (e IdiomEntry) Valid() bool {
	return e.ID != "" && e.SessionID != "" && e.Text != "" && e.Tail != ""
}
func (s ChainSession) Valid() bool { return s.ID != "" && s.Name != "" }
