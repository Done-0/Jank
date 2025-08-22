package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// CategoryService 分类服务接口
type CategoryService interface {
	GetCategory(c *app.RequestContext, req *dto.GetCategoryRequest) (*vo.GetCategoryResponse, error)          // 获取单个分类
	ListCategories(c *app.RequestContext, req *dto.ListCategoriesRequest) (*vo.ListCategoriesResponse, error) // 获取分类列表
	Create(c *app.RequestContext, req *dto.CreateCategoryRequest) (*vo.CreateCategoryResponse, error)         // 创建分类
	Update(c *app.RequestContext, req *dto.UpdateCategoryRequest) (*vo.UpdateCategoryResponse, error)         // 更新分类
	Delete(c *app.RequestContext, req *dto.DeleteCategoryRequest) (*vo.DeleteCategoryResponse, error)         // 删除分类
}
