package dto

// UpdatePostRequest 更新文章请求参数结构体
// @Param   id   			  body    int	    	true      "文章 ID"
// @Param   title		      body 	  string        false	  "文章标题"
// @Param   image			  body 	  string        false     "文章图片(可选)"
// @Param   contentMarkdown	  body    string 		false     "文章内容(markdown格式)"
type UpdatePostRequest struct {
	ID              int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Title           string `json:"title" xml:"title" form:"title" query:"title"`
	Image           string `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      string `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" default:"private"`
	ContentMarkdown string `json:"contentMarkdown" xml:"contentMarkdown" form:"contentMarkdown" query:"contentMarkdown"`
}
