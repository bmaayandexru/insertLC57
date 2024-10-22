package main

import "fmt"

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	interval := []int{2, 5}
	fmt.Println(insert(intervals, interval))
}
func insert(intervals [][]int, newInterval []int) [][]int {
	// правая граница i дожна быть быть больше или равна левой границы ni
	// перед или на каком отрезке каким отрезком начинается новый отрезок
	i := 0
	for intervals[i][1] < newInterval[0] {
		i++
	}
	// перед каким заканчивается
	j := 0
	for intervals[j][0] < newInterval[1] {
		j++
	}
	// i == j
	// ni[1] < i[i][0] - вставить ni перед i[i] и выйти
	// иначе
	// в i[i][0] = min(i[i][0],ni[0])
	// в i[i][1] = max(i[i][1],ni[1])
	// i-й исходный для слияния следующих интервалов
	return intervals
}
