package repository

import (
	"errors"
	"go_backend/helper"
	"go_backend/model"
	"gorm.io/gorm"
)

type UsrRulesRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsrRulesRepositoryImpl(Db *gorm.DB) UsrRulesRepository {
	return &UsrRulesRepositoryImpl{Db: Db}
}

func (r *UsrRulesRepositoryImpl) Save(rules model.UsrRules) error {
	result := r.Db.Create(&rules)
	return result.Error
}

func (r *UsrRulesRepositoryImpl) Update(rules model.UsrRules) error {
	result := r.Db.Updates(&rules)
	return result.Error
}

func (r *UsrRulesRepositoryImpl) Delete(id int) error {
	var rules model.UsrRules
	result := r.Db.Where("id = ?", id).Delete(&rules)
	return result.Error
}

func (r *UsrRulesRepositoryImpl) FindId(id int) (rules model.UsrRules, err error) {
	var rule model.UsrRules
	result := r.Db.Find(&rule, id)
	if result != nil {
		return rule, nil
	} else {
		return rule, errors.New("rule is not found")
	}
}

func (r *UsrRulesRepositoryImpl) FindAll() []model.UsrRules {
	var rules []model.UsrRules
	result := r.Db.Find(&rules)
	helper.ErrorPanic(result.Error)
	return rules
}
