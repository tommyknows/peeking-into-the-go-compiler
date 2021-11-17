package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

// START OMIT
func quicksort(p []int) []int {
	if len(p) == 0 {
		return []int{}
	}

	lesser := filter(func(x int) bool { return p[0] > x }, p[1:])
	greater := filter(func(x int) bool { return p[0] <= x }, p[1:])
	return append(
		quicksort(lesser),
		prepend(p[0], quicksort(greater))...,
	)
}

func main() {
	fmt.Println(quicksort([]int{1, 8, 5, 3, 4, 9}))
	funcheck()

	res := fmap(func(x int) string { return strconv.Itoa(x) }, []int{1, 2, 3})
	fmt.Printf("%T, %v\n", res, res)
}

// END OMIT
func funcheck() {
	//cmd := exec.Command("/Users/ramon/Documents/go/bin/funcheck", "./code/quicksort/main.go")
	cmd := exec.Command("/Users/ramon/Documents/go/bin/funcheck", ".")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println("funcheck: no reported errors")
		return
	}
	fmt.Printf("%s\n%s\n", out, err)
}
