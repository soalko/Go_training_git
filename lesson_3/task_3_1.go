package main

import (
	"fmt"
)

func main() {
	var a = 0
	fmt.Scan(&a)
	if a%4 == 0 && a%100 != 0 {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}

}
