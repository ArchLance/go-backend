package service

import (
	"go_backend/data/request"
	"go_backend/data/response"
)

type ProcessRulesService interface {
	Create(rules request.CreateProcessRulesRequest)
	Update(rules request.UpdateProcessRulesRequest)
	Delete(id int)
	FindId(id int) response.ProcessRuleResponse
	FindAll() []response.ProcessRuleResponse
}
