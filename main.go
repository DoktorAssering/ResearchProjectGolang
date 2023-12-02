package main

import (
	"fmt"
)

func main() {
	var array = [...]int{432, 6453, 123, 355, 500}

	var min = array[0]
	var max = array[0]

	for i := 1; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
		if max < array[i] {
			max = array[i]
		}
	}

	difference := max - min
	fmt.Print("Пока что!")
	fmt.Printf("Разница между максимальным и минимальным значениями: %d\n", difference)
}
