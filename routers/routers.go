package routers

import (
"github.com/kataras/iris/v12"
v1 "queue/routers/apis"

)

func InitRouter(enable_log bool) *iris.Application {
	app := iris.New()

	app.Options("/*", func(ctx iris.Context) {
		ctx.Next()
	})

	// 健康检测接口
	app.Get("/hs", func(ctx iris.Context) {
		// ctx.StatusCode(200)
		ctx.WriteString("OK")
	})
	userOpt := app.Party("/user")
	{
		userOpt.Get("", v1.SaveInfo)
	}
	return app
}