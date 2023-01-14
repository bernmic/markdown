package markdown

import "fmt"

type List struct {
	items   []string
	ordered bool
}

func (m *Markdown) AddList(items []string, ordered bool) *Markdown {
	m.elements = append(m.elements, NewList(items, ordered))
	return m
}

func NewList(items []string, ordered bool) *List {
	return &List{
		items:   items,
		ordered: ordered,
	}
}

func (l *List) render(m *Markdown) ([]byte, error) {
	result := ""
	for i, s := range l.items {
		if l.ordered {
			result += fmt.Sprintf("%d. %s%s", i+1, s, m.lineBreak)
		} else {
			result += fmt.Sprintf("* %s%s", s, m.lineBreak)
		}
	}
	return []byte(result), nil
}
