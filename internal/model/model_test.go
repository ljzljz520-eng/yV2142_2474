package model

import (
	"testing"
	"time"
)

func TestEntities(t *testing.T) {
	now := time.Unix(1, 0)
	if !NewSession("s", "name", now).Valid() {
		t.Fatal()
	}
	if !NewEntry("e", "s", "天天", "天", now).Valid() {
		t.Fatal()
	}
	if !IsHanText("天天") {
		t.Fatal()
	}
}
func TestStableID(t *testing.T) {
	if StableID("a") != StableID("a") {
		t.Fatal()
	}
}
