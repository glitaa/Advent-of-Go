package main

import "testing"

type passwordTest struct {
	name     string
	password Password
	valid    bool
}

var passwordTests = []passwordTest{
	{password: Password{Value: "abcde", ContainsLetter: 'a', FirstPosition: 1, SecondPosition: 3}},
	{password: Password{Value: "cdefg", ContainsLetter: 'b', FirstPosition: 1, SecondPosition: 3}},
	{password: Password{Value: "ccccccccc", ContainsLetter: 'c', FirstPosition: 2, SecondPosition: 9}},
}

func TestOldPolicyValidation(t *testing.T) {

	passwordTests[0].name = "valid with the old policy"
	passwordTests[0].valid = true
	passwordTests[1].name = "invalid with the old policy"
	passwordTests[1].valid = false
	passwordTests[2].name = "valid  with the old policy"
	passwordTests[2].valid = true

	for _, test := range passwordTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.password.OldPolicyValidation()
			assertValidity(t, got, test)
		})
	}

}

func TestNewPolicyValidation(t *testing.T) {

	passwordTests[0].name = "valid with the new policy"
	passwordTests[0].valid = true
	passwordTests[1].name = "invalid with the new policy"
	passwordTests[1].valid = false
	passwordTests[2].name = "invalid  with the new policy"
	passwordTests[2].valid = false

	for _, test := range passwordTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.password.NewPolicyValidation()
			assertValidity(t, got, test)
		})
	}

}

func assertValidity(t *testing.T, got bool, test passwordTest) {
	t.Helper()
	if got != test.valid {
		t.Errorf("%#v got %t want %t", test.password, got, test.valid)
	}
}
