package controller

import (
	"github.com/gin-gonic/gin"
	"go_backend/data/request"
	"go_backend/data/response"
	"go_backend/helper"
	"go_backend/service"
	"net/http"
	"strconv"
)

type UsrRulesController struct {
	UsrRulesService service.UsrRulesService
}

func NewUsrRulesController(service service.UsrRulesService) *UsrRulesController {
	return &UsrRulesController{
		UsrRulesService: service,
	}
}
func (controller *UsrRulesController) Create(ctx *gin.Context) {
	createUsrRulesRequest := request.CreateUsrRulesRequest{}
	err := ctx.ShouldBindJSON(&createUsrRulesRequest)
	helper.ErrorPanic(err)

	controller.UsrRulesService.Create(createUsrRulesRequest)
	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsrRulesController) Update(ctx *gin.Context) {
	updateUsrRulesRequest := request.UpdateUsrRulesRequest{}
	err := ctx.ShouldBindJSON(&updateUsrRulesRequest)
	helper.ErrorPanic(err)

	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)
	updateUsrRulesRequest.Id = id

	controller.UsrRulesService.Update(updateUsrRulesRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsrRulesController) Delete(ctx *gin.Context) {
	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)
	controller.UsrRulesService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsrRulesController) FindId(ctx *gin.Context) {
	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)

	ruleResponse := controller.UsrRulesService.FindId(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ruleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsrRulesController) FindAll(ctx *gin.Context) {
	ruleResponse := controller.UsrRulesService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ruleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
