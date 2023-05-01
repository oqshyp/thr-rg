package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	editedData := strings.Split(data, ",")
	age, _ := strconv.Atoi(strings.TrimSpace(editedData[1]))
	return User{
		Name:    strings.TrimSpace(editedData[0]),
		Age:     age,
		Address: strings.TrimSpace(editedData[2]),
	}
	//your code here!
}

func main() {
	data1 := "Edo,20,Jakarta"
	user1 := ConvertData(data1)
	fmt.Printf("Name: %s, age: %d, Address: %s", user1.Name, user1.Age, user1.Address)

	fmt.Println()

	data2 := "Budi,30,Bandung"
	user2 := ConvertData(data2)
	fmt.Printf("Name: %s, age: %d, Address: %s", user2.Name, user2.Age, user2.Address)
}
