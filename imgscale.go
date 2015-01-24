package main

import (
	"flag"
	"fmt"
	"github.com/vanng822/gopid"
	"github.com/vanng822/imgscale/imgscale"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		configPath string
		host       string
		port       int
		pidFile    string
		force      bool
	)

	flag.StringVar(&host, "h", "127.0.0.1", "Host to listen on")
	flag.IntVar(&port, "p", 8080, "Port number to listen on")
	flag.StringVar(&configPath, "c", "./config/formats.json", "Path to configurations")
	flag.StringVar(&pidFile, "pid", "imgscale.pid", "Pid file")
	flag.BoolVar(&force, "f", false, "Force and remove pid file")
	flag.Parse()

	if pidFile != "" {		
		gopid.CheckPid(pidFile, force)
		gopid.CreatePid(pidFile)
		defer gopid.CleanPid(pidFile)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Kill, os.Interrupt, syscall.SIGTERM)

	handler := imgscale.Configure(configPath)
	defer handler.Cleanup()
	http.Handle("/", handler)
	log.Printf("listening to address %s:%d", host, port)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)

	sig := <-sigc
	log.Printf("Got signal: %s", sig)
}
