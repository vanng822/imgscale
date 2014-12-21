package main

import (
	"net/http"
	"middleware"
	"fmt"
	"github.com/go-martini/martini"
	"os"
	"encoding/json"
)

func ReadConfig(filename string) *middleware.Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := middleware.Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func main() {
	app := martini.Classic()
	config := ReadConfig("./config/formats.json")
	fmt.Println(config)
	middleware.Martini(config, app)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), app)
}

