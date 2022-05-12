package main

import (
	"fmt"
)

//sin concurrencia

type Coord2 struct {
	x int
	y int
}

func Abs2(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func attackCheckSolution2(x Coord, y Coord) bool {
	return x.x == y.x || x.y == y.y || (Abs(x.y-y.y) == Abs(x.x-y.x))
}

func feasibilityCheck2(curr Coord, list []Coord) bool {
	for _, coord := range list {
		if attackCheckSolution2(curr, coord) {
			return false
		}
	}
	return true
}

func nQueensFunc2(n int, x int, ch chan []Coord2, list []Coord2) {
	if x == n {
		ch <- list
		return
	}
	for i := 0; i < n; i++ {
		curCoord := Coord2{x, i}
		if feasibilityCheck2(curCoord, list) {
			currList := append(list, curCoord)
			nQueensFunc2(n, x+1, ch, currList)
		}
	}
}

func nQueens2(n int, ch chan []Coord2) {
	var list = make([]Coord2, n)
	nQueensFunc2(n, 0, ch, list)
}

//concurrente
type Coord struct {
	x int
	y int
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func attackCheckSolution(x Coord, y Coord) bool {
	return x.x == y.x || x.y == y.y || (Abs(x.y-y.y) == Abs(x.x-y.x))
}

func feasibilityCheck(curr Coord, list []Coord) bool {
	for _, coord := range list {
		if attackCheckSolution(curr, coord) {
			return false
		}
	}
	return true
}

func nQueensFunc(n int, x int, ch chan []Coord, list []Coord) {
	if x == n {
		ch <- list
		return
	}
	for i := 0; i < n; i++ {
		curCoord := Coord{x, i}
		if feasibilityCheck(curCoord, list) {
			currList := append(list, curCoord)
			go nQueensFunc(n, x+1, ch, currList)
		}
	}
}

func nQueens(n int, ch chan []Coord) {
	var list = make([]Coord, n)
	go nQueensFunc(n, 0, ch, list)
}

func main() {
	n := 5
	ch := make(chan []Coord, n)

	// con concurrencia
	nQueens(n, ch)
	for i := 0; i < n; i++ {
		j := <-ch
		fmt.Println(j)
	}
	// sin concurrencia
	//nQueens2(n,ch)
}
