package main

//数据转换：假设有一个字符串切片，需要将其中的每个字符串转换为大写。
import (
	"fmt"
	"strings"
)

func main() {
	word := []string{"banana", "apple", "orange"}
	var result []string

	for _, wor := range word {
		upper := strings.ToUpper(wor)
		result = append(result, upper)
	}
	fmt.Println(result)
}
