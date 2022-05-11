package dom

type props map[string]string

func (p props) set(k, v string) {
	p[k] = v
}

func (p props) del(k string) {
	delete(p, k)
}

func p() props {
	return make(props)
}
