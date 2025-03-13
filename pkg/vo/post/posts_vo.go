package post

// PostsVO    获取帖子的响应结构
// @Description	获取帖子时返回的响应数据
// @Property			id			    	body	int64	true	"帖子唯一标识"
// @Property			title			    body	string	true	"帖子标题"
// @Property			image			    body	string	true	"帖子封面图片 URL"
// @Property			visibility		    body	bool	true	"帖子可见性状态"
// @Property			content_html		body	string	true	"帖子 HTML 格式内容"
// @Property			category_ids	    body	[]int64	true	"帖子所属分类 ID 列表"
type PostsVO struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Visibility bool   `json:"visibility"`
	//ContentMarkdown string  `json:"content_markdown"`
	ContentHTML string  `json:"content_html"`
	CategoryIDs []int64 `json:"category_ids"`
}
