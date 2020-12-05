package main

import "strings"

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
