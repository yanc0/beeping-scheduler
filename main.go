package main

import (
	"flag"
	"github.com/yanc0/beeping-scheduler/scheduler"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		go func() {
			time.Sleep(time.Second * 30)
			pprof.StopCPUProfile()
		}()
	}
	s := scheduler.NewScheduler()
	s.Run()
}
