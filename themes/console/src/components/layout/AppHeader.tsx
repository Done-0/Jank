/**
 * 应用头部组件
 */

import { SidebarTrigger } from "@/components/ui/sidebar";
import { ROUTE_TITLES } from "@/constants";

interface AppHeaderProps {
  activeTab: string;
}

export function AppHeader({ activeTab }: AppHeaderProps) {
  const title =
    ROUTE_TITLES[activeTab as keyof typeof ROUTE_TITLES] || "控制台";

  return (
    <header className="sticky top-0 z-50 w-full bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="flex h-14 items-center px-4">
        <div className="flex items-center gap-3">
          <SidebarTrigger className="h-8 w-8 hover:bg-accent/50 transition-colors" />
          <h1 className="text-base font-medium text-foreground">{title}</h1>
        </div>
      </div>
    </header>
  );
}
