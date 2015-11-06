package main

import  (
	"fmt"
	"citest/finder"
)

func main() {
	fmt.Println("FAT GOPHER TESTING")


	isit, err := finder.IsItGopher("gopher")
	if (err == nil && isit) {
		fmt.Println("FAT GOPHER DETECTED")
	}

	noit, err := finder.IsItGopher("weasel")
	if (err != nil && !noit) {
		fmt.Println(err)
	}
}
