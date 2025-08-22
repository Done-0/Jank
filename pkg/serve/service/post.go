package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// PostService 文章服务接口
type PostService interface {
	GetPost(c *app.RequestContext, req *dto.GetPostRequest) (*vo.GetPostResponse, error)                         // 获取单篇文章
	ListPublishedPosts(c *app.RequestContext, req *dto.ListPublishedPostsRequest) (*vo.ListPostsResponse, error) // 获取已发布文章列表
	ListPostsByStatus(c *app.RequestContext, req *dto.ListPostsByStatusRequest) (*vo.ListPostsResponse, error)   // 根据状态获取文章列表，支持管理员查询所有文章
	Create(c *app.RequestContext, req *dto.CreatePostRequest) (*vo.CreatePostResponse, error)                    // 创建文章
	Update(c *app.RequestContext, req *dto.UpdatePostRequest) (*vo.UpdatePostResponse, error)                    // 更新文章
	Delete(c *app.RequestContext, req *dto.DeletePostRequest) (*vo.DeletePostResponse, error)                    // 删除文章
}
