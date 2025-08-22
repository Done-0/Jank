/**
 * 分类服务
 */

import { CATEGORY_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  CreateCategoryRequest,
  CreateCategoryResponse,
  DeleteCategoryRequest,
  DeleteCategoryResponse,
  GetCategoryRequest,
  GetCategoryResponse,
  UpdateCategoryRequest,
  UpdateCategoryResponse,
  ListCategoriesRequest,
  ListCategoriesResponse,
} from "@/types";

class CategoryService {
  // ===== 分类管理 =====

  // 创建分类
  async createCategory(
    request: CreateCategoryRequest
  ): Promise<CreateCategoryResponse> {
    const response = await apiClient.post<ApiResponse<CreateCategoryResponse>>(
      CATEGORY_ENDPOINTS.CREATE_CATEGORY,
      request
    );
    return response.data.data!;
  }

  // 获取分类详情
  async getCategory(request: GetCategoryRequest): Promise<GetCategoryResponse> {
    const response = await apiClient.get<ApiResponse<GetCategoryResponse>>(
      CATEGORY_ENDPOINTS.GET_CATEGORY,
      { params: request }
    );
    return response.data.data!;
  }

  // 更新分类
  async updateCategory(
    request: UpdateCategoryRequest
  ): Promise<UpdateCategoryResponse> {
    const response = await apiClient.post<ApiResponse<UpdateCategoryResponse>>(
      CATEGORY_ENDPOINTS.UPDATE_CATEGORY,
      request
    );
    return response.data.data!;
  }

  // 删除分类
  async deleteCategory(
    request: DeleteCategoryRequest
  ): Promise<DeleteCategoryResponse> {
    const response = await apiClient.post<ApiResponse<DeleteCategoryResponse>>(
      CATEGORY_ENDPOINTS.DELETE_CATEGORY,
      request
    );
    return response.data.data!;
  }

  // 获取分类列表
  async listCategories(
    request: ListCategoriesRequest
  ): Promise<ListCategoriesResponse> {
    const response = await apiClient.get<ApiResponse<ListCategoriesResponse>>(
      CATEGORY_ENDPOINTS.LIST_CATEGORIES,
      { params: request }
    );
    return response.data.data!;
  }
}

export const categoryService = new CategoryService();
