package dto

// UpdateOnePostRequest       更新文章请求参数结构体
// @Param   id   			  body    int	    	true      "文章 ID"
// @Param   title		      body    string        false	  "文章标题"
// @Param   image 		      body 	  string        false     "文章图片(可选)"
// @Param   visibility 	      body 	  string        false     "文章可见性(可选)"
// @Param   content_markdown  body    string 		false     "文章内容(markdown格式)"
// @Param   category_ids 	  body    interface{}       false     "文章分类ID列表(可选)"
type UpdateOnePostRequest struct {
	ID              int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"min=0,max=255" default:""`
	Image           string `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      string `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" default:"private"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown" default:""`
	CategoryIDs     string `json:"category_ids" xml:"category_ids" form:"category_ids" query:"category_ids" default:""`
}
