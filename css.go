package dom

import (
	"fmt"
)

type CSSElement interface {
	Set(string, string) CSSElement
	Del(string) CSSElement
	Inline() string
	Element
}

type css struct {
	selector string
	props    props
}

func (e *css) Set(k, v string) CSSElement {
	e.props.set(k, v)
	return e
}

func (e *css) Del(k string) CSSElement {
	e.props.del(k)
	return e
}

func (e *css) Inline() string {
	b := getBuffer()
	defer putBuffer(b)
	for k, v := range e.props {
		b.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	return b.String()
}

func (e *css) SetProp(k, v string) Element {
	return e.Set(k, v)
}

func (e *css) DelProp(k string) Element {
	return e.Del(k)
}

func (e *css) Render() []byte {
	b := getBuffer()
	defer putBuffer(b)
	b.WriteString(fmt.Sprintf("%s{%s}", e.selector, e.Inline()))
	return b.Bytes()
}

func CSS(selector string) CSSElement {
	return &css{
		selector: selector,
		props:    p(),
	}
}
