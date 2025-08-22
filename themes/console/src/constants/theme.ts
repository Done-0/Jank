/**
 * 主题相关常量定义
 */

// ===== 主题类型 =====
export const THEME_TYPE = {
  FRONTEND: "frontend",
  CONSOLE: "console",
} as const;

export type ThemeType = (typeof THEME_TYPE)[keyof typeof THEME_TYPE];

// ===== 主题状态 =====
export const THEME_STATUS = {
  READY: "ready",
  ACTIVE: "active",
  INACTIVE: "inactive",
  ERROR: "error",
} as const;

export type ThemeStatus = (typeof THEME_STATUS)[keyof typeof THEME_STATUS];

// ===== 主题查询键 =====
export const THEME_QUERY_KEYS = {
  THEMES: "themes",
  ACTIVE_THEME: "activeTheme",
} as const;
