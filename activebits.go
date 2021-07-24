package student

func ActiveBits(n int) uint {

	var active uint

	for n != 0 {
		mod := n % 2
		if mod == 1 {
			active++
		}
		n = n / 2
	}
	return active
}
