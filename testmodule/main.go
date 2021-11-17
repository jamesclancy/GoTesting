package main

import (
	"net/http"

	"github.com/jamesclancy/GoTestApi/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
