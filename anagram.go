package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"unicode"
)

func main() {
	//read file
	dat, err := readFile("/usr/share/dict/words")
	if err != nil {
		fmt.Println(err)
	}

	//split file into slice of words
	s := strings.Split(string(dat), "\n")

	//sort letters of each word in dic alphabetically
	sortedStringMap := secretAlgo(s)

	//get user input
	userInput := getUserInput()

	//verify user input
	isAlpha := verifyUserInput(userInput)

	//compare userInput to sorted dictionary words
	if isAlpha {
		sortedUserInput := SortString(userInput)
		for k, v := range sortedStringMap {
			if sortedUserInput == v && userInput != k {
				fmt.Println("Anagram: ", k)
			}
		}
	} else {
		fmt.Println("input must be alpha characters only")
	}
}

func readFile(file string) ([]byte, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	} else {
		return dat, nil
	}
}

func getUserInput() string {
	var userInput string
	fmt.Print("Enter text: ")
	fmt.Scanln(&userInput)
	return userInput
}

func verifyUserInput(userInput string) bool {
	for _, r := range userInput {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func secretAlgo(s []string) map[string]string {
	var m map[string]string
	m = make(map[string]string)
	for _, v := range s {
		m[v] = SortString(v)
	}
	return m
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
