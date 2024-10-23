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
	lpi := -1  // индекс отрезка внутри которого левая точка интервала. -1 точка не принадлежит никакому отрезку
	rpi := -1  // индекс отрезка внутри которого правая точка интервала
	lppi := -1 // индекс отрезка перед которым находится левая точка интервала
	rppi := -1 // индекс отрезка перед которым находится правая точка
	// цикл по отрезкам
	// проверяются все 4ре условия
	// окончание цикла по определению rppi или концу списка
	for i, v := range intervals {
		// принадлежность левой точки i-му отрезку
		if newInterval[0] >= v[0] && newInterval[0] <= v[1] {
			lpi = i
		}
		// левая точка перед i-м отрезком
		if newInterval[0] < v[0] && newInterval[0] < v[1] {
			lppi = i
		}
		// принадлежность правой точке i-му отрезку
		if newInterval[1] >= v[0] && newInterval[1] <= v[1] {
			rpi = i
		}
		// левая точка перед i-м отрезком
		if newInterval[1] < v[0] && newInterval[1] < v[1] {
			rppi = i
			break
		}
	}
	if lpi == -1 {
		// левая точка не на отрезке
		if lppi == -1 {
			// левая точка после всех отрезков
			// добавляем отрезок в конец списка
			return append(intervals, newInterval)
		}
		if lppi == 0 {
			// левая точка перед всеми отрезками
			if rpi == -1 {
				// правая точка не принадлежит никаком утрезку
				if rppi == 0 {
					// интервал перед всеми отрезками
					ret = append(ret, newInterval)
					ret = append(ret, intervals...)
					return ret
				}
				if rppi == -1 {
					// интервал включает все отрезки и остается только он
					return append(ret, newInterval)
				}
				// интервал захватывает хотя бы один отрезок
				// (rppi-1) последний отрезок который захватывает интервал
				ret = append(ret, newInterval)
				ret = append(ret, intervals[rppi:]...)
				return ret
			}
		}
		// внутри lppi > 0 < len(intervals)

	} else {
		// левая точка на отрезке lpi
		// левая граника отрезка остается
		intervals[lpi][1] = max(newInterval[1], intervals[lpi][1])
		// отрезок  поглащает отрезки справа пока правая граница больше левых точек
		// правая граница изменяется на последнем отрезке если newInterval[1] < intervals[][1]
	}
	return ret
}
