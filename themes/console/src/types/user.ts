/**
 * 用户认证相关类型
 */

// ===== 请求类型 (Request) =====

// LoginRequest 用户登录请求
export interface LoginRequest {
  email: string; // 用户邮箱
  password: string; // 用户密码
}

// RegisterRequest 用户注册请求
export interface RegisterRequest {
  email: string; // 用户邮箱
  password: string; // 用户密码
  nickname: string; // 用户昵称
  email_verification_code: string; // 邮箱验证码
}

// RefreshTokenRequest 刷新token请求
export interface RefreshTokenRequest {
  refresh_token: string; // refresh token
}

// UpdateRequest 更新用户信息请求
export interface UpdateRequest {
  nickname?: string; // 用户昵称
  avatar?: string; // 用户头像
}

// ResetPasswordRequest 重置密码请求
export interface ResetPasswordRequest {
  old_password: string; // 原密码
  new_password: string; // 新密码
  email_verification_code: string; // 邮箱验证码
}

// ListUsersRequest 获取用户列表请求
export interface ListUsersRequest {
  page_no: number; // 页码
  page_size: number; // 每页数量
  keyword?: string; // 搜索关键词（邮箱、昵称）
  role?: string; // 角色筛选
}

// UpdateUserRoleRequest 管理员更新用户角色请求
export interface UpdateUserRoleRequest {
  id: string; // 目标用户 ID
  role: string; // 新角色
}

// AdminUpdateUserRequest 管理员更新用户信息请求
export interface AdminUpdateUserRequest {
  id: string; // 目标用户 ID
  nickname?: string; // 昵称
  avatar?: string; // 头像
  email?: string; // 邮箱
}

// DeleteUserRequest 删除用户请求
export interface DeleteUserRequest {
  id: string; // 目标用户 ID
}

// AddUserRoleRequest 为用户添加角色请求
export interface AddUserRoleRequest {
  id: string; // 目标用户 ID
  role: string; // 要添加的角色
}

// RemoveUserRoleRequest 移除用户角色请求
export interface RemoveUserRoleRequest {
  id: string; // 目标用户 ID
  role: string; // 要移除的角色
}

// BatchUpdateUserRolesRequest 批量更新用户角色请求
export interface BatchUpdateUserRolesRequest {
  id: string; // 目标用户 ID
  roles: string[]; // 新的角色列表（完全替换）
}

// ===== 响应类型 (Response) =====

// RegisterResponse 用户注册响应
export interface RegisterResponse {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  roles: string[]; // 用户角色列表
}

// LoginResponse 用户登录响应
export interface LoginResponse {
  access_token: string; // 访问令牌
  refresh_token: string; // 刷新令牌
}

// LogoutResponse 用户登出响应
export interface LogoutResponse {
  message: string; // 登出结果消息
}

// RefreshTokenResponse 刷新token响应
export interface RefreshTokenResponse {
  access_token: string; // 新的访问令牌
  refresh_token: string; // 新的刷新令牌
}

// UserItem 用户列表项
export interface UserItem {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  avatar: string; // 用户头像
  roles: string[]; // 用户角色列表
}

// GetProfileResponse 获取用户资料响应
export interface GetProfileResponse {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  avatar: string; // 用户头像
  roles: string[]; // 用户角色列表
}

// UpdateResponse 更新用户信息响应
export interface UpdateResponse {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  avatar: string; // 用户头像
  roles: string[]; // 用户角色列表
}

// ResetPasswordResponse 重置密码响应
export interface ResetPasswordResponse {
  message: string; // 重置结果消息
}

// UpdateUserRoleResponse 管理员更新用户角色响应
export interface UpdateUserRoleResponse {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  roles: string[]; // 更新后的角色列表
  message: string; // 操作结果消息
}

// UserRoleOperationResponse 用户角色操作响应（通用）
export interface UserRoleOperationResponse {
  id: string; // 用户 ID
  email: string; // 用户邮箱
  nickname: string; // 用户昵称
  roles: string[]; // 更新后的角色列表
  operation: "add" | "remove" | "update" | "batch_update"; // 操作类型
  affected_role?: string; // 受影响的角色（单个操作时）
  message: string; // 操作结果消息
}

// ListUsersResponse 用户列表响应
export interface ListUsersResponse {
  total: number; // 总数量
  page_no: number; // 当前页码
  page_size: number; // 每页数量
  list: UserItem[]; // 用户列表
}

// ===== 客户端状态类型 =====

// UserState 用户状态（客户端使用）
export interface UserState {
  user: GetProfileResponse | null; // 当前用户信息
  isAuthenticated: boolean; // 是否已认证
  loading: boolean; // 加载状态
}

// ===== 角色权限相关类型 =====

// RoleHierarchy 角色层级定义
export interface RoleHierarchy {
  role: string; // 角色名称
  level: number; // 角色级别（数字越大权限越高）
  inherits?: string[]; // 继承的角色列表
  permissions: string[]; // 直接权限列表
  description: string; // 角色描述
}

// UserPermissions 用户权限信息
export interface UserPermissions {
  roles: string[]; // 用户角色列表
  permissions: string[]; // 用户所有权限（基于角色计算）
  computed_at: string; // 权限计算时间
}

// PermissionCheck 权限检查结果
export interface PermissionCheck {
  permission: string; // 被检查的权限
  granted: boolean; // 是否授予
  source_roles: string[]; // 授予该权限的角色列表
  reason?: string; // 拒绝原因（如果被拒绝）
}
