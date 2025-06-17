package main

// Хороший опыт работы через SQLX
// SQLX лучше чем базовая библиотека Golang

// go mod init sql_test
// go get -u github.com/lib/pq
// требуется для установки драйвера коннекта к БД postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Имптор для сторонних эффектов
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	fmt.Println("Hi")
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=goLANGn1nja")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success Open connection to PSQL")
	}
	defer db.Close()

	if err2 := db.Ping(); err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("Success Ping to PSQL")
	}

	fmt.Println("CONNECTED")

	// err = insertUser(db, User{Name: "Andrew", Email: "andrew@mail.com", Password: "somepassword"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	users, err := getAllusers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

	user, err := getUserByID(1, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	// err = deletetUser(db, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// users, err = getAllusers(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(users)
}

func getAllusers(db *sql.DB) ([]User, error) {
	//	ВАРИАНТ СТРОГОЙ ТИПИЗАЦИИ
	rows, err3 := db.Query("select * from users")
	if err3 != nil {
		return nil, err3
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		u := User{}
		// сканируем и записываем все поля из БД в Структуру
		errf := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if errf != nil {
			return nil, errf
		}
		users = append(users, u)
	}

	errR := rows.Err()

	if errR != nil {
		return nil, errR
	}

	return users, nil
}

func getUserByID(id int, db *sql.DB) (User, error) {
	// ВАРИАНТ ЕДИНИЧНОЙ ВЫГРУЗКИ
	var u User
	err := db.QueryRow("select * from users where id = $1", 2).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	// запрос заполняется через $ Для POSTGRES , для MySQL используют ?
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("NO ROWS")
			return u, err
		}
		return u, err
	}

	return u, nil
}

// В данной функции показан стиль с транзакциями и RollBack если транзакция неудачная.
func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3)", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into logs (entity, action) values ($1, $2)", "user", "created")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func deletetUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)

	return err
}
