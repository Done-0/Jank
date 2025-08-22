/**
 * 文章相关类型定义
 */

import type { PostStatus } from "@/constants/post";

// ===== 请求类型 (Request) =====

// CreatePostRequest 创建文章请求
export interface CreatePostRequest {
  title: string; // 文章标题
  description?: string; // 文章描述/摘要
  image?: string; // 文章封面图片
  status?: PostStatus; // 文章状态
  category_id?: string; // 分类 ID
  markdown?: string; // Markdown 内容
}

// DeletePostRequest 删除文章请求
export interface DeletePostRequest {
  id: string; // 文章 ID
}

// GetPostRequest 获取文章请求
export interface GetPostRequest {
  id: string; // 文章 ID
}

// UpdatePostRequest 更新文章请求
export interface UpdatePostRequest {
  id: string; // 文章 ID
  title?: string; // 文章标题
  description?: string; // 文章描述/摘要
  image?: string; // 文章封面图片
  status?: PostStatus; // 文章状态
  category_id?: string; // 分类 ID
  markdown?: string; // Markdown内容
}

// ListPublishedPostsRequest 获取已发布文章列表请求
export interface ListPublishedPostsRequest {
  page_no: number; // 页码，从1开始
  page_size: number; // 每页数量
  category_id?: number; // 分类 ID，为空时不按分类筛选
}

// ListPostsByStatusRequest 根据状态获取文章列表请求
export interface ListPostsByStatusRequest {
  page_no: number; // 页码，从1开始
  page_size: number; // 每页数量
  status?: PostStatus; // 文章状态，为空时获取所有文章
  category_id?: number; // 分类 ID，为空时不按分类筛选
}

// ===== 响应类型 (Response) =====

// CreatePostResponse 创建文章响应
export interface CreatePostResponse {
  id: string; // 文章 ID
  title: string; // 文章标题
  description: string; // 文章描述/摘要
  image: string; // 文章封面图片
  status: string; // 文章状态
  category_id: string; // 分类 ID
  category_name: string; // 分类名称
  markdown: string; // Markdown内容
  message: string; // 创建结果消息
}

// GetPostResponse 获取文章响应
export interface GetPostResponse {
  id: string; // 文章 ID
  title: string; // 文章标题
  description: string; // 文章描述/摘要
  image: string; // 文章封面图片
  status: string; // 文章状态
  category_id: string; // 分类 ID
  category_name: string; // 分类名称
  markdown: string; // Markdown 内容
  html: string; // 渲染后的 HTML
  created_at: string; // 创建时间
  updated_at: string; // 更新时间
}

// UpdatePostResponse 更新文章响应
export interface UpdatePostResponse {
  id: string; // 文章 ID
  title: string; // 文章标题
  description: string; // 文章描述/摘要
  image: string; // 文章封面图片
  status: string; // 文章状态
  category_id: string; // 分类 ID
  category_name: string; // 分类名称
  markdown: string; // Markdown内容
  message: string; // 更新结果消息
}

// DeletePostResponse 删除文章响应
export interface DeletePostResponse {
  message: string; // 删除结果消息
}

// PostItem 文章列表项
export interface PostItem {
  id: string; // 文章 ID
  title: string; // 文章标题
  description: string; // 文章描述/摘要
  image: string; // 文章封面图片
  status: string; // 文章状态
  category_id: string; // 分类 ID
  category_name: string; // 分类名称
  created_at: string; // 创建时间
  updated_at: string; // 更新时间
}

// ListPostsResponse 文章列表响应
export interface ListPostsResponse {
  total: number; // 总数量
  page_no: number; // 当前页码
  page_size: number; // 每页数量
  list: PostItem[]; // 文章列表
}

// ===== 客户端状态类型 =====

// PostEditState 文章编辑状态（客户端使用）
export interface PostEditState {
  post: GetPostResponse | null;
  isEditing: boolean;
  isDirty: boolean;
  loading: boolean;
}
