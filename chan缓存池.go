package quick_sort_test

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() chan interface{} {
	ch := make(chan interface{}, 1)
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}



func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	<-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)


}
