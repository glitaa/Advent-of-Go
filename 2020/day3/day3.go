package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Slope [][]byte

func (s Slope) Check() (trees int) {
	x := 0
	for y := 1; y < len(s); y++ {
		x += 3
		if s[y][x%(len(s[0])-1)] == '#' {
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

	trees := slope.Check()

	fmt.Println("Number of trees on the slope:", trees)
}
