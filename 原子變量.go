package quick_sort_test

import (
	"fmt"
	"sync/atomic"
)

var seed int64

func main() {
	atomic.AddInt64(&seed, 1)
	fmt.Println(seed)
}
