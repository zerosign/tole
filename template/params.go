package template

type Options map[string]interface{}

type Params struct {
	inner   []interface{}
	options Options
}

func EmptyParams() Params {
	return Params{make([]interface{}, 0), make(Options)}
}

func (p Params) ParamSize() int {
	return len(p.inner)
}

func (p Params) OptionSize() int {
	return len(p.options)
}
