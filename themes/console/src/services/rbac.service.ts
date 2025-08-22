/**
 * RBAC 权限管理服务
 */

import { RBAC_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  CreatePermissionRequest,
  PolicyOpResponse,
  DeletePermissionRequest,
  AssignPermissionRequest,
  PermissionOpResponse,
  RevokePermissionRequest,
  PermissionListResponse,
  CreateRoleRequest,
  CreateRoleResponse,
  DeleteRoleRequest,
  DeleteRoleResponse,
  RoleListResponse,
  GetRolePermissionsRequest,
  RolePermissionsResponse,
  AssignRoleRequest,
  RoleOpResponse,
  RevokeRoleRequest,
  GetUserRolesRequest,
  UserRolesResponse,
  CheckPermissionRequest,
  CheckResponse,
} from "@/types";

class RbacService {
  // ===== 权限管理 =====

  // 创建权限策略
  async createPermission(
    request: CreatePermissionRequest
  ): Promise<PolicyOpResponse> {
    const response = await apiClient.post<ApiResponse<PolicyOpResponse>>(
      RBAC_ENDPOINTS.CREATE_PERMISSION,
      request
    );
    return response.data.data!;
  }

  // 删除权限策略
  async deletePermission(
    request: DeletePermissionRequest
  ): Promise<PolicyOpResponse> {
    const response = await apiClient.post<ApiResponse<PolicyOpResponse>>(
      RBAC_ENDPOINTS.DELETE_PERMISSION,
      request
    );
    return response.data.data!;
  }

  // 为角色分配权限
  async assignPermission(
    request: AssignPermissionRequest
  ): Promise<PermissionOpResponse> {
    const response = await apiClient.post<ApiResponse<PermissionOpResponse>>(
      RBAC_ENDPOINTS.ASSIGN_PERMISSION,
      request
    );
    return response.data.data!;
  }

  // 撤销角色权限
  async revokePermission(
    request: RevokePermissionRequest
  ): Promise<PermissionOpResponse> {
    const response = await apiClient.post<ApiResponse<PermissionOpResponse>>(
      RBAC_ENDPOINTS.REVOKE_PERMISSION,
      request
    );
    return response.data.data!;
  }

  // 获取权限列表
  async listPermissions(): Promise<PermissionListResponse> {
    const response = await apiClient.get<ApiResponse<PermissionListResponse>>(
      RBAC_ENDPOINTS.LIST_PERMISSIONS
    );
    return response.data.data!;
  }

  // ===== 角色管理 =====

  // 创建角色
  async createRole(request: CreateRoleRequest): Promise<CreateRoleResponse> {
    const response = await apiClient.post<ApiResponse<CreateRoleResponse>>(
      RBAC_ENDPOINTS.CREATE_ROLE,
      request
    );
    return response.data.data!;
  }

  // 删除角色
  async deleteRole(request: DeleteRoleRequest): Promise<DeleteRoleResponse> {
    const response = await apiClient.post<ApiResponse<DeleteRoleResponse>>(
      RBAC_ENDPOINTS.DELETE_ROLE,
      request
    );
    return response.data.data!;
  }

  // 获取角色列表
  async listRoles(): Promise<RoleListResponse> {
    const response = await apiClient.get<ApiResponse<RoleListResponse>>(
      RBAC_ENDPOINTS.LIST_ROLES
    );
    return response.data.data!;
  }

  // 获取角色权限
  async getRolePermissions(
    request: GetRolePermissionsRequest
  ): Promise<RolePermissionsResponse> {
    const response = await apiClient.get<ApiResponse<RolePermissionsResponse>>(
      RBAC_ENDPOINTS.GET_ROLE_PERMISSIONS,
      { params: request }
    );
    return response.data.data!;
  }

  // ===== 用户角色管理 =====

  // 为用户分配角色
  async assignRole(request: AssignRoleRequest): Promise<RoleOpResponse> {
    const response = await apiClient.post<ApiResponse<RoleOpResponse>>(
      RBAC_ENDPOINTS.ASSIGN_ROLE,
      request
    );
    return response.data.data!;
  }

  // 撤销用户角色
  async revokeRole(request: RevokeRoleRequest): Promise<RoleOpResponse> {
    const response = await apiClient.post<ApiResponse<RoleOpResponse>>(
      RBAC_ENDPOINTS.REVOKE_ROLE,
      request
    );
    return response.data.data!;
  }

  // 获取用户角色
  async getUserRoles(request: GetUserRolesRequest): Promise<UserRolesResponse> {
    const backendParams = {
      user_id: request.user_id,
    };
    const response = await apiClient.get<ApiResponse<UserRolesResponse>>(
      RBAC_ENDPOINTS.GET_USER_ROLES,
      { params: backendParams }
    );
    return response.data.data!;
  }

  // ===== 权限检查 =====

  // 权限检查
  async checkPermission(
    request: CheckPermissionRequest
  ): Promise<CheckResponse> {
    const response = await apiClient.post<ApiResponse<CheckResponse>>(
      RBAC_ENDPOINTS.CHECK_PERMISSION,
      request
    );
    return response.data.data!;
  }
}

export const rbacService = new RbacService();
