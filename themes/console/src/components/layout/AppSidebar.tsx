/**
 * 侧边栏组件
 */

import {
  LayoutDashboard,
  Users,
  Shield,
  Palette,
  Puzzle,
  FileText,
  FolderOpen,
} from "lucide-react";
import { Link } from "@tanstack/react-router";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
  useSidebar,
} from "@/components/ui/sidebar";
import type { GetProfileResponse } from "@/types";
import { CONSOLE_ROUTES } from "@/constants";
import { ConsoleUserMenu } from "@/components/layout/ConsoleUserMenu";

interface AppSidebarProps {
  activeTab: string;
  user: GetProfileResponse | null;
  onLogout: () => void;
}

export function AppSidebar({ activeTab, user, onLogout }: AppSidebarProps) {
  const { isMobile, setOpenMobile } = useSidebar();

  // 简洁的分组导航配置
  const menuGroups = [
    {
      label: "概览",
      items: [
        {
          id: CONSOLE_ROUTES.ROOT,
          label: "控制台",
          icon: LayoutDashboard,
          route: CONSOLE_ROUTES.ROOT,
        },
      ],
    },
    {
      label: "内容",
      items: [
        {
          id: CONSOLE_ROUTES.POSTS,
          label: "文章",
          icon: FileText,
          route: CONSOLE_ROUTES.POSTS,
        },
        {
          id: CONSOLE_ROUTES.CATEGORIES,
          label: "分类",
          icon: FolderOpen,
          route: CONSOLE_ROUTES.CATEGORIES,
        },
      ],
    },
    {
      label: "权限",
      items: [
        {
          id: CONSOLE_ROUTES.USERS,
          label: "用户",
          icon: Users,
          route: CONSOLE_ROUTES.USERS,
        },
        {
          id: CONSOLE_ROUTES.RBAC,
          label: "角色",
          icon: Shield,
          route: CONSOLE_ROUTES.RBAC,
        },
      ],
    },
    {
      label: "系统",
      items: [
        {
          id: CONSOLE_ROUTES.THEMES,
          label: "主题",
          icon: Palette,
          route: CONSOLE_ROUTES.THEMES,
        },
        {
          id: CONSOLE_ROUTES.PLUGINS,
          label: "插件",
          icon: Puzzle,
          route: CONSOLE_ROUTES.PLUGINS,
        },
      ],
    },
  ] as const;

  // 移动端点击链接后自动关闭侧边栏
  const handleLinkClick = () => {
    if (isMobile) {
      setOpenMobile(false);
    }
  };

  return (
    <Sidebar variant="inset">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild className="hover:bg-accent/50">
              <Link to="/">
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-primary text-primary-foreground">
                  <span className="font-bold text-sm">J</span>
                </div>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-semibold">Jank</span>
                  <span className="truncate text-xs text-muted-foreground">
                    管理控制台
                  </span>
                </div>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent>
        {menuGroups.map((group, groupIndex) => (
          <SidebarGroup key={group.label} className="py-0">
            {groupIndex > 0 && <div className="mx-3 my-1 h-px bg-border/30" />}
            <SidebarGroupContent>
              <SidebarMenu className="gap-0.5">
                {group.items.map((item) => {
                  const Icon = item.icon;
                  return (
                    <SidebarMenuItem key={item.id}>
                      <SidebarMenuButton
                        asChild
                        tooltip={item.label}
                        isActive={activeTab === item.id}
                        className="h-9 hover:bg-accent/50 data-[active=true]:bg-accent data-[active=true]:text-accent-foreground transition-colors"
                      >
                        <Link to={item.route} onClick={handleLinkClick}>
                          <Icon className="h-4 w-4" />
                          <span className="font-medium">{item.label}</span>
                        </Link>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  );
                })}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>

      <SidebarFooter>
        <SidebarMenu>
          <ConsoleUserMenu user={user} onLogout={onLogout} />
        </SidebarMenu>
      </SidebarFooter>

      <SidebarRail />
    </Sidebar>
  );
}
