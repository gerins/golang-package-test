package golangpackagetest

import "fmt"

func PrintTenTimes(value string) {
	for i := 0; i < 10; i++ {
		fmt.Println(value)
	}
	fmt.Println("done")
}
