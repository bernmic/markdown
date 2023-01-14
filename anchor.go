package markdown

import "fmt"

type Anchor struct {
	name  string
	url   string
	title string
}

func NewAnchor(name string, url string, title string) *Anchor {
	return &Anchor{name: name, url: url, title: title}
}

func (m *Markdown) AddAnchor(name string, url string, title string) *Markdown {
	m.elements = append(m.elements, NewAnchor(name, url, title))
	return m
}

func (l *Anchor) render(m *Markdown) ([]byte, error) {
	s := fmt.Sprintf("[%s] %s \"%s\"", l.name, l.url, l.title)
	return []byte(s), nil
}
