package dom

type rawText string

func (e rawText) SetProp(string, string) Element { return e }

func (e rawText) DelProp(string) Element { return e }

func (e rawText) Render() []byte { return []byte(e) }

func RawText(s string) Element { return rawText(s) }
