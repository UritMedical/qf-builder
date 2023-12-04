/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/1 9:08
 */

package main

import "fmt"

const cmd string = "quickly framework code builder"

func showHelp() {
	showVersion()
	fmt.Println(cmd, "qf need a command like -b build | -h help | -v version")
}
func showVersion() {
	fmt.Println(cmd, "version", version)
}
