/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:51
 */
package server

import (
	"config-service/config"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	gin.SetMode(config.Mode)
	app := gin.New()
	app.MaxMultipartMemory = 1024 << 20 // 1GB
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	_ = app.SetTrustedProxies(nil)

	apiGroup := app.Group("/api")
	RegisterApiRoute(apiGroup)

	if err := app.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatalf("start server failed! error:%s\n", err.Error())
	}
}
