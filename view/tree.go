package view

type ProductTreeResponse struct {
	Data     string                 `json:"data"`
	Children []*ProductTreeResponse `json:"children"`
	ParentId *int                   `json:"parent_id"`
	Id       *int                   `json:"id"`
	Text     string                 `json:"text"`
	Type     string                 `json:"type"`
}
