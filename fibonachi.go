package main

import (
	"fmt"
)

func Fibonachi(num int) int  {
	if num <= 1 {
		return num
	}
	return Fibonachi(num - 1) + Fibonachi(num - 2)

}


func main ()  {

	var num int

	fmt.Println("Enter the numbers of line")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Print(Fibonachi(i), "\t")
	}

}
