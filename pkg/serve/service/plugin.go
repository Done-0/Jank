package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// PluginService 插件服务接口
type PluginService interface {
	RegisterPlugin(c *app.RequestContext, req *dto.RegisterPluginRequest) (*vo.RegisterPluginResponse, error)
	UnregisterPlugin(c *app.RequestContext, req *dto.UnregisterPluginRequest) (*vo.UnregisterPluginResponse, error)
	ExecutePlugin(c *app.RequestContext, req *dto.ExecutePluginRequest) (*vo.ExecutePluginResponse, error)
	GetPlugin(c *app.RequestContext, req *dto.GetPluginRequest) (*vo.GetPluginResponse, error)
	ListPlugins(c *app.RequestContext, req *dto.ListPluginsRequest) (*vo.ListPluginsResponse, error)
}
