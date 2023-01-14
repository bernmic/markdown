package markdown

import "fmt"

type Markdown struct {
	elements  []MarkdownObject
	lineBreak string
}

type MarkdownObject interface {
	render(m *Markdown) ([]byte, error)
}

func NewMarkdown() *Markdown {
	m := Markdown{lineBreak: "\n"}
	m.elements = make([]MarkdownObject, 0)
	return &m
}

func (m *Markdown) LineBreak(l string) {
	m.lineBreak = l
}

func (m *Markdown) Append(mo MarkdownObject) *Markdown {
	m.elements = append(m.elements, mo)
	return m
}

func (m *Markdown) Render() ([]byte, error) {
	result := make([]byte, 0)
	for _, e := range m.elements {
		o, err := e.render(m)
		if err != nil {
			return nil, fmt.Errorf("error render element: %v", err)
		}
		result = append(result, o...)
		if m.lineBreak != "" {
			result = append(result, []byte(m.lineBreak)...)
		}
	}
	return result, nil
}

type Linebreak struct {
}

func (n *Linebreak) render(m *Markdown) ([]byte, error) {
	return []byte("<br/>"), nil
}

func (m *Markdown) AddLinebreak() *Markdown {
	m.elements = append(m.elements, &Linebreak{})
	return m
}
