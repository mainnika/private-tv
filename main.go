package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lidouf/glib"
	"github.com/lidouf/gst"
	"github.com/mainnika/private-tv/backend"
	"github.com/mainnika/private-tv/server"
)

func usage() {

	flag.Usage()
	os.Exit(1)
}

func main() {

	var src string
	flag.StringVar(&src, "src", "", "media file path")
	flag.Parse()

	info, err := os.Stat(src)
	switch true {
	case len(src) == 0:
		usage()
	case err != nil:
		log.Fatal(err)
	case info.IsDir():
		log.Fatal(fmt.Errorf("file %s is a directory", src))
	}

	err, pipeline := backend.NewPipeline()
	if err != nil {
		log.Fatal(err)
	}

	err, _ = server.NewServer(9000)
	if err != nil {
		log.Fatal(err)
	}

	pipeline.SetSource(src)
	pipeline.SetState(gst.STATE_PLAYING)
	pipeline.GetState(gst.CLOCK_TIME_NONE)

	glib.NewMainLoop(nil).Run()
}
