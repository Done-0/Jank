package comment

// GetAllCommentsVo 获取所有评论响应
// @Description 获取所有评论时返回的响应数据
// @Param comments body []GetOneCommentVo true "评论列表"
type GetAllCommentsVo struct {
	Comments []*GetOneCommentVo `json:"comments" xml:"comments" form:"comments" query:"comments"`
}
