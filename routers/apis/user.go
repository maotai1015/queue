package apis

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"queue/models"
)


func SaveInfo(ctx iris.Context)  {
	phone := ctx.PostValueDefault("phone", "123")
	err := models.SaveInfo(phone)
	if err != nil{
		fmt.Println("信息存储失败！")
	}

}
