package accumulate

func Accumulate(sl []string, operator func(string) string) []string {
	nl := []string{}
	for _, v := range sl {
		nl = append(nl, operator(v))
	}
	return nl
}
