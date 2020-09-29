package router

import (
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
	"goshop/front-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("member", new(controller.Member), r, middleware.VerifyToken())
		g.Get("/info")
		g.Post("/update")
	})
}
