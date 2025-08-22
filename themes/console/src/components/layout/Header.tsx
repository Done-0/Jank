import { useNavigate } from "@tanstack/react-router";

import { Button } from "@/components/ui/button";
import { AUTH_ROUTES } from "@/constants";
import { useUserStore, useAuthStore } from "@/stores";

export function Header() {
  const { user } = useUserStore();
  const { logout } = useAuthStore();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate({ to: AUTH_ROUTES.LOGIN });
  };

  return (
    <header className="h-16 border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="flex h-16 items-center justify-between px-6">
        <div className="flex items-center gap-4">
          <h1 className="text-xl font-semibold">Jank Console</h1>
        </div>

        <div className="flex items-center gap-4">
          {user && (
            <span className="text-sm text-muted-foreground">
              欢迎，{user.nickname || user.email}
            </span>
          )}

          <Button variant="outline" size="sm" onClick={handleLogout}>
            退出登录
          </Button>
        </div>
      </div>
    </header>
  );
}
