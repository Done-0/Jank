//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/Done-0/jank/pkg/serve/controller"
)

// NewPluginController 使用 Wire 初始化插件控制器
func NewPluginController() (*controller.PluginController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewPluginController,
	))
}

// NewRBACController 使用 Wire 初始化RBAC控制器
func NewRBACController() (*controller.RBACController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewRBACController,
	))
}

// NewThemeController 使用 Wire 初始化主题控制器
func NewThemeController() (*controller.ThemeController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewThemeController,
	))
}

// NewUserController 使用 Wire 初始化用户控制器
func NewUserController() (*controller.UserController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewUserController,
	))
}

// NewVerificationController 使用 Wire 初始化验证码控制器
func NewVerificationController() (*controller.VerificationController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewVerificationController,
	))
}

// NewPostController 使用 Wire 初始化文章控制器
func NewPostController() (*controller.PostController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewPostController,
	))
}

// NewCategoryController 使用 Wire 初始化分类控制器
func NewCategoryController() (*controller.CategoryController, error) {
	panic(wire.Build(
		AllProviderSet,
		controller.NewCategoryController,
	))
}
