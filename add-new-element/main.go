package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!
	if position == "up" {
		newArr := make([]int, len(data)+1)
		newArr[0] = newData
		copy(newArr[1:], data)
		return newArr

	}

	if position == "down" {
		return append(data, newData)
		// last := len(data) - 1
		// data := append(data, data[newData])
		// data = append(data, newData)

	} else {
		return nil
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6
	position := "up"
	output := AddElement(data, newData, position)
	fmt.Println(output)

	data1 := []int{1, 2, 3, 4, 5}
	newData1 := 6
	position1 := "down"
	output1 := AddElement(data1, newData1, position1)
	fmt.Println(output1)
}
