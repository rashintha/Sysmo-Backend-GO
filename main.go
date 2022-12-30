package main

import (
	"sysmo/lib/env"
	"sysmo/lib/log"
	"sysmo/route"
)

func init() {
	log.Defaultln("Starting server on " + env.CONF["HOST"] + ":" + env.CONF["PORT"])
}

func main() {
	route.Router.Run(env.CONF["HOST"] + ":" + env.CONF["PORT"])
}
