package markdown

import (
	"fmt"
	"strings"
)

type Table struct {
	header []string
	data   [][]string
}

func NewTable(header []string, data [][]string) *Table {
	return &Table{
		header: header,
		data:   data,
	}
}

func (m *Markdown) AddTable(header []string, data [][]string) *Markdown {
	m.elements = append(m.elements, NewTable(header, data))
	return m
}

func (t *Table) render(m *Markdown) ([]byte, error) {
	result := ""
	cols := len(t.header)
	if cols > 0 {
		l1 := "|"
		l2 := "|"
		for i := 0; i < cols; i++ {
			s, l, r := justification(t.header[i])
			l1 += fmt.Sprintf(" %s |", s)
			if l && r {
				l2 += " :---: |"
			} else if l {
				l2 += " :--- |"
			} else if r {
				l2 += " ---: |"
			} else {
				l2 += " --- |"
			}

		}
		result += fmt.Sprintf("%s%s%s%s", l1, m.lineBreak, l2, m.lineBreak)
		for n, row := range t.data {
			if len(row) < cols {
				return nil, fmt.Errorf("error in render table. Line %d has not enough elements", n+2)
			}
			l := "|"
			for i := 0; i < cols; i++ {
				l += fmt.Sprintf(" %s |", row[i])
			}
			result += fmt.Sprintf("%s%s", l, m.lineBreak)
		}
	}
	return []byte(result), nil
}

func justification(s string) (string, bool, bool) {
	left := false
	right := false
	start := 0
	end := len(s)
	if strings.HasPrefix(s, ":") {
		left = true
		start++
	}
	if strings.HasSuffix(s, ":") {
		right = true
		end--
	}
	return s[start:end], left, right
}
