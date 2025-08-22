/**
 * 主题管理页面
 */

import { ThemesContent } from "@/components/console/themes/ThemesContent";
import { useThemes } from "@/hooks/use-themes";

export function ConsoleThemesPage() {
  // ===== 数据获取 =====
  const {
    data: themesData,
    isLoading,
    error,
  } = useThemes({
    page_no: 1,
    page_size: 100,
  });

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-destructive mb-2">加载主题失败</p>
          <p className="text-sm text-muted-foreground">{error.message}</p>
        </div>
      </div>
    );
  }

  return (
    <ThemesContent themes={themesData?.list || []} isLoading={isLoading} />
  );
}
