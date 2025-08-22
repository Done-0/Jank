/**
 * 主题管理内容组件
 * 负责主题列表展示、搜索、切换等核心功能
 */

import { useState, useMemo } from "react";
import {
  Search,
  MoreHorizontal,
  Eye,
  Palette,
  AlertCircle,
} from "lucide-react";
import { toast } from "sonner";
import { useQueryClient } from "@tanstack/react-query";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ImageLoader } from "@/components/ui/image-loader";

import { useSwitchTheme, themeKeys } from "@/hooks/use-themes";
import type { GetThemeResponse } from "@/types";

interface ThemesContentProps {
  themes: GetThemeResponse[];
  isLoading: boolean;
}

export function ThemesContent({ themes, isLoading }: ThemesContentProps) {
  // ===== State =====
  const [searchQuery, setSearchQuery] = useState("");
  const [configDialogOpen, setConfigDialogOpen] = useState(false);
  const [selectedTheme, setSelectedTheme] = useState<GetThemeResponse | null>(
    null
  );

  // ===== Hooks =====
  const queryClient = useQueryClient();

  const switchThemeMutation = useSwitchTheme({
    onSuccess: () => {
      toast.success("主题切换成功");
      // 刷新主题列表缓存，更新UI状态
      queryClient.invalidateQueries({ queryKey: themeKeys.all });
    },
    onError: (error: any) => {
      console.error("Theme switch error:", error);
      const errorMessage = error?.response?.data?.message || error?.message;
      if (errorMessage) {
        console.error("Error details:", errorMessage);
      }
      toast.error("主题切换失败，请稍后重试");
    },
  });

  // ===== Event Handlers =====
  const handleSwitchTheme = async (themeId: string, themeType: string) => {
    try {
      await switchThemeMutation.mutateAsync({
        id: themeId,
        theme_type: themeType as "frontend" | "console",
      });
    } catch (error) {
      // 错误已在 onError 回调中处理
    }
  };

  const handleViewConfig = (theme: GetThemeResponse) => {
    setSelectedTheme(theme);
    setConfigDialogOpen(true);
  };

  // ===== Utils =====
  const getPreviewImageUrl = (theme: GetThemeResponse): string | null => {
    if (!theme.preview) return null;
    if (theme.preview.startsWith("http")) return theme.preview;
    const backendBaseUrl = import.meta.env.VITE_API_BASE_URL;
    const cleanPath = theme.preview.startsWith("/")
      ? theme.preview.slice(1)
      : theme.preview;
    return `${backendBaseUrl}/${cleanPath}?theme_type=${theme.type}`;
  };

  // 状态和类型配置 - 使用 useMemo 优化性能
  const badgeConfigs = useMemo(
    () => ({
      status: {
        active: { variant: "default" as const, label: "激活" },
        ready: { variant: "outline" as const, label: "就绪" },
        inactive: { variant: "secondary" as const, label: "未激活" },
        error: { variant: "destructive" as const, label: "错误" },
      },
      type: {
        frontend: { variant: "default" as const, label: "前端" },
        console: { variant: "secondary" as const, label: "后台" },
      },
    }),
    []
  );

  const getStatusBadge = (status: string) =>
    badgeConfigs.status[status as keyof typeof badgeConfigs.status] ||
    badgeConfigs.status.inactive;

  const getTypeBadge = (type: string) =>
    badgeConfigs.type[type as keyof typeof badgeConfigs.type] || {
      variant: "outline" as const,
      label: type,
    };

  // ===== Computed =====
  const filteredThemes = useMemo(() => {
    if (!searchQuery.trim()) return themes;
    const query = searchQuery.toLowerCase();
    return themes.filter(
      (theme) =>
        theme.name.toLowerCase().includes(query) ||
        theme.id.toLowerCase().includes(query) ||
        (theme.description && theme.description.toLowerCase().includes(query))
    );
  }, [themes, searchQuery]);

  // ===== Loading State =====
  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-2 border-primary border-t-transparent"></div>
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      {/* Header */}
      <div className="px-4 py-4 border-b">
        <div className="relative flex-1 sm:w-80">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="搜索主题..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-9 h-10 rounded-full"
          />
        </div>
      </div>

      {/* Content */}
      <div className="flex-1 scrollbar-hover">
        {filteredThemes.length === 0 ? (
          <div className="p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <AlertCircle className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">
                {searchQuery ? "未找到匹配的主题" : "暂无主题"}
              </h3>
              <p className="text-muted-foreground">
                {searchQuery ? "尝试使用其他关键词搜索" : "还没有安装任何主题"}
              </p>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 p-4">
            {filteredThemes.map((theme) => {
              const previewUrl = getPreviewImageUrl(theme);

              return (
                <div
                  key={theme.id}
                  className="group relative bg-card rounded-xl border border-border hover:border-border/80 hover:shadow-md transition-all duration-300 ease-out overflow-hidden"
                >
                  {/* Preview Image */}
                  <div className="relative">
                    <ImageLoader
                      src={previewUrl}
                      alt={`${theme.name} 预览`}
                      aspectRatio="video"
                      fallbackIcon={<Palette className="h-12 w-12" />}
                      className="group-hover:scale-105 transition-transform duration-300 ease-out"
                    />
                    {/* Status indicator overlay */}
                    {theme.status === "active" && (
                      <div className="absolute top-2 right-2 w-3 h-3 bg-green-500 rounded-full border-2 border-white shadow-sm" />
                    )}
                  </div>

                  {/* Content */}
                  <div className="p-4">
                    <div className="flex items-start justify-between gap-2 mb-2">
                      <h3 className="font-semibold text-base line-clamp-1 group-hover:text-primary transition-colors duration-200">
                        {theme.name}
                      </h3>
                      <Badge
                        variant={getStatusBadge(theme.status).variant}
                        className="text-xs shrink-0"
                      >
                        {getStatusBadge(theme.status).label}
                      </Badge>
                    </div>

                    <p className="text-sm text-muted-foreground line-clamp-2 mb-3 min-h-[2.5rem]">
                      {theme.description || "暂无描述"}
                    </p>

                    <div className="flex items-center gap-2 text-xs text-muted-foreground mb-4">
                      <div className="flex items-center gap-1.5">
                        <Badge
                          variant={getTypeBadge(theme.type).variant}
                          className="text-xs"
                        >
                          {getTypeBadge(theme.type).label}
                        </Badge>
                        <span>v{theme.version}</span>
                      </div>
                      {theme.author && (
                        <div className="flex items-center gap-1.5">
                          <div className="w-1.5 h-1.5 rounded-full bg-slate-400" />
                          <span className="truncate">{theme.author}</span>
                        </div>
                      )}
                    </div>

                    {/* Actions */}
                    <div className="flex items-center gap-2">
                      {theme.status !== "active" ? (
                        <Button
                          variant="default"
                          size="sm"
                          className="flex-1 h-8"
                          disabled={switchThemeMutation.isPending}
                          onClick={() =>
                            handleSwitchTheme(theme.id, theme.type)
                          }
                        >
                          {switchThemeMutation.isPending
                            ? "切换中..."
                            : "切换主题"}
                        </Button>
                      ) : (
                        <Button
                          variant="outline"
                          size="sm"
                          className="flex-1 h-8"
                          disabled
                        >
                          当前主题
                        </Button>
                      )}
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button
                            variant="outline"
                            size="sm"
                            className="h-8 w-8 p-0"
                          >
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="w-32">
                          <DropdownMenuItem
                            onClick={() => handleViewConfig(theme)}
                            className="text-sm"
                          >
                            <Eye className="mr-2 h-4 w-4" />
                            查看详情
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* Footer */}
      <div className="px-4 py-3 border-t">
        <p className="text-sm text-muted-foreground">
          共 {filteredThemes.length} 个主题
          {searchQuery && ` · 显示 ${filteredThemes.length} 个匹配结果`}
        </p>
      </div>

      {/* Theme Config Dialog */}
      <Dialog open={configDialogOpen} onOpenChange={setConfigDialogOpen}>
        <DialogContent className="w-[95vw] max-w-md max-h-[85vh] overflow-hidden rounded-2xl">
          {/* Header */}
          <div className="px-4 py-3">
            <DialogTitle className="text-xl font-bold mb-2">
              {selectedTheme?.name}
            </DialogTitle>
            <div className="flex items-center gap-3 text-sm text-muted-foreground">
              <span>v{selectedTheme?.version}</span>
              {selectedTheme?.author && <span>by {selectedTheme.author}</span>}
              {selectedTheme?.is_active && (
                <div className="flex items-center gap-1 text-green-600 font-medium">
                  <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                  <span>当前主题</span>
                </div>
              )}
            </div>
          </div>

          {/* Content */}
          <div className="flex-1 overflow-y-auto px-4">
            {/* Preview Image */}
            {selectedTheme?.preview && (
              <div className="mb-6">
                <div className="w-full h-48 bg-muted rounded-xl overflow-hidden">
                  <img
                    src={
                      getPreviewImageUrl(selectedTheme) || selectedTheme.preview
                    }
                    alt={`${selectedTheme.name} 预览`}
                    className="w-full h-full object-cover"
                    loading="lazy"
                  />
                </div>
              </div>
            )}

            {/* Description */}
            {selectedTheme?.description && (
              <div className="mb-6">
                <p className="text-sm text-muted-foreground leading-relaxed">
                  {selectedTheme.description}
                </p>
              </div>
            )}

            {/* Info List */}
            <div className="space-y-4 mb-6">
              <div className="flex items-center justify-between">
                <span className="text-sm font-medium text-foreground/70">
                  类型
                </span>
                <Badge
                  variant={getTypeBadge(selectedTheme?.type || "").variant}
                  className="text-xs"
                >
                  {getTypeBadge(selectedTheme?.type || "").label}
                </Badge>
              </div>

              <div className="flex items-center justify-between">
                <span className="text-sm font-medium text-foreground/70">
                  状态
                </span>
                <Badge
                  variant={getStatusBadge(selectedTheme?.status || "").variant}
                  className="text-xs"
                >
                  {getStatusBadge(selectedTheme?.status || "").label}
                </Badge>
              </div>

              {selectedTheme?.repository && (
                <div className="flex items-start justify-between gap-4">
                  <span className="text-sm font-medium text-foreground/70 flex-shrink-0">
                    仓库
                  </span>
                  <a
                    href={selectedTheme.repository}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-sm text-blue-600 hover:text-blue-800 hover:underline text-right break-all"
                  >
                    {selectedTheme.repository.replace(/^https?:\/\//, "")}
                  </a>
                </div>
              )}

              <div className="flex items-start justify-between gap-4">
                <span className="text-sm font-medium text-foreground/70 flex-shrink-0">
                  主题 ID
                </span>
                <code className="text-xs bg-muted px-2 py-1 rounded font-mono text-right break-all">
                  {selectedTheme?.id}
                </code>
              </div>

              {selectedTheme?.loaded_at && (
                <div className="flex items-center justify-between">
                  <span className="text-sm font-medium text-foreground/70">
                    加载时间
                  </span>
                  <span className="text-sm text-muted-foreground">
                    {new Date(
                      selectedTheme.loaded_at * 1000
                    ).toLocaleDateString("zh-CN")}
                  </span>
                </div>
              )}
            </div>
          </div>

          {/* Footer */}
          <div className="px-4 py-3">
            <div className="flex gap-3">
              <Button
                variant="outline"
                onClick={() => setConfigDialogOpen(false)}
                className="flex-1"
              >
                关闭
              </Button>
              {selectedTheme && !selectedTheme.is_active && (
                <Button
                  onClick={() => {
                    handleSwitchTheme(selectedTheme.id, selectedTheme.type);
                    setConfigDialogOpen(false);
                  }}
                  disabled={switchThemeMutation.isPending}
                  className="flex-1"
                >
                  {switchThemeMutation.isPending ? "切换中..." : "切换主题"}
                </Button>
              )}
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
