//go:generate ../scripts/protoc-generate.sh

package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(os.Getenv("DB_HOST"))
}
