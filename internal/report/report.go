package report

import (
	"fmt"
	"idiomchain/internal/model"
	"strings"
)

func RenderHistory(session model.ChainSession, entries []model.IdiomEntry, rejections []model.Rejection) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Session: %s\n", session.Name)
	for i, e := range entries {
		fmt.Fprintf(&b, "%d. %s [%s]\n", i+1, e.Text, e.CreatedAt.Format("2006-01-02T15:04:05Z"))
	}
	fmt.Fprintf(&b, "Accepted: %d\nRejected: %d\n", len(entries), len(rejections))
	return b.String()
}
func RenderCompact(entries []model.IdiomEntry) string {
	a := make([]string, len(entries))
	for i, e := range entries {
		a[i] = e.Text
	}
	return strings.Join(a, " -> ")
}
