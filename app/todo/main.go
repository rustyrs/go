package main 

import "fmt"

type Task struct {
	Title  string 
	IsDone bool
}

func (task *Task) CompleteTask() {
	task.IsDone = true
}

func main() {
	todo1 := Task {
		Title:  "買い物",
		IsDone: false,
	}

	todo2 := Task {
		Title:  "洗濯",
		IsDone: false,
	}

	fmt.Println(todo1)	
	fmt.Println(todo2)

	todo1.CompleteTask()
	fmt.Println(todo1)	
}