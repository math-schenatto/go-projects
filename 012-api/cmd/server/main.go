package main

import "github.com/math-schenatto/go-projects/api/configs"

func main() {
	_, err := configs.LOadConfig(".")
	if err != nil {
		panic(err)
	}
}
