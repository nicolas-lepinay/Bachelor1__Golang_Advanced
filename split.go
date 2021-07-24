package student

func Split(s, sep string) []string {
	g := 0
	result := make([]string, Size(s, sep))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(sep); j++ {
			if i+j < len(s) && s[i+j] == sep[j] {
				if j == len(sep)-1 {
					i += len(sep)
					g++
				}
			}
		}
		result[g] += string(s[i])

	}
	return result
}
func Size(s, sep string) int {
	var count int = 1
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(sep); j++ {
			if i+j < len(s) && s[i+j] == sep[j] {
				if j == len(sep)-1 {
					count++
				}
			}
		}
	}
	return count
}
