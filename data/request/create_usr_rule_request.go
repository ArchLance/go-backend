package request

type CreateUsrRulesRequest struct {
	Path string `validate:"required,min=1,max=255" json:"path"`
	Usr  string `validate:"required,min=1,max=255" json:"usr"`
}
