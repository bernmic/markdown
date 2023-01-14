package markdown

type HorizontalLine struct {
}

func (m *Markdown) AddHorizontalLine() *Markdown {
	m.elements = append(m.elements, &HorizontalLine{})
	return m
}
func (o *HorizontalLine) render(m *Markdown) ([]byte, error) {
	return []byte("---"), nil
}
