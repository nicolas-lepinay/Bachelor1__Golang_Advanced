package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	var result int64

	// S'il n'y a pas 3 arguments, renvoyer 0 :
	if len(args) != 3 {
		result = 0
		printInt(result)
		return
	}

	// Si l'un des deux arguments n'est pas un nombre, renvoyer 0 :
	if isValid(args[0]) == false || isValid(args[2]) == false {
		result = 0
		printInt(result)
		return
	}

	// Si l'addition/soustraction des deux arguments n'est pas valide, renvoyer 0 :
	if areValidInt64(atoi(args[0]), atoi(args[2])) == false {
		result = 0
		printInt(result)
		return
	}

	// Si l'un des deux arguments n'est pas un chiffre compris entre -9223372036854775808 et 9223372036854775807, renvoyer 0 :
	if isValidInt64(args[0]) == false || isValidInt64(args[2]) == false {
		result = 0
		printInt(result)
		return
	}

	// Si l'opérateur n'est pas un opérateur valide, renvoyer 0 :
	if !(args[1] == "+" || args[1] == "-" || args[1] == "*" || args[1] == "/" || args[1] == "%") {
		result = 0
		printInt(result)
		return
	}

	if args[1] == "+" {
		result = add(atoi(args[0]), atoi(args[2]))
		printInt(result)
		return
	}

	if args[1] == "-" {
		result = subtract(atoi(args[0]), atoi(args[2]))
		printInt(result)
		return
	}

	if args[1] == "*" {
		result = multiply(atoi(args[0]), atoi(args[2]))
		printInt(result)
		return
	}

	if args[1] == "/" {
		result = divide(atoi(args[0]), atoi(args[2]))
		printInt(result)
		return
	}

	if args[1] == "%" {
		result = modulo(atoi(args[0]), atoi(args[2]))
		printInt(result)
		return
	}

}

func add(a, b int64) int64 {
	return a + b
}

func subtract(a, b int64) int64 {
	return a - b
}

func multiply(a, b int64) int64 {
	return a * b
}

func divide(a, b int64) int64 {

	errorZero := "No division by "

	if b == 0 {
		for _, letters := range errorZero {
			z01.PrintRune(letters)
		}
		return 0
	}
	return a / b
}

func modulo(a, b int64) int64 {

	errorModulo := "No modulo by "

	if b == 0 {
		for _, letters := range errorModulo {
			z01.PrintRune(letters)
		}
		return 0
	}
	return a % b
}

// Fonction pour convertir une string en int :
func atoi(s string) int64 {
	sCopy := s
	if sCopy[0] == '-' {
		sCopy = s[1:]
	}

	res := 0
	for _, myrune := range sCopy {
		res = res*10 + (int(myrune) - 48)
	}

	if s[0] == '-' {
		res = res * -1
	}

	var result int64
	result = int64(res)

	return result
}

// Fonction qui affiche un entier avec PrintRune :
func printInt(n int64) {
	tableau := []int64{}

	if n == 0 {
		z01.PrintRune('0')
	}

	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}

	for n > 0 {
		tableau = append(tableau, n%10)
		n = n / 10
	}

	for i := 0; i < len(tableau); i++ {
		for j := i + 1; j < len(tableau); j++ {
			tableau[i], tableau[j] = tableau[j], tableau[i]
		}
	}

	for _, element := range tableau {
		z01.PrintRune(rune(element + 48))
	}
	z01.PrintRune('\n')
}

// Vérifie qu'une string ne comporte que des chiffres (pas de lettres) :
func isValid(str string) bool {

	isValid := false

	for _, character := range str {
		if character >= 48 && character <= 57 {
			isValid = true
		}
	}
	return isValid
}

// Vérifie que l'addition/soustraction de 2 nombres est bien valide :
func areValidInt64(arg1, arg2 int64) bool {

	areValidInt64 := true

	// Si 2 chiffres sont positifs, que l'opérateur est un +, mais que l'addition renvoit un nombre négatif, il y a une erreur :
	if arg1 > 0 && arg2 > 0 && os.Args[1] == "+" && add(arg1, arg2) < 0 {
		areValidInt64 = false
	}

	// Si 2 chiffres sont négatifs, que l'opérateur est un -, mais que la soustraction renvoit un nombre positif, il y a une erreur :
	if arg1 < 0 && arg2 < 0 && os.Args[1] == "-" && subtract(arg1, arg2) > 0 {
		areValidInt64 = false
	}
	return areValidInt64
}

// Vérifie que l'entier est bien un int64 compris entre -9223372036854775808 et 9223372036854775807 :
func isValidInt64(arg string) bool {

	isValid := true

	// Si un chiffre n'est pas négatif (le 1er caractère de l'argument n'est PAS un '-') mais que atoi renvoit quand même un chiffre négatif, il y a une erreur d'overflow :
	if arg[0] != '-' && atoi(arg) < 0 {
		isValid = false
	}

	// Si un chiffre est négatif (le 1er caractère de l'argument est un '-') mais que atoi renvoit quand même un chiffre positif, il y a une erreur d'overflow :
	if arg[0] == '-' && atoi(arg) > 0 {
		isValid = false
	}
	return isValid

}
