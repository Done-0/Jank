/**
 * 插件服务
 */

import { PLUGIN_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  RegisterPluginRequest,
  RegisterPluginResponse,
  UnregisterPluginRequest,
  UnregisterPluginResponse,
  ExecutePluginRequest,
  ExecutePluginResponse,
  GetPluginRequest,
  GetPluginResponse,
  ListPluginsRequest,
  ListPluginsResponse,
} from "@/types";

class PluginService {
  // ===== 插件管理 =====

  // 注册插件
  async registerPlugin(
    request: RegisterPluginRequest
  ): Promise<RegisterPluginResponse> {
    const response = await apiClient.post<ApiResponse<RegisterPluginResponse>>(
      PLUGIN_ENDPOINTS.REGISTER_PLUGIN,
      request
    );
    return response.data.data!;
  }

  // 注销插件
  async unregisterPlugin(
    request: UnregisterPluginRequest
  ): Promise<UnregisterPluginResponse> {
    const response = await apiClient.post<
      ApiResponse<UnregisterPluginResponse>
    >(PLUGIN_ENDPOINTS.UNREGISTER_PLUGIN, request);
    return response.data.data!;
  }

  // 执行插件
  async executePlugin(
    request: ExecutePluginRequest
  ): Promise<ExecutePluginResponse> {
    const response = await apiClient.post<ApiResponse<ExecutePluginResponse>>(
      PLUGIN_ENDPOINTS.EXECUTE_PLUGIN,
      request
    );
    return response.data.data!;
  }

  // 获取插件详情
  async getPlugin(request: GetPluginRequest): Promise<GetPluginResponse> {
    const response = await apiClient.get<ApiResponse<GetPluginResponse>>(
      PLUGIN_ENDPOINTS.GET_PLUGIN,
      { params: request }
    );
    return response.data.data!;
  }

  // 插件列表
  async listPlugins(request: ListPluginsRequest): Promise<ListPluginsResponse> {
    const response = await apiClient.get<ApiResponse<ListPluginsResponse>>(
      PLUGIN_ENDPOINTS.LIST_PLUGINS,
      { params: request }
    );
    return response.data.data!;
  }
}

export const pluginService = new PluginService();
