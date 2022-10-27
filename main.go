package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"queue/models"
)

func main() {
	if err:= models.InitDB();err != nil{
		fmt.Println("连接数据库失败！")
	}
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.Run(iris.Addr(":8080"))
}
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}