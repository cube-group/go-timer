package main

import (
	"cube/utils"
	"time"
	"fmt"
)

func main() {
	t := utils.SetInterval(1 * time.Second, func(args...interface{}) {
		fmt.Println(args)
	}, "hello")

	time.Sleep(10 * time.Second)

	t.Stop()

	utils.SetTimeout(2 * time.Second, func(args...interface{}) {
		fmt.Println("setTimeout")
	}, nil)

	time.Sleep(5 * time.Second)
}
