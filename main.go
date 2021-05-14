package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/jevans40/go-Boids/boidobjects"
	"github.com/jevans40/psychic-spork/event"
	"github.com/jevans40/psychic-spork/game"
)

var cpuprofile = flag.String("cpuprofile", "cputest", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "memtest", "write memory profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var newGame game.Game
	err := newGame.Init()
	if err != nil {
		panic(err)
	}

	go func() {
		events := make([]event.UpdateEvent, 200000)
		for i := 0; i < 200000; i++ {
			var newSquare boidobjects.Square
			newSquare.Init(i)

			events[i] = event.UpdateEvent{EventCode: event.UpdateEvent_NewObject,
				Receiver: -1,
				Sender:   -1,
				Event:    event.UpdateEvent_NewObjectEvent{&newSquare}}
		}
		newGame.EventChannel <- events
	}()

	go newGame.Start()

	time.Sleep(time.Second * 180)
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
