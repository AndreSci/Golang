package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// func main() {

// 	// go fmt.Println("This guy is not concurent.")
// 	// go fmt.Println("This guy is not concurent.")
// 	// go fmt.Println("This guy is not concurent.")

// 	time.Sleep(3 * time.Second)

// 	fmt.Println("This guy is the best applicant.")
// }

var numberUsers int = 10


// СЛАЙС для шаблонных данных
var actions = []string {
	"logged in",
	"logged out",
	"create record",
	"delete record",
}

// Структура для логов
type logItem struct {
	action string
	timestamp time.Time
}
// Структура для пользователя
type User struct {
	id int
	email string
	logs []logItem
}
// Функция для структуры User
func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}
	return out
}

// func main () {
// 	u := User {
// 		id: 1, 
// 		email: "Andrew@go.ru",
// 		logs: []logItem {
// 			{actions[0], time.Now()},
// 			{actions[3], time.Now()},
// 			{actions[2], time.Now()},
// 			{actions[1], time.Now()},
// 			{actions[0], time.Now()},
// 			{actions[3], time.Now()},
// 		},
// 	}
// 	fmt.Println(u.getActivityInfo())
// }


func main() {
	// rand.Seed(time.Now().Unix())

	users := generateUsers(numberUsers)

	// for _, user := range users {
	// 	fmt.Println(user.getActivityInfo())
	// }

	for _, user := range users {
		saveUserInfo(user)
	}
}

// Функция обработчик данных для СЛАЙСа Users (в данном случаи генератор данных)
func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User {
			id: i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs: generateLogs(rand.Intn(10)),				// Генератор чисем (rand.Intn(100))
		}
	}
	return users
}
// Функция обработчик данных для СЛАЙСа logItem(в данном случаи генератор данных)
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem {
			timestamp: time.Now(),
			action: actions[rand.Intn(len(actions) - 1)],
		}
	}

	return logs
}

// Функция сохраняет данные в файл
func saveUserInfo(user User) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())

	return err
}
