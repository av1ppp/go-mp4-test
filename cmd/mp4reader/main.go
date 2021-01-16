package main

import (
	"fmt"
	"time"

	"github.com/AviParampampam/go-mp4/pkg/mp4"
)

func main() {
	v, err := mp4.NewVideo("example/videos/small.mp4")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go v.ParseAtoms()

	time.Sleep(time.Second * 5)

	fmt.Println("End")
}
