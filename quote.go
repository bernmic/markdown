package markdown

import (
	"fmt"
)

type Quote struct {
	text string
}

func (m *Markdown) AddQuote(text string) *Markdown {
	m.elements = append(m.elements, NewQuote(text))
	return m
}

func NewQuote(text string) *Quote {
	return &Quote{
		text: text,
	}
}

func (q *Quote) render(m *Markdown) ([]byte, error) {
	s := fmt.Sprintf("> %s%s", q.text, m.lineBreak)
	return []byte(s), nil
}
