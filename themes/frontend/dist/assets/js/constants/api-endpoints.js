/**
 * API 端点常量定义
 */

// ===== 文章相关端点 =====
export const POST_ENDPOINTS = {
  LIST_PUBLISHED: "/api/v1/post/list-published",
  GET_POST: "/api/v1/post/get",
};

// ===== 分类相关端点 =====
export const CATEGORY_ENDPOINTS = {
  LIST_CATEGORIES: "/api/v1/category/list",
  GET_CATEGORY: "/api/v1/category/get",
};

// ===== 所有端点汇总 =====
export const API_ENDPOINTS = {
  ...POST_ENDPOINTS,
  ...CATEGORY_ENDPOINTS,
};
