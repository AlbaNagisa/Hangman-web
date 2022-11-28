package HangmanModule

import (
	"bufio"
	"log"
	"os"
)

/*
lit n'importe quel fichier
*/
func FileReader(w *[]string, args string) {
	readFile, err := os.Open("./" + args)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	*w = fileLines
}

func encode(input, key string) string {
	var output string
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range input {
		output += string(key[index(letter, alphabet)])
	}

	return output
}

func decode(input, key string) string {
	var output string
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range input {
		output += string(alphabet[index(letter, key)])
	}

	return output
}

func index(letter rune, word string) int {
	for index, value := range word {
		if value == letter {
			return index
		}
	}
	return -1
}
