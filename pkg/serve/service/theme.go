package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// ThemeService 主题服务接口
type ThemeService interface {
	// SwitchTheme 切换主题（需要指定主题类型以支持跨类型管理）
	SwitchTheme(c *app.RequestContext, req *dto.SwitchThemeRequest) (*vo.SwitchThemeResponse, error)
	// GetActiveTheme 获取当前激活的主题（根据路由前缀自动判断主题类型）
	GetActiveTheme(c *app.RequestContext) (*vo.GetActiveThemeResponse, error)
	// ListThemes 列举所有主题
	ListThemes(c *app.RequestContext, req *dto.ListThemesRequest) (*vo.ListThemesResponse, error)

	// 路由服务方法（根据路径自动识别主题类型）
	ServeHomePage(c *app.RequestContext) (string, error)
	ServeStaticResource(c *app.RequestContext, requestPath string) (string, error)
}
