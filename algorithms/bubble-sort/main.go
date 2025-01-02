package main

import "fmt"

func main() {
	oldList := []int{7, 3, 10, 2, 6, 8, 4, 1, 0, 5}
	fmt.Printf("Old list: %d\n", oldList)

	newList := bubbleSort(oldList)
	fmt.Printf("New list: %d\n", newList)
}

func bubbleSort(list []int) []int {
	qtd := len(list)

	for i := 0; i < qtd; i++ {
		for j := 0; j < qtd-1-i; j++ {
			if list[j] > list[j+1] {
				aux := list[j]
				list[j] = list[j+1]
				list[j+1] = aux
			}
		}
	}

	return list
}
