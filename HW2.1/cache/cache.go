package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Обьявляем структуру
type cache struct {
	m map[string]int
	expireAt map[string]time.Time
	ttlHis map[string]time.Duration
	mu sync.Mutex
}

// Создаем образец структуры
func New() *cache {
	return &cache{
		m: make(map[string]int),
		expireAt: make(map[string]time.Time),
		ttlHis: make(map[string]time.Duration),
	}
}

// Добавляем элемент в массивы и при необходимости создаем горутину для самоудалении 
func (c *cache) Set(str string, value any, ttl time.Duration) {
	switch id := value.(type) {
	case int:
		c.mu.Lock()

		_, ok := c.m[str]

		if !ok {
			go c.timerDel(str)
		}
		c.m[str] = id
		c.expireAt[str] = time.Now()
		c.ttlHis[str] = ttl
		c.mu.Unlock()
		
	default:
		fmt.Println("Wrong type of id")
	}

}

// Получаем элемент из массива
func (c *cache) Get(str string) (int, error) {

	c.mu.Lock()
	value, exist := c.m[str]
	c.mu.Unlock()

	if exist {
		return value, nil
	} 

	return 0, errors.New("not found key")
}

// Метод удаления из массива
func (c *cache) Delete(str string) {
	c.mu.Lock()
	delete(c.m, str)
	delete(c.expireAt, str)
	delete(c.ttlHis, str)
	c.mu.Unlock()
}

// Функция для горутинны которая удаляет элемент
func (c * cache) timerDel(str string) {
	for {
		time.Sleep(100 * time.Millisecond)

		c.mu.Lock()
		if time.Now().After(c.expireAt[str].Add(c.ttlHis[str])) {
			delete(c.m, str)
			delete(c.expireAt, str)
			delete(c.ttlHis, str)
			c.mu.Unlock()
			break
		}else{
			c.mu.Unlock()
		}
	}
}
