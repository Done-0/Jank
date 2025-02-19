package category

// CategoriesVo 获取类目响应
// @Description 获取类目响应
// @Property		id			body	int64	true	"类目唯一标识"
// @Property		name		body	string	true	"类目名称"
// @Property		description	body	string	true	"类目描述"
// @Property		parent_id	body	int64	true	"父类目ID"
// @Property		path		body	string	true	"类目路径"
// @Property		children	body	[]CategoriesVo	true	"子类目列表"
type CategoriesVo struct {
	ID          int64           `json:"id" xml:"id" form:"id" query:"id"`
	Name        string          `json:"name" xml:"name" form:"name" query:"name"`
	Description string          `json:"description" xml:"description" form:"description" query:"description"`
	ParentID    int64           `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id"`
	Path        string          `json:"path" xml:"path" form:"path" query:"path"`
	Children    []*CategoriesVo `json:"children" xml:"children" form:"children" query:"children"`
}
