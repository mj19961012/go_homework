package main

import "fmt"

func printArr[T any](arr []T) {
	fmt.Printf("arr:%v,len:%d,cap:%d\n", arr, len(arr), cap(arr))
}

func main() {
	numArr := []int{1, 2, 3, 4, 5}
	strArr := []string{"1", "2", "3", "4", "5"}
	printArr(numArr)
	numArr = DeleteWithIndex(numArr, 3)
	printArr(numArr)
	printArr(strArr)
	strArr = DeleteWithIndex(strArr, 3)
	printArr(strArr)
	numArr1 := []int{1, 2}
	printArr(numArr1)
	numArr1 = DeleteWithIndex(numArr1, 3)
	printArr(numArr1)
	numArr1 = DeleteWithIndex(numArr1, 1)
	printArr(numArr1)
	numArr2 := make([]int, 3, 9)
	printArr(numArr2)
	numArr2 = DeleteWithIndex(numArr2, 1)
	printArr(numArr2)
}
