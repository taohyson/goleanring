package main

import "sync"

type UserAgs struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAgs) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAgs) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAgs{}
	ua.ages = make(map[string]int)
	for i := 0; i < 10000; i++ {
		go ua.Add("a",1)
		go ua.Get("a")
		i++

	}

}
