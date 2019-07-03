package quick_sort_test

import "fmt"

type zoo struct {
	name string
}

var (
	pig   = zoo{"pig"}
	sheep = zoo{"sheep"}

	a = 1
	b = 2
)

func main() {
	pig, sheep = sheep, pig
	{
		pig.name = "pig1"
	}
	a, b = b, a
	fmt.Println(sheep)
	fmt.Println(pig)
	{{
		 a := 3
		fmt.Println(a)
	}

	}
	fmt.Println(a)
	fmt.Println(b)
}
