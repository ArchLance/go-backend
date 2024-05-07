package repository

import "go_backend/model"

type ProcessRulesRepository interface {
	Save(rules model.ProcessRules) (err error)
	Update(rules model.ProcessRules) (err error)
	Delete(id int) (err error)
	FindId(id int) (rules model.ProcessRules, err error)
	FindAll() []model.ProcessRules
}
