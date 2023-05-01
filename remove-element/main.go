package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		arr = append(arr[:0], arr[0+1:]...)
		return arr

	}
	if position == "right" {
		if len(arr) > 0 {
			arr = arr[:len(arr)-1]
		}

	} else {
		return nil
	}

	return arr
}

func main() {
	arr0 := []any{1, 2, 3, 4, 5}
	position0 := "left"
	output0 := removeLeftRight(arr0, position0)
	fmt.Println(output0)

	arr1 := []any{1, 2, 3, 4, 5}
	position1 := "right"
	output1 := removeLeftRight(arr1, position1)
	fmt.Println(output1)

	arr2 := []any{"Edo", "Budi", "Joko", "Tono"}
	position2 := "left"
	output2 := removeLeftRight(arr2, position2)
	fmt.Println(output2)

	arr3 := []any{"Edo", "Budi", "Joko", "Tono"}
	position3 := "right"
	output3 := removeLeftRight(arr3, position3)
	fmt.Println(output3)
}
