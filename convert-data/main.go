package main

type User struct {
	Name    string
	Age     intpackage main

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
	
		parts := strings.Split(data, ",")
		age, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		//your code here!
		return User{
			Name:    strings.TrimSpace(parts[0]),
			Age:     age,
			Address: strings.TrimSpace(parts[2]),
		}
	}
	
	func main() {
		data := "Edo, 20, Jakarta"
		user := ConvertData(data)
		fmt.Printf("{name: %q, age: %d, address: %q }\n", user.Name, user.Age, user.Address)
	}
	
	Address string
}

func ConvertData(data string) User {
	//your code here!

	return User{}
}

func main() {}
