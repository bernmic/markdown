package markdown

import (
	"fmt"
	"strings"
)

type Headline struct {
	text  string
	level int
}

func (m *Markdown) AddHeadline(text string, level int) *Markdown {
	m.elements = append(m.elements, NewHeadline(text, level))
	return m
}

func NewHeadline(text string, level int) *Headline {
	return &Headline{
		text:  text,
		level: level,
	}
}

func (p *Headline) render(m *Markdown) ([]byte, error) {
	s := fmt.Sprintf("%s %s", strings.Repeat("#", p.level), p.text)
	return []byte(s), nil
}
