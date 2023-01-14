package markdown

import (
	"fmt"
)

type Link struct {
	text  string
	url   string
	title string
}

func NewLink(text string, url string, title string) *Link {
	return &Link{text: text, url: url, title: title}
}

func (m *Markdown) AddLink(text string, url string, title string) *Markdown {
	m.elements = append(m.elements, NewLink(text, url, title))
	return m
}

func (l *Link) render(m *Markdown) ([]byte, error) {
	s := ""
	if l.title == "" {
		s = fmt.Sprintf("[%s](%s)", l.text, l.url)
	} else {
		s = fmt.Sprintf("[%s](%s \"%s\")", l.text, l.url, l.title)
	}
	return []byte(s), nil
}
