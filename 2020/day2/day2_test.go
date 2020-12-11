package main

import "testing"

func TestOldPolicyValidation(t *testing.T) {

	passwordTests := []struct {
		name     string
		password Password
		valid    bool
	}{
		{name: "valid", password: Password{Value: "abcde", ContainsLetter: "a", MinLetterOccurrences: 1, MaxLetterOccurrences: 3}, valid: true},
		{name: "invalid", password: Password{Value: "cdefg", ContainsLetter: "b", MinLetterOccurrences: 1, MaxLetterOccurrences: 3}, valid: false},
		{name: "valid", password: Password{Value: "ccccccccc", ContainsLetter: "c", MinLetterOccurrences: 2, MaxLetterOccurrences: 9}, valid: true},
	}

	for _, test := range passwordTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.password.OldPolicyValidation()
			if got != test.valid {
				t.Errorf("%#v got %t want %t", test.password, got, test.valid)
			}
		})
	}

}
