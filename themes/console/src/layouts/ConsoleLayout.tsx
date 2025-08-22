import { ReactNode } from "react";
import { useLocation } from "@tanstack/react-router";

import { SidebarProvider } from "@/components/ui/sidebar";
import { AppHeader } from "@/components/layout/AppHeader";
import { AppSidebar } from "@/components/layout/AppSidebar";
import { useUserStore, useAuthStore } from "@/stores";

interface ConsoleLayoutProps {
  children: ReactNode;
}

export function ConsoleLayout({ children }: ConsoleLayoutProps) {
  const location = useLocation();
  const { user } = useUserStore();
  const { logout } = useAuthStore();

  const handleLogout = async () => {
    await logout();
  };

  return (
    <SidebarProvider>
      <div className="flex h-screen w-full">
        <AppSidebar
          activeTab={location.pathname}
          user={user}
          onLogout={handleLogout}
        />
        <div className="flex flex-1 flex-col overflow-hidden">
          <AppHeader activeTab={location.pathname} />
          <main className="flex-1 overflow-auto p-6">{children}</main>
        </div>
      </div>
    </SidebarProvider>
  );
}
