package clock

import "testing"

func TestFixed(t *testing.T) {
	f := Fixed{Value: At(2024, 1, 1)}
	if f.Now().Year() != 2024 {
		t.Fatal()
	}
}
