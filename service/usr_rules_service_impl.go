package service

import (
	"github.com/go-playground/validator/v10"
	"go_backend/data/request"
	"go_backend/data/response"
	"go_backend/helper"
	"go_backend/model"
	"go_backend/repository"
)

type UsrRulesServiceImpl struct {
	UsrRulesRepository repository.UsrRulesRepository
	Validate           *validator.Validate
}

func NewUsrRulesServiceImpl(ruleRepository repository.UsrRulesRepository, validate *validator.Validate) UsrRulesService {
	return &UsrRulesServiceImpl{
		UsrRulesRepository: ruleRepository,
		Validate:           validate,
	}
}

func (r *UsrRulesServiceImpl) Create(rules request.CreateUsrRulesRequest) {
	err := r.Validate.Struct(rules)
	helper.ErrorPanic(err)
	ruleModel := model.UsrRules{
		Path: rules.Path,
		Usr:  rules.Usr,
	}
	r.UsrRulesRepository.Save(ruleModel)
}

func (r *UsrRulesServiceImpl) Update(rules request.UpdateUsrRulesRequest) {
	// 先看是否存在
	ruleData, err := r.UsrRulesRepository.FindId(rules.Id)
	helper.ErrorPanic(err)
	// 存在则更新
	ruleData.Path = rules.Path
	ruleData.Usr = rules.Usr
	r.UsrRulesRepository.Update(ruleData)
}

func (r *UsrRulesServiceImpl) Delete(id int) {
	r.UsrRulesRepository.Delete(id)
}

func (r *UsrRulesServiceImpl) FindId(id int) response.UsrRuleResponse {
	ruleData, err := r.UsrRulesRepository.FindId(id)
	helper.ErrorPanic(err)
	ruleResponse := response.UsrRuleResponse{
		Id:   ruleData.Id,
		Path: ruleData.Path,
		Usr:  ruleData.Usr,
	}
	return ruleResponse
}

func (r *UsrRulesServiceImpl) FindAll() []response.UsrRuleResponse {
	result := r.UsrRulesRepository.FindAll()

	var rules []response.UsrRuleResponse
	for _, value := range result {
		rule := response.UsrRuleResponse{
			Id:   value.Id,
			Path: value.Path,
			Usr:  value.Usr,
		}
		rules = append(rules, rule)
	}
	return rules
}
