package src

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func RetrieveWordFromDB(SelectDiff string) (string, []string) {
	wordComplete := ""
	var result []string
	rand.Seed(time.Now().UTC().UnixNano())
	var file string
	//ce block de code permet de choisir le fichier en fonction de la difficulté
	if SelectDiff == "EASY" {
		file = "txt/WordsEasy.txt"
	} else if SelectDiff == "MEDIUM" {
		file = "txt/WordsMedium.txt"
	} else if SelectDiff == "HARD" {
		file = "txt/WordsHard.txt"
	}
	//ce block de code permet de lire le fichier et de choisir un mot aléatoire

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	RandomIntegerWithinRange := rand.Intn(len(lines)-0) + 0
	wordComplete = lines[RandomIntegerWithinRange]

	nLetter := len(wordComplete)/2 - 1
	var wordTable []string
	for _, val := range wordComplete {
		wordTable = append(wordTable, string(val)) //transforme le mot en tableau de string
	}
	for range wordComplete {
		result = append(result, "_") //rempli le tableau de string du resultat avec des "_"
	}
	for i := 0; i <= nLetter; i++ {
		index := rand.Intn(len(result))
		result[index] = wordTable[index]
	}
	return wordComplete, result

}

func Gamemode(letter string, wordComplete string, result []string, error int) ([]string, int) {
	//tant que le joueur n'a pas fait 10 erreur ou qu'il a pas gagner
	fmt.Println("choose a letter:")

	if IsLetter(letter) == true { //verifie si l'input est une lettre
		if ValidInput(letter, wordComplete) == false { //verifie si l'input est dans le mot
			error += 1
		} else {
			for i, value := range wordComplete { //remplace les "_" par la lettre si elle est dans le mot
				if string(value) == letter {
					result[i] = letter
				}
			}

		}
	}
	return result, error

}

func IsLetter(input string) bool {
	//ce block verifie si l'input est bien une lettre
	if input == "a" || input == "b" || input == "c" || input == "d" || input == "e" || input == "f" || input == "g" || input == "h" || input == "i" || input == "j" || input == "k" || input == "l" || input == "m" || input == "n" || input == "o" || input == "p" || input == "q" || input == "r" || input == "s" || input == "t" || input == "u" || input == "v" || input == "w" || input == "x" || input == "y" || input == "z" || input == "é" || input == "è" || input == "ê" || input == "à" || input == "ù" || input == "ç" || input == "î" || input == "ï" || input == "ô" || input == "û" || input == "â" {
		return true
	}
	return false
}

func ValidInput(input string, word string) bool {
	//ce block verifie si l'input n'est pas dans le mot
	i := 0
	for _, value := range word {
		if string(value) == input {
			i++
		}
	}
	if i > 0 {
		return true
	} else {
		return false
	}
}

func CheckWin(result []string) bool {
	//ce block verifie si le joueur a gagné
	for _, value := range result {
		if value == "_" {
			return true
		}
	}
	return false
}

func AddNewWord(word string, SelectDiff string) {
	var byt []string
	//ce bloc de code choisi le fichier dans lequel ecrire
	var fileChoose string
	var dbString string
	if SelectDiff == "EASY" {
		fileChoose = "txt/WordsEasy.txt"
	} else if SelectDiff == "MEDIUM" {
		fileChoose = "txt/WordsMedium.txt"
	} else if SelectDiff == "HARD" {
		fileChoose = "txt/WordsHard.txt"

	}
	//ce bloc de code ajoute le mot dans le fichier
	database, err := os.Open(fileChoose)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(database)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	for _, val := range lines {
		byt = append(byt, val)
	}
	input := word
	i := 0
	for _, value := range input {
		if IsLetter(string(value)) == true {
			i++
		}

	}
	if i == len(input) {
		byt = append(byt, input)

		database.Close()
		for _, val := range byt {
			dbString += val + "\n"
		}
		bytes := []byte(dbString)
		ioutil.WriteFile(fileChoose, bytes, 0664)
	}
}
