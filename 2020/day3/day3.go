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

	fmt.Println("Number of trees on the slope:", trees)
}
