package dom

type doctype struct{}

func (e *doctype) SetProp(string, string) Element { return e }

func (e *doctype) DelProp(string) Element { return e }

func (e *doctype) Render() []byte { return []byte("<!DOCTYPE html>") }

func Doctype() Element {
	return &doctype{}
}

type document struct {
	html Element
}

func (e *document) SetProp(string, string) Element { return e }

func (e *document) DelProp(string) Element { return e }

func (e *document) Render() []byte {
	b := getBuffer()
	defer putBuffer(b)
	b.Write(Doctype().Render())
	b.Write(e.html.Render())
	return b.Bytes()
}

func Document(html Element) Element {
	return &document{html: html}
}
