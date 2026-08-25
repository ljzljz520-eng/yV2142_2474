package report

import "idiomchain/internal/catalog"

func SuggestionCount() int {
	return catalog.Count1() + catalog.Count2() + catalog.Count3() + catalog.Count4() + catalog.Count5() + catalog.Count6()
}
func Suggestions(prefix string) []string {
	out := catalog.Lookup1(prefix)
	out = append(out, catalog.Lookup2(prefix)...)
	out = append(out, catalog.Lookup3(prefix)...)
	return out
}
