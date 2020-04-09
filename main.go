package main

import (
	"net/http"

	"github.com/djigoio/traininggo/controllers"
)

//Web service

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
