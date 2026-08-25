package report

import (
	"idiomchain/internal/model"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	s := model.ChainSession{Name: "x"}
	out := RenderHistory(s, []model.IdiomEntry{{Text: "天天"}}, nil)
	if !strings.Contains(out, "Accepted: 1") {
		t.Fatal(out)
	}
}
