/**
 * 分类相关常量
 */

// ===== React Query 查询键 =====

export const CATEGORY_QUERY_KEYS = {
  LIST: "categories",
  ALL: "categories-all",
  DETAIL: "category-detail",
} as const;

// ===== 分类颜色配置 =====

export const CATEGORY_COLORS = [
  "bg-blue-500",
  "bg-green-500",
  "bg-purple-500",
  "bg-red-500",
  "bg-yellow-500",
  "bg-indigo-500",
  "bg-pink-500",
  "bg-teal-500",
] as const;

// ===== 分类状态 =====

export const CATEGORY_STATUS = {
  ACTIVE: true,
  INACTIVE: false,
} as const;
