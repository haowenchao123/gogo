//Page 6

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
