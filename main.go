package main

import "fmt"

func main()  {
	var n uint64
	fmt.Scanln(&n)
	s := make([]int64,n)
	result := int64(0)
	for i := 0; i < int(n); i++ {
		fmt.Scanln(&s[i])
		result += s[i]
	}
	fmt.Println(result)


}