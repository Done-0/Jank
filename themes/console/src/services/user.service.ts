/**
 * 用户服务
 */

import { USER_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import { useAuthStore } from "@/stores/auth.store";
import { useUserStore } from "@/stores/user.store";
import type {
  ApiResponse,
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
  UpdateRequest,
  UpdateResponse,
  GetProfileResponse,
  ListUsersRequest,
  ListUsersResponse,
  ResetPasswordRequest,
  ResetPasswordResponse,
  RefreshTokenRequest,
  RefreshTokenResponse,
  UpdateUserRoleRequest,
  UpdateUserRoleResponse,
} from "@/types";

class UserService {
  // ===== 用户相关 =====

  // 用户登录
  async login(request: LoginRequest): Promise<LoginResponse> {
    const response = await apiClient.post<ApiResponse<LoginResponse>>(
      USER_ENDPOINTS.LOGIN,
      request
    );
    const loginData = response.data.data!;
    useAuthStore.getState().login(loginData);

    try {
      const userProfile = await this.getProfile();
      useUserStore.getState().setUser(userProfile);
    } catch (error) {
      console.error("Failed to fetch user profile after login:", error);
      useAuthStore.getState().logout();
      throw error;
    }

    return loginData;
  }

  // 用户注册
  async register(request: RegisterRequest): Promise<RegisterResponse> {
    const response = await apiClient.post<ApiResponse<RegisterResponse>>(
      USER_ENDPOINTS.REGISTER,
      request
    );
    return response.data.data!;
  }

  // 用户登出
  async logout(): Promise<void> {
    await apiClient.post(USER_ENDPOINTS.LOGOUT);
    useAuthStore.getState().clearAuth();
    useUserStore.getState().clearUser();
  }

  // 刷新 Token
  async refreshToken(
    request: RefreshTokenRequest
  ): Promise<RefreshTokenResponse> {
    const response = await apiClient.post<ApiResponse<RefreshTokenResponse>>(
      USER_ENDPOINTS.REFRESH_TOKEN,
      request
    );
    const tokens = response.data.data!;
    useAuthStore
      .getState()
      .refreshTokens(tokens.access_token, tokens.refresh_token);
    return tokens;
  }

  // ===== 用户信息管理 =====

  // 获取用户资料
  async getProfile(): Promise<GetProfileResponse> {
    const response = await apiClient.get<ApiResponse<GetProfileResponse>>(
      USER_ENDPOINTS.GET_PROFILE
    );
    return response.data.data!;
  }

  // 更新用户信息
  async updateProfile(request: UpdateRequest): Promise<UpdateResponse> {
    const response = await apiClient.post<ApiResponse<UpdateResponse>>(
      USER_ENDPOINTS.UPDATE,
      request
    );
    const updatedUser = response.data.data!;

    // 更新 User Store 中的用户信息
    useUserStore.getState().updateUser({
      nickname: updatedUser.nickname,
      avatar: updatedUser.avatar,
    });

    return updatedUser;
  }

  // 重置密码
  async resetPassword(
    request: ResetPasswordRequest
  ): Promise<ResetPasswordResponse> {
    const response = await apiClient.post<ApiResponse<ResetPasswordResponse>>(
      USER_ENDPOINTS.RESET_PASSWORD,
      request
    );
    return response.data.data!;
  }

  // ===== 管理员功能 =====

  // 获取用户列表
  async listUsers(request: ListUsersRequest): Promise<ListUsersResponse> {
    const response = await apiClient.get<ApiResponse<ListUsersResponse>>(
      USER_ENDPOINTS.LIST_USERS,
      { params: request }
    );
    return response.data.data!;
  }

  // 更新用户角色
  async updateUserRole(
    request: UpdateUserRoleRequest
  ): Promise<UpdateUserRoleResponse> {
    const response = await apiClient.post<ApiResponse<UpdateUserRoleResponse>>(
      USER_ENDPOINTS.UPDATE_USER_ROLE,
      request
    );
    return response.data.data!;
  }
}

export const userService = new UserService();
