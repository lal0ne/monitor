package main

import (
	// "fmt"
	"monitor/lib"
)

func main() {
	// lib.CFG.Monitor.IGNOREFLAG = true
	// fmt.Println(lib.CFG)
	// if len(lib.CFG.Monitor.IGNOREPATH) != 0 {
	// 	fmt.Println(lib.CFG.Monitor.IGNOREPATH)
	// } else {
	// 	fmt.Println("no")
	// }
	lib.Run()
}
