package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	// S'il n'y a pas d'argument, on récupère la saisie de l'utilisateur (qu'on stocke dans une variable "bytes") et on la print :
	if len(args) < 1 {
		bytes, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			return
		}

		for _, char := range bytes {
			z01.PrintRune(rune(char))
		}
		return
	}

	// On lit chaque argument (fichier) et stocke leur contenu (texte) dans une variable "content" :
	for _, arg := range args {
		content, err := ioutil.ReadFile(arg)

		// Si l'argument est invalide (ne renvoit à aucun fichier existant), on print une erreur "No such file or directory" :
		if err != nil {
			str := ""
			str = str + "ERROR: open " + arg + " : no such file or directory\n"

			for _, char := range str {
				z01.PrintRune(rune(char))
			}
			os.Exit(1)

			// Sinon, on affiche le contenu du fichier :
		} else {
			for _, char := range content {
				z01.PrintRune(rune(char))
			}
		}
	}

}
