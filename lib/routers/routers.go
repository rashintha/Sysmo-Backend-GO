package routers

import (
	"fmt"
	"sysmo/lib/env"

	"github.com/gin-gonic/gin"
)

var GetRoutes = map[string]gin.HandlerFunc{}
var NoAuthGetPaths = []string{}
var PostRoutes = map[string]gin.HandlerFunc{}
var NoAuthPostPaths = []string{}
var PutRoutes = map[string]gin.HandlerFunc{}
var NoAuthPutPaths = []string{}
var DeleteRoutes = map[string]gin.HandlerFunc{}
var NoAuthDeletePaths = []string{}

func Get(path string, handler gin.HandlerFunc, authRequired ...bool) {
	authReq := true
	finalPath := fmt.Sprintf("%s%s", env.CONF["API_PREFIX"], path)

	if len(authRequired) > 0 {
		authReq = authRequired[0]
	}

	if !authReq {
		NoAuthGetPaths = append(NoAuthGetPaths, finalPath)
	}

	GetRoutes[finalPath] = handler
}

func Post(path string, handler gin.HandlerFunc, authRequired ...bool) {
	authReq := true
	finalPath := fmt.Sprintf("%s%s", env.CONF["API_PREFIX"], path)

	if len(authRequired) > 0 {
		authReq = authRequired[0]
	}

	if !authReq {
		NoAuthPostPaths = append(NoAuthPostPaths, finalPath)
	}

	PostRoutes[finalPath] = handler
}

func Put(path string, handler gin.HandlerFunc, authRequired ...bool) {
	authReq := true
	finalPath := fmt.Sprintf("%s%s", env.CONF["API_PREFIX"], path)

	if len(authRequired) > 0 {
		authReq = authRequired[0]
	}

	if !authReq {
		NoAuthPutPaths = append(NoAuthPutPaths, finalPath)
	}

	PutRoutes[finalPath] = handler
}

func Delete(path string, handler gin.HandlerFunc, authRequired ...bool) {
	authReq := true
	finalPath := fmt.Sprintf("%s%s", env.CONF["API_PREFIX"], path)

	if len(authRequired) > 0 {
		authReq = authRequired[0]
	}

	if !authReq {
		NoAuthDeletePaths = append(NoAuthDeletePaths, finalPath)
	}

	DeleteRoutes[finalPath] = handler
}
