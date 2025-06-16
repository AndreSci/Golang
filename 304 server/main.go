package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Требуется изучение JSON стиля в GO
type User struct {
	ID   int    `json:"id`
	Name string `json:"name,omitempty"`
}

// GLOBAL переменные
var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main() {
	// Каждый запрос обрабатывается в отдельной ГОрутине
	fmt.Println("START SERVER")

	// Классная функция определяет ROUTE
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsersGET)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Этот route Не является каким то методом GET/POST/DELETE /...
// Функция HANDLE для обработки url -> :8080/users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Marshal запаковываев в json(byte) какую то структуру
	resp, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// передавать нужно набор байт или структуры которые предназначены для этого
	// для передачи строки нужно преобразовывать в байт код []byte("HELLO")
	w.WriteHeader(http.StatusOK)
	RawContentLength, err := w.Write(resp)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RawContentLength:", RawContentLength)
}

// Этот route Не является каким то методом GET/POST/DELETE /...
// Функция HANDLE для обработки url -> :8080/users
func handleUsersGET(w http.ResponseWriter, r *http.Request) {

	fmt.Println("REQUEST INCOME")
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Marshal запаковываев в json(byte) какую то структуру
	resp, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// передавать нужно набор байт или структуры которые предназначены для этого
	// для передачи строки нужно преобразовывать в байт код []byte("HELLO")

	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
}

// Middleware ЛОГИРОВАНИЯ ДАННЫХ В ЗАПРОСЕ
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idFromCx := r.Context().Value("id")
		userID, ok := idFromCx.(string)

		if !ok {
			log.Printf("[%s] %s - error: userID is invalid", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("[%s] %s by userID %s\n", r.Method, r.URL, userID)
		next(w, r)
	}
}

// Middleware ЛОГИРОВАНИЯ ДАННЫХ В ЗАПРОСЕ
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usedID := r.Header.Get("x-id")

		if usedID =="" {
			log.Printf("[%s] %s - error UserID not provided\n", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", usedID)

		r = r.WithContext(ctx)

		next(w, r)
	}
}
