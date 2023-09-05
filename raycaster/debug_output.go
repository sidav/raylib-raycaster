package raycaster

import "fmt"

func debugPrint(line string) {
	fmt.Println(line)
}

func debugPrintf(line string, args ...interface{}) {
	fmt.Printf(line, args...)
}
