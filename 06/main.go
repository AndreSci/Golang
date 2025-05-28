package main

import "fmt"

func main() {
	fmt.Println("Hi")

	users :=map[string]int{
		"Andrew": 32,
		"Dima": 22,
		"Kostya": 48,
	}


	fmt.Println(users)


	forMap(users)
}


func forMap(users map[string]int) bool {

	age, exists := users["Andrew"]
	fmt.Println(age, exists)

	
	age, exists = users["Andrew2"]
	fmt.Println(age, exists)

	for name, value := range users {
		fmt.Print(name, ": ", value, "\n")
	}

	delete(users, "Andrew")
	users["Dddd"] = 12

	fmt.Println(users)
	fmt.Println(len(users))

	return true
}
