package main

import (
	"fmt"
	"strings"
)

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	keys := strings.Split(key, ",")
	for _, n := range keys {
		delete(dataMap, n)
	}
	return dataMap
}

func main() {
	data0 := map[string]any{
		"name":    "Edo",
		"age":     "20",
		"address": "Jakarta",
	}

	data0edit := removeUnrelated(data0, "address")
	fmt.Println(data0edit)
}
