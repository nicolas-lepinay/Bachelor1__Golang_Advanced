package student

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
