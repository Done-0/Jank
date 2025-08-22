/**
 * API 端点配置
 */

// ===== 验证码相关 =====
export const VERIFICATION_ENDPOINTS = {
  SEND_EMAIL_CODE: "/api/v1/verification/email",
} as const;

// ===== 用户相关 =====
export const USER_ENDPOINTS = {
  REGISTER: "/api/v1/user/register",
  LOGIN: "/api/v1/user/login",
  LOGOUT: "/api/v1/user/logout",
  REFRESH_TOKEN: "/api/v1/user/refresh-token",
  UPDATE: "/api/v1/user/update",
  RESET_PASSWORD: "/api/v1/user/reset-password",
  GET_PROFILE: "/api/v1/user/profile",
  LIST_USERS: "/api/v1/user/list",
  UPDATE_USER_ROLE: "/api/v1/user/role",
} as const;

// ===== RBAC权限管理相关 =====
export const RBAC_ENDPOINTS = {
  // 权限管理
  CREATE_PERMISSION: "/api/v1/rbac/create-permission",
  DELETE_PERMISSION: "/api/v1/rbac/delete-permission",
  ASSIGN_PERMISSION: "/api/v1/rbac/assign-permission",
  REVOKE_PERMISSION: "/api/v1/rbac/revoke-permission",
  LIST_PERMISSIONS: "/api/v1/rbac/list-permissions",
  
  // 角色管理
  CREATE_ROLE: "/api/v1/rbac/create-role",
  DELETE_ROLE: "/api/v1/rbac/delete-role",
  LIST_ROLES: "/api/v1/rbac/list-roles",
  GET_ROLE_PERMISSIONS: "/api/v1/rbac/get-role-permissions",
  
  // 用户角色管理
  ASSIGN_ROLE: "/api/v1/rbac/assign-role",
  REVOKE_ROLE: "/api/v1/rbac/revoke-role",
  GET_USER_ROLES: "/api/v1/rbac/get-user-roles",
  
  // 权限检查
  CHECK_PERMISSION: "/api/v1/rbac/check-permission",
} as const;

// ===== 主题相关 =====
export const THEME_ENDPOINTS = {
  SWITCH_THEME: "/api/v1/theme/switch",
  GET_ACTIVE_THEME: "/api/v1/theme/get",
  LIST_THEMES: "/api/v1/theme/list",
} as const;

// ===== 插件相关 =====
export const PLUGIN_ENDPOINTS = {
  REGISTER_PLUGIN: "/api/v1/plugin/register",
  UNREGISTER_PLUGIN: "/api/v1/plugin/unregister",
  EXECUTE_PLUGIN: "/api/v1/plugin/execute",
  GET_PLUGIN: "/api/v1/plugin/get",
  LIST_PLUGINS: "/api/v1/plugin/list",
} as const;

// ===== 文章相关 =====
export const POST_ENDPOINTS = {
  GET_POST: "/api/v1/post/get",
  LIST_PUBLISHED_POSTS: "/api/v1/post/list-published",
  LIST_POSTS_BY_STATUS: "/api/v1/post/list-by-status",
  CREATE_POST: "/api/v1/post/create",
  UPDATE_POST: "/api/v1/post/update",
  DELETE_POST: "/api/v1/post/delete",
} as const;

// ===== 分类相关 =====
export const CATEGORY_ENDPOINTS = {
  GET_CATEGORY: "/api/v1/category/get",
  LIST_CATEGORIES: "/api/v1/category/list",
  CREATE_CATEGORY: "/api/v1/category/create",
  UPDATE_CATEGORY: "/api/v1/category/update",
  DELETE_CATEGORY: "/api/v1/category/delete",
} as const;

