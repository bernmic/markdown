package markdown

import "fmt"

type Code struct {
	lines   []string
	ordered bool
}

func (m *Markdown) AddCode(lines []string) *Markdown {
	m.elements = append(m.elements, NewCode(lines))
	return m
}

func NewCode(lines []string) *Code {
	return &Code{
		lines: lines,
	}
}

func (l *Code) render(m *Markdown) ([]byte, error) {
	result := ""
	for _, s := range l.lines {
		result += fmt.Sprintf("\t%s%s", s, m.lineBreak)
	}
	return []byte(result), nil
}
