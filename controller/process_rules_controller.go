package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go_backend/data/request"
	"go_backend/data/response"
	"go_backend/helper"
	"go_backend/service"
	"net/http"
	"strconv"
)

type ProcessRulesController struct {
	ProcessRulesService service.ProcessRulesService
}

func NewProcessRulesController(service service.ProcessRulesService) *ProcessRulesController {
	return &ProcessRulesController{
		ProcessRulesService: service,
	}
}
func (controller *ProcessRulesController) Create(ctx *gin.Context) {
	log.Info().Msg("create usr rules")
	createProcessRulesRequest := request.CreateProcessRulesRequest{}
	err := ctx.ShouldBindJSON(&createProcessRulesRequest)
	helper.ErrorPanic(err)

	controller.ProcessRulesService.Create(createProcessRulesRequest)
	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProcessRulesController) Update(ctx *gin.Context) {
	updateProcessRulesRequest := request.UpdateProcessRulesRequest{}
	err := ctx.ShouldBindJSON(&updateProcessRulesRequest)
	helper.ErrorPanic(err)

	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)
	updateProcessRulesRequest.Id = id

	controller.ProcessRulesService.Update(updateProcessRulesRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProcessRulesController) Delete(ctx *gin.Context) {
	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)
	controller.ProcessRulesService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProcessRulesController) FindId(ctx *gin.Context) {
	ruleId := ctx.Param("id")
	id, err := strconv.Atoi(ruleId)
	helper.ErrorPanic(err)

	ruleResponse := controller.ProcessRulesService.FindId(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ruleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *ProcessRulesController) FindAll(ctx *gin.Context) {
	ruleResponse := controller.ProcessRulesService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ruleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
