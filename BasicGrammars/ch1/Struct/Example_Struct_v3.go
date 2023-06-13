package main

import "fmt"

//简单的图书管理系统采用机构体和方法实现

// book 结构体
type Book struct {
	Title  string
	Author string
}

// 定义图书馆机构体
type Library struct {
	Books []Book
}

// 定义方法
func (lib *Library) AddBook(book Book) {
	lib.Books = append(lib.Books, book)
}

//
func (lib *Library) PrintBooks() {
	fmt.Println("Library Books:")
	for _, book := range lib.Books {
		fmt.Printf("Titie: %s, Author: %s\n", book.Title, book.Author)
	}
}

func main() {
	// 创建一个Library结构体实例
	library := Library{}
	//添加图书
	library.AddBook(
		Book{
			Title:  "The Great Gatsby",
			Author: "abbott",
		})
	library.AddBook(
		Book{
			Title:  "JAVA",
			Author: "abbott",
		})
	library.AddBook(
		Book{
			Title:  "C",
			Author: "abbott",
		})
	//打印图书管中的图书
	library.PrintBooks()
}
