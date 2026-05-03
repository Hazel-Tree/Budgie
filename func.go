// extra functions
package main

import (
	"fmt"
)

func clear() {
	fmt.Print("\033[2J\033[H")
}

func ulb(str string) string { // underline bold
	return "\033[4m\033[1m" + str + "\033[m"
}
