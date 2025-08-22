import { createRbacGuardedRoute } from "@/router/guards/RbacGuard";
import { CONSOLE_ROUTES, RBAC_ACTION } from "@/constants";

import { ConsoleDashboardPage } from "@/pages/console/ConsoleDashboardPage";
import { ConsoleUsersPage } from "@/pages/console/ConsoleUsersPage";
import { ConsoleRBACPage } from "@/pages/console/ConsoleRBACPage";
import { ConsoleThemesPage } from "@/pages/console/ConsoleThemesPage";
import { ConsolePluginsPage } from "@/pages/console/ConsolePluginsPage";
import { ConsolePostsPage } from "@/pages/console/ConsolePostsPage";
import { PostEditorPage } from "@/pages/console/PostEditorPage";
import { ConsoleCategoriesPage } from "@/pages/console/ConsoleCategoriesPage";

export const createConsoleRoutes = (rootRoute: any) => {
  return [
    // Console 主页面 - 需要 Console 访问权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.ROOT, ConsoleDashboardPage, {
      requireConsoleAccess: true,
    }),

    // 用户管理页面 - 需要用户管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.USERS, ConsoleUsersPage, {
      requireConsoleAccess: true,
      resource: '/api/users/*',
      action: RBAC_ACTION.GET,
    }),

    // 权限管理页面 - 需要 RBAC 管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.RBAC, ConsoleRBACPage, {
      requireConsoleAccess: true,
      resource: '/api/rbac/*',
      action: RBAC_ACTION.GET,
    }),

    // 主题管理页面 - 需要主题管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.THEMES, ConsoleThemesPage, {
      requireConsoleAccess: true,
      resource: '/api/themes/*',
      action: RBAC_ACTION.GET,
    }),

    // 插件管理页面 - 需要插件管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.PLUGINS, ConsolePluginsPage, {
      requireConsoleAccess: true,
      resource: '/api/plugins/*',
      action: RBAC_ACTION.GET,
    }),

    // 文章管理页面 - 需要文章管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.POSTS, ConsolePostsPage, {
      requireConsoleAccess: true,
      resource: '/api/posts/*',
      action: RBAC_ACTION.GET,
    }),

    // 文章编辑器页面 - 需要文章管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.POST_EDITOR, PostEditorPage, {
      requireConsoleAccess: true,
      resource: '/api/posts/*',
      action: RBAC_ACTION.POST,
    }),

    // 分类管理页面 - 需要分类管理权限
    createRbacGuardedRoute(rootRoute, CONSOLE_ROUTES.CATEGORIES, ConsoleCategoriesPage, {
      requireConsoleAccess: true,
      resource: '/api/categories/*',
      action: RBAC_ACTION.GET,
    })
  ];
};
