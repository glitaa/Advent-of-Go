package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Slope [][]byte

func (s Slope) Check(right, down int) (trees int) {
	x := 0
	for y := down; y < len(s); y += down {
		x += right
		if s[y][x%len(s[0])] == '#' {
			trees++
		}
	}
	return
}

func MakeSlope(str string) (slope Slope) {
	rows := strings.Split(strings.Replace(str, "\r\n", "\n", -1), "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		slope = append(slope, []byte(row))
	}
	return slope
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	slope := MakeSlope(string(input))

	trees := slope.Check(3, 1)

	fmt.Println("Number of trees on the slope (right 3, down 1):", trees)

	multiplication := slope.Check(1, 1) * slope.Check(3, 1) * slope.Check(5, 1) * slope.Check(7, 1) * slope.Check(1, 2)

	fmt.Println("Multiplication of the number of trees encountered on each of the slopes (right, down: (1, 1), (3, 1), (5, 1), (7, 1), (1, 2)):", multiplication)
}
