//

package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("RUNNING ... ")
	robotgo.MoveMouse(1419, 763)
	robotgo.Sleep(1)
	robotgo.MouseClick("left", true)
	robotgo.Sleep(1)
	robotgo.MouseClick("left", true)
	robotgo.MoveClick(1419, 763)
	fmt.Println("AWESOME!!! ... ")
}

// pos:  1419 763
