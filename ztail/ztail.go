package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

// La fonction affiche les X derniers bytes d'un fichier texte lorsque qu'on tape "go run ztail.go -cX nom_du_fichier.txt" :
func main() {

	// S'il n'y a pas exactement 2 arguments, on affiche un message d'erreur :
	args := os.Args[1:]
	error := "ERROR : tail requires a valid parameter."

	if len(args) != 2 {
		for _, char := range error {
			z01.PrintRune(char)
		}
		return
	}

	// On vérifie que le 1er argument commence bien par "tail-c" :
	arg1 := os.Args[1]
	var str []byte

	for i := 0; i <= 1; i++ {
		str = append(str, arg1[i])
	}

	if string(str) != "-c" {
		for _, char := range error {
			z01.PrintRune(char)
		}
		return
	}

	// On vérifie que le 1er argument se termine bien par un nombre :
	var number []byte

	for i := 2; i <= len(arg1)-1; i++ {
		number = append(number, arg1[i])
	}

	for _, byte := range number {
		if string(byte) < "0" || string(byte) > "9" {
			for _, char := range error {
				z01.PrintRune(char)
			}
			return

		}
	}

	// On lit le texte contenu dans le 2nd argument et on le stocke dans une variable "content" :
	arg2 := os.Args[2]

	content, err := ioutil.ReadFile(arg2)

	// Si l'argument est invalide (ne renvoit à aucun fichier existant), on print une erreur "No such file or directory" :
	if err != nil {
		str := ""
		str = str + "ERROR: open " + arg2 + " : no such file or directory\n"

		for _, char := range str {
			z01.PrintRune(rune(char))
		}
		os.Exit(1)

		// Sinon, on affiche le contenu du fichier :
	} else {

		// On convertit le chiffre (à la fin de arg2) en int :
		var n int

		for _, mybyte := range number {
			n = n*10 + (int(mybyte - 48))
		}

		// On affiche les n derniers bytes (les lettres) du fichier texte, à condition que :
		// - (1) le fichier texte ne soit pas vide,
		// - (2) n ne dépasse pas la longueur du texte,
		// - (3) n ne soit pas égal à 0.

		// Condition (1) (le fichier texte ne doit pas être vide) :
		if len(content) == 0 {
			msg := ""
			msg = msg + "ERROR: The text file is empty."

			for _, char := range msg {
				z01.PrintRune(rune(char))
			}
			return
		}

		// Condition (2) (n ne doit pas être supérieur à la longueur du texte) :
		if n > len(content) {
			msg := ""
			msg = msg + "ERROR: <" + itoa(n) + "> is not a valid value. The number of bytes cannot exceed <" + itoa(len(content)) + ">."

			for _, char := range msg {
				z01.PrintRune(rune(char))
			}
			return
		}

		// Condition (3) (n ne doit pas être égal à 0) :
		if n == 0 {
			msg := ""
			msg = msg + "ERROR: <0> or <nil> are not a valid values. The number of bytes must be superior or equal to 1."

			for _, char := range msg {
				z01.PrintRune(rune(char))
			}
			return

		} else if n <= len(content) {
			for i := len(content) - n; i < len(content); i++ {
				z01.PrintRune(rune(content[i]))
			}
		}
	}
}

func itoa(n int) string {
	result := ""

	if n == 0 {
		result = "0"

	} else {
		for n > 0 {
			s := n % 10
			result = result + string(s+48)
			n = n / 10
		}
	}
	// On inverse le résultat obtenu :
	res := ""
	for i := len(result) - 1; i > -1; i-- {
		res = res + string(result[i])
	}
	return res
}
