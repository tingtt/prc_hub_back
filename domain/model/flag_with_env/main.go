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
