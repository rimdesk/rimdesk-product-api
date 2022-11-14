package main

import (
	"github.com/rimdesk/product-api/api/bootstrap"
	"github.com/rimdesk/product-api/api/config"
)

func main() {
	app := bootstrap.New()
	app.StartAndListen("0.0.0.0", config.Get("SERVER_PORT"))
}
