/**
 * 仪表盘主内容组件
 */

import { useMemo } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";
import { Users, FileText, Palette, Plug } from "lucide-react";

import type { UserItem } from "@/types/user";

interface DashboardContentProps {
  allUsers: UserItem[];
  allPosts: any[];
  allThemes: any[];
  allPlugins: any[];
  isLoading: boolean;
}

export default function DashboardContent({
  allUsers,
  allPosts,
  allThemes,
  allPlugins,
  isLoading,
}: DashboardContentProps) {
  // ===== 统计数据计算 =====
  const stats = useMemo(() => {
    const totalUsers = allUsers.length;

    // 动态统计角色分布，不使用硬编码
    const roleDistribution: Record<string, number> = {};
    allUsers.forEach((user) => {
      user.roles.forEach((role) => {
        roleDistribution[role] = (roleDistribution[role] || 0) + 1;
      });
    });

    // 获取最常见的角色作为显示
    const topRoles = Object.entries(roleDistribution)
      .sort(([, a], [, b]) => b - a)
      .slice(0, 2);

    const [primaryRole, secondaryRole] = topRoles;

    const totalPosts = allPosts.length;
    const publishedPosts = allPosts.filter(
      (post) => post.status === "published"
    ).length;
    const draftPosts = allPosts.filter(
      (post) => post.status === "draft"
    ).length;
    const archivedPosts = allPosts.filter(
      (post) => post.status === "archived"
    ).length;

    const totalThemes = allThemes.length;
    const activeThemes = allThemes.filter((theme) => theme.is_active).length;

    const totalPlugins = allPlugins.length;
    const activePlugins = allPlugins.filter(
      (plugin) => plugin.is_active
    ).length;

    return {
      totalUsers,
      primaryRoleName: primaryRole?.[0] || "暂无角色",
      primaryRoleCount: primaryRole?.[1] || 0,
      secondaryRoleName: secondaryRole?.[0] || "暂无角色",
      secondaryRoleCount: secondaryRole?.[1] || 0,
      totalPosts,
      publishedPosts,
      draftPosts,
      archivedPosts,
      totalThemes,
      activeThemes,
      totalPlugins,
      activePlugins,
    };
  }, [allUsers, allPosts, allThemes, allPlugins]);

  // ===== 加载状态 =====
  if (isLoading) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">仪表盘</h1>
          <p className="text-muted-foreground">系统数据概览</p>
        </div>
        <Separator />

        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <Card key={i}>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <Skeleton className="h-4 w-20" />
                <Skeleton className="h-4 w-4" />
              </CardHeader>
              <CardContent>
                <Skeleton className="h-8 w-16 mb-2" />
                <Skeleton className="h-3 w-24" />
              </CardContent>
            </Card>
          ))}
        </div>

        <div className="grid gap-4 md:grid-cols-2">
          {Array.from({ length: 2 }).map((_, i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-6 w-32" />
              </CardHeader>
              <CardContent>
                <div className="space-y-3">
                  {Array.from({ length: 3 }).map((_, j) => (
                    <div key={j} className="flex items-center justify-between">
                      <Skeleton className="h-4 w-20" />
                      <Skeleton className="h-4 w-8" />
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* ===== 关键指标卡片 ===== */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {/* 用户统计 */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">用户</CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stats.totalUsers}</div>
            <p className="text-xs text-muted-foreground">
              {stats.primaryRoleCount} {stats.primaryRoleName}
              {stats.secondaryRoleCount > 0 &&
                ` · ${stats.secondaryRoleCount} ${stats.secondaryRoleName}`}
            </p>
          </CardContent>
        </Card>

        {/* 内容统计 */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">内容</CardTitle>
            <FileText className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stats.totalPosts}</div>
            <p className="text-xs text-muted-foreground">
              {stats.publishedPosts} 已发布 · {stats.draftPosts} 草稿
            </p>
          </CardContent>
        </Card>

        {/* 主题统计 */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">主题</CardTitle>
            <Palette className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stats.totalThemes}</div>
            <p className="text-xs text-muted-foreground">
              {stats.activeThemes} 活跃
            </p>
          </CardContent>
        </Card>

        {/* 插件统计 */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">插件</CardTitle>
            <Plug className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stats.totalPlugins}</div>
            <p className="text-xs text-muted-foreground">
              {stats.activePlugins} 活跃
            </p>
          </CardContent>
        </Card>
      </div>

      {/* ===== 详细统计 ===== */}
      <div className="grid gap-4 md:grid-cols-2">
        {/* 用户分布 */}
        <Card>
          <CardHeader>
            <CardTitle>用户分布</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <span className="text-sm">{stats.primaryRoleName}</span>
                <span className="text-sm font-medium">
                  {stats.primaryRoleCount}
                </span>
              </div>
              {stats.secondaryRoleCount > 0 && (
                <div className="flex items-center justify-between">
                  <span className="text-sm">{stats.secondaryRoleName}</span>
                  <span className="text-sm font-medium">
                    {stats.secondaryRoleCount}
                  </span>
                </div>
              )}
            </div>
          </CardContent>
        </Card>

        {/* 内容分布 */}
        <Card>
          <CardHeader>
            <CardTitle>内容分布</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <span className="text-sm">已发布</span>
                <span className="text-sm font-medium">
                  {stats.publishedPosts}
                </span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm">草稿</span>
                <span className="text-sm font-medium">{stats.draftPosts}</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-sm">已归档</span>
                <span className="text-sm font-medium">
                  {stats.archivedPosts}
                </span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
