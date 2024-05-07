package service

import (
	"go_backend/data/request"
	"go_backend/data/response"
)

type UsrRulesService interface {
	Create(rules request.CreateUsrRulesRequest)
	Update(rules request.UpdateUsrRulesRequest)
	Delete(id int)
	FindId(id int) response.UsrRuleResponse
	FindAll() []response.UsrRuleResponse
}
