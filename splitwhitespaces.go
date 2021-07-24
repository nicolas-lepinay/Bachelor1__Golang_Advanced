package student

func SplitWhiteSpaces(s string) []string {

	word := ""
	tableau := make([]string, Size(s))
	var i int

	// On parcourt chaque lettre de s.
	// Tant qu'on ne rencontre pas d'espace/tabulation/saut de ligne, cela signifie qu'on est en train de parcourir un mot (word).
	// Dans ce cas, on ajoute les lettres à la variable word pour former le mot.
	for _, lettre := range s {
		if lettre != ' ' && lettre != '\t' && lettre != '\n' {
			word = word + string(lettre)

			// Sinon (si on rencontre un espace ou tout autre séparateur), cela signifie qu'on a fini de parcourir le mot.
			// Donc on ajoute le mot ainsi formé au tableau, et on réinitialise la variable word pour pouvoir créer un nouveau mot ensuite.
		} else if word != "" {
			tableau[i] = word
			i++
			word = ""
		}
	}
	tableau[i] = word

	return tableau

}

func Size(s string) int {

	var countWord int

	// Si jamais une string commence par un espace (ou une suite d'espaces, par exemple : "    Hello how are you?"), on les saute en faisant index++ :
	index := 0
	for s[index] == ' ' || s[index] == '\t' || s[index] == '\n' {
		index++
	}

	// On commence la boucle à partir de index (au cas où les indices précédents correspondraient à des espaces) :
	for i := index; i < len(s); i++ {
		// Si s[i] est un espace, et si le caractère suivant n'est pas aussi un espace (donc si c'est une lettre), on incrémente le compteur :
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if s[i] != s[i+1] {
				countWord++
			}
		}
	}
	// On incrémente le compteur une dernière fois pour prendre en compte le dernier mot :
	countWord++
	return countWord
}
