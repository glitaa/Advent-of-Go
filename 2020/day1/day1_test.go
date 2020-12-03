package main

import (
	"reflect"
	"sort"
	"testing"
)

var input []int = []int{1721, 979, 366, 299, 675, 1456}
var unsuitableInput = []int{1, 2, 3}
var got Year2020

func TestFindWithTwoNumbers(t *testing.T) {

	t.Run("with suitable numbers", func(t *testing.T) {
		got.FindWithTwoNumbers(input)
		want := Year2020{514579, []int{1721, 299}}

		sortNumbers(t, &got, &want)
		assertYears(t, got, want)
	})

	t.Run("no suitable numbers", func(t *testing.T) {
		err := got.FindWithTwoNumbers(unsuitableInput)

		assertError(t, err, Err2020NotFound)
	})

}

func TestFind2020WithThreeNumbers(t *testing.T) {

	t.Run("with suitable numbers", func(t *testing.T) {
		got.FindWithThreeNumbers(input)
		want := Year2020{241861950, []int{979, 366, 675}}

		sortNumbers(t, &got, &want)
		assertYears(t, got, want)
	})

	t.Run("no suitable numbers", func(t *testing.T) {
		err := got.FindWithThreeNumbers(unsuitableInput)

		assertError(t, err, Err2020NotFound)
	})

}

func sortNumbers(t *testing.T, got, want *Year2020) {
	t.Helper()
	sort.Ints(got.Numbers)
	sort.Ints(want.Numbers)
}

func assertYears(t *testing.T, got, want Year2020) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v but wanted %#v", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}
