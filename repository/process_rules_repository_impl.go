package repository

import (
	"errors"
	"go_backend/helper"
	"go_backend/model"
	"gorm.io/gorm"
)

type ProcessRulesRepositoryImpl struct {
	Db *gorm.DB
}

func NewProcessRulesRepositoryImpl(Db *gorm.DB) ProcessRulesRepository {
	return &ProcessRulesRepositoryImpl{Db: Db}
}
func (r *ProcessRulesRepositoryImpl) Save(rules model.ProcessRules) (err error) {
	result := r.Db.Create(&rules)
	return result.Error
}

func (r *ProcessRulesRepositoryImpl) Update(rules model.ProcessRules) (err error) {
	result := r.Db.Updates(&rules)
	return result.Error
}

func (r *ProcessRulesRepositoryImpl) Delete(id int) (err error) {
	var rules model.ProcessRules
	result := r.Db.Where("id = ?", id).Delete(&rules)
	return result.Error
}

func (r *ProcessRulesRepositoryImpl) FindId(id int) (rules model.ProcessRules, err error) {
	var rule model.ProcessRules
	result := r.Db.Find(&rule, id)
	if result != nil {
		return rule, nil
	} else {
		return rule, errors.New("rule is not found")
	}
}

func (r *ProcessRulesRepositoryImpl) FindAll() []model.ProcessRules {
	var rules []model.ProcessRules
	result := r.Db.Find(&rules)
	helper.ErrorPanic(result.Error)
	return rules
}
