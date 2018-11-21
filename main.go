package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Print(uuid.New().String())
}
