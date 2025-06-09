package cache

import (
	"errors"
	"fmt"
	"time"
)

func New() *cache {
	return &cache{
		m: make(map[string]int),
	}
}

type cache struct {
	m map[string]int
}

func (c *cache) Set(str string, value any, ttl time.Duration) {
	switch id := value.(type) {
	case int:
		c.m[str] = id
		go func(){
			time.Sleep(ttl)
			c.Delete(str)
		}()
	default:
		fmt.Println("Wrong type of id")
	}

}

func (c cache) Get(str string) (int, error) {
	value, exist := c.m[str]

	if exist {
		return value, nil
	} 

	return 0, errors.New("not found key")
}

func (c *cache) Delete(str string) {
	delete(c.m, str)
}
