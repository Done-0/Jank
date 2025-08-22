import { Shield, Key, UserCheck, Settings } from "lucide-react";

interface RBACStats {
  totalRoles: number;
  totalPermissions: number;
  totalUserRoles: number;
}

interface RBACSidebarProps {
  selectedSection: "roles" | "permissions" | "role-permissions" | "user-roles";
  stats: RBACStats;
  onSectionChange: (
    section: "roles" | "permissions" | "role-permissions" | "user-roles"
  ) => void;
  className?: string;
}

export function RBACSidebar({
  selectedSection,
  stats,
  onSectionChange,
}: RBACSidebarProps) {
  const sectionItems = [
    {
      section: "roles" as const,
      icon: Shield,
      label: "角色管理",
      count: stats.totalRoles,
    },
    {
      section: "permissions" as const,
      icon: Key,
      label: "权限管理",
      count: stats.totalPermissions,
    },
    {
      section: "role-permissions" as const,
      icon: Settings,
      label: "角色权限",
      count: stats.totalRoles,
    },
    {
      section: "user-roles" as const,
      icon: UserCheck,
      label: "用户角色",
      count: stats.totalUserRoles,
    },
  ];

  return (
    <div className="w-64 lg:w-72 border-r bg-background hidden md:flex">
      <div className="h-full flex flex-col w-full">
        <div className="px-4 py-4">
          <div className="space-y-0.5">
            {sectionItems.map(({ section, icon: Icon, label, count }) => (
              <button
                key={section}
                onClick={() => onSectionChange(section)}
                className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-full text-left transition-colors ${
                  selectedSection === section
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
      </div>
    </div>
  );
}
