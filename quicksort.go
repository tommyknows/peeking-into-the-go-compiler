package main

import (
	"fmt"
	"os/exec"
)

// START OMIT
func quicksort(p []int) []int {
	if len(p) == 0 {
		return []int{}
	}

	lesser := filter(func(x int) bool { return p[0] > x }, p[1:])   // HLcustom
	greater := filter(func(x int) bool { return p[0] <= x }, p[1:]) // HLcustom
	return append(
		quicksort(lesser),
		prepend(p[0], quicksort(greater))..., // HLcustom
	)
}

// END OMIT

func main() {
	fmt.Println(quicksort([]int{1, 8, 5, 3, 4, 9}))
	funcheck()
}

func funcheck() {
	cmd := exec.Command("/go/bin/funcheck", "./code/quicksort/main.go")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println("funcheck: no reported errors")
		return
	}
	fmt.Printf("%s\n%s\n", out, err)
}
