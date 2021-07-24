package student

func Compare(a, b string) int {

	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	for i := 0; i < len(a)-1; i++ {
		// Si on passe a[i] et a[i+1] dans la fonction f (Compare) et que le résultat est 1, cela signifie que a[i] est supérieur à a[i+1].
		// Donc on les inverse.
		if f(a[i], a[i+1]) == 1 {
			a[i], a[i+1] = a[i+1], a[i]
			// On réinitialise i à 0 (i = -1 car i++) pour que le programme recommence le tri depuis le début et que tous les éléments soient bien triés.
			// Si deux éléments sont déjà triés, il passe aux suivants et les trie, puis revient au début, jusqu'à avoir parcouru tout le tableau.
			i = -1
		}

	}
}
