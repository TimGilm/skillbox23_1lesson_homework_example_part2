// вторая часть дз, поиск индекса первого вхождения заданного числа в сортированном массиве
package main

import "fmt"

const size = 6

func main() {
	var a [size]int
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}
	n := 0
	fmt.Scan(&n)

	index := find(a, n)
	fmt.Println("Индекс", index)
}

func find(a [size]int, n int) (index int) {
	index = -1
	min := 0
	max := size - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == n {
			index = middle
			break
		} else if a[middle] > n {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	for index > 0 && a[index-1] == n {
		index--
	}
	return
}
