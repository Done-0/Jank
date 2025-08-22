/**
 * 存储相关常量定义
 */

// ===== localStorage 存储键 =====
export const STORAGE_KEYS = {
  // 认证相关
  ACCESS_TOKEN: "jank-access-token",
  REFRESH_TOKEN: "jank-refresh-token",
  USER_INFO: "jank-user-info",

  // Zustand 持久化存储键
  AUTH_STORE: "jank-auth-store",
  USER_STORE: "jank-user-store",
} as const;

// ===== 类型定义 =====
export type StorageKey = (typeof STORAGE_KEYS)[keyof typeof STORAGE_KEYS];
