package main 

import (
	"fmt"
	"net/http"
)

func main() {
	// Routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server: http://localhost:8080")

	// Serve 
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("!Error!: %s\n", err)
	}
}