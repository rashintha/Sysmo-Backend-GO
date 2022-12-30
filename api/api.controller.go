package api

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matishsiao/goInfo"
)

func getAPI(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	gi, _ := goInfo.GetInfo()

	c.JSON(http.StatusOK, map[string]interface{}{
		"os":        runtime.GOOS,
		"arch":      runtime.GOARCH,
		"mem_usage": fmt.Sprintf("%v", m.Sys/1024/1024) + " MB",
		"kernel":    gi.Kernel,
		"core":      gi.Core,
		"host":      gi.Hostname,
		"cpu_count": gi.CPUs,
		"time":      time.Now(),
	})
}

func getAPIStatus(c *gin.Context) {
	fmt.Println("Hello, world.")
}
