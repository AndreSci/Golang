package main

import "fmt"

func main() {
	fmt.Println("hi")

	user := User{"Andrew", 12, "male", 65, 175}

	new_user := NewUser("Dima", "male", 33, 88, 180)

	user.printUserInfo()
	new_user.printUserInfoChangeName("TestChange")
	new_user.printUserInfo()

	// КАСТОМНЫЕ ТИПЫ ДАННЫХ В КОТОРЫЕ МОЖНО ДОБАВЛЯТЬ МЕТОДЫ
	var age_new Age
	age_new = 18
	fmt.Println(age_new.isAdult())

}

type User struct {
		name string
		age int
		sex string
		weight int
		height int
}
func (user User) printUserInfo() {		// рессивер для структур

	fmt.Println(user.age, user.name, user.height)
}
func (user *User) printUserInfoChangeName(name string) {		// рессивер для структур
	user.name = name
	fmt.Println(user.age, user.name, user.height)
}

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) getName() string {
	return user.name
}

func NewUser(name, sex string, age, weight, height int) User {

	return User{
		name: name,
		sex: sex,
		age: age,
		weight: weight,
		height: height,
	}
}

type DUmbDataBase struct {
	m map[string]string
}

func NewDumpDataBase() *DUmbDataBase {
	return &DUmbDataBase{
		m: make(map[string]string),
	}
}

// КАСТОМНЫЕ ТИПЫ ДАННЫХ В КОТОРЫЕ МОЖНО ДОБАВЛЯТЬ МЕТОДЫ
type Age int

func (a Age) isAdult() bool {
	return a >= 18
}
