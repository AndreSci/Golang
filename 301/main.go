package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Базовый Middleware для логгирования запросов
type loggingRoundTripper struct {
	logger io.Writer
	next http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %sn", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	fmt.Println("HI")

	// COOKIES данные
	// jar, errJ := cookiejar.New(nil)

	// if errJ != nil {
	// 	log.Fatal(errJ)
	// }

	// jar.SetCookies(нужен слайс данных)

	// Пример создание клиента с Редиректом если сайт перенаправляет на другую ссылку
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECT")
			return nil
		},
		// Определяем наш middleware для logger
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next: http.DefaultTransport,
		},
		// Кастомный тайм-аут
		Timeout: time.Second * 10,
		// COOKIES данные
		//Jar: jar,
	}


	respons, err := client.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}

	defer respons.Body.Close()

	fmt.Println("Respons status:", respons.StatusCode)

	body, errB := io.ReadAll(respons.Body)

	if errB != nil{
		log.Fatal(errB)
	}

	fmt.Println(string(body))	// приводим типы - string и есть байт код
}
