/**
 * 路由常量定义
 */

// ===== 认证路由（公开访问） =====
export const AUTH_ROUTES = {
  LOGIN: "/console/login",
  REGISTER: "/console/register",
} as const;

// ===== Console 管理路由（需要登录） =====
export const CONSOLE_ROUTES = {
  ROOT: "/console",
  USERS: "/console/users",
  RBAC: "/console/rbac",
  THEMES: "/console/themes",
  PLUGINS: "/console/plugins",
  POSTS: "/console/posts",
  POST_EDITOR: "/console/posts/editor",
  CATEGORIES: "/console/categories",
  SYSTEM: "/console/system",
  PROFILE: "/console/profile",
} as const;

// ===== 路由标题映射 =====
export const ROUTE_TITLES = {
  // 认证路由标题
  [AUTH_ROUTES.LOGIN]: "登录",
  [AUTH_ROUTES.REGISTER]: "注册",

  // Console 路由标题
  [CONSOLE_ROUTES.ROOT]: "控制台",
  [CONSOLE_ROUTES.USERS]: "用户管理",
  [CONSOLE_ROUTES.RBAC]: "权限管理",
  [CONSOLE_ROUTES.THEMES]: "主题管理",
  [CONSOLE_ROUTES.PLUGINS]: "插件管理",
  [CONSOLE_ROUTES.POSTS]: "文章管理",
  [CONSOLE_ROUTES.POST_EDITOR]: "文章编辑器",
  [CONSOLE_ROUTES.CATEGORIES]: "分类管理",
  [CONSOLE_ROUTES.SYSTEM]: "系统设置",
  [CONSOLE_ROUTES.PROFILE]: "个人资料",
} as const;

// ===== 路由类型定义 =====
export type AuthRoute = (typeof AUTH_ROUTES)[keyof typeof AUTH_ROUTES];
export type ConsoleRoute = (typeof CONSOLE_ROUTES)[keyof typeof CONSOLE_ROUTES];
export type RouteTitle = (typeof ROUTE_TITLES)[keyof typeof ROUTE_TITLES];
