package dom

import (
	"fmt"
	"strings"
)

type Element interface {
	SetProp(string, string) Element
	DelProp(string) Element
	Render() []byte
}

type element struct {
	tag      string
	open     bool
	props    props
	contents []Element
}

func (e *element) SetProp(k, v string) Element {
	if e.props != nil {
		e.props.set(k, v)
	}
	return e
}

func (e *element) DelProp(k string) Element {
	if e.props != nil {
		e.props.del(k)
	}
	return e
}

func (e *element) Render() []byte {
	b := getBuffer()
	defer putBuffer(b)
	tags := []string{e.tag}
	for k, v := range e.props {
		tags = append(tags, fmt.Sprintf(`%s="%s"`, k, v))
	}
	b.WriteString(fmt.Sprintf(`<%s>`, strings.Join(tags, " ")))
	if e.open {
		return b.Bytes()
	}
	for _, c := range e.contents {
		b.Write(c.Render())
	}
	b.WriteString(fmt.Sprintf(`</%s>`, e.tag))
	return b.Bytes()
}

func e(tag string, open bool, props props, contents ...Element) *element {
	return &element{
		tag:      tag,
		open:     open,
		props:    props,
		contents: contents,
	}
}

func HTML(contents ...Element) Element {
	return e("html", false, p(), contents...)
}

func Body(contents ...Element) Element {
	return e("body", false, p(), contents...)
}

func Br() Element {
	return e("br", true, nil)
}

func Div(contents ...Element) Element {
	return e("div", false, p(), contents...)
}

func H1(contents ...Element) Element {
	return e("h1", false, p(), contents...)
}

func Head(contents ...Element) Element {
	return e("head", false, p(), contents...)
}

func Img() Element {
	return e("img", true, p())
}

func Meta() Element {
	return e("meta", true, p())
}

func P(contents ...Element) Element {
	return e("p", false, p(), contents...)
}

func Style(contents ...Element) Element {
	return e("style", false, nil, contents...)
}
