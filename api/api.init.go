package api

import (
	_ "sysmo/api/auth"
	_ "sysmo/api/users"
	"sysmo/lib/routers"
)

func init() {
	// API
	routers.Get("/api", getAPI, false)
	routers.Get("/api/status", getAPIStatus, false)
}
