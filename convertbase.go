package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, base string) string {
	if nbr == 0 {
		return "0"
	}
	str := ""
	res := ""
	for nbr > 0 {
		str += string(base[nbr%len(base)])
		nbr /= len(base)
	}
	for i := len(str) - 1; i > -1; i-- {
		res += string(str[i])
	}
	return res
}
