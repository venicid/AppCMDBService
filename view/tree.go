package view

type GlobalTreeResponse struct {
	Data     map[string]int        `json:"data"`
	Children []*GlobalTreeResponse `json:"children"`
	ParentId *int                  `json:"parent_id"`
	Id       *int                  `json:"id"`
	Text     string                `json:"text"`
	Type     string                `json:"type"`
}
