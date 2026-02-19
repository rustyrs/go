package main 

import (
	"fmt"
)

type Summary interface {
	Summarize() string 
}

type Article struct {
	Content string 
}

func (a Article) Summarize() string {
	return a.Content 
}

func PrintSummary[T Summary](item T) {
	fmt.Println(item.Summarize())
}

func main() {
	a := Article {
		Content: "Golang",
	}

	PrintSummary(a)
}
