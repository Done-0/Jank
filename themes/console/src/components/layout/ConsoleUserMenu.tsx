import { useState } from "react";
import { ChevronsUpDown, LogOut, Moon, Sun } from "lucide-react";
import { GetProfileResponse } from "@/types";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { SidebarMenuButton, SidebarMenuItem } from "@/components/ui/sidebar";
import { useTheme } from "@/components/theme";

interface ConsoleUserMenuProps {
  user: GetProfileResponse | null;
  onLogout: () => void;
}

export function ConsoleUserMenu({ user, onLogout }: ConsoleUserMenuProps) {
  const userInitial = user?.nickname?.charAt(0)?.toUpperCase();
  const userName = user?.nickname;
  // 动态显示用户角色，不硬编码判断
  const userRole = user?.roles?.length ? user.roles.join(", ") : "普通用户";
  const { theme, toggleTheme } = useTheme();
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  const isDark = theme === "dark";

  return (
    <SidebarMenuItem>
      <DropdownMenu open={isDropdownOpen} onOpenChange={setIsDropdownOpen}>
        <DropdownMenuTrigger asChild>
          <SidebarMenuButton
            size="lg"
            className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
          >
            <Avatar className="h-8 w-8 rounded-lg">
              {user?.avatar && <AvatarImage src={user.avatar} alt={userName} />}
              <AvatarFallback className="rounded-lg">
                {userInitial}
              </AvatarFallback>
            </Avatar>
            <div className="grid flex-1 text-left text-sm leading-tight">
              <span className="truncate font-semibold">{userName}</span>
              <span className="truncate text-xs text-muted-foreground">
                {userRole}
              </span>
            </div>
            <ChevronsUpDown className="ml-auto size-4" />
          </SidebarMenuButton>
        </DropdownMenuTrigger>
        <DropdownMenuContent
          className="w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg"
          side="bottom"
          align="end"
          sideOffset={4}
        >
          <DropdownMenuLabel className="p-0 font-normal">
            <div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
              <Avatar className="h-8 w-8 rounded-lg">
                {user?.avatar && (
                  <AvatarImage src={user.avatar} alt={userName} />
                )}
                <AvatarFallback className="rounded-lg">
                  {userInitial}
                </AvatarFallback>
              </Avatar>
              <div className="grid flex-1 text-left text-sm leading-tight">
                <span className="truncate font-semibold">{userName}</span>
                <span className="truncate text-xs text-muted-foreground">
                  {userRole}
                </span>
              </div>
            </div>
          </DropdownMenuLabel>

          <DropdownMenuSeparator />
          <DropdownMenuItem
            className="gap-2 cursor-pointer hover:bg-accent/30 transition-colors"
            onClick={() => {
              setIsDropdownOpen(false);
              toggleTheme();
            }}
          >
            {isDark ? (
              <>
                <Sun className="size-4" />
                明亮模式
              </>
            ) : (
              <>
                <Moon className="size-4" />
                深色模式
              </>
            )}
          </DropdownMenuItem>
          <DropdownMenuItem
            className="gap-2 cursor-pointer hover:bg-accent/30 transition-colors"
            onClick={() => {
              setIsDropdownOpen(false);
              onLogout();
            }}
          >
            <LogOut className="size-4" />
            退出登录
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </SidebarMenuItem>
  );
}
