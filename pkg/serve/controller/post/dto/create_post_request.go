package dto

// CreateOnePostRequest         发布文章的请求结构体
// @Param	title				body	string	true	"文章标题"
// @Param	image				body	string	true	"文章图片(可选)"
// @Param	visibility			body	string	true	"文章可见性(可选,默认 private)"
// @Param	content_html	    body	string	true	"文章内容(markdown格式)"
// @Param	category_ids		body	[]int64	true	"文章分类ID列表"
type CreateOnePostRequest struct {
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"required,min=1,max=225"`
	Image           string `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      string `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" default:"private"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown" default:""`
	CategoryIDs     string `json:"category_ids" xml:"category_ids" form:"category_ids" query:"category_ids"`
}
