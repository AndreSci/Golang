package main

import (
	"fmt"
	"log"
	"test/httpclient/coincap"
	"time"
)

func main() {
	fmt.Println("HI")

	// COOKIES данные
	// jar, errJ := cookiejar.New(nil)

	// if errJ != nil {
	// 	log.Fatal(errJ)
	// }

	// jar.SetCookies(нужен слайс данных)

	// Пример создание клиента с Редиректом если сайт перенаправляет на другую ссылку
	// client := &http.Client{
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 		fmt.Println(req.Response.Status)
	// 		fmt.Println("REDIRECT")
	// 		return nil
	// 	},
	// 	// Определяем наш middleware для logger
	// 	Transport: &loggingRoundTripper{
	// 		logger: os.Stdout,
	// 		next:   http.DefaultTransport,
	// 	},
	// 	// Кастомный тайм-аут
	// 	Timeout: time.Second * 10,
	// 	// COOKIES данные
	// 	//Jar: jar,
	// }

	coincapClient, err := coincap.NewClient(time.Second * 10)

	if err != nil {
		log.Fatal(err)
	}

	assets, errorAs := coincapClient.GetAssets()

	if errorAs != nil {
		log.Fatal(errorAs)
	}

	index := 0
	for _, asset := range assets {
		fmt.Println(asset.Info())
		index++

		if index > 200 {
			break
		}
	}
}
