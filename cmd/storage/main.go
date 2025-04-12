package main

import (
	"fmt"

	"github.com/eselse/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("It works.", st)
}
