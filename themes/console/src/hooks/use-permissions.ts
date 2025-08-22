/**
 * 权限检查 - Hooks
 * 提供灵活的权限检查，避免硬编码
 */

import { useHasPermission } from "./use-rbac";
import type { RbacAction } from "@/constants";

// ===== 通用权限检查 Hook =====
/**
 * 检查特定资源的权限
 * @param resource 资源路径，如 "/api/users/*"
 * @param action 操作类型，如 "GET", "POST", "PUT", "DELETE"
 * @returns 权限检查结果
 */
export const usePermission = (resource: string, action: RbacAction) => {
  return useHasPermission(resource, action);
};

// ===== 常用权限检查（仅保留真正需要的）=====

/**
 * 检查 Console 访问权限
 */
export const useConsoleAccess = () => {
  return useHasPermission("/console", "GET");
};
