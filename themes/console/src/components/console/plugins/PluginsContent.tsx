/**
 * 插件管理内容组件
 * 负责插件列表展示、搜索、启动/停止等核心功能
 */

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Input } from "@/components/ui/input";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Search,
  MoreHorizontal,
  Play,
  Square,
  Settings,
  AlertCircle,
} from "lucide-react";
import { useInstallPlugin, useUninstallPlugin } from "@/hooks/use-plugins";
import type { GetPluginResponse } from "@/types";

interface PluginsContentProps {
  plugins: GetPluginResponse[];
  isLoading: boolean;
}

export function PluginsContent({ plugins, isLoading }: PluginsContentProps) {
  // ===== State =====
  const [searchQuery, setSearchQuery] = useState("");
  const [configDialogOpen, setConfigDialogOpen] = useState(false);
  const [selectedPlugin, setSelectedPlugin] =
    useState<GetPluginResponse | null>(null);

  // ===== Hooks =====
  const installMutation = useInstallPlugin();
  const uninstallMutation = useUninstallPlugin();

  // ===== Event Handlers =====
  const handleStartPlugin = async (pluginId: string) => {
    try {
      await installMutation.mutateAsync({ id: pluginId });
    } catch (error) {
      console.error("启动插件失败:", error);
    }
  };

  const handleStopPlugin = async (pluginId: string) => {
    try {
      await uninstallMutation.mutateAsync({ id: pluginId });
    } catch (error) {
      console.error("停止插件失败:", error);
    }
  };

  const handleConfigurePlugin = (plugin: GetPluginResponse) => {
    setSelectedPlugin(plugin);
    setConfigDialogOpen(true);
  };

  // ===== Utils =====
  const getStatusBadge = (status: string) => {
    const statusConfig = {
      running: { variant: "default" as const, label: "运行中" },
      stopped: { variant: "secondary" as const, label: "已停止" },
      error: { variant: "destructive" as const, label: "错误" },
      available: { variant: "outline" as const, label: "可用" },
      source_only: { variant: "secondary" as const, label: "仅源码" },
    };
    return (
      statusConfig[status as keyof typeof statusConfig] || {
        variant: "secondary" as const,
        label: status,
      }
    );
  };

  const getTypeBadge = (type: string) => {
    const typeConfig = {
      provider: { variant: "default" as const, label: "提供者" },
      filter: { variant: "secondary" as const, label: "过滤器" },
      handler: { variant: "outline" as const, label: "处理器" },
      notifier: { variant: "destructive" as const, label: "通知器" },
    };
    return (
      typeConfig[type as keyof typeof typeConfig] || {
        variant: "secondary" as const,
        label: type,
      }
    );
  };

  // ===== Filtered Data =====
  const filteredPlugins = plugins.filter(
    (plugin) =>
      plugin.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      plugin.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
      plugin.description?.toLowerCase().includes(searchQuery.toLowerCase())
  );

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-primary"></div>
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
            placeholder="搜索插件..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-9 h-10 rounded-full"
          />
        </div>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-y-auto scrollbar-hidden">
        {filteredPlugins.length === 0 ? (
          <div className="p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <AlertCircle className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">
                {searchQuery ? "未找到匹配的插件" : "暂无插件"}
              </h3>
              <p className="text-muted-foreground">
                {searchQuery ? "尝试使用其他关键词搜索" : "还没有安装任何插件"}
              </p>
            </div>
          </div>
        ) : (
          <div className="divide-y">
            {filteredPlugins.map((plugin) => (
              <div
                key={plugin.id}
                className="px-4 py-4 hover:bg-accent/50 transition-colors"
              >
                <div className="flex flex-col gap-3">
                  <div className="flex items-start justify-between gap-3">
                    <div className="flex-1 min-w-0">
                      <h3 className="font-medium text-lg line-clamp-2">
                        {plugin.name}
                      </h3>
                    </div>
                    <Badge
                      variant={getStatusBadge(plugin.status).variant}
                      className="shrink-0"
                    >
                      {getStatusBadge(plugin.status).label}
                    </Badge>
                  </div>

                  <div className="min-h-[1.25rem]">
                    <p className="text-sm text-muted-foreground line-clamp-2">
                      {plugin.description || "暂无描述"}
                    </p>
                  </div>

                  <div className="flex items-center justify-between gap-3">
                    <div className="flex items-center gap-2 text-xs text-muted-foreground">
                      <div className="flex items-center gap-1.5">
                        <Badge
                          variant={getTypeBadge(plugin.type).variant}
                          className="text-xs"
                        >
                          {getTypeBadge(plugin.type).label}
                        </Badge>
                        <span>v{plugin.version}</span>
                      </div>
                      {plugin.author && (
                        <div className="flex items-center gap-1.5">
                          <div className="w-1.5 h-1.5 rounded-full bg-slate-400" />
                          <span>{plugin.author}</span>
                        </div>
                      )}
                    </div>

                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button
                          variant="ghost"
                          size="sm"
                          className="h-8 w-8 p-0 rounded-full"
                        >
                          <MoreHorizontal className="h-4 w-4" />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent
                        align="end"
                        className="w-36 rounded-xl"
                      >
                        <DropdownMenuItem
                          onClick={() => handleConfigurePlugin(plugin)}
                          className="py-2.5"
                        >
                          <Settings className="mr-2 h-4 w-4" />
                          配置插件
                        </DropdownMenuItem>
                        {plugin.status === "running" ? (
                          <DropdownMenuItem
                            disabled={uninstallMutation.isPending}
                            onClick={() => handleStopPlugin(plugin.id)}
                            className="py-2.5"
                          >
                            <Square className="mr-2 h-4 w-4" />
                            {uninstallMutation.isPending
                              ? "停止中..."
                              : "停止插件"}
                          </DropdownMenuItem>
                        ) : (
                          <DropdownMenuItem
                            disabled={installMutation.isPending}
                            onClick={() => handleStartPlugin(plugin.id)}
                            className="py-2.5"
                          >
                            <Play className="mr-2 h-4 w-4" />
                            {installMutation.isPending
                              ? "启动中..."
                              : "启动插件"}
                          </DropdownMenuItem>
                        )}
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Footer */}
      <div className="px-4 py-3 border-t">
        <p className="text-sm text-muted-foreground">
          共 {filteredPlugins.length} 个插件
          {searchQuery && ` · 显示 ${filteredPlugins.length} 个匹配结果`}
        </p>
      </div>

      {/* Plugin Config Dialog */}
      <Dialog open={configDialogOpen} onOpenChange={setConfigDialogOpen}>
        <DialogContent className="w-[95vw] max-w-md max-h-[85vh] overflow-hidden rounded-2xl">
          {/* Header */}
          <div className="px-4 py-3">
            <DialogTitle className="text-xl font-bold mb-2">
              {selectedPlugin?.name}
            </DialogTitle>
            <div className="flex items-center gap-3 text-sm text-muted-foreground">
              <span>v{selectedPlugin?.version}</span>
              {selectedPlugin?.author && (
                <span>by {selectedPlugin.author}</span>
              )}
              {selectedPlugin?.status === "running" && (
                <div className="flex items-center gap-1 text-green-600 font-medium">
                  <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                  <span>运行中</span>
                </div>
              )}
            </div>
          </div>

          {/* Content */}
          <div className="flex-1 overflow-y-auto px-4">
            {/* Description */}
            {selectedPlugin?.description && (
              <div className="mb-6">
                <p className="text-sm text-muted-foreground leading-relaxed">
                  {selectedPlugin.description}
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
                  variant={getTypeBadge(selectedPlugin?.type || "").variant}
                  className="text-xs"
                >
                  {getTypeBadge(selectedPlugin?.type || "").label}
                </Badge>
              </div>

              <div className="flex items-center justify-between">
                <span className="text-sm font-medium text-foreground/70">
                  状态
                </span>
                <Badge
                  variant={getStatusBadge(selectedPlugin?.status || "").variant}
                  className="text-xs"
                >
                  {getStatusBadge(selectedPlugin?.status || "").label}
                </Badge>
              </div>

              {selectedPlugin?.repository && (
                <div className="flex items-start justify-between gap-4">
                  <span className="text-sm font-medium text-foreground/70 flex-shrink-0">
                    仓库
                  </span>
                  <a
                    href={selectedPlugin.repository}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-sm text-blue-600 hover:text-blue-800 hover:underline text-right break-all"
                  >
                    {selectedPlugin.repository.replace(/^https?:\/\//, "")}
                  </a>
                </div>
              )}

              <div className="flex items-center justify-between">
                <span className="text-sm font-medium text-foreground/70">
                  自动启动
                </span>
                <span className="text-sm text-muted-foreground">
                  {selectedPlugin?.auto_start ? "启用" : "禁用"}
                </span>
              </div>

              <div className="flex items-start justify-between gap-4">
                <span className="text-sm font-medium text-foreground/70 flex-shrink-0">
                  插件 ID
                </span>
                <code className="text-xs bg-muted px-2 py-1 rounded font-mono text-right break-all">
                  {selectedPlugin?.id}
                </code>
              </div>
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
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
