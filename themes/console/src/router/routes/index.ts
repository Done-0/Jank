import { AUTH_ROUTES } from "@/constants/routes";
import { CONSOLE_ROUTES } from "@/constants";

// 导出各模块路由创建函数
export { createAuthRoutes } from "./auth.routes";
export { createConsoleRoutes } from "./console.routes";

// 统一路由常量
export const ROUTES = {
  HOME: "/", // 首页 - 重定向到登录或控制台
  ...AUTH_ROUTES,
  ...CONSOLE_ROUTES,
} as const;
