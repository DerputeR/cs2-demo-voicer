package main

import (
	"fmt"
	"log"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
)

func main() {
	f, err := os.Open("G:/Games/CS2 Demos/2024-10-17_Anubis_11-13.dem")
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)

	var eventsInCurrentFrame []interface{}

	// register handler on all events
	p.RegisterEventHandler(func(e interface{}) {
		eventsInCurrentFrame = append(eventsInCurrentFrame, e)
	})

	// parse frame by frame
	for ok := true; ok; ok, err = p.ParseNextFrame() {
		checkError(err)

		// iterate over events in frame
		for _, event := range eventsInCurrentFrame {
			handleEvent(event)
		}

		// reset event list
		eventsInCurrentFrame = eventsInCurrentFrame[:0]
	}
}

func handleEvent(event interface{}) {
	fmt.Println(event)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
