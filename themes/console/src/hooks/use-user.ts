/**
 * 用户相关 Query Hooks
 */

import {
  useQuery,
  useMutation,
  useQueryClient,
  type UseQueryOptions,
  type UseMutationOptions,
} from "@tanstack/react-query";
import { userService } from "@/services";
import { QUERY_KEYS } from "@/constants/user";
import type {
  GetProfileResponse,
  UpdateRequest,
  UpdateResponse,
  ResetPasswordRequest,
  ResetPasswordResponse,
  ListUsersRequest,
  ListUsersResponse,
  UpdateUserRoleRequest,
  UpdateUserRoleResponse,
} from "@/types";

// Query Keys
export const userKeys = {
  all: [QUERY_KEYS.CURRENT_USER] as const,
  profile: () => [...userKeys.all, "profile"] as const,
  lists: () => [...userKeys.all, "list"] as const,
  list: (params: ListUsersRequest) => [...userKeys.lists(), params] as const,
} as const;

// 获取用户资料
export function useProfile(
  options?: Omit<UseQueryOptions<GetProfileResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: userKeys.profile(),
    queryFn: () => userService.getProfile(),
    ...options,
  });
}

// 更新用户信息
export function useUpdateProfile(
  options?: Omit<
    UseMutationOptions<UpdateResponse, Error, UpdateRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateRequest) => userService.updateProfile(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: userKeys.profile() });
    },
    ...options,
  });
}

// 重置密码
export function useResetPassword(
  options?: Omit<
    UseMutationOptions<ResetPasswordResponse, Error, ResetPasswordRequest>,
    "mutationFn"
  >
) {
  return useMutation({
    mutationFn: (data: ResetPasswordRequest) => userService.resetPassword(data),
    ...options,
  });
}

// 获取用户列表（管理员功能）
export function useUsers(
  params: ListUsersRequest,
  options?: Omit<UseQueryOptions<ListUsersResponse>, "queryKey" | "queryFn">
) {
  return useQuery({
    queryKey: userKeys.list(params),
    queryFn: () => userService.listUsers(params),
    ...options,
  });
}

// 更新用户角色（管理员功能）
export function useUpdateUserRole(
  options?: Omit<
    UseMutationOptions<UpdateUserRoleResponse, Error, UpdateUserRoleRequest>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateUserRoleRequest) =>
      userService.updateUserRole(data),
    onSuccess: () => {
      // 失效用户列表缓存
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      // 失效所有用户相关查询
      queryClient.invalidateQueries({ queryKey: userKeys.all });
      // 失效RBAC相关缓存，因为角色变更会影响权限
      queryClient.invalidateQueries({ queryKey: ["rbac"] });
    },
    ...options,
  });
}
