package main

import (
	"fmt"
	"os"

	"github.com/ezratameno/wolt/internal/wolt"
)

func main() {
	w := wolt.New(os.Getenv("token"))
	err := w.GetSummary()
	if err != nil {
		fmt.Println(err)
	}
}
