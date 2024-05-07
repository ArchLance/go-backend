package service

import (
	"github.com/go-playground/validator/v10"
	"go_backend/data/request"
	"go_backend/data/response"
	"go_backend/helper"
	"go_backend/model"
	"go_backend/repository"
)

type ProcessRulesServiceImpl struct {
	ProcessRulesRepository repository.ProcessRulesRepository
	Validate               *validator.Validate
}

func NewProcessRulesServiceImpl(ruleRepository repository.ProcessRulesRepository, validate *validator.Validate) ProcessRulesService {
	return &ProcessRulesServiceImpl{
		ProcessRulesRepository: ruleRepository,
		Validate:               validate,
	}
}

func (r *ProcessRulesServiceImpl) Create(rules request.CreateProcessRulesRequest) {
	err := r.Validate.Struct(rules)
	helper.ErrorPanic(err)
	ruleModel := model.ProcessRules{
		Path:    rules.Path,
		Process: rules.Process,
	}
	r.ProcessRulesRepository.Save(ruleModel)
}

func (r *ProcessRulesServiceImpl) Update(rules request.UpdateProcessRulesRequest) {
	// 先看是否存在
	ruleData, err := r.ProcessRulesRepository.FindId(rules.Id)
	helper.ErrorPanic(err)
	// 存在则更新
	ruleData.Path = rules.Path
	ruleData.Process = rules.Process
	r.ProcessRulesRepository.Update(ruleData)
}

func (r *ProcessRulesServiceImpl) Delete(id int) {
	r.ProcessRulesRepository.Delete(id)
}

func (r *ProcessRulesServiceImpl) FindId(id int) response.ProcessRuleResponse {
	ruleData, err := r.ProcessRulesRepository.FindId(id)
	helper.ErrorPanic(err)
	ruleResponse := response.ProcessRuleResponse{
		Id:      ruleData.Id,
		Path:    ruleData.Path,
		Process: ruleData.Process,
	}
	return ruleResponse
}

func (r *ProcessRulesServiceImpl) FindAll() []response.ProcessRuleResponse {
	result := r.ProcessRulesRepository.FindAll()

	var rules []response.ProcessRuleResponse
	for _, value := range result {
		rule := response.ProcessRuleResponse{
			Id:      value.Id,
			Path:    value.Path,
			Process: value.Process,
		}
		rules = append(rules, rule)
	}
	return rules
}
