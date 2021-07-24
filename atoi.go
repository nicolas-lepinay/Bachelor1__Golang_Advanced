package student

func Atoi(s string) int {

	if len(s) == 0 {
		return 0
	}
	result := 0

	isNegative := s[0] == '-'

	for i := 0; i < len(s); i++ {

		isNotNumber := s[i] < '0' || s[i] > '9'

		if i != 0 && isNotNumber == true {
			return 0
		}

		if !isNotNumber == true {
			result = result*10 + int(s[i]-48)
		}
	}

	if isNegative == true {
		result = result * -1
	}

	return result

}
