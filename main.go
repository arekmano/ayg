package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		x, y := robotgo.GetMousePos()
		fmt.Println(time.Now(), "pos: ", x, y)
		time.Sleep(time.Second)
	}
}
