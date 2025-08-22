/**
 * 文章服务
 */

import { POST_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  CreatePostRequest,
  CreatePostResponse,
  DeletePostRequest,
  DeletePostResponse,
  GetPostRequest,
  GetPostResponse,
  UpdatePostRequest,
  UpdatePostResponse,
  ListPublishedPostsRequest,
  ListPostsByStatusRequest,
  ListPostsResponse,
} from "@/types";

class PostService {
  // ===== 文章管理 =====

  // 创建文章
  async createPost(request: CreatePostRequest): Promise<CreatePostResponse> {
    const response = await apiClient.post<ApiResponse<CreatePostResponse>>(
      POST_ENDPOINTS.CREATE_POST,
      request
    );
    return response.data.data!;
  }

  // 获取文章详情
  async getPost(request: GetPostRequest): Promise<GetPostResponse> {
    const response = await apiClient.get<ApiResponse<GetPostResponse>>(
      POST_ENDPOINTS.GET_POST,
      { params: request }
    );
    return response.data.data!;
  }

  // 更新文章
  async updatePost(request: UpdatePostRequest): Promise<UpdatePostResponse> {
    const response = await apiClient.post<ApiResponse<UpdatePostResponse>>(
      POST_ENDPOINTS.UPDATE_POST,
      request
    );
    return response.data.data!;
  }

  // 删除文章
  async deletePost(request: DeletePostRequest): Promise<DeletePostResponse> {
    const response = await apiClient.post<ApiResponse<DeletePostResponse>>(
      POST_ENDPOINTS.DELETE_POST,
      request
    );
    return response.data.data!;
  }

  // 获取已发布文章列表
  async listPublishedPosts(
    request: ListPublishedPostsRequest
  ): Promise<ListPostsResponse> {
    const response = await apiClient.get<ApiResponse<ListPostsResponse>>(
      POST_ENDPOINTS.LIST_PUBLISHED_POSTS,
      { params: request }
    );
    return response.data.data!;
  }

  // 根据状态获取文章列表
  async listPostsByStatus(
    request: ListPostsByStatusRequest
  ): Promise<ListPostsResponse> {
    const response = await apiClient.get<ApiResponse<ListPostsResponse>>(
      POST_ENDPOINTS.LIST_POSTS_BY_STATUS,
      { params: request }
    );
    return response.data.data!;
  }
}

export const postService = new PostService();
