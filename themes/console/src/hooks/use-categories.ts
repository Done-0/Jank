/**
 * 分类相关的 React Query hooks
 */

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { categoryService } from "@/services";
import { CATEGORY_QUERY_KEYS } from "@/constants";
import type {
  ListCategoriesRequest,
  CreateCategoryRequest,
  UpdateCategoryRequest,
  DeleteCategoryRequest,
} from "@/types/category";

// ===== 查询 hooks =====

/**
 * 获取分类列表
 */
export function useCategories(params: ListCategoriesRequest) {
  return useQuery({
    queryKey: [CATEGORY_QUERY_KEYS.LIST, params],
    queryFn: () => categoryService.listCategories(params),
    staleTime: 5 * 60 * 1000, // 5分钟缓存
  });
}

/**
 * 获取所有分类（不分页）
 */
export function useAllCategories() {
  return useQuery({
    queryKey: [CATEGORY_QUERY_KEYS.ALL],
    queryFn: () =>
      categoryService.listCategories({ page_no: 1, page_size: 100 }), // 修正为后端最大限制100
    staleTime: 10 * 60 * 1000, // 10分钟缓存
    select: (data) => data.list || [], // 只返回分类列表
  });
}

// ===== 变更 hooks =====

/**
 * 创建分类
 */
export function useCreateCategory() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: CreateCategoryRequest) =>
      categoryService.createCategory(request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.LIST] });
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.ALL] });
    },
  });
}

/**
 * 更新分类
 */
export function useUpdateCategory() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: UpdateCategoryRequest) =>
      categoryService.updateCategory(request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.LIST] });
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.ALL] });
    },
  });
}

/**
 * 删除分类
 */
export function useDeleteCategory() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: DeleteCategoryRequest) =>
      categoryService.deleteCategory(request),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.LIST] });
      queryClient.invalidateQueries({ queryKey: [CATEGORY_QUERY_KEYS.ALL] });
    },
  });
}
