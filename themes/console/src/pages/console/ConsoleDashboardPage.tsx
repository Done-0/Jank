/**
 * 控制台仪表盘页面
 */

import DashboardContent from "@/components/console/dashboard/DashboardContent";
import { useUsers } from "@/hooks/use-user";
import { usePostsByStatus } from "@/hooks/use-posts";
import { useThemes } from "@/hooks/use-themes";
import { usePlugins } from "@/hooks/use-plugins";

export function ConsoleDashboardPage() {
  // ===== 数据获取 =====
  const { data: usersData, isLoading: usersLoading } = useUsers({
    page_no: 1,
    page_size: 100,
  });

  const { data: postsData, isLoading: postsLoading } = usePostsByStatus({
    page_no: 1,
    page_size: 100,
  });

  const { data: themesData, isLoading: themesLoading } = useThemes({
    page_no: 1,
    page_size: 100,
  });

  const { data: pluginsData, isLoading: pluginsLoading } = usePlugins({
    page_no: 1,
    page_size: 100,
  });

  // ===== 提取数组数据 =====
  const users = usersData?.list || [];
  const posts = postsData?.list || [];
  const themes = themesData?.list || [];
  const plugins = pluginsData?.list || [];

  // ===== 加载状态 =====
  const isLoading =
    usersLoading || postsLoading || themesLoading || pluginsLoading;

  return (
    <DashboardContent
      allUsers={users}
      allPosts={posts}
      allThemes={themes}
      allPlugins={plugins}
      isLoading={isLoading}
    />
  );
}
