package main

import (
	"flag"
	"fmt"
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
		if !force {
			if _, err := os.Stat(pidFile); err == nil {
				log.Printf("Could not create pid file, error: %v", err)
				panic(fmt.Sprintf("Pidfile %s exist", pidFile))
			}
		}
		pid := syscall.Getpid()
		pidf, err := os.Create(pidFile)
		if err != nil {
			log.Printf("Could not create pid file, error: %v", err)
			panic("Could not create pid file")
		}
		pidf.WriteString(fmt.Sprintf("%d", pid))
		pidf.Close()
		defer func() {
			log.Println("Cleaning up pid file")
			err := os.Remove(pidFile)
			if err != nil {
				log.Printf("Fail to clean up pid file %s", pidFile)
			}
		}()
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
