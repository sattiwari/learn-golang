package main

import "fmt"

func fun2(arr []int) {
	for num, _ := range arr {
		fmt.Println(num)
	}
}

func fun1(arr []int) {
	fun2(arr)
}

func main() {
	arr := [2]int{1, 2}
	fun1(arr[:])
}
