package dto

// GetOnePostRequest        获取文章的请求结构体
//
//	@Param	id		path	string	true	"文章 ID"
//	@Param	title	query	string	false	"文章标题"
type GetOnePostRequest struct {
	ID    int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Title string `json:"title" xml:"title" form:"title" query:"title" validate:"min=1,max=225"`
}
