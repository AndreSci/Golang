package cache

import "fmt"

func New() *cache {
	return &cache{
		m: make(map[string]int),
	}
}

type cache struct {
	m map[string]int
}

func (c *cache) Set(str string, value interface{}) {
	switch id := value.(type) {
	case int:
		c.m[str] = id
	default:
		fmt.Println("Wrong type of id")
	}

}

func (c cache) Get(str string) int {
	return c.m[str]
}

func (c *cache) Delete(str string) {
	delete(c.m, str)
}
