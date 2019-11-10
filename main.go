package main

import (
	"github.com/gin-gonic/gin"
	"ihomebj5q/model"
)

func main() {
	router:=gin.Default()
	r1:=router.Group("/v1")
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

	model.InitModel()
	router.Run(":8099")

}
