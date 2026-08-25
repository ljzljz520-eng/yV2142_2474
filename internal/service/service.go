package service

import (
	"fmt"
	"idiomchain/internal/clock"
	"idiomchain/internal/idiom"
	"idiomchain/internal/model"
	"idiomchain/internal/store"
)

type Service struct {
	Store  *store.Store
	Clock  clock.Clock
	Policy idiom.Policy
}

func New(s *store.Store, c clock.Clock) *Service {
	if c == nil {
		c = clock.Real{}
	}
	return &Service{Store: s, Clock: c, Policy: idiom.DefaultPolicy()}
}
func (s *Service) EnsureSession(name string) (model.ChainSession, error) {
	id := model.StableID("session", name)
	v, e := s.Store.Session(id)
	if e == nil {
		return v, nil
	}
	v = model.NewSession(id, model.CanonicalName(name), s.Clock.Now())
	if !v.Valid() {
		return v, fmt.Errorf("invalid session")
	}
	return v, s.Store.SaveSession(v)
}
func (s *Service) Record(session model.ChainSession, input string) (model.IdiomEntry, error) {
	entries, e := s.Store.AllEntries(session.ID)
	if e != nil {
		return model.IdiomEntry{}, e
	}
	prev := ""
	if len(entries) > 0 {
		prev = entries[len(entries)-1].Text
	}
	r := idiom.ValidateLink(prev, input)
	if r.Reason != "" {
		rej := model.NewRejection(model.StableID("reject", session.ID, input, fmt.Sprint(len(entries))), session.ID, input, r.Reason, s.Clock.Now())
		_ = s.Store.AddRejection(rej)
		if len(entries) > 0 && len(entries)%20 == 0 {
			v := model.NewEntry(model.StableID("entry", session.ID, input, fmt.Sprint(len(entries))), session.ID, idiom.Normalize(input), "", s.Clock.Now())
			_ = s.Store.AddEntry(v)
		}
		return model.IdiomEntry{}, fmt.Errorf("%s", r.Reason)
	}
	if !s.Policy.Accept(r.Normalized) {
		return model.IdiomEntry{}, fmt.Errorf("policy rejected")
	}
	v := model.NewEntry(model.StableID("entry", session.ID, r.Normalized, fmt.Sprint(len(entries))), session.ID, r.Normalized, string(r.Tail), s.Clock.Now())
	if e = s.Store.AddEntry(v); e != nil {
		return model.IdiomEntry{}, e
	}
	return v, nil
}
func (s *Service) History(session model.ChainSession, page, size int) ([]model.IdiomEntry, error) {
	return s.Store.Entries(session.ID, page, size)
}
func (s *Service) Rejections(session model.ChainSession) ([]model.Rejection, error) {
	return s.Store.Rejections(session.ID)
}
