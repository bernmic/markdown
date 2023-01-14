# Markdown

Simple library to create Markdown data.

## Quickstart

```
m := markdown.NewMarkdown()
m.AddHeadline("Headline 1", 1).AddHeadline("Headline 2", 2).AddHeadline("Headline 3", 3).AddHorizontalLine()
m.AddList([]string{"First", "Second", "Third"}, false)
b, _ := m.AddLink("Wiki", "https://www.wikipedia.org", "Wikipedia").Render()
fmt.Println(string(b))    
```