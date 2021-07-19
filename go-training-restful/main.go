package main

import (
	"fmt"
	"os"
)

func main() {
	connectionString := os.Getenv("CONNECTION_STRING")
	fmt.Println(connectionString)
}
