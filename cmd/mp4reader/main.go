package main

import (
	"fmt"

	"github.com/AviParampampam/go-mp4/pkg/mp4"
)

func main() {
	v, err := mp4.NewVideo("example/videos/small.mp4")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(v)

	fmt.Println("End")
}
