package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

var Err2020NotFound = errors.New("couldn't find numbers that summed give 2020")
var ErrNotANumber = errors.New("didn't get a number but wanted one")

type Year2020 struct {
	Multiplication int
	Numbers        []int
}

func (y *Year2020) FindWithTwoNumbers(numbers []int) error {
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			if i == j {
				continue
			}
			if num1+num2 == 2020 {
				*y = Year2020{num1 * num2, []int{num1, num2}}
				return nil
			}
		}
	}
	return Err2020NotFound
}

func (y *Year2020) FindWithThreeNumbers(numbers []int) error {
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			for k, num3 := range numbers {
				if i == j {
					break
				}
				if i == k || j == k {
					continue
				}
				if num1+num2+num3 == 2020 {
					*y = Year2020{num1 * num2 * num3, []int{num1, num2, num3}}
					return nil
				}
			}
		}
	}
	return Err2020NotFound
}

func main() {
	myYear2020 := Year2020{}

	fmt.Println("Enter numbers to look through to find 2020 with two of them:")
	numbers := readNumbers()
	findErr := myYear2020.FindWithTwoNumbers(numbers)
	if findErr == nil {
		fmt.Printf("Multiplication of two numbers that summed give 2020: %d \nThe two numbers are: %d, %d\n", myYear2020.Multiplication, myYear2020.Numbers[0], myYear2020.Numbers[1])
	} else {
		log.Fatalf("%s, given numbers: %v", findErr, numbers)
	}

	fmt.Println("Enter numbers to look through to find 2020 with three of them:")
	numbers = readNumbers()
	findErr2 := myYear2020.FindWithThreeNumbers(numbers)
	if findErr2 == nil {
		fmt.Printf("Multiplication of three numbers that summed give 2020: %d \nThe three numbers are: %d, %d, %d\n", myYear2020.Multiplication, myYear2020.Numbers[0], myYear2020.Numbers[1], myYear2020.Numbers[2])
	} else {
		log.Printf("%s, given numbers: %v", findErr2, numbers)
	}
}

func readNumbers() (numbers []int) {
	var number int
	for {
		_, err := fmt.Scan(&number)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("%s, error: %q", ErrNotANumber, err)
			}
			break
		}
		numbers = append(numbers, number)
	}
	return
}
