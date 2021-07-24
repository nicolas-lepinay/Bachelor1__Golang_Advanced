package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {

	str := ""

	// On calcule la longueur de la base :
	lengthBase := 0
	for range base {
		lengthBase++
	}

	// Si la base a moins de 2 caractères, elle est invalide :
	if lengthBase < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Si chaque caractère de la base n'est pas unique, elle est invalide :
	for i := 0; i < lengthBase; i++ {
		for j := i + 1; j < lengthBase; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	// Si nbr est négatif, on affiche '-' et on le transforme en entier positif.
	// Si nbr = 0, on affiche le 1er caractère de la base.
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	} else if nbr == 0 {
		z01.PrintRune(rune(base[0]))
	}

	for nbr > 0 {
		str = str + string(base[nbr%lengthBase])
		nbr = nbr / lengthBase
	}

	// On affiche la string str à l'envers :
	for i := len(str) - 1; i > -1; i-- {
		z01.PrintRune(rune(str[i]))
	}

}
