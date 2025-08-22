/**
 * 权限守卫组件 - UI Component
 * 用于按钮/操作级权限控制，根据后端 RBAC 权限表动态控制组件显示/禁用
 */

import * as React from "react";
import { useHasPermission } from "@/hooks/use-rbac";

// ===== 类型定义 =====
import type { RbacAction } from "@/constants/rbac";

// ===== 接口定义 =====
interface PermissionGuardProps {
  /** 子组件 */
  children: React.ReactNode;
  /** 资源路径，对应后端权限表 v1 字段 */
  resource: string;
  /** 操作类型，对应后端权限表 v2 字段 */
  action: RbacAction;
  /** 无权限时的行为 */
  fallback?: "hide" | "disable" | "custom";
  /** 无权限时的自定义内容 */
  fallbackContent?: React.ReactNode;
  /** 无权限时的提示信息 */
  noPermissionMessage?: string;
}

// ===== 权限守卫组件 =====
const PermissionGuard = React.forwardRef<HTMLDivElement, PermissionGuardProps>(
  (
    {
      children,
      resource,
      action,
      fallback = "hide",
      fallbackContent,
      noPermissionMessage = "暂无权限",
    },
    ref
  ) => {
    const { hasPermission, isLoading } = useHasPermission(resource, action);

    // 加载中显示原组件（避免闪烁）
    if (isLoading) {
      return <>{children}</>;
    }

    // 完全依赖后端动态权限检查结果
    if (hasPermission) {
      return <>{children}</>;
    }

    // 无权限时的处理
    switch (fallback) {
      case "hide":
        return null;

      case "disable":
        // 如果子组件支持 disabled 属性，则禁用
        if (React.isValidElement(children)) {
          return (
            <div ref={ref} title={noPermissionMessage}>
              {React.cloneElement(children, { 
                ...children.props, 
                disabled: true 
              } as any)}
            </div>
          );
        }
        return <>{children}</>;

      case "custom":
        return <>{fallbackContent}</>;

      default:
        return null;
    }
  }
);

PermissionGuard.displayName = "PermissionGuard";

export { PermissionGuard };
export type { PermissionGuardProps };
