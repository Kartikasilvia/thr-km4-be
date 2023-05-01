package main

import (
	"fmt"
	"strings"
)

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	keys := strings.Split(key, ",")
	for _, n := range keys {
		delete(dataMap, n)
	}
	//yout code here
	return dataMap
}

func main() {
	myData := map[string]any{
		"name":    "Edo",
		"age":     "20",
		"address": "Jakarta",
	}

	m := removeUnrelated(myData, "address")
	fmt.Println(m)
	// /h := []any{"Edo", "Budi", "Joko", "Tono"}
	// i := removeLeftRight(h, "left")
	// fmt.Println(i, "left")

	// j := []any{"Edo", "Budi", "Joko", "Tono"}
	// k := removeLeftRight(j, "right")
	// fmt.Println(k, "right")
}
