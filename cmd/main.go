package main

import (
	"github.com/YurcheuskiRadzivon/to_do_app"
	"log"
)

func main() {
	srv := new(to_do_app.Server)
	err := srv.Run("8080")
	if err != nil {
		log.Fatalf("error occured while running http server", err.Error())
	}

}
