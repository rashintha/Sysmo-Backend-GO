package apiAuth

import "sysmo/lib/routers"

func init() {
	routers.Post("/api/auth/login", login, false)
}
