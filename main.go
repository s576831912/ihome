package main

import (
	"github.com/gin-gonic/gin"
	"ihomebj5q/model"
	"fmt"
)

func main() {
	router:=gin.Default()
	r1:=router.Group("/v1")
	err:=model.InitDb()
	if err != nil{
		//将错误打印到日志文件中
		fmt.Println(err)
		return
	}
	{
		r1.GET("/efg", func(ctx *gin.Context) {
			ctx.Writer.WriteString("adbcdw")
		})
	}
	r2:=router.Group("/v2")
	{
		r2.GET("/efg", func(ctx *gin.Context) {
			ctx.Writer.WriteString("11112334")
		})
	}

	router.Run(":8099")

}
