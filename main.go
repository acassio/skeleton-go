package main

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/macaron.v1"

	config "github.com/acassio/base-app-go/conf"
	conf "github.com/acassio/base-app-go/conf/app"
)

// application entrypoint
func main() {

	app := macaron.Classic()
	//app.Use(auth.Basic(config.Cfg.Section("").Key("user").String(), config.Cfg.Section("").Key("password").String()))
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	app.Run(port())

}

// configure http port
func port() int {
	port, err := config.Cfg.Section("").Key("http_port").Int()
	if err != nil {
		log.Fatal(err)
	}

	if forceLocal, _ := config.Cfg.Section("").Key("force_local_http_port").Bool(); forceLocal == false {
		if i, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
			port = i
		}
	}

	return port

}
