package main

import "strings"

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
	rows := strings.Split(str, "\n")
	for _, row := range rows {
		slope = append(slope, []byte(row))
	}
	return slope
}

func main() {
}
