package main

import "fmt"

func main() {
	intervals := [][]int{{1, 5}, {17, 19}}
	//	interval := []int{2, 5}
	interval := []int{2, 4}
	fmt.Println(insert(intervals, interval))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{}
	lpi := -1 // индекс отрезка внутри которого левая точка интервала
	rpi := -1 // индекс отрезка внутри которого правая точка интервала
	lppi := 0 // индекс отрезка перед которым находится левая точка интервала
	rppi := 0 // индекс отрезка перед которым находится правая точка
	return ret
}
