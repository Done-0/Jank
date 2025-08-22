/**
 * RBAC权限管理页面
 */

import { useState } from "react";

import { RBACSidebar } from "@/components/console/rbac/RBACsidebar";
import { RBACMainContent } from "@/components/console/rbac/RBACMainContent";

import { useRoles, usePermissions } from "@/hooks/use-rbac";

export function ConsoleRBACPage() {
  // ===== 状态管理 =====
  const [selectedSection, setSelectedSection] = useState<
    "roles" | "permissions" | "role-permissions" | "user-roles"
  >("roles");
  const [searchQuery, setSearchQuery] = useState("");

  // ===== 数据获取 =====
  const { data: rolesData } = useRoles();
  const { data: permissionsData } = usePermissions();

  // ===== 计算数据 =====
  const roles = rolesData?.list || [];
  const permissions = permissionsData?.list || [];

  const stats = {
    totalRoles: roles.length,
    totalPermissions: permissions.length,
    totalUserRoles: 0, // TODO: 实际统计用户角色数量
  };

  // ===== 事件处理 =====
  const handleSectionChange = (
    section: "roles" | "permissions" | "role-permissions" | "user-roles"
  ) => setSelectedSection(section);
  const handleSearchChange = (query: string) => setSearchQuery(query);

  // ===== 渲染 =====
  return (
    <div className="flex h-full">
      <div className="hidden md:block">
        <RBACSidebar
          selectedSection={selectedSection}
          stats={stats}
          onSectionChange={handleSectionChange}
        />
      </div>

      <div className="flex-1 flex flex-col h-full">
        <RBACMainContent
          selectedSection={selectedSection}
          searchTerm={searchQuery}
          onSearchChange={handleSearchChange}
        />
      </div>
    </div>
  );
}
