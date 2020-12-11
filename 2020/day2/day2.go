package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	Value          string
	ContainsLetter string
	FirstPosition  int
	SecondPosition int
}

func (p Password) OldPolicyValidation() bool {
	LetterOccurrences := strings.Count(p.Value, p.ContainsLetter)
	return LetterOccurrences >= p.FirstPosition && LetterOccurrences <= p.SecondPosition
}

func main() {
	passwords := readPasswordsFromFile("input.txt")

	validPasswordsCount := 0
	for _, password := range passwords {
		if password.OldPolicyValidation() {
			validPasswordsCount++
		}
	}

	fmt.Println("Number of passwords validated with the old policy:", validPasswordsCount)
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
		password.FirstPosition, _ = strconv.Atoi(re.Split(scanner.Text(), -1)[0])
		password.SecondPosition, _ = strconv.Atoi(re.Split(scanner.Text(), -1)[1])
		password.ContainsLetter = re.Split(scanner.Text(), -1)[2]
		password.Value = re.Split(scanner.Text(), -1)[3]
		passwords = append(passwords, password)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}
