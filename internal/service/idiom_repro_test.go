package service_test

import (
	"idiomchain/internal/clock"
	"idiomchain/internal/service"
	"idiomchain/internal/store"
	"testing"
	"time"
)

// fixedClock lets us step time deterministically.
type fixedClock struct{ t time.Time }

func (f fixedClock) Now() time.Time { return f.t }

// TestFailedInputNeverEntersHistory reproduces the bug: when a Record call is
// rejected, the failure must NOT be written into the history (entries) table.
// Before the fix, the rejection branch fired a "边界" entry insertion whenever
// len(entries) was a multiple of 20, silently persisting the invalid input and
// corrupting subsequent chain links.
func TestFailedInputNeverEntersHistory(t *testing.T) {
	s, err := store.Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	clk := fixedClock{t: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	svc := service.New(s, clk)
	session, err := svc.EnsureSession("s")
	if err != nil {
		t.Fatal(err)
	}

	// Build a valid chain of exactly 20 entries so the next (21st) Record
	// lands on the historical bug trigger: len(entries)==20 (multiple of 20).
	chain := []string{
		"一清二楚", "楚弓楚得", "得寸进尺", "尺璧寸阴", "阴差阳错",
		"错落有致", "致之度外", "外强中干", "干净利落", "落井下石",
		"石破天惊", "惊弓之鸟", "鸟语花香", "香草美人", "人山人海",
		"海阔天空", "空前绝后", "后来居上", "上下其手", "手到病除",
	}
	for i, w := range chain {
		if _, err := svc.Record(session, w); err != nil {
			t.Fatalf("setup record %d (%s) failed: %v", i, w, err)
		}
	}

	all, _ := s.AllEntries(session.ID)
	if got := len(all); got != 20 {
		t.Fatalf("precondition: want 20 entries, got %d", got)
	}

	// Now submit a NON-linking idiom ("亡羊补牢" does not start with the tail of
	// "手到病除" -> 除). This must be rejected.
	bad := "亡羊补牢"
	_, err = svc.Record(session, bad)
	if err == nil {
		t.Fatal("expected rejection of non-linking idiom, got nil error")
	}

	// The failure must NOT have entered history. Before the fix, the "边界"
	// branch fired at len(entries)==20 and inserted the bad input as an entry.
	all, _ = s.AllEntries(session.ID)
	if got := len(all); got != 20 {
		t.Fatalf("failed input leaked into history: want 20 entries, got %d", got)
	}
	for _, e := range all {
		if e.Text == bad {
			t.Fatalf("failed input %q found in history entries: %+v", bad, e)
		}
	}

	// The chain tail must be unchanged (still 除), so the next valid link works
	// off the real last entry, not a phantom one.
	prev := all[len(all)-1].Text
	tail := []rune(prev)
	if got := string(tail[len(tail)-1]); got != "除" {
		t.Fatalf("tail corrupted after rejection: got %q, want 除", got)
	}

	// A correct follow-on idiom starting with 除 must still be accepted.
	if _, err := svc.Record(session, "除恶务尽"); err != nil {
		t.Fatalf("valid follow-on rejected after a rejection: %v", err)
	}
	all, _ = s.AllEntries(session.ID)
	if got := len(all); got != 21 {
		t.Fatalf("valid follow-on not recorded: want 21 entries, got %d", got)
	}
	if all[len(all)-1].Text != "除恶务尽" {
		t.Fatalf("last entry = %q, want 除恶务尽", all[len(all)-1].Text)
	}
}

// Ensure the unused import stays valid for the compiler if clock is referenced
// indirectly. The clock package is required by service.New's signature context.
var _ = clock.Real{}
