package markdown

import "fmt"

type Text struct {
	text   string
	bold   bool
	italic bool
	quoted bool
	sub    bool
	super  bool
}

func NewText(s string) Text {
	t := Text{
		text:   s,
		bold:   false,
		italic: false,
		quoted: false,
	}
	return t
}

func BoldText(s string) Text {
	t := Text{
		text:   s,
		bold:   true,
		italic: false,
		quoted: false,
	}
	return t
}

func (m *Markdown) AddText(text string, bold bool, italic bool, quoted bool) *Markdown {
	m.elements = append(m.elements, &Text{
		text:   text,
		bold:   bold,
		italic: italic,
		quoted: quoted,
	})
	return m
}

func ItalicText(s string) Text {
	t := Text{
		text:   s,
		bold:   false,
		italic: true,
		quoted: false,
	}
	return t
}

func QuotedText(s string) *Text {
	t := Text{
		text:   s,
		bold:   false,
		italic: false,
		quoted: true,
	}
	return &t
}

func (t *Text) SetText(s string) *Text {
	t.text = s
	return t
}

func (t *Text) SetBold(b bool) *Text {
	t.bold = b
	return t
}

func (t *Text) SetItalic(b bool) *Text {
	t.italic = b
	return t
}

func (t *Text) SetQuote(b bool) *Text {
	t.quoted = b
	return t
}

func (t *Text) SetSub(b bool) *Text {
	t.sub = b
	return t
}

func (t *Text) SetSuper(b bool) *Text {
	t.super = b
	return t
}

func (t *Text) IsBold() bool {
	return t.bold
}

func (t *Text) IsItalic() bool {
	return t.italic
}

func (t *Text) IsQuoted() bool {
	return t.quoted
}

func (t *Text) IsSub() bool {
	return t.sub
}

func (t *Text) IsSuper() bool {
	return t.super
}

func (t *Text) render(m *Markdown) ([]byte, error) {
	s := t.text
	if t.bold {
		s = fmt.Sprintf("**%s**", s)
	}
	if t.italic {
		s = fmt.Sprintf("*%s*", s)
	}
	if t.quoted {
		s = fmt.Sprintf("`%s`", s)
	}
	if t.sub {
		s = fmt.Sprintf("<sub>%s</sub>", s)
	}
	if t.super {
		s = fmt.Sprintf("<sup>%s</sup>", s)
	}
	return []byte(s), nil
}
