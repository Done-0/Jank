package dto

import (
	"jank.com/jank_blog/pkg/vo/category"
)

// UpdateOneCategoryRequest    更新类目请求
// @Param id          path     int64   true  "类目ID"
// @Param name        formData string  true  "类目名称"
// @Param description formData string  false "类目描述"
// @Param parent_id   formData int64   false "父类目ID"
// @Param path        formData string  false "类目路径"
// @Param children    formData array   false "子类目"
type UpdateOneCategoryRequest struct {
	ID          int64                        `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Name        string                       `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string                       `json:"description" xml:"description" form:"description" query:"description" validate:"required,min=0,max=255" default:""`
	ParentID    int64                        `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id"`
	Path        string                       `json:"path" xml:"path" form:"path" query:"path"`
	Children    []*category.GetOneCategoryVo `json:"children" xml:"children" form:"children" query:"children"`
}
