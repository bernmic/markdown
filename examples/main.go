package main

import (
	"fmt"
	"markdown"
)

func main() {
	m := markdown.NewMarkdown()
	m.AddHeadline("Headline 1", 1).AddHeadline("Headline 2", 2).AddHeadline("Headline 3", 3).AddHorizontalLine()
	m.AddText("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", false, false, false)
	m.AddImage("ITA", "https://upload.wikimedia.org/wikipedia/en/thumb/0/03/Flag_of_Italy.svg/188px-Flag_of_Italy.svg.png", "Italy")
	m.AddQuote("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.")
	m.Append(table())
	m.AddList([]string{"First", "Second", "Third"}, false)
	m.AddList([]string{"First", "Second", "Third"}, true)
	m.AddCode([]string{"for _, v := range l {", "  doSomething(v)", "}"})
	b, err := m.AddText("BOLD TEXT", true, false, false).AddLinebreak().AddLink("Wiki", "https://www.wikipedia.org", "Wikipedia").Render()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(string(b))
}

func table() *markdown.Table {
	h := []string{"Country", ":Capital", "Population:", "Area:"}
	d := [][]string{
		{"Great Britain", "London", "60,800,000", "209,331"},
		{"France", "Paris", "67,897,000", "643,801"},
		{"Spain", "Madrid", "47,325,360", "505,990"},
		{"Germany", "Berlin", "83,695,430", "357,592"},
		{"Italy", "Rome", "58,853,482", "301,230"},
		{"Luxembourg", "Luxembourg", "645,397", "2,586"},
		{"Belgium", "Bruselles", "11,584,008", "30,528"},
		{"Netherlands", "Amsterdam", "17,785,200", "41,850"},
	}
	return markdown.NewTable(h, d)
}
