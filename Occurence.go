package main

import "fmt"

func find_occurences(arr []int) {
	freq := make(map[int]int)
	for _, num := range arr {
		if freq[num] == 0 {
			freq[num] = 1
		} else {
			freq[num]++
		}
	}
	for num, c := range freq {
		fmt.Printf("Element %d : %d times.\n", num, c)
	}
}
func main() {
	var n int
	fmt.Print("Enter size of array: ")
    fmt.Scan(&n)  
	fmt.Println("Enter array elements")
	var arr = make([]int, n)
	
	for i:=0;i<n;i++ {
		fmt.Scan(&arr[i])
	}

	find_occurences(arr)
}