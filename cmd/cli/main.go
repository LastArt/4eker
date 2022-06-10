package main

import (
	"4eker/pkg"
	"4eker/set"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(string(set.Yellow), set.Headlogo, string(set.ResetColor))
	fmt.Println(string(set.Blue), set.Logo)
	fmt.Println(string(set.Yellow), set.Sublogo, string(set.ResetColor))
	pkg.GetStart()
}
