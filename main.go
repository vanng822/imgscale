package main

import (
	"net/http"
	"middleware"
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	app := martini.Classic()
	formats := make([]*middleware.Format, 0)
	formats = append(formats, &middleware.Format{"100x100", 100, 100})
	formats = append(formats, &middleware.Format{"original", 0, 0})
	config := middleware.Config{"./data", "img", formats}	
	middleware.Martini(&config, app)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8080), app)
}

