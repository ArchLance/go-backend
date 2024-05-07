package response

type ProcessRuleResponse struct {
	Id      int    `json:"id"`
	Path    string `json:"path"`
	Process string `json:"process"`
}
