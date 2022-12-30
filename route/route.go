package route

import (
	_ "sysmo/api"
	"sysmo/lib/env"
	"sysmo/lib/log"
	"sysmo/lib/routers"
	corsUtil "sysmo/middleware/cors"
	"sysmo/middleware/jwtauth"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	if env.CONF["MODE"] == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()
	Router.SetTrustedProxies(nil)
	Router.Use(corsUtil.CORS())
	Router.Use(jwtauth.AuthorizeJWT())

	log.Defaultln("Initializing routes")

	for key, value := range routers.GetRoutes {
		Router.GET(key, value)
	}

	for key, value := range routers.PostRoutes {
		Router.POST(key, value)
	}

	for key, value := range routers.PutRoutes {
		Router.PUT(key, value)
	}

	for key, value := range routers.DeleteRoutes {
		Router.DELETE(key, value)
	}
}
