/**
 * React Query 配置
 */

import { QueryClient } from "@tanstack/react-query";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // 5分钟缓存
      staleTime: 5 * 60 * 1000,
      // 10分钟垃圾回收
      gcTime: 10 * 60 * 1000,
      // 失败重试1次
      retry: 1,
      // 窗口聚焦时重新获取
      refetchOnWindowFocus: false,
    },
    mutations: {
      // 失败重试1次
      retry: 1,
    },
  },
});
