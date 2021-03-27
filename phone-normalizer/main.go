package main

import (
	"fmt"
	"regexp"
)

func normalize(n string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(n, "")
}

func main() {
	fmt.Println("vim-go")
}
