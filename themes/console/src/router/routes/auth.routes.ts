import { createGuardedRoute } from "@/router/guards/AuthGuard";
import { AUTH_ROUTES } from "@/constants";
import { LoginPage } from "@/pages/auth/LoginPage";
import { RegisterPage } from "@/pages/auth/RegisterPage";

export const createAuthRoutes = (rootRoute: any) => {
  return [
    // 登录页面 (公开路由)
    createGuardedRoute(rootRoute, AUTH_ROUTES.LOGIN, LoginPage, true),
    // 注册页面 (公开路由)
    createGuardedRoute(rootRoute, AUTH_ROUTES.REGISTER, RegisterPage, true),
  ];
};
