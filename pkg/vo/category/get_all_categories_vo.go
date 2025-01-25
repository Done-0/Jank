package category

// GetAllCategoriesVo       获取所有类目响应
// @Description 获取所有类目时返回的响应数据
// @Param		categories	body	[]GetOneCategoryVo	true	"类目列表"
type GetAllCategoriesVo struct {
	Categories []*GetOneCategoryVo `json:"categories" xml:"categories" form:"categories" query:"categories"`
}
