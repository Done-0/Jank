import { createRoute, redirect } from "@tanstack/react-router";

import { useAuthStore } from "@/stores";
import { ConsoleLayout } from "@/layouts/ConsoleLayout";
import { AUTH_ROUTES } from "@/constants";

// 基础认证路由守卫
export function createGuardedRoute(
  parentRoute: any,
  path: string,
  component: React.ComponentType,
  isPublic: boolean = false
) {
  return createRoute({
    getParentRoute: () => parentRoute,
    path,
    component: () => {
      const Component = component;
      return isPublic ? (
        <Component />
      ) : (
        <ConsoleLayout>
          <Component />
        </ConsoleLayout>
      );
    },
    beforeLoad: isPublic ? undefined : () => {
      const authStore = useAuthStore.getState();
      if (!authStore.isAuthenticated) {
        throw redirect({ to: AUTH_ROUTES.LOGIN });
      }
    },
  });
}
