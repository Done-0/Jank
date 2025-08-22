/**
 * 文章相关 Query Hooks
 */

import {
  useQuery,
  useMutation,
  useQueryClient,
  type UseQueryOptions,
  type UseMutationOptions,
} from "@tanstack/react-query";
import { postService } from "@/services";
import { POST_QUERY_KEYS } from "@/constants";
import type {
  GetPostRequest,
  GetPostResponse,
  CreatePostRequest,
  CreatePostResponse,
  UpdatePostRequest,
  UpdatePostResponse,
  DeletePostRequest,
  DeletePostResponse,
  ListPublishedPostsRequest,
  ListPostsByStatusRequest,
  ListPostsResponse,
} from "@/types";

// Query Keys
export const postKeys = {
  all: [POST_QUERY_KEYS.POSTS] as const,
  lists: () => [...postKeys.all, "list"] as const,
  list: (params: ListPublishedPostsRequest | ListPostsByStatusRequest) =>
    [...postKeys.lists(), params] as const,
  details: () => [...postKeys.all, POST_QUERY_KEYS.POST_DETAIL] as const,
  detail: (id: string) => [...postKeys.details(), id] as const,
} as const;

// 获取文章详情
export function usePost(
  params: GetPostRequest,
  options?: Omit<UseQueryOptions<GetPostResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: postKeys.detail(params.id),
    queryFn: () => postService.getPost(params),
    enabled: !!params.id,
    ...options,
  });
}

// 获取已发布文章列表
export function usePublishedPosts(
  params: ListPublishedPostsRequest,
  options?: Omit<UseQueryOptions<ListPostsResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: postKeys.list(params),
    queryFn: () => postService.listPublishedPosts(params),
    ...options,
  });
}

// 根据状态获取文章列表
export function usePostsByStatus(
  params: ListPostsByStatusRequest,
  options?: Omit<UseQueryOptions<ListPostsResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: postKeys.list(params),
    queryFn: () => postService.listPostsByStatus(params),
    ...options,
  });
}

// 创建文章
export function useCreatePost(
  options?: Omit<
    UseMutationOptions<CreatePostResponse, Error, CreatePostRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreatePostRequest) => postService.createPost(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: postKeys.lists() });
    },
    ...options,
  });
}

// 更新文章
export function useUpdatePost(
  options?: Omit<
    UseMutationOptions<UpdatePostResponse, Error, UpdatePostRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdatePostRequest) => postService.updatePost(data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({
        queryKey: postKeys.detail(variables.id),
      });
      queryClient.invalidateQueries({ queryKey: postKeys.lists() });
    },
    ...options,
  });
}

// 删除文章
export function useDeletePost(
  options?: Omit<
    UseMutationOptions<DeletePostResponse, Error, DeletePostRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: DeletePostRequest) => postService.deletePost(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: postKeys.lists() });
    },
    ...options,
  });
}
