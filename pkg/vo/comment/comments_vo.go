package comment

// CommentsVo 获取评论响应
// @Description 获取单个评论的响应
// @Property id                  body int64  			 true  "评论唯一标识"
// @Property content             body string  			 true  "评论内容"
// @Property user_id             body int64             true  "评论所属用户ID"
// @Property post_id             body int64             true  "评论所属文章ID"
// @Property reply_to_comment_id body int64             false "回复的目标评论ID"
// @Property replies             body []CommentsVo true  "子评论列表"
type CommentsVo struct {
	ID               int64         `json:"id" xml:"id" form:"id" query:"id"`
	Content          string        `json:"content" xml:"content" form:"content" query:"content"`
	UserId           int64         `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	PostId           int64         `json:"post_id" xml:"post_id" form:"post_id" query:"post_id"`
	ReplyToCommentId int64         `json:"reply_to_comment_id" xml:"reply_to_comment_id" form:"reply_to_comment_id" query:"reply_to_comment_id"`
	Reply            []*CommentsVo `json:"replies" xml:"replies" form:"replies" query:"replies"`
}
