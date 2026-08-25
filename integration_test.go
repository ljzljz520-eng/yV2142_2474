package idiomchain

import (
	"idiomchain/internal/clock"
	"idiomchain/internal/report"
	"idiomchain/internal/service"
	"idiomchain/internal/store"
	"testing"
)

func TestWorkflowCreateAndRecord(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	v := service.New(s, clock.Fixed{})
	sess, e := v.EnsureSession("demo")
	if e != nil {
		t.Fatal(e)
	}
	if _, e = v.Record(sess, "天高云淡"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowRejectsInvalidLink(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	v := service.New(s, clock.Fixed{})
	sess, _ := v.EnsureSession("demo")
	_, _ = v.Record(sess, "天高云淡")
	if _, e := v.Record(sess, "海阔天空"); e == nil {
		t.Fatal()
	}
}
func TestWorkflowReport(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	v := service.New(s, clock.Fixed{})
	sess, _ := v.EnsureSession("demo")
	_, _ = v.Record(sess, "天高云淡")
	e, _ := v.History(sess, 1, 20)
	r, _ := v.Rejections(sess)
	if report.RenderHistory(sess, e, r) == "" {
		t.Fatal()
	}
}
