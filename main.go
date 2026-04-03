package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter folder path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	OrganizeFiles(path)
	defer fmt.Println("Files organized successfully!")
}
