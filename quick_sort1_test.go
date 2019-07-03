package quick_sort1_test

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	datas := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(datas, 0, len(datas)-1)
	fmt.Println(datas)
}

func quickSort(datas []int, start, end int) {
	i, j := start, end
	key := datas[(start+end)/2]
	for i <= j {
		for datas[i] < key {
			i++
		}
		for datas[j] > key {
			j--
		}
		if i <= j {
			datas[i], datas[j] = datas[j], datas[i]
			i++
			j--
		}
	}
	if start < j {
		quickSort(datas, start, j)
	}
	if end > i {
		quickSort(datas, i, end)
	}

}
