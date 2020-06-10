package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

import (
	"fmt"
	"strings"
)

// "time"
func main() {

	message := "TOTAL: 5 FAILED, 7 SUCCESS"

	if strings.Index(message, "TOTAL:") >= 0 {

		message
		fmt.Printf("Total")
	}

	// fmt.Printf("Total %s\n")
}
