import { createRoute, redirect } from "@tanstack/react-router";

import { useAuthStore } from "@/stores";
import { rbacService } from "@/services";
import { ConsoleLayout } from "@/layouts/ConsoleLayout";
import { AUTH_ROUTES } from "@/constants";
import { RBAC_ACTION, type RbacAction } from "@/constants";
import { STORAGE_KEYS } from "@/constants";

// 动态 Console 权限检查
async function checkConsoleAccess(): Promise<boolean> {
  try {
    const userInfoStr = localStorage.getItem(STORAGE_KEYS.USER_INFO);
    if (!userInfoStr) return false;
    
    const userInfo = JSON.parse(userInfoStr);
    
    const result = await rbacService.checkPermission({
      user_id: userInfo.id.toString(),
      resource: '/console',
      action: RBAC_ACTION.GET,
    });
    
    return result.allowed;
  } catch (error) {
    console.error('Console 权限检查失败:', error);
    return false;
  }
}

// RBAC 权限路由守卫 - 负责角色和动态权限检查
export function createRbacGuardedRoute(
  parentRoute: any,
  path: string,
  component: React.ComponentType,
  options: {
    requireConsoleAccess?: boolean;
    resource?: string;
    action?: RbacAction;
  } = {}
) {
  const { requireConsoleAccess = false, resource, action } = options;

  return createRoute({
    getParentRoute: () => parentRoute,
    path,
    component: () => {
      const Component = component;
      return (
        <ConsoleLayout>
          <Component />
        </ConsoleLayout>
      );
    },
    beforeLoad: async () => {
      // 先进行基础认证检查
      const authStore = useAuthStore.getState();
      if (!authStore.isAuthenticated) {
        throw redirect({ to: AUTH_ROUTES.LOGIN });
      }

      // Console 权限检查
      if (requireConsoleAccess) {
        const canAccess = await checkConsoleAccess();
        if (!canAccess) {
          throw redirect({ to: AUTH_ROUTES.LOGIN });
        }
      }

      // 特定资源权限检查
      if (resource && action) {
        try {
          const userInfoStr = localStorage.getItem(STORAGE_KEYS.USER_INFO);
          if (!userInfoStr) {
            throw redirect({ to: AUTH_ROUTES.LOGIN });
          }
          
          const userInfo = JSON.parse(userInfoStr);
          
          // 完全依赖后端动态权限判断，不在前端写死任何角色逻辑
          const result = await rbacService.checkPermission({
            user_id: userInfo.id.toString(),
            resource,
            action,
          });
          
          if (!result.allowed) {
            throw redirect({ to: AUTH_ROUTES.LOGIN });
          }
        } catch (error) {
          console.error('权限检查失败:', error);
          throw redirect({ to: AUTH_ROUTES.LOGIN });
        }
      }
    },
  });
}
