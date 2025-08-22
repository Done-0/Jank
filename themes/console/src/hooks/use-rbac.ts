/**
 * RBAC 权限管理
 */

import { useMemo, useEffect, useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

import { useAuthStore } from "@/stores/auth.store";
import { useUserStore } from "@/stores/user.store";
import { rbacService } from "@/services/rbac.service";
import { RBAC_QUERY_KEYS, RBAC_ACTION, type RbacAction } from "@/constants";
import type {
  CreateRoleRequest,
  DeleteRoleRequest,
  CreatePermissionRequest,
  DeletePermissionRequest,
  AssignPermissionRequest,
  AssignRoleRequest,
  RevokeRoleRequest,
  CheckPermissionRequest,
  GetUserRolesRequest,
  GetRolePermissionsRequest,
} from "@/types/rbac";

// ===== 类型定义 =====

// 权限项定义
export interface PermissionItem {
  action: RbacAction;
  resource: string;
}

// RBAC Query Keys - 统一查询键管理
export const rbacKeys = {
  all: ["rbac"] as const,
  roles: () => [...rbacKeys.all, "roles"] as const,
  permissions: () => [...rbacKeys.all, "permissions"] as const,
  userRoles: (userId: string) =>
    [...rbacKeys.all, "user-roles", userId] as const,
};

// ===== 权限检查 Hooks =====

// 批量权限检查
export const useBatchPermissions = (permissionItems: PermissionItem[]) => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();
  const [permissions, setPermissions] = useState<Record<string, boolean>>({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!isAuthenticated || !user?.id || permissionItems.length === 0) {
      setIsLoading(false);
      return;
    }

    const checkPermissions = async () => {
      try {
        setIsLoading(true);
        setError(null);

        // 批量检查所有权限
        const permissionResults = await Promise.all(
          permissionItems.map(({ resource, action }) =>
            rbacService.checkPermission({
              user_id: user.id.toString(),
              resource,
              action,
            })
          )
        );

        // 构建权限映射对象
        const perms = permissionItems.reduce(
          (acc, { action, resource }, index) => {
            acc[`${action}:${resource}`] =
              permissionResults[index]?.allowed || false;
            return acc;
          },
          {} as Record<string, boolean>
        );

        setPermissions(perms);
      } catch (err) {
        console.error("批量权限检查失败:", err);
        setError("权限检查失败");
      } finally {
        setIsLoading(false);
      }
    };

    checkPermissions();
  }, [isAuthenticated, user?.id, permissionItems]);

  return { permissions, isLoading, error };
};

// 单个权限检查
export const usePermission = (resource: string, action: RbacAction) => {
  const permissionItems = useMemo(
    () => [{ resource, action }],
    [resource, action]
  );
  const { permissions, isLoading, error } =
    useBatchPermissions(permissionItems);

  const hasPermission = permissions[`${action}:${resource}`] || false;

  return { hasPermission, isLoading, error };
};

// 便捷的权限检查函数
export const useHasPermission = (resource: string, action: RbacAction) => {
  return usePermission(resource, action);
};

// 权限检查 Hook
export const useCheckPermission = (resource: string, action: RbacAction) => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();

  const request: CheckPermissionRequest = useMemo(
    () => ({
      user_id: user?.id?.toString() || "",
      resource,
      action,
    }),
    [user?.id, resource, action]
  );

  return useQuery({
    queryKey: [RBAC_QUERY_KEYS.CHECK_PERMISSION, request],
    queryFn: () => rbacService.checkPermission(request),
    enabled: isAuthenticated && !!user?.id,
    staleTime: 5 * 60 * 1000, // 5分钟缓存
    gcTime: 10 * 60 * 1000, // 10分钟垃圾回收
  });
};

// 用户角色检查
export const useUserRoles = () => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();

  // 构建用户角色查询请求参数
  const request: GetUserRolesRequest = useMemo(
    () => ({
      user_id: user?.id?.toString() || "",
    }),
    [user?.id]
  );

  return useQuery({
    queryKey: [RBAC_QUERY_KEYS.USER_ROLES, user?.id],
    queryFn: () => rbacService.getUserRoles(request),
    enabled: isAuthenticated && !!user?.id,
    staleTime: 10 * 60 * 1000, // 10分钟缓存
    gcTime: 30 * 60 * 1000, // 30分钟垃圾回收
  });
};

// Console 访问权限检查
export const useCanAccessConsole = () => {
  const { user } = useUserStore();
  const { isAuthenticated } = useAuthStore();

  const { data: consolePermission, isLoading } = useCheckPermission(
    "console",
    RBAC_ACTION.GET
  );

  const canAccess = useMemo(() => {
    if (!isAuthenticated || !user) return false;

    return consolePermission?.allowed || false;
  }, [isAuthenticated, user, consolePermission?.allowed]);

  return {
    canAccess,
    isLoading,
    user,
  };
};

// ===== 数据查询 Hooks =====

// 获取角色列表
export const useRoles = () => {
  return useQuery({
    queryKey: rbacKeys.roles(),
    queryFn: () => rbacService.listRoles(),
    staleTime: 5 * 60 * 1000, // 5分钟缓存
  });
};

// 获取权限列表
export const usePermissions = () => {
  return useQuery({
    queryKey: rbacKeys.permissions(),
    queryFn: () => rbacService.listPermissions(),
    staleTime: 5 * 60 * 1000, // 5分钟缓存
  });
};

// 获取角色权限
export const useRolePermissions = (role: string, options?: any) => {
  const request: GetRolePermissionsRequest = useMemo(
    () => ({
      role,
    }),
    [role]
  );

  return useQuery({
    queryKey: [RBAC_QUERY_KEYS.GET_ROLE_PERMISSIONS, role],
    queryFn: () => rbacService.getRolePermissions(request),
    enabled: !!role,
    staleTime: 5 * 60 * 1000, // 5分钟缓存
    ...options,
  });
};

// ===== 角色管理 Hooks =====

// 创建角色
export const useCreateRole = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: CreateRoleRequest) => rbacService.createRole(request),
    onSuccess: () => {
      toast.success("角色创建成功");
      queryClient.invalidateQueries({ queryKey: rbacKeys.roles() });
    },
    onError: (error: any) => {
      console.error("Create role error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "角色创建失败";
      toast.error(errorMessage);
    },
  });
};

// 删除角色
export const useDeleteRole = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: DeleteRoleRequest) => rbacService.deleteRole(request),
    onSuccess: () => {
      toast.success("角色删除成功");
      queryClient.invalidateQueries({ queryKey: rbacKeys.roles() });
    },
    onError: (error: any) => {
      console.error("Delete role error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "角色删除失败";
      toast.error(errorMessage);
    },
  });
};

// ===== 权限管理 Hooks =====

// 创建权限
export const useCreatePermission = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: CreatePermissionRequest) =>
      rbacService.createPermission(request),
    onSuccess: () => {
      toast.success("权限创建成功");
      queryClient.invalidateQueries({ queryKey: rbacKeys.permissions() });
    },
    onError: (error: any) => {
      console.error("Create permission error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "权限创建失败";
      toast.error(errorMessage);
    },
  });
};

// 删除权限
export const useDeletePermission = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: DeletePermissionRequest) =>
      rbacService.deletePermission(request),
    onSuccess: () => {
      toast.success("权限删除成功");
      queryClient.invalidateQueries({ queryKey: rbacKeys.permissions() });
    },
    onError: (error: any) => {
      console.error("Delete permission error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "权限删除失败";
      toast.error(errorMessage);
    },
  });
};

// ===== 权限分配 Hooks =====

// 分配权限
export const useAssignPermission = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: AssignPermissionRequest) =>
      rbacService.assignPermission(request),
    onSuccess: () => {
      toast.success("权限分配成功");
      queryClient.invalidateQueries({ queryKey: rbacKeys.permissions() });
    },
    onError: (error: any) => {
      console.error("Assign permission error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "权限分配失败";
      toast.error(errorMessage);
    },
  });
};

// ===== 用户角色分配 Hooks =====

// 分配角色
export const useAssignRole = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: AssignRoleRequest) => rbacService.assignRole(request),
    onSuccess: (_, variables) => {
      toast.success("角色分配成功");
      queryClient.invalidateQueries({
        queryKey: rbacKeys.userRoles(variables.user_id),
      });
    },
    onError: (error: any) => {
      console.error("Assign role error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "角色分配失败";
      toast.error(errorMessage);
    },
  });
};

// 撤销角色
export const useRevokeRole = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (request: RevokeRoleRequest) => rbacService.revokeRole(request),
    onSuccess: (_, variables) => {
      toast.success("角色撤销成功");
      queryClient.invalidateQueries({
        queryKey: rbacKeys.userRoles(variables.user_id),
      });
    },
    onError: (error: any) => {
      console.error("Revoke role error:", error);
      const errorMessage =
        error?.response?.data?.message || error?.message || "角色撤销失败";
      toast.error(errorMessage);
    },
  });
};
