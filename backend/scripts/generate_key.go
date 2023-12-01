package main

import (
	"fmt"
	"niya-voyage/backend/pkg/auth"
)

func main() {
	fmt.Println(auth.GenerateRandomKey())
}
