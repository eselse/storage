package main

import (
	"fmt"
	"log"

	"github.com/eselse/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It works.", file)
}
