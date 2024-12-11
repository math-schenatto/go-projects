package main

import (
	"fmt"

	"github.com/math-schenatto/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
