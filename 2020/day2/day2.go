package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ErrCannotOpenFile = errors.New("could not open the input file")
var ErrCannotConvert = errors.New("could not convert given string to int")

const REGEX_SCANNER_PATTERN = "-|\\s|:\\s"

type Password struct {
	Value          string
	ContainsLetter byte
	FirstPosition  int
	SecondPosition int
}

func (p Password) OldPolicyValidation() bool {
	LetterOccurrences := strings.Count(p.Value, string(p.ContainsLetter))
	return LetterOccurrences >= p.FirstPosition && LetterOccurrences <= p.SecondPosition
}

func (p Password) NewPolicyValidation() bool {
	return (p.Value[p.FirstPosition-1] == p.ContainsLetter) != (p.Value[p.SecondPosition-1] == p.ContainsLetter)
}

func main() {
	passwords, err := readPasswordsFromFile("input.txt")
	if err != nil {
		log.Println(err)
	}

	validPasswordsCount := 0
	for _, password := range passwords {
		if password.OldPolicyValidation() {
			validPasswordsCount++
		}
	}
	fmt.Println("Number of passwords validated with the old policy:", validPasswordsCount)

	validPasswordsCount = 0
	for _, password := range passwords {
		if password.NewPolicyValidation() {
			validPasswordsCount++
		}
	}
	fmt.Println("Number of passwords validated with the new policy:", validPasswordsCount)
}

func readPasswordsFromFile(path string) (passwords []Password, readErr error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, ErrCannotOpenFile
	}
	defer file.Close()

	re := regexp.MustCompile(REGEX_SCANNER_PATTERN)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var password Password
		password.FirstPosition, err = strconv.Atoi(re.Split(scanner.Text(), -1)[0])
		if err != nil {
			fmt.Println(err)
			readErr = ErrCannotConvert
			continue
		}
		password.SecondPosition, err = strconv.Atoi(re.Split(scanner.Text(), -1)[1])
		if err != nil {
			fmt.Println(err)
			readErr = ErrCannotConvert
			continue
		}
		password.ContainsLetter = []byte(re.Split(scanner.Text(), -1)[2])[0]
		password.Value = re.Split(scanner.Text(), -1)[3]
		passwords = append(passwords, password)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return
}
