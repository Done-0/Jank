/**
 * 分类相关类型定义
 */

// ===== 请求类型 (Request) =====

// CreateCategoryRequest 创建分类请求
export interface CreateCategoryRequest {
  name: string; // 分类名称
  description?: string; // 分类描述
  parent_id?: string; // 父分类 ID，为空表示顶级分类
  sort?: number; // 排序权重（int64），数字越大越靠前
  is_active?: boolean; // 是否启用，默认为 true
}

// DeleteCategoryRequest 删除分类请求
export interface DeleteCategoryRequest {
  id: string; // 分类 ID
}

// GetCategoryRequest 获取分类请求
export interface GetCategoryRequest {
  id: string; // 分类 ID
}

// UpdateCategoryRequest 更新分类请求
export interface UpdateCategoryRequest {
  id: string; // 分类 ID
  name?: string; // 分类名称
  description?: string; // 分类描述
  parent_id?: string; // 父分类 ID，为空表示顶级分类
  sort?: number; // 排序权重（int64），数字越大越靠前
  is_active?: boolean; // 是否启用，默认为 true
}

// ListCategoriesRequest 获取分类列表请求
export interface ListCategoriesRequest {
  page_no: number; // 页码（int64），从1开始
  page_size: number; // 每页数量（int64）
  parent_id?: string; // 父分类 ID，为空时获取顶级分类
  is_active?: boolean; // 是否启用，为空时获取所有分类
}

// ===== 响应类型 (Response) =====

// CreateCategoryResponse 创建分类响应
export interface CreateCategoryResponse {
  id: string; // 分类 ID
  name: string; // 分类名称
  description: string; // 分类描述
  parent_id: string; // 父分类 ID
  sort: number; // 排序权重（int64）
  is_active: boolean; // 是否启用
  message: string; // 创建结果消息
}

// GetCategoryResponse 获取分类响应
export interface GetCategoryResponse {
  id: string; // 分类 ID
  name: string; // 分类名称
  description: string; // 分类描述
  parent_id: string; // 父分类 ID
  sort: number; // 排序权重（int64）
  is_active: boolean; // 是否启用
  created_at: string; // 创建时间
  updated_at: string; // 更新时间
}

// UpdateCategoryResponse 更新分类响应
export interface UpdateCategoryResponse {
  id: string; // 分类 ID
  name: string; // 分类名称
  description: string; // 分类描述
  parent_id: string; // 父分类 ID
  sort: number; // 排序权重（int64）
  is_active: boolean; // 是否启用
  message: string; // 更新结果消息
}

// DeleteCategoryResponse 删除分类响应
export interface DeleteCategoryResponse {
  message: string; // 删除结果消息
}

// CategoryItem 分类列表项
export interface CategoryItem {
  id: string; // 分类 ID
  name: string; // 分类名称
  description: string; // 分类描述
  parent_id: string; // 父分类 ID
  sort: number; // 排序权重（int64）
  is_active: boolean; // 是否启用
  created_at: string; // 创建时间
  updated_at: string; // 更新时间
}

// ListCategoriesResponse 分类列表响应
export interface ListCategoriesResponse {
  total: number; // 总数量（int64）
  page_no: number; // 当前页码（int64）
  page_size: number; // 每页数量（int64）
  list: CategoryItem[]; // 分类列表
}

// ===== 客户端状态类型 =====

// CategoryTreeState 分类树状态（客户端使用）
export interface CategoryTreeState {
  categories: CategoryItem[];
  selectedCategory: CategoryItem | null;
  expandedKeys: string[];
  loading: boolean;
}
