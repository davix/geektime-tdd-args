package args

import "strconv"

type Args struct {
	args []string
	res  map[string]any
}

func New(args ...string) Args {
	return Args{
		args: args,
		res:  make(map[string]any),
	}
}

func (a Args) Parse() {
	for i, arg := range a.args {
		if result, ok := a.res[arg]; ok {
			switch v := result.(type) {
			case *bool:
				*v = true
			case *int:
				value, err := strconv.ParseInt(a.args[i+1], 10, 32)
				if err != nil {
					continue
				}
				*v = int(value)
			case *string:
				*v = a.args[i+1]
			}
		}
	}
}

func (a Args) Bool(name string) *bool {
	var result bool
	a.res["-"+name] = &result
	return &result
}

func (a Args) Int(name string) *int {
	var result int
	a.res["-"+name] = &result
	return &result
}

func (a Args) String(name string) *string {
	var result string
	a.res["-"+name] = &result
	return &result
}
