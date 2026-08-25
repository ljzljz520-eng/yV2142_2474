package idiom_test

import (
	"idiomchain/internal/clock"
	"idiomchain/internal/idiom"
	"idiomchain/internal/service"
	"idiomchain/internal/store"
	"testing"
)

func TestValidateLink(t *testing.T) {
	if !idiom.IsValid("", "天高云淡") {
		t.Fatal()
	}
	if idiom.IsValid("天高云淡", "海阔天空") {
		t.Fatal()
	}
}
func TestInvalidLinkNotRecorded(t *testing.T) {
	s, _ := store.Open(":memory:")
	defer s.Close()
	svc := service.New(s, clock.Fixed{})
	session, _ := svc.EnsureSession("boundary")
	for i := 0; i < 20; i++ {
		if _, err := svc.Record(session, "天天"); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := svc.Record(session, "海海"); err == nil {
		t.Fatal("invalid input should fail")
	}
	entries, _ := svc.History(session, 1, 100)
	if len(entries) != 20 {
		t.Fatalf("invalid input entered history: %d", len(entries))
	}
}
func TestPolicy(t *testing.T) {
	if !idiom.DefaultPolicy().Accept("天天") {
		t.Fatal()
	}
	if idiom.PageOffset(2, 20) != 20 {
		t.Fatal()
	}
}
