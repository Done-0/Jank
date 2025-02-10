package dto

// CreateOnePostRequest         发布文章的请求结构体
//
//	@Param	title				body	string	true	"文章标题"
//	@Param	image				body	string	true	"文章图片(可选)"
//	@Param	visibility			body	string	true	"文章可见性(可选,默认 private)"
//	@Param	contentMarkdown	    body	string	true	"文章内容(markdown格式)"
type CreateOnePostRequest struct {
	Title           string  `json:"title" xml:"title" form:"title" query:"title" validate:"required,min=1,max=225"`
	Image           string  `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      string  `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" default:"private"`
	ContentMarkdown string  `json:"contentMarkdown" xml:"contentMarkdown" form:"contentMarkdown" query:"contentMarkdown" validate:"required"`
	CategoryIDs     []int64 `json:"categoryIds" xml:"categoryIds" form:"categoryIds" query:"categoryIds" default:""`
}
