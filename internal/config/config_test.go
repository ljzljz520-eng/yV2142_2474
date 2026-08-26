package config

import "testing"

func TestConfig(t *testing.T) {
	if Load().DBPath == "" {
		t.Fatal()
	}
	if ForPath("", "x").Validate() != nil {
		t.Fatal()
	}
}
