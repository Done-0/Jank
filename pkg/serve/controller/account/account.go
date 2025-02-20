package account

import (
	"github.com/labstack/echo/v4"
	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/controller/verification"
	"jank.com/jank_blog/pkg/serve/service"
	"jank.com/jank_blog/pkg/vo"
	"net/http"
)

const (
	LocalsUserIdKey = "Locals_User_Id"
)

// GetAccount godoc
// @Summary      获取账户信息
// @Description  根据提供的邮箱获取对应用户的详细信息
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetAccountRequest  true  "获取账户请求参数"
// @Success      200     {object}   vo.Result{data=account.GetAccountVo}  "获取成功"
// @Failure      400     {object}   vo.Result              "请求参数错误"
// @Failure      404     {object}   vo.Result              "用户不存在"
// @Router       /account/getAccount [post]
func GetAccount(c echo.Context) error {
	req := new(dto.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	response, err := service.GetAccount(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(response, c))
}

// RegisterAcc godoc
// @Summary      用户注册
// @Description  注册新用户账号，支持图形验证码和邮箱验证码校验
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "注册信息"
// @Param        ImgVerificationCode  query   string  true  "图形验证码"
// @Param        EmailVerificationCode  query   string  true  "邮箱验证码"
// @Success      200     {object}   vo.Result{data=dto.RegisterRequest}  "注册成功"
// @Failure      400     {object}   vo.Result         "参数错误，验证码校验失败"
// @Failure      500     {object}   vo.Result         "服务器错误"
// @Router       /account/registerAccount [post]
func RegisterAcc(c echo.Context) error {
	req := new(dto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	if !verification.VerifyImgCode(req.ImgVerificationCode, req.Email, c) {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.SendImgVerificationCodeFail, "图形验证码校验失败"), c))
	}

	if !verification.VerifyEmailCode(req.EmailVerificationCode, req.Email, c) {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.SendEmailVerificationCodeFail, "邮箱验证码校验失败"), c))
	}

	user, err := service.RegisterUser(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(user, c))
}

// LoginAccount godoc
// @Summary      用户登录
// @Description  用户登录并获取访问令牌，支持图形验证码校验
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "登录信息"
// @Param        ImgVerificationCode  query   string  true  "图形验证码"
// @Success      200     {object}   vo.Result{data=account.LoginVo}  "登录成功，返回访问令牌"
// @Failure      400     {object}   vo.Result         "参数错误，验证码校验失败"
// @Failure      401     {object}   vo.Result         "登录失败，凭证无效"
// @Router       /account/loginAccount [post]
func LoginAccount(c echo.Context) error {
	req := new(dto.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	if !verification.VerifyImgCode(req.ImgVerificationCode, req.Email, c) {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.SendImgVerificationCodeFail, "图形验证码校验失败"), c))
	}

	response, err := service.LoginUser(req, c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, vo.Fail(err, bizErr.New(bizErr.UnKnowErr, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(response, c))
}

// LogoutAccount godoc
// @Summary      用户登出
// @Description  退出当前用户登录状态
// @Tags         账户
// @Produce      json
// @Success      200  {object}  vo.Result{data=string}  "登出成功"
// @Failure      401  {object}  vo.Result  "未授权"
// @Failure      500  {object}  vo.Result  "服务器错误"
// @Security     BearerAuth
// @Router       /account/logoutAccount [post]
func LogoutAccount(c echo.Context) error {
	userId, ok := c.Get(LocalsUserIdKey).(int64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, vo.Fail(nil, bizErr.New(bizErr.UnKnowErr, "用户未登录"), c))
	}

	if err := service.LogoutUser(userId, c); err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("用户注销成功", c))
}

// ResetPassword godoc
// @Summary      重置密码
// @Description  重置用户账户密码，支持邮箱验证码校验
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        request  body      dto.ResetPwdRequest  true  "重置密码信息"
// @Param        EmailVerificationCode  query   string  true  "邮箱验证码"
// @Success      200     {object}   vo.Result{data=string}  "密码重置成功"
// @Failure      400     {object}   vo.Result         "参数错误，验证码校验失败"
// @Failure      401     {object}   vo.Result         "未授权，用户未登录"
// @Failure      500     {object}   vo.Result         "服务器错误"
// @security     BearerAuth
// @Router       /account/resetPassword [post]
func ResetPassword(c echo.Context) error {
	accountID, ok := c.Get(LocalsUserIdKey).(int64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, vo.Fail(nil, bizErr.New(bizErr.UnKnowErr, "用户未登录"), c))
	}

	req := new(dto.ResetPwdRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	if !verification.VerifyEmailCode(req.EmailVerificationCode, req.Email, c) {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.SendEmailVerificationCodeFail, "邮箱验证码校验失败"), c))
	}

	err := service.ResetPassword(req, accountID, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("密码重置成功", c))
}

// CreateRole 创建角色
// @Summary      创建角色
// @Description  创建一个新的角色，角色信息包括代码和描述
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateRoleRequest  true  "角色信息"
// @Success      200     {object}  vo.Result{data=account.RoleVo}  "创建成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /role/createOneRole [post]
func CreateRole(c echo.Context) error {
	req := new(dto.CreateRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	role, err := service.CreateRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(role, c))
}

// UpdateRole 更新角色
// @Summary      更新角色
// @Description  更新角色的代码和描述
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.UpdateRoleRequest  true  "角色信息"
// @Success      200     {object}  vo.Result{data=account.RoleVo}  "更新成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /role/updateOneRole [post]
func UpdateRole(c echo.Context) error {
	req := new(dto.UpdateRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	role, err := service.UpdateRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(role, c))
}

// DeleteRole 删除角色
// @Summary      删除角色
// @Description  根据角色ID删除角色
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeleteRoleRequest  true  "角色ID"
// @Success      200     {object}  vo.Result{data=string}  "删除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role/deleteOneRole [post]
func DeleteRole(c echo.Context) error {
	req := new(dto.DeleteRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.DeleteRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("角色删除成功", c))
}

// ListRoles 获取所有角色
// @Summary      获取所有角色
// @Description  获取系统中所有角色的信息
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Success      200     {object}  vo.Result{data=[]account.RoleVo}  "获取成功"
// @Failure      500     {object}  vo.Result{message=string}     "服务器错误"
// @Router       /role/listAllRoles [post]
func ListRoles(c echo.Context) error {
	roles, err := service.ListRoles(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}
	return c.JSON(http.StatusOK, vo.Success(roles, c))
}

// CreatePermission 创建权限
// @Summary      创建权限
// @Description  创建新的权限，权限信息包括代码和描述
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreatePermissionRequest  true  "权限信息"
// @Success      200     {object}  vo.Result{data=account.PermissionVo}  "创建成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /permission/createOnePermission [post]
func CreatePermission(c echo.Context) error {
	req := new(dto.CreatePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	permission, err := service.CreatePermission(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(permission, c))
}

// UpdatePermission 更新权限
// @Summary      更新权限
// @Description  更新权限的代码和描述
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.UpdatePermissionRequest  true  "权限信息"
// @Success      200     {object}  vo.Result{data=account.PermissionVo}  "更新成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /permission/updateOnePermission [post]
func UpdatePermission(c echo.Context) error {
	req := new(dto.UpdatePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	permission, err := service.UpdatePermission(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(permission, c))
}

// DeletePermission 删除权限
// @Summary      删除权限
// @Description  根据权限ID删除权限
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeletePermissionRequest  true  "权限ID"
// @Success      200     {object}  vo.Result{data=string}  "删除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /permission/deleteOnePermission [post]
func DeletePermission(c echo.Context) error {
	req := new(dto.DeletePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.DeletePermission(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("权限删除成功", c))
}

// ListPermissions 获取所有权限
// @Summary      获取所有权限
// @Description  获取系统中所有权限的信息
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Success      200     {object}  vo.Result{data=[]account.PermissionVo}  "获取成功"
// @Failure      500     {object}  vo.Result{message=string}     "服务器错误"
// @Router       /permission/listAllPermissions [post]
func ListPermissions(c echo.Context) error {
	permissions, err := service.ListPermissions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}
	return c.JSON(http.StatusOK, vo.Success(permissions, c))
}

// AssignRoleToAcc 为用户分配角色
// @Summary      为用户分配角色
// @Description  根据用户ID和角色ID为用户分配角色
// @Tags         用户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignRoleRequest  true  "分配角色信息"
// @Success      200     {object}  vo.Result{data=string}  "角色分配成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /acc-role/assignRoleToAcc [post]
func AssignRoleToAcc(c echo.Context) error {
	req := new(dto.AssignRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.AssignRoleToAcc(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("角色分配成功", c))
}

// AssignPermissionToRole 为角色分配权限
// @Summary      为角色分配权限
// @Description  根据角色ID和权限ID为角色分配权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignPermissionRequest  true  "分配权限信息"
// @Success      200     {object}  vo.Result{data=string}  "权限分配成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/assignPermissionToRole [post]
func AssignPermissionToRole(c echo.Context) error {
	req := new(dto.AssignPermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.AssignPermissionToRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("权限分配成功", c))
}

// DeleteRoleFromAcc 移除用户角色
// @Summary      移除用户角色
// @Description  根据用户ID和角色ID移除用户的角色
// @Tags         用户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignRoleRequest  true  "移除角色信息"
// @Success      200     {object}  vo.Result{data=string}  "角色移除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /acc-role/deleteRoleFromAcc [post]
func DeleteRoleFromAcc(c echo.Context) error {
	req := new(dto.AssignRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.RemoveRoleFromAcc(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("角色移除成功", c))
}

// DeletePermissionFromRole 移除角色权限
// @Summary      移除角色权限
// @Description  根据角色ID和权限ID移除角色的权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignPermissionRequest  true  "移除权限信息"
// @Success      200     {object}  vo.Result{data=string}  "权限移除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/deletePermissionFromRole [post]
func DeletePermissionFromRole(c echo.Context) error {
	req := new(dto.AssignPermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.RemovePermissionFromRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("权限移除成功", c))
}

// UpdateRoleForAcc 更新用户角色
// @Summary      更新用户角色
// @Description  根据用户ID和角色ID更新用户角色
// @Tags         用户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignRoleRequest  true  "更新角色信息"
// @Success      200     {object}  vo.Result{data=string}  "角色更新成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /acc-role/updateRoleForAcc [post]
func UpdateRoleForAcc(c echo.Context) error {
	req := new(dto.AssignRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.UpdateRoleForAcc(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("用户角色更新成功", c))
}

// UpdatePermissionForRole 更新角色权限
// @Summary      更新角色权限
// @Description  根据角色ID和权限ID更新角色权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignPermissionRequest  true  "更新权限信息"
// @Success      200     {object}  vo.Result{data=string}  "权限更新成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/updatePermissionForRole [post]
func UpdatePermissionForRole(c echo.Context) error {
	req := new(dto.AssignPermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	err := service.UpdatePermissionForRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("角色权限更新成功", c))
}

// GetRolesByAcc 获取用户的角色
// @Summary      获取用户角色
// @Description  根据用户ID获取用户的所有角色
// @Tags         用户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetRolesByAccRequest  true  "获取用户角色信息"
// @Success      200     {object}  vo.Result{data=[]account.RoleVo}  "角色列表"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /acc-role/getRolesByAcc [POST]
func GetRolesByAcc(c echo.Context) error {
	req := new(dto.GetRolesByAccRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	roles, err := service.GetRolesByAcc(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(roles, c))
}

// GetPermissionsByRole 获取角色的权限
// @Summary      获取角色权限
// @Description  根据角色ID获取角色的所有权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetPermissionsByRoleRequest  true  "获取角色权限信息"
// @Success      200     {object}  vo.Result{data=[]account.PermissionVo}  "权限列表"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/getPermissionsByRole [POST]
func GetPermissionsByRole(c echo.Context) error {
	req := new(dto.GetPermissionsByRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BadRequest, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BadRequest, "请求参数校验失败"), c))
	}

	permissions, err := service.GetPermissionsByRole(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(permissions, c))
}
