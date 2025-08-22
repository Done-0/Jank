/**
 * 主题相关类型定义
 */

import type { ThemeType } from "@/constants/theme";

// ===== 请求类型 (Request) =====

// SwitchThemeRequest 切换主题请求
export interface SwitchThemeRequest {
  id: string; // 主题 ID
  theme_type: ThemeType; // 主题类型：frontend/console
  rebuild?: boolean; // 是否重新构建主题，默认 false
}

// ListThemesRequest 获取主题列表请求
export interface ListThemesRequest {
  status?: string; // 主题状态筛选
  page_no: number; // 页码，从1开始，必填
  page_size: number; // 每页数量，1-100，必填
}

// ===== 响应类型 (Response) =====

// SwitchThemeResponse 切换主题响应
export interface SwitchThemeResponse {
  message: string; // 切换结果消息
}

// GetThemeResponse 主题信息
export interface GetThemeResponse {
  // 基本信息
  id: string; // 主题唯一标识
  name: string; // 主题显示名称
  version: string; // 主题版本号
  author: string; // 主题作者
  description?: string; // 主题描述
  repository?: string; // 主题仓库地址
  preview?: string; // 主题预览图片地址

  // 配置信息
  type: string; // 主题类型（frontend/console）
  index_file_path?: string; // 首页文件相对路径
  static_dir_path?: string; // 静态资源目录相对路径

  // 运行时信息
  status: string; // 当前状态（ready/active/inactive/error）
  loaded_at?: number; // 加载时间戳
  path: string; // 主题文件路径
  is_active: boolean; // 是否为当前激活主题
}

// GetActiveThemeResponse 获取当前主题响应
export interface GetActiveThemeResponse {
  theme: GetThemeResponse | null; // 当前激活主题信息
}

// ListThemesResponse 主题列表响应
export interface ListThemesResponse {
  total: number; // 总数
  page_no: number; // 当前页码
  page_size: number; // 每页大小
  list: GetThemeResponse[]; // 主题列表
}

// ===== 客户端状态类型 =====

// ThemeState 主题状态（客户端使用）
export interface ThemeState {
  currentTheme: GetThemeResponse | null;
  themes: GetThemeResponse[];
  loading: boolean;
}
