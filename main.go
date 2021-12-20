package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/jevans40/Ruthenium/component"
	"github.com/jevans40/Ruthenium/game"
	"github.com/jevans40/Ruthenium/world"
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

	newGame := game.NewGameECS()
	err := newGame.Init()
	if err != nil {
		panic(err)
	}
	channel := newGame.GetRenderChannel()

	dispatcher := world.NewSimpleDispatcher()
	transforms := component.NewDenseStorage[world.Transform]()
	transwrite, _ := component.GetWriteStorage[world.Transform](transforms)
	sprites := component.NewDenseStorage[world.Sprite]()
	spritewrite, _ := component.GetWriteStorage[world.Sprite](sprites)
	rendererService := world.NewRenderService(channel)
	dispatcher.AddService(rendererService)
	dispatcher.AddStorage(transforms)
	dispatcher.AddStorage(sprites)
	testt := world.Transform{10, 10, 10, 100, 100}
	tests := world.Sprite{0, 0, 0, 0, 0, [4]uint8{255, 0, 255, 255}}
	transwrite.AddEntity(component.EntityID(1), testt)
	spritewrite.AddEntity(component.EntityID(1), tests)

	go update(dispatcher)

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

func update(dis world.Dispatcher) {
	for {
		time.Sleep(time.Millisecond * 10)
		dis.Maintain()

	}
}
