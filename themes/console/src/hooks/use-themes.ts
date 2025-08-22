/**
 * 主题管理相关 Query Hooks
 */

import {
  useQuery,
  useMutation,
  useQueryClient,
  type UseQueryOptions,
  type UseMutationOptions,
} from "@tanstack/react-query";
import { themeService } from "@/services/theme.service";
import type {
  SwitchThemeRequest,
  SwitchThemeResponse,
  ListThemesRequest,
  ListThemesResponse,
  GetActiveThemeResponse,
} from "@/types";

// Query Keys
export const themeKeys = {
  all: ["themes"] as const,
  lists: () => [...themeKeys.all, "list"] as const,
  list: (params: ListThemesRequest) => [...themeKeys.lists(), params] as const,
  active: () => [...themeKeys.all, "active"] as const,
} as const;

// 获取主题列表
export function useThemes(
  params: ListThemesRequest,
  options?: Omit<UseQueryOptions<ListThemesResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: themeKeys.list(params),
    queryFn: () => themeService.listThemes(params),
    staleTime: 5 * 60 * 1000, // 5分钟
    ...options,
  });
}

// 获取当前激活主题
export function useActiveTheme(
  options?: Omit<
    UseQueryOptions<GetActiveThemeResponse>,
    "queryKey" | "queryFn"
  >
) {
  return useQuery({
    queryKey: themeKeys.active(),
    queryFn: () => themeService.getActiveTheme(),
    staleTime: 10 * 60 * 1000, // 10分钟
    ...options,
  });
}

// 切换主题
export function useSwitchTheme(
  options?: Omit<
    UseMutationOptions<SwitchThemeResponse, Error, SwitchThemeRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: SwitchThemeRequest) =>
      themeService.switchTheme(request),
    onSuccess: (data) => {
      console.log("主题切换成功，后端响应:", data);
      // 刷新相关查询
      queryClient.invalidateQueries({ queryKey: themeKeys.all });
    },
    onError: (error: any) => {
      console.error("主题切换失败，错误详情:", error);
    },
    ...options,
  });
}
