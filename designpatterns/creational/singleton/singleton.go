package singleton

import "sync"

var once sync.Once

type single struct {
	CreateCount           int
	MethodInvocationCount int
}

var Singleton *single

func getSingleton() *single {
	if Singleton == nil {
		once.Do(
			func() {
				Singleton = &single{}
				Singleton.CreateCount++
			},
		)
	}
	Singleton.MethodInvocationCount++
	return Singleton
}
