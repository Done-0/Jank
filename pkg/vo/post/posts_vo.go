package post

// PostsVo    获取帖子的响应结构
// @Description	获取帖子时返回的响应数据
// @Property			id			    	body	int64	true	"帖子唯一标识"
// @Property			title			    body	string	true	"帖子标题"
// @Property			image			    body	string	true	"帖子封面图片 URL"
// @Property			visibility		    body	string	true	"帖子可见性状态"
// @Property			content_html		body	string	true	"帖子 HTML 格式内容"
// @Property			category_ids	    body	[]int64	true	"帖子所属分类 ID 列表"
type PostsVo struct {
	ID          int64   `json:"id" xml:"id" form:"id" query:"id"`
	Title       string  `json:"title" xml:"title" form:"title" query:"title"`
	Image       string  `json:"image" xml:"image" form:"image" query:"image"`
	Visibility  string  `json:"visibility" xml:"visibility" form:"visibility" query:"visibility"`
	ContentHTML string  `json:"content_html" xml:"content_html" form:"content_html" query:"content_html"`
	CategoryIDs []int64 `json:"category_ids" xml:"category_ids" form:"category_ids" query:"category_ids"`
}
