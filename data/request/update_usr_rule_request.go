package request

type UpdateUsrRulesRequest struct {
	Id   int
	Path string `validate:"required,min=1,max=255" json:"path"`
	Usr  string `validate:"required,min=1,max=255" json:"usr"`
}
