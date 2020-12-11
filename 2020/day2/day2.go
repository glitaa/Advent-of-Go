package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	Value                string
	ContainsLetter       string
	MinLetterOccurrences int
	MaxLetterOccurrences int
}

func (p Password) IsValid() bool {
	LetterOccurrences := strings.Count(p.Value, p.ContainsLetter)
	if LetterOccurrences >= p.MinLetterOccurrences && LetterOccurrences <= p.MaxLetterOccurrences {
		return true
	}
	return false
}

func main() {
}

func readPasswordsFromFile(path string) (passwords []Password) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("-|\\s|:\\s")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var password Password
		password.MinLetterOccurrences, _ = strconv.Atoi(re.Split(scanner.Text(), -1)[0])
		password.MaxLetterOccurrences, _ = strconv.Atoi(re.Split(scanner.Text(), -1)[1])
		password.ContainsLetter = re.Split(scanner.Text(), -1)[2]
		password.Value = re.Split(scanner.Text(), -1)[3]
		passwords = append(passwords, password)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}
