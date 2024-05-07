package repository

import "go_backend/model"

type UsrRulesRepository interface {
	Save(rules model.UsrRules) error
	Update(rules model.UsrRules) error
	Delete(id int) error
	FindId(id int) (rules model.UsrRules, err error)
	FindAll() []model.UsrRules
}
