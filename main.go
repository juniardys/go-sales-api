package main

import (
	"fmt"
	db "go-sales-api/config"
)

func main() {
	fmt.Println("Go sales api started...")
	db.Connect()
}
