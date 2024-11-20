package main

import "github.com/math-schenatto/go-projects/api/configs"

func main() {
	configs, err := configs.LOadConfig(".")
	if err != nil {
		panic(err)
	}
}
