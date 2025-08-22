import { FileText, Edit, Eye, Archive } from "lucide-react";

import { CATEGORY_COLORS } from "@/constants";
import type { PostStatus } from "@/constants/post";
import type { CategoryItem } from "@/types/category";

interface PostsSidebarProps {
  selectedStatus: PostStatus | null;
  selectedCategory: string | null;
  stats: {
    total: number;
    published: number;
    draft: number;
    archived: number;
  };
  categories: CategoryItem[];
  onStatusChange: (status: PostStatus | null) => void;
  onCategoryChange: (category: string | null) => void;
  className?: string;
}

export function PostsSidebar({
  selectedStatus,
  selectedCategory,
  stats,
  categories,
  onStatusChange,
  onCategoryChange,
}: PostsSidebarProps) {
  // 获取分类颜色
  const getCategoryColor = (index: number) => {
    return CATEGORY_COLORS[index % CATEGORY_COLORS.length];
  };

  const statusItems = [
    { status: null, icon: FileText, label: "全部", count: stats.total },
    {
      status: "published" as PostStatus,
      icon: Eye,
      label: "已发布",
      count: stats.published,
    },
    {
      status: "draft" as PostStatus,
      icon: Edit,
      label: "草稿",
      count: stats.draft,
    },
    {
      status: "archived" as PostStatus,
      icon: Archive,
      label: "已归档",
      count: stats.archived,
    },
  ];

  return (
    <div className="w-64 lg:w-72 border-r bg-background hidden md:flex">
      <div className="h-full flex flex-col w-full">
        {/* 状态筛选 */}
        <div className="px-4 py-4">
          <div className="space-y-0.5">
            {statusItems.map(({ status, icon: Icon, label, count }) => (
              <button
                key={status || "all"}
                onClick={() => onStatusChange(status)}
                className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-full text-left transition-colors ${
                  selectedStatus === status
                    ? "bg-accent text-accent-foreground font-medium"
                    : "hover:bg-accent/50 text-foreground"
                }`}
              >
                <Icon className="h-5 w-5 flex-shrink-0" />
                <span className="flex-1 text-sm">{label}</span>
                <span className="text-xs text-muted-foreground font-medium">
                  {count}
                </span>
              </button>
            ))}
          </div>
        </div>

        {/* 分类筛选 */}
        <div className="flex-1 flex flex-col min-h-0 border-t">
          <div className="flex-1 overflow-y-auto">
            <div className="px-4 py-4 space-y-0.5">
              <button
                onClick={() => onCategoryChange(null)}
                className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-full text-left transition-colors ${
                  selectedCategory === null
                    ? "bg-accent text-accent-foreground font-medium"
                    : "hover:bg-accent/50 text-foreground"
                }`}
              >
                <div className="w-5 h-5 flex items-center justify-center">
                  <div className="w-2 h-2 rounded-full bg-muted-foreground" />
                </div>
                <span className="flex-1 text-sm">全部分类</span>
              </button>

              {categories.length > 0 ? (
                categories
                  .filter((category) => category.is_active)
                  .map((category, index) => (
                    <button
                      key={category.id}
                      onClick={() => onCategoryChange(category.id)}
                      className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-full text-left transition-colors ${
                        selectedCategory === category.id
                          ? "bg-accent text-accent-foreground font-medium"
                          : "hover:bg-accent/50 text-foreground"
                      }`}
                    >
                      <div className="w-5 h-5 flex items-center justify-center">
                        <div
                          className={`w-2 h-2 rounded-full ${getCategoryColor(
                            index
                          )}`}
                        />
                      </div>
                      <span className="flex-1 text-sm truncate">
                        {category.name}
                      </span>
                    </button>
                  ))
              ) : (
                <div className="px-3 py-2 text-sm text-muted-foreground">
                  暂无分类
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
