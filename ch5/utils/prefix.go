package utils

func Prefixer(prefix string) func(string) string {
	return func(name string) string {
		return prefix + " " + name
	}
}
