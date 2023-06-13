package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{1, 2, 3, 4, 5, 10, 20, 15, 19}
	var result []int
	for _, num := range number {
		if num > 10 {
			result = append(result, num)
		}
	}
	//排序从小到大  //升序
	sort.Ints(result)
	fmt.Println(result)
	//字符串排序：默认按照字母顺序
	words := []string{"banana", "apple", "orange"}
	sort.Strings(words)
	fmt.Println(words)

}
