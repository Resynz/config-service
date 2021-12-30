/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:58
 */
package api

import (
	"config-service/code"
	"config-service/common"
	"config-service/config"
)

func GetConfig(ctx *common.Context) {
	type formValidate struct {
		Env  string `form:"env" binding:"required" json:"env"`
		Name string `form:"name" binding:"" json:"name"`
	}
	var form formValidate
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	if form.Name != "" {
		common.HandleResponse(ctx, code.SuccessCode, config.ConfMaps.Get(form.Env, form.Name))
		return
	}
	common.HandleResponse(ctx, code.SuccessCode, config.ConfMaps.GetByEnv(form.Env))
	return
}

func GetEnvs(ctx *common.Context) {
	common.HandleResponse(ctx, code.SuccessCode, config.Envs)
}
