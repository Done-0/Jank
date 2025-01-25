package post

// GetAllPostsVo    获取所有帖子的响应结构
//
//	@Description	获取所有帖子时返回的响应数据
//	@Param			id				body	int64	true	"帖子唯一标识"
//	@Param			title			body	string	true	"帖子标题"
//	@Param			image			body	string	true	"帖子封面图片 URL"
//	@Param			visibility		body	string	true	"帖子可见性状态"
//	@Param			contentMarkdown	body	string	true	"帖子 Markdown 格式内容"
//	@Param			contentHtml		body	string	true	"帖子 HTML 格式内容"
//	@Param			category_ids	body	[]int64	true	"帖子所属分类 ID 列表"
type GetAllPostsVo struct {
	ID              int64   `json:"id" xml:"id" form:"id" query:"id"`
	Title           string  `json:"title" xml:"title" form:"title" query:"title"`
	Image           string  `json:"image" xml:"image" form:"image" query:"image"`
	Visibility      string  `json:"visibility" xml:"visibility" form:"visibility" query:"visibility"`
	ContentMarkdown string  `json:"contentMarkdown" xml:"contentMarkdown" form:"contentMarkdown" query:"contentMarkdown"`
	ContentHTML     string  `json:"contentHtml" xml:"contentHtml" form:"contentHtml" query:"contentHtml"`
	CategoryIDs     []int64 `json:"category_ids" xml:"category_ids" form:"category_ids" query:"category_ids"`
}
