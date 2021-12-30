/**
 * @Author: Resynz
 * @Date: 2021/7/19 13:48
 */
package common

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(ctx *Context)

func AuthDetection(next HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := new(Context)
		context.Context = ctx
		next(context)
	}
}
