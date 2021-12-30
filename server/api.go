/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:57
 */
package server

import (
	"config-service/common"
	"config-service/controller/api"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoute(route *gin.RouterGroup) {
	route.GET("/config", common.AuthDetection(api.GetConfig))
	route.GET("/envs", common.AuthDetection(api.GetEnvs))
}
