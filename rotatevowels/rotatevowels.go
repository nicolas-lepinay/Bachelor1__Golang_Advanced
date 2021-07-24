package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args[1:]

	var args string
	var vowels []string
	var result string

	// S'il n'y a aucun argument, on imprime un saut à la ligne :
	if len(arguments) < 1 {
		z01.PrintRune('\n')
	}

	// On remplit une string ARGS avec chaque argument (+ un espace après chaque mot) :
	for i := 0; i < len(arguments); i++ {
		args = args + arguments[i] + " "
	}

	// On crée un tableau VOWELS contenant les voyelles (bytes castées en string) contenues dans ARGS :
	for i := 0; i < len(args); i++ {
		if args[i] == 'a' || args[i] == 'e' || args[i] == 'i' || args[i] == 'o' || args[i] == 'u' || args[i] == 'A' || args[i] == 'E' || args[i] == 'I' || args[i] == 'O' || args[i] == 'U' {
			vowels = append(vowels, string(args[i]))
		}
	}

	// On inverse la tableau des voyelles :
	for i := 0; i < len(vowels); i++ {
		for j := i + 1; j < len(vowels); j++ {
			vowels[i], vowels[j] = vowels[j], vowels[i]
		}
	}

	// On parcourt la string ARGS. Si la lettre est une voyelle, on ajoute sa voyelle "miroir" (inversée) à la string vide RESULT.
	// Sinon, on ajoute simplement la lettre (consonne) args[i].
	j := 0
	for i := 0; i < len(args); i++ {
		if args[i] == 'a' || args[i] == 'e' || args[i] == 'i' || args[i] == 'o' || args[i] == 'u' || args[i] == 'A' || args[i] == 'E' || args[i] == 'I' || args[i] == 'O' || args[i] == 'U' {
			result = result + vowels[j]
			j++

		} else {
			result = result + string(args[i])
		}
	}

	// On affiche chaque lettre de result :
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
	}

}
