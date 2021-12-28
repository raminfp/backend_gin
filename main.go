package main

import (
	"github.com/raminfp/backend_gin/routes"
)


func main()  {

	r := routes.Urls()
	err := r.Run()
	if err != nil {
		return
	}
}
