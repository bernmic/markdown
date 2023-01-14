package markdown

import (
	"fmt"
)

type Image struct {
	text  string
	url   string
	title string
}

func NewImage(text string, url string, title string) *Image {
	return &Image{text: text, url: url, title: title}
}

func (m *Markdown) AddImage(text string, url string, title string) *Markdown {
	m.elements = append(m.elements, NewImage(text, url, title))
	return m
}

func (l *Image) render(m *Markdown) ([]byte, error) {
	s := ""
	if l.title == "" {
		s = fmt.Sprintf("![%s](%s)", l.text, l.url)
	} else {
		s = fmt.Sprintf("%s![%s](%s \"%s\")", m.lineBreak, l.text, l.url, l.title)
	}
	return []byte(s), nil
}
