import { createElement } from "react";

import type { QueryClient } from "@tanstack/react-query";
import { createRouter, createRootRoute, Outlet } from "@tanstack/react-router";

import { createAuthRoutes, createConsoleRoutes } from "@/router/routes";
import { NotFoundPage } from "@/pages/NotFoundPage";

// 根路由
const rootRoute = createRootRoute({
  component: () => createElement(Outlet),
  notFoundComponent: NotFoundPage,
});

// 创建路由器
export const createAppRouter = (queryClient: QueryClient) => {
  const routes = [
    ...createAuthRoutes(rootRoute),
    ...createConsoleRoutes(rootRoute),
  ];

  return createRouter({
    routeTree: rootRoute.addChildren(routes),
    context: { queryClient },
  });
};
