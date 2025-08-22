import { Toaster } from "sonner";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider } from "@tanstack/react-router";

import { ThemeProvider } from "@/components/theme/theme-provider";
import { createAppRouter } from "@/router";
import "@/styles/globals.css";

// 创建 QueryClient 实例
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1, // 失败重试次数
      refetchOnWindowFocus: false, // 窗口聚焦时不重新获取
      staleTime: 5 * 60 * 1000, // 数据过期时间：5分钟
    },
    mutations: {
      retry: 1, // 变更失败重试次数
    },
  },
});

// 创建路由器实例
const router = createAppRouter(queryClient);

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="jank-theme">
      <QueryClientProvider client={queryClient}>
        {/* 主应用容器 */}
        <div className="min-h-screen w-full overflow-x-hidden bg-background font-sans antialiased">
          <RouterProvider router={router} />
        </div>

        {/* 全局通知组件 */}
        <Toaster
          position="top-right"
          richColors
          closeButton
          duration={4000}
          toastOptions={{
            style: {
              background: "hsl(var(--background))",
              color: "hsl(var(--foreground))",
              border: "1px solid hsl(var(--border))",
            },
          }}
        />
      </QueryClientProvider>
    </ThemeProvider>
  );
}

export { App };
