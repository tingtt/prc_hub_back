package flag_with_env

import "flag"

func Parse() {
	flag.Parse()
}

func Uint(paramName string, envName string, fallback uint, desc string) *uint {
	return flag.Uint(paramName, getUintEnv(envName, fallback), desc)
}

func String(paramName string, envName string, fallback string, desc string) *string {
	return flag.String(paramName, getEnv(envName, fallback), desc)
}

type arrayParam []string

// Implements for flag.Value
func (i *arrayParam) String() string {
	str := "["
	for idx, p := range *i {
		str += p
		if len(*i) != idx+1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

// Implements for flag.Value
func (i *arrayParam) Set(v string) error {
	*i = append(*i, v)
	return nil
}

func Array(paramName string, desc string) *arrayParam {
	array := new(arrayParam)
	flag.Var(array, paramName, desc)
	return array
}
