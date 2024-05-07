package request

type UpdateProcessRulesRequest struct {
	Id      int
	Path    string `validate:"required,max=255,min=1" json:"path"`
	Process string `validate:"required,max=255,min=1" json:"process"`
}
