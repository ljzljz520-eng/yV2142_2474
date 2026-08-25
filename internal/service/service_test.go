package service

import (
	"idiomchain/internal/clock"
	"idiomchain/internal/store"
	"testing"
)

func TestRecord(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	v := New(s, clock.Fixed{})
	sess, _ := v.EnsureSession("x")
	if _, e := v.Record(sess, "天高云淡"); e != nil {
		t.Fatal(e)
	}
	if _, e := v.Record(sess, "淡泊明志"); e != nil {
		t.Fatal(e)
	}
}
func TestReject(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	v := New(s, clock.Fixed{})
	sess, _ := v.EnsureSession("x")
	_, _ = v.Record(sess, "天高云淡")
	if _, e := v.Record(sess, "海阔天空"); e == nil {
		t.Fatal()
	}
}
