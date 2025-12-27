// Singleton is a creational design pattern that lets you
// ensure that a class has only one instance, while providing a global access point to this instance

package singelton

import (
	"sync"
)

type singleton struct {
	value int
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			value: 42,
		}
	})
	return instance
}
