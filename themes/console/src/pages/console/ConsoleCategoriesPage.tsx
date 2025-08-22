/**
 * 分类管理页面
 */

import { useAllCategories } from "@/hooks/use-categories";
import { CategoriesContent } from "@/components/console/categories/CategoriesContent";

export function ConsoleCategoriesPage() {
  const { data: categories = [], isLoading, error } = useAllCategories();

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-destructive mb-2">加载分类失败</p>
          <p className="text-sm text-muted-foreground">{error.message}</p>
        </div>
      </div>
    );
  }

  return <CategoriesContent categories={categories} isLoading={isLoading} />;
}
