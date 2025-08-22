/**
 * 用户相关常量定义
 */

// ===== 用户角色标签 =====
export const USER_ROLE_LABELS = {
  super_admin: "超级管理员",
  user: "普通用户",
} as const;

// ===== 用户查询键 =====
export const QUERY_KEYS = {
  CURRENT_USER: "currentUser",
  USERS: "users",
} as const;
