package serviceR

func concantenateProducts(sl []string, slp []string) []string {
	for _, v := range sl {
		slp = append(slp, v)
	}
	return slp
}
