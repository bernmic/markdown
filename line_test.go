package markdown

import "testing"

func Test_line(t *testing.T) {
	m := NewMarkdown()
	l := HorizontalLine{}
	m.Append(&l)
	b, err := m.Render()
	if err != nil {
		t.Errorf("error rendering line: %v", err)
	}
	if string(b) != "---" {
		t.Errorf("expected '---', got '%s'", string(b))
	}
}
