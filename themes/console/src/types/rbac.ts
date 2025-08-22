/**
 * RBAC 权限管理相关类型定义
 */

import type { RbacAction } from "../constants/rbac";

// ===== 请求类型 (Request) =====

// CreatePermissionRequest 创建权限策略请求
export interface CreatePermissionRequest {
  name: string; // 权限名称
  description?: string; // 权限描述（可选）
  role: string; // 角色名称
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// DeletePermissionRequest 删除权限策略请求
export interface DeletePermissionRequest {
  role: string; // 角色名称
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// AssignPermissionRequest 为角色分配权限请求
export interface AssignPermissionRequest {
  role: string; // 角色名称
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// RevokePermissionRequest 撤销角色权限请求
export interface RevokePermissionRequest {
  role: string; // 角色名称
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// CreateRoleRequest 创建角色请求
export interface CreateRoleRequest {
  name: string; // 角色名称
  description?: string; // 角色描述（可选）
  role: string; // 角色标识符
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// CreateRoleOnlyRequest 仅创建角色请求（不包含权限）
export interface CreateRoleOnlyRequest {
  name: string; // 角色名称
  description?: string; // 角色描述（可选）
}

// DeleteRoleRequest 删除角色请求
export interface DeleteRoleRequest {
  role: string; // 角色名称
}

// GetRolePermissionsRequest 获取角色权限请求
export interface GetRolePermissionsRequest {
  role: string; // 角色名称
}

// AssignRoleRequest 为用户分配角色请求
export interface AssignRoleRequest {
  user_id: string; // 用户 ID (前端使用下划线格式)
  role: string; // 角色名称
}

// RevokeRoleRequest 撤销用户角色请求
export interface RevokeRoleRequest {
  user_id: string; // 用户 ID
  role: string; // 角色名称
}

// GetUserRolesRequest 获取用户角色请求
export interface GetUserRolesRequest {
  user_id: string; // 用户 ID
}

// CheckPermissionRequest 权限检查请求
export interface CheckPermissionRequest {
  user_id: string; // 用户 ID
  resource: string; // 资源路径
  action: RbacAction; // 操作方法
}

// ===== 响应类型 (Response) =====

// PolicyResponse 权限策略响应
export interface PolicyResponse {
  role: string; // 角色名称
  resource: string; // 资源路径
  action: string; // 操作方法
}

// PermissionResponse 权限响应
export interface PermissionResponse {
  name: string; // 权限名称
  description: string; // 权限描述
  resource: string; // 资源路径
  action: string; // 操作方法
}

// PermissionListResponse 权限列表响应
export interface PermissionListResponse {
  total: number; // 总条数
  list: PermissionResponse[]; // 权限列表
}

// PermissionOpResponse 权限分配操作响应
export interface PermissionOpResponse {
  success: boolean; // 操作是否成功
  role: string; // 角色名称
  permission: string; // 权限名称
  message: string; // 操作结果消息
}

// PolicyOpResponse 策略操作响应
export interface PolicyOpResponse {
  success: boolean; // 操作是否成功
  name: string; // 权限名称
  description: string; // 权限描述
  role: string; // 角色名称
  resource: string; // 资源路径
  action: string; // 操作方法
  message: string; // 操作消息
}

// RoleListResponse 角色列表响应
export interface RoleListResponse {
  total: number; // 总条数
  list: string[]; // 角色列表
}

// RolePermissionsResponse 角色权限响应
export interface RolePermissionsResponse {
  role: string; // 角色名称
  total: number; // 权限总数
  permissions: PermissionResponse[]; // 权限列表
}

// UserRolesResponse 用户角色响应
export interface UserRolesResponse {
  user_id: string; // 用户 ID
  roles: string[]; // 角色列表
}

// ListUserResponse 用户列表响应
export interface ListUserResponse {
  total: number; // 总条数
  list: UserInfo[]; // 用户列表
}

// UserInfo 用户信息
export interface UserInfo {
  id: string; // 用户 ID
  nickname: string; // 用户昵称
  email: string; // 用户邮箱
  roles: string[]; // 用户角色列表
}

// RoleOpResponse 用户角色操作响应
export interface RoleOpResponse {
  success: boolean; // 操作是否成功
  user_id: string; // 用户 ID
  role: string; // 角色名称
  roles: string[]; // 用户所有角色列表（成功时返回）
  message: string; // 操作结果消息
}

// CheckResponse 权限检查响应
export interface CheckResponse {
  allowed: boolean; // 是否允许
  reason: string; // 原因说明
}

// CreateRoleResponse 创建角色响应
export interface CreateRoleResponse {
  success: boolean; // 操作是否成功
  name: string; // 角色名称
  description: string; // 角色描述
  role: string; // 角色标识符
  resource: string; // 资源路径
  action: string; // 操作方法
  message: string; // 操作消息
}

// DeleteRoleResponse 删除角色响应
export interface DeleteRoleResponse {
  role: string; // 角色名称
  success: boolean; // 删除是否成功
  message: string; // 操作结果消息
}

// ===== 客户端状态类型 =====

// RbacState 权限管理状态（客户端使用）
export interface RbacState {
  roles: string[];
  permissions: PermissionResponse[];
  userRoles: { [userId: string]: string[] };
  loading: boolean;
}
