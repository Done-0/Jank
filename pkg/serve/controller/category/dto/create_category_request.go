package dto

// CreateOneCategoryRequest       创建类目请求
// @Param name        body string true  "类目名称"
// @Param description body string false "类目描述"
// @Param parent_id   body int64  false "父类目ID"
type CreateOneCategoryRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id" validate:"gte=0"`
}
