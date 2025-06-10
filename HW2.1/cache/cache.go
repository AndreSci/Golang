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
	// Проверяем переменную
	c._timerDel(str)

	// Практикуем медленный mutex
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exist := c.m[str]

	if exist {
		return value.value, nil
	}

	return 0, errors.New("not found key")
}

// Метод удаления из массива
func (c *cache) Delete(str string) {
	// Практикуем медленный mutex
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.m, str)
}

// Функция которая удаляет элемент
func (c *cache) _timerDel(str string) {

	// Практикуем медленный mutex
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.m[str]

	if ok && time.Now().After(value.expireAt) {
		delete(c.m, str)
	}
}
