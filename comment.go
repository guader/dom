package dom

type comment struct {
	contents []Element
}

func (e *comment) SetProp(string, string) Element { return e }

func (e *comment) DelProp(string) Element { return e }

func (e *comment) Render() []byte {
	b := getBuffer()
	defer putBuffer(b)
	b.WriteString("<!-- ")
	for _, c := range e.contents {
		b.Write(c.Render())
	}
	b.WriteString(" -->")
	return b.Bytes()
}

func Comment(contents ...Element) Element {
	return &comment{contents: contents}
}
