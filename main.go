package main

import (
	"monitor/cmd"
	// "fmt"
	_ "monitor/conf"
)

func main() {
	// lib.CFG.Monitor.IGNOREFLAG = true
	// fmt.Println(lib.CFG)
	// if len(lib.CFG.Monitor.IGNOREPATH) != 0 {
	// 	fmt.Println(lib.CFG.Monitor.IGNOREPATH)
	// } else {
	// 	fmt.Println("no")
	// }
	cmd.Run()
}
