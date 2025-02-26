package dto

// UpdateOneCategoryRequest    更新类目请求
// @Param id          path     int64   true  "类目ID"
// @Param name        body     string  true  "类目名称"
// @Param description body     string  false "类目描述"
// @Param parent_id   body     int64   false "父类目ID"
// @Param path        body     string  false "类目路径"
// @Param children    body     array   false "子类目"
type UpdateOneCategoryRequest struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id" validate:"gte=0"`
}
