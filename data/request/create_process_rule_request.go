package request

type CreateProcessRulesRequest struct {
	Path    string `validate:"required,max=255,min=1" json:"path"`
	Process string `validate:"required,max=255,min=1" json:"process"`
}
