/**
 * 插件管理页面
 */

import { PluginsContent } from "@/components/console/plugins/PluginsContent";
import { usePlugins } from "@/hooks/use-plugins";

export function ConsolePluginsPage() {
  // ===== 数据获取 =====
  const {
    data: pluginsData,
    isLoading,
    error,
  } = usePlugins({
    page_no: 1,
    page_size: 100,
  });

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-destructive mb-2">加载插件失败</p>
          <p className="text-sm text-muted-foreground">{error.message}</p>
        </div>
      </div>
    );
  }

  return (
    <PluginsContent plugins={pluginsData?.list || []} isLoading={isLoading} />
  );
}
