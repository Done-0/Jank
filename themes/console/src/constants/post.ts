/**
 * 文章相关常量定义
 */

// ===== 文章状态 =====
export const POST_STATUS = {
  DRAFT: "draft",
  PUBLISHED: "published",
  PRIVATE: "private",
  ARCHIVED: "archived",
} as const;

export type PostStatus = (typeof POST_STATUS)[keyof typeof POST_STATUS];

// ===== 文章查询键 =====
export const POST_QUERY_KEYS = {
  POSTS: "posts",
  POST_DETAIL: "postDetail",
} as const;
