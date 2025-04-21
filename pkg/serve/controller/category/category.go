package category

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/category/dto"
	service "jank.com/jank_blog/pkg/serve/service/category"
	"jank.com/jank_blog/pkg/vo"
)

// GetOneCategory godoc
// @Summary      获取单个类目详情
// @Description  根据类目 ID 获取单个类目的详细信息
// @Tags         类目
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "类目ID"
// @Success      200   {object} vo.Result{data=category.CategoriesVO}  "获取成功"
// @Failure      400   {object} vo.Result  "请求参数错误"
// @Failure      404   {object} vo.Result  "类目不存在"
// @Router       /category/getOneCategory [get]
func GetOneCategory(c echo.Context) error {
	req := new(dto.GetOneCategoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	category, err := service.GetCategoryByID(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(category, c))
}

// GetCategoryTree godoc
// @Summary      获取类目树
// @Description  获取类目树
// @Tags         类目
// @Accept       json
// @Produce      json
// @Success      200  {object}  vo.Result{data=[]category.CategoriesVO}  "获取成功"
// @Failure      500  {object}  vo.Result                 "服务器错误"
// @Router       /category/getCategoryTree [get]
func GetCategoryTree(c echo.Context) error {
	categories, err := service.GetCategoryTree(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(categories, c))
}

// GetCategoryChildrenTree godoc
// @Summary      获取子类目树
// @Description  根据类目 ID 获取子类目树
// @Tags         类目
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "类目ID"
// @Success      200   {object} vo.Result{data=[]category.CategoriesVO}  "获取成功"
// @Failure      400   {object} vo.Result  "请求参数错误"
// @Failure      404   {object} vo.Result  "类目不存在"
// @Failure      500   {object} vo.Result  "服务器错误"
// @Router       /category/getCategoryChildrenTree [post]
func GetCategoryChildrenTree(c echo.Context) error {
	req := new(dto.GetOneCategoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	childrenCategories, err := service.GetCategoryChildrenByID(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(childrenCategories, c))
}

// CreateOneCategory     godoc
// @Summary      创建类目
// @Description  创建新的类目
// @Tags         类目
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOneCategoryRequest  true  "创建类目请求参数"
// @Success      200     {object}   vo.Result{data=category.CategoriesVO}  "创建成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Security     BearerAuth
// @Router       /category/createOneCategory [post]
func CreateOneCategory(c echo.Context) error {
	req := new(dto.CreateOneCategoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	createdCategory, err := service.CreateCategory(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(createdCategory, c))
}

// UpdateOneCategory     godoc
// @Summary      更新类目
// @Description  更新已存在的类目信息
// @Tags         类目
// @Accept       json
// @Produce      json
// @Param        id       path      int                       true  "类目ID"
// @Param        request  body      dto.UpdateOneCategoryRequest true  "更新类目请求参数"
// @Success      200     {object}   vo.Result{data=category.CategoriesVO}  "更新成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "类目不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /category/updateOneCategory [post]
func UpdateOneCategory(c echo.Context) error {
	req := new(dto.UpdateOneCategoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	updatedCategory, err := service.UpdateCategory(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(updatedCategory, c))
}

// DeleteOneCategory   godoc
// @Summary      删除类目
// @Description  根据类目 ID 删除类目
// @Tags         类目
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "类目ID"
// @Success      200   {object} vo.Result{data=category.CategoriesVO}  "删除成功"
// @Failure      400   {object} vo.Result  "请求参数错误"
// @Failure      404   {object} vo.Result  "类目不存在"
// @Failure      500   {object} vo.Result  "服务器错误"
// @Security     BearerAuth
// @Router       /category/deleteOneCategory [post]
func DeleteOneCategory(c echo.Context) error {
	req := new(dto.DeleteOneCategoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	category, err := service.DeleteCategory(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(category, c))
}
