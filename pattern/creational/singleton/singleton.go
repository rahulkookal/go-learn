package singleton

import "sync"

type Singleton struct {
	Name string
}

var instance *Singleton
var once sync.Once

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "Rahul"}
	})
	return instance
}
