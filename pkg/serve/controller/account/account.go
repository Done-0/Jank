package account

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/controller/verification"
	"jank.com/jank_blog/pkg/serve/service/account"
	"jank.com/jank_blog/pkg/vo"
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
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
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
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
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
		return c.JSON(http.StatusUnauthorized, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
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
	if err := service.LogoutUser(c); err != nil {
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
// @Success      200     {object}   vo.Result{data=string}  "密码重置成功"
// @Failure      400     {object}   vo.Result         "参数错误，验证码校验失败"
// @Failure      401     {object}   vo.Result         "未授权，用户未登录"
// @Failure      500     {object}   vo.Result         "服务器错误"
// @security     BearerAuth
// @Router       /account/resetPassword [post]
func ResetPassword(c echo.Context) error {
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

	err := service.ResetPassword(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.ServerError, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("密码重置成功", c))
}
