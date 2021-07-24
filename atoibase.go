package student

func AtoiBase(s string, base string) int {

	result := 0
	mult := RecursivePower(len(base), len(s)-1)

	// Si la base a moins de 2 caractères, on renvoie 0 :
	if len(base) < 2 {
		return 0
	}

	// Si 2 caractères de la base sont identiques, ou si la base contient un - ou un +, on renvoie 0 :
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] || base[i] == '+' || base[i] == '-' {
				return 0
			}
		}
	}
	for i := 0; i < len(s); {
		for j := 0; j < len(base); j++ {
			if i == len(s) {
				j = len(base)
			} else if s[i] == base[j] {
				i++
				result = result + (j * mult)
				mult = mult / len(base)
				j = len(base)
			}
		}
	}
	return result

}

func RecursivePower(nb int, power int) int {
	result := nb

	if power == 0 {
		result = 1
	}
	if power > 0 {
		result = result * RecursivePower(nb, power-1)
	}
	return result
}
