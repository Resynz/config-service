/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:24
 */
package common

import (
	"config-service/code"
	"net/http"
)

func HandleResponse(ctx *Context, c code.ResponseCode, d interface{}, msg ...string) {
	m := code.GetCodeMsg(c)
	if len(msg) > 0 {
		m = msg[0]
	}
	data := map[string]interface{}{
		"code":    c,
		"message": m,
	}
	if d != nil {
		data["data"] = d
	}
	ctx.JSON(http.StatusOK, data)
}
