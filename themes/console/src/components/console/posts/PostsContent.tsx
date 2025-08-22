/**
 * 文章列表内容组件
 * 负责文章列表展示、搜索、分页等核心功能
 */

import React, { useState, useMemo } from "react";

import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Input } from "@/components/ui/input";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import {
  Search,
  Plus,
  Edit,
  Trash2,
  Eye,
  ChevronLeft,
  ChevronRight,
  FileText,
  MoreHorizontal,
} from "lucide-react";

import type { PostStatus } from "@/constants/post";
import type { PostItem } from "@/types/post";
import type { CategoryItem } from "@/types/category";

interface PostsContentProps {
  // 状态数据
  selectedStatus: PostStatus | null;
  selectedCategory: string | null;
  searchQuery: string;
  categories: CategoryItem[];
  allPosts: PostItem[];
  isLoading: boolean;

  // 操作方法
  onStatusChange: (status: PostStatus | null) => void;
  onCategoryChange: (categoryId: string | null) => void;
  onSearchChange: (query: string) => void;
  onNewPost: () => void;
  onEditPost: (postId: string) => void;
}

export function PostsContent({
  selectedStatus,
  selectedCategory,
  searchQuery,
  categories,
  allPosts,
  isLoading,
  onSearchChange,
  onNewPost,
  onStatusChange,
  onEditPost,
}: PostsContentProps) {
  // ===== 状态管理 =====
  const [currentPage, setCurrentPage] = useState(1);
  const pageSize = 10;

  // ===== 副作用 =====
  React.useEffect(() => {
    setCurrentPage(1);
  }, [selectedStatus, selectedCategory, searchQuery]);

  // ===== 工具函数 =====
  const getCategoryName = (categoryId: string): string => {
    if (!categoryId) return "暂未分类";
    const category = categories.find((cat) => cat.id === categoryId);
    return category?.name || "未知分类";
  };

  const getStatusBadge = (status: PostStatus) => {
    const statusConfig: Record<
      PostStatus,
      { variant: "default" | "secondary" | "outline"; label: string }
    > = {
      published: { variant: "default", label: "已发布" },
      draft: { variant: "secondary", label: "草稿" },
      archived: { variant: "outline", label: "已归档" },
      private: { variant: "outline", label: "私有" },
    };
    return statusConfig[status] || { variant: "outline", label: status };
  };

  const formatDate = (dateStr: string) => {
    return new Date(dateStr).toLocaleDateString();
  };

  const getCategoryColor = (categoryId?: string) => {
    if (!categoryId) return "bg-muted-foreground";
    const colors = [
      "bg-blue-500",
      "bg-green-500",
      "bg-yellow-500",
      "bg-red-500",
      "bg-purple-500",
      "bg-pink-500",
      "bg-indigo-500",
      "bg-orange-500",
    ];
    return colors[parseInt(categoryId) % colors.length];
  };

  // ===== 分页器工具函数 =====
  const createPageButton = (pageNum: number, isActive = false) => (
    <Button
      key={pageNum}
      variant={isActive ? "default" : "outline"}
      size="sm"
      className="h-8 w-8 p-0"
      onClick={() => setCurrentPage(pageNum)}
    >
      {pageNum}
    </Button>
  );

  const generatePageNumbers = () => {
    const pages = [];
    const isMobile = typeof window !== "undefined" && window.innerWidth < 640;
    const maxPages = isMobile ? 3 : 5;

    if (totalPages <= maxPages) {
      // 显示所有页码
      for (let i = 1; i <= totalPages; i++) {
        pages.push(createPageButton(i, i === currentPage));
      }
    } else {
      // 省略号逻辑
      pages.push(createPageButton(1, currentPage === 1));

      if (currentPage > 3) {
        pages.push(
          <span key="ellipsis1" className="px-1 text-muted-foreground text-xs">
            …
          </span>
        );
      }

      // 当前页周围的页码
      const start = Math.max(2, currentPage - 1);
      const end = Math.min(totalPages - 1, currentPage + 1);

      for (let i = start; i <= end; i++) {
        if (i !== 1 && i !== totalPages) {
          pages.push(createPageButton(i, i === currentPage));
        }
      }

      if (currentPage < totalPages - 2) {
        pages.push(
          <span key="ellipsis2" className="px-1 text-muted-foreground text-xs">
            …
          </span>
        );
      }

      if (totalPages > 1) {
        pages.push(createPageButton(totalPages, currentPage === totalPages));
      }
    }

    return pages;
  };

  // ===== 数据计算 =====
  const filteredPosts = useMemo(() => {
    let posts = allPosts || [];

    if (selectedCategory) {
      posts = posts.filter((post) => post.category_id === selectedCategory);
    }

    if (searchQuery.trim()) {
      const query = searchQuery.toLowerCase();
      posts = posts.filter(
        (post) =>
          post.title.toLowerCase().includes(query) ||
          (post.description && post.description.toLowerCase().includes(query))
      );
    }

    return posts;
  }, [allPosts, selectedCategory, searchQuery]);

  const { posts, total, totalPages, hasPrevPage, hasNextPage } = useMemo(() => {
    const total = filteredPosts.length;
    const totalPages = Math.ceil(total / pageSize);
    const startIndex = (currentPage - 1) * pageSize;
    const posts = filteredPosts.slice(startIndex, startIndex + pageSize);

    return {
      posts,
      total,
      totalPages,
      hasPrevPage: currentPage > 1,
      hasNextPage: currentPage < totalPages,
    };
  }, [filteredPosts, currentPage, pageSize]);

  // ===== 工具函数 =====
  const statusTabs = [
    { value: null, label: "全部" },
    { value: "published", label: "已发布" },
    { value: "draft", label: "草稿" },
    { value: "archived", label: "已归档" },
  ] as const;

  const getTabClassName = (status: PostStatus | null) => {
    const isActive = selectedStatus === status;
    return `flex-1 py-3 text-sm font-medium border-b-2 transition-colors ${
      isActive
        ? "border-primary text-primary"
        : "border-transparent text-muted-foreground hover:text-foreground"
    }`;
  };

  return (
    <div className="flex-1 flex flex-col h-full">
      {/* 移动端状态导航 */}
      <div className="md:hidden border-b bg-background">
        <div className="flex">
          {statusTabs.map(({ value, label }) => (
            <button
              key={value || "all"}
              className={getTabClassName(value)}
              onClick={() => onStatusChange(value)}
            >
              {label}
            </button>
          ))}
        </div>
      </div>

      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索文章..."
              className="pl-9 h-10 rounded-full"
              value={searchQuery}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
          <Button
            onClick={onNewPost}
            className="h-10 px-4 shrink-0 rounded-full"
          >
            <Plus className="mr-2 h-4 w-4" />
            <span className="hidden sm:inline">新建文章</span>
            <span className="sm:hidden">新建</span>
          </Button>
        </div>
      </div>

      {/* 文章列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {isLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : posts.length > 0 ? (
          <>
            <div className="flex-1 overflow-y-scroll scrollbar-hidden">
              <div className="divide-y divide-border">
                {posts.map((post) => (
                  <div
                    key={post.id}
                    className="px-4 py-4 hover:bg-accent/50 transition-colors cursor-pointer"
                    onClick={() => onEditPost(post.id)}
                  >
                    <div className="flex flex-col gap-3">
                      <div className="flex items-start justify-between gap-3">
                        <h3 className="font-medium text-lg line-clamp-2 flex-1 min-w-0 hover:text-primary transition-colors">
                          {post.title}
                        </h3>
                        <Badge
                          variant={
                            getStatusBadge(post.status as PostStatus).variant
                          }
                          className="shrink-0"
                        >
                          {getStatusBadge(post.status as PostStatus).label}
                        </Badge>
                      </div>

                      <div className="min-h-[1.25rem]">
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          {post.description || "暂无描述"}
                        </p>
                      </div>

                      <div className="flex items-center justify-between gap-3">
                        <div className="flex items-center gap-2 text-xs text-muted-foreground">
                          <div className="flex items-center gap-1.5">
                            <div
                              className={`w-1.5 h-1.5 rounded-full ${getCategoryColor(
                                post.category_id
                              )}`}
                            />
                            <span>{getCategoryName(post.category_id)}</span>
                          </div>
                          <div className="flex items-center gap-1.5">
                            <div className="w-1.5 h-1.5 rounded-full bg-slate-400" />
                            <span>{formatDate(post.created_at)}</span>
                          </div>
                        </div>

                        <DropdownMenu>
                          <DropdownMenuTrigger asChild>
                            <Button
                              variant="ghost"
                              size="sm"
                              className="h-8 w-8 p-0 rounded-full"
                              onClick={(e: React.MouseEvent) => {
                                e.stopPropagation(); // 防止触发父级点击事件
                              }}
                            >
                              <MoreHorizontal className="h-4 w-4" />
                            </Button>
                          </DropdownMenuTrigger>
                          <DropdownMenuContent align="end" className="w-40">
                            <DropdownMenuItem
                              className="cursor-pointer"
                              onClick={(e: React.MouseEvent) => {
                                e.stopPropagation();
                                onEditPost(post.id);
                              }}
                            >
                              <Edit className="mr-2 h-4 w-4" />
                              编辑文章
                            </DropdownMenuItem>
                            <DropdownMenuItem className="cursor-pointer">
                              <Eye className="mr-2 h-4 w-4" />
                              查看文章
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem className="cursor-pointer text-destructive focus:text-destructive">
                              <Trash2 className="mr-2 h-4 w-4" />
                              删除文章
                            </DropdownMenuItem>
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>

            {/* 分页控件 */}
            <div className="flex flex-col sm:flex-row items-center justify-between px-3 sm:px-4 lg:px-5 py-3 border-t gap-3 sm:gap-0">
              <div className="text-sm text-muted-foreground order-2 sm:order-1">
                共 {total} 条记录
              </div>

              <div className="flex items-center gap-1 order-1 sm:order-2">
                {/* 上一页按钮 */}
                <Button
                  variant="outline"
                  size="sm"
                  className="h-8 w-8 p-0"
                  onClick={() =>
                    setCurrentPage((prev) => Math.max(1, prev - 1))
                  }
                  disabled={!hasPrevPage}
                >
                  <ChevronLeft className="h-4 w-4" />
                </Button>

                {/* 页码按钮组 */}
                <div className="flex items-center gap-1 mx-2">
                  {generatePageNumbers()}
                </div>

                {/* 下一页按钮 */}
                <Button
                  variant="outline"
                  size="sm"
                  className="h-8 w-8 p-0"
                  onClick={() =>
                    setCurrentPage((prev) => Math.min(totalPages, prev + 1))
                  }
                  disabled={!hasNextPage}
                >
                  <ChevronRight className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </>
        ) : (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <FileText className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无文章</h3>
              <p className="text-muted-foreground mb-4">
                {selectedStatus === null
                  ? "暂无任何文章内容"
                  : selectedCategory
                  ? `在"${getCategoryName(selectedCategory)}"分类下暂无${
                      selectedStatus === "published"
                        ? "已发布"
                        : selectedStatus === "draft"
                        ? "草稿"
                        : "已归档"
                    }文章`
                  : searchQuery
                  ? `未找到包含"${searchQuery}"的文章`
                  : `暂无${
                      selectedStatus === "published"
                        ? "已发布"
                        : selectedStatus === "draft"
                        ? "草稿"
                        : "已归档"
                    }文章`}
              </p>
              <p className="text-sm text-muted-foreground">
                <span className="hidden sm:inline">
                  点击上方&ldquo;新建文章&rdquo;按钮开始创作
                </span>
                <span className="sm:hidden">点击上方按钮开始创作</span>
              </p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
