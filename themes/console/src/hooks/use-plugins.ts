/**
 * 插件相关 React Query Hooks
 */

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { pluginService } from "@/services/plugin.service";
import type {
  ListPluginsRequest,
  RegisterPluginRequest,
  UnregisterPluginRequest,
  ExecutePluginRequest,
} from "@/types";

// ===== Query Keys =====
export const pluginKeys = {
  all: ["plugins"] as const,
  lists: () => [...pluginKeys.all, "list"] as const,
  list: (params: ListPluginsRequest) =>
    [...pluginKeys.lists(), params] as const,
  details: () => [...pluginKeys.all, "detail"] as const,
  detail: (id: string) => [...pluginKeys.details(), id] as const,
};

// ===== Query Hooks =====

/**
 * 获取插件列表
 */
export function usePlugins(params: ListPluginsRequest) {
  return useQuery({
    queryKey: pluginKeys.list(params),
    queryFn: () => pluginService.listPlugins(params),
  });
}

/**
 * 获取可安装插件列表
 * 显示状态为 'available' 或 'source_only' 的插件（未安装的插件）
 */
export function useAvailablePlugins(
  params: Omit<ListPluginsRequest, "status">
) {
  return useQuery({
    queryKey: [...pluginKeys.lists(), "available", params],
    queryFn: async () => {
      // 获取所有插件
      const allPlugins = await pluginService.listPlugins(params);

      // 过滤出未安装的插件（可以安装的插件）
      const availablePlugins =
        allPlugins.list?.filter(
          (plugin) =>
            plugin.status === "available" || plugin.status === "source_only"
        ) || [];

      return {
        ...allPlugins,
        list: availablePlugins,
        total: availablePlugins.length,
      };
    },
  });
}

/**
 * 获取插件详情
 */
export function usePlugin(id: string) {
  return useQuery({
    queryKey: pluginKeys.detail(id),
    queryFn: () => pluginService.getPlugin({ id }),
    enabled: !!id,
  });
}

// ===== Mutation Hooks =====

/**
 * 安装插件（注册插件）
 * 将未安装的插件安装到系统中
 */
export function useInstallPlugin() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: RegisterPluginRequest) =>
      pluginService.registerPlugin(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: pluginKeys.lists() });
    },
  });
}

/**
 * 卸载插件（注销插件）
 * 将已安装的插件从系统中移除
 */
export function useUninstallPlugin() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UnregisterPluginRequest) =>
      pluginService.unregisterPlugin(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: pluginKeys.lists() });
    },
  });
}

/**
 * 执行插件
 */
export function useExecutePlugin() {
  return useMutation({
    mutationFn: (data: ExecutePluginRequest) =>
      pluginService.executePlugin(data),
  });
}
