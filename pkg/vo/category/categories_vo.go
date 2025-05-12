// Package category 提供类目相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package category

// CategoriesVO 获取类目响应
// @Description 获取类目响应
// @Property		id			body	int64	true	"类目唯一标识"
// @Property		name		body	string	true	"类目名称"
// @Property		description	body	string	true	"类目描述"
// @Property		parent_id	body	int64	true	"父类目ID"
// @Property		path		body	string	true	"类目路径"
// @Property		children	body	[]*CategoriesVO	true	"子类目列表"
type CategoriesVO struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ParentID    int64           `json:"parent_id"`
	Path        string          `json:"path"`
	Children    []*CategoriesVO `json:"children"`
}
