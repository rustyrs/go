package main 

import "fmt"

type Status int 

const (
	Todo Status = 0
	InProgress  
	Done
)

func main() {
	var taskStatus Status = InProgress

	if taskStatus == InProgress {
		fmt.Println("進行中")
	}
}