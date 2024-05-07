package router

import (
	"github.com/gin-gonic/gin"
	"go_backend/controller"
	"net/http"
)

func NewRouter(usrRulesController *controller.UsrRulesController, processRulesController *controller.ProcessRulesController) *gin.Engine {
	routers := gin.Default()
	routers.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome go")
	})
	baseRouter := routers.Group("/api")
	usrRulesRouter := baseRouter.Group("/usr")
	usrRulesRouter.GET("", usrRulesController.FindAll)
	usrRulesRouter.GET("/:id", usrRulesController.FindId)
	usrRulesRouter.POST("/create", usrRulesController.Create)
	usrRulesRouter.PATCH("/:id", usrRulesController.Update)
	usrRulesRouter.DELETE("/:id", usrRulesController.Delete)

	processRulesRouter := baseRouter.Group("/process")
	processRulesRouter.GET("", processRulesController.FindAll)
	processRulesRouter.GET("/:id", processRulesController.FindId)
	processRulesRouter.POST("/create", processRulesController.Create)
	processRulesRouter.PATCH("/:id", processRulesController.Update)
	processRulesRouter.DELETE("/:id", processRulesController.Delete)
	return routers
}
