package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type data struct {
	value int
	expireAt time.Time
}
// Обьявляем структуру
type cache struct {
	m        map[string]data
	mu       sync.Mutex
}

// Создаем образец структуры
func New() *cache {
	return &cache{
		m:        make(map[string]data),
	}
}

// Добавляем элемент в массивы и при необходимости создаем горутину для самоудалении
func (c *cache) Set(str string, value any, ttl time.Duration) {
	switch id := value.(type) {
	case int:
		c.mu.Lock()
		c.m[str] = data{value: id, expireAt: time.Now().Add(ttl)}
		c.mu.Unlock()

	default:
		fmt.Println("Wrong type of id")
	}

}

// Получаем элемент из массива
func (c *cache) Get(str string) (int, error) {

	c.mu.Lock()
	// Проверяем переменную
	c._timerDel(str)
	value, exist := c.m[str]
	c.mu.Unlock()

	if exist {
		return value.value, nil
	}

	return 0, errors.New("not found key")
}

// Метод удаления из массива
func (c *cache) Delete(str string) {
	c.mu.Lock()
	delete(c.m, str)
	c.mu.Unlock()
}

// Функция которая удаляет элемент
func (c *cache) _timerDel(str string) {

	value, ok := c.m[str]

	if ok && time.Now().After(value.expireAt) {
		delete(c.m, str)
	}
}
