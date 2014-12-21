package main

import (
	"net/http"
	"imgscale"
	"fmt"
	"github.com/go-martini/martini"
	"os"
	"encoding/json"
)

func ReadConfig(filename string) *imgscale.Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := imgscale.Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func main() {
	config := ReadConfig("./config/formats.json")
	fmt.Println(config)
	// Martini
	app := martini.Classic()
	app.Use(imgscale.Middleware(config))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), app)
	// http.HandleFunc
	/*
	http.HandleFunc("/", middleware.Middleware(config))
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), nil)*/
}

