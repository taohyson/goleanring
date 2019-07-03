package quick_sort_test

import (
	"fmt"
)

const cccc = 1

func DoTheThing(reallyDoIt bool) (err int) {
	//fmt.Printf("%v\n",&cccc)
	if reallyDoIt {
		err :=1
		fmt.Printf("%v\n",&err)
	}
	return err
}

func tryTheThing() (string, int) {
	return "", cccc
}

func main() {
	fmt.Println(DoTheThing(true))
	//fmt.Println(DoTheThing(false))
}
