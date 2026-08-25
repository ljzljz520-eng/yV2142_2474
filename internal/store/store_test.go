package store

import (
	"idiomchain/internal/model"
	"testing"
	"time"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(":memory:")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v := model.NewSession("s", "demo", time.Now())
	if e = s.SaveSession(v); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Session("s"); e != nil {
		t.Fatal(e)
	}
}
func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db.sqlite"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	v := model.NewSession("s", "demo", time.Now())
	_ = s.SaveSession(v)
	_ = s.AddEntry(model.NewEntry("e", "s", "天天", "天", time.Now()))
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	a, e := s.AllEntries("s")
	if e != nil || len(a) != 1 {
		t.Fatalf("%v %d", e, len(a))
	}
}
