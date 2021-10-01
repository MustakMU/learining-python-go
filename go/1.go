package main

import "fmt"

func main() {
	var arr = []int{}
	arr = append(arr, 11)
	println(len(arr))
	for i := 0; i <= 10; i++ {
		fmt.Println("Even")

		arr = append(arr, i)
	}
	for i, a := range arr {
		println(i, a)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			println("ODD")
		}

	}
	t := func(fdfv string) {
		fmt.Println("Even ", fdfv)

	}
	t("sdfdf")
}
