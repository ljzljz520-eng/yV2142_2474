package idiom

type Policy struct {
	MaxLength   int
	AllowRepeat bool
}

func DefaultPolicy() Policy { return Policy{MaxLength: 12} }
func (p Policy) Accept(text string) bool {
	if p.MaxLength > 0 && len([]rune(text)) > p.MaxLength {
		return false
	}
	return text != ""
}
func (p Policy) DuplicateAllowed() bool { return p.AllowRepeat }
func PageOffset(page, size int) int {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	return (page - 1) * size
}
