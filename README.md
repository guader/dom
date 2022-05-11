## Dom

Write HTML document with golang.

## Example

```go
package main

import (
	"net/http"

	"github.com/guader/dom"
)

func main() {
	doc := dom.Document(dom.HTML(
		dom.Head(
			dom.Meta().SetProp(dom.CHARSET, "utf-8"),
			dom.Style(
				dom.CSS(".p1").SetProp(dom.COLOR, "red"),
				dom.CSS("#world").SetProp(dom.COLOR, "blue"),
			),
		),
		dom.Body(
			dom.H1(dom.RawText("Test")).SetProp(dom.STYLE,
				dom.CSS("").Set(dom.COLOR, "green").Inline()),
			dom.Comment(dom.RawText("this is a paragraph")),
			dom.P(
				dom.Comment(dom.RawText("raw text here")),
				dom.RawText("hello"),
				dom.Br(),
				dom.P(dom.RawText("world")).SetProp(dom.ID, "world"),
				dom.Img().SetProp(dom.SRC, "https://baidu.com/favicon.ico"),
			).SetProp(dom.CLASS, "p1"),
		),
	).SetProp("lang", "en"))

	s := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write(doc.Render())
		}),
	}
	s.ListenAndServe()
}
```
