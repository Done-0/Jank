import { create } from "zustand";

export interface UIState {
  // UI 状态
  sidebarOpen: boolean;
  globalLoading: boolean;
  // 设置侧边栏状态
  setSidebarOpen: (open: boolean) => void;
  // 切换侧边栏
  toggleSidebar: () => void;
  // 设置全局加载状态
  setGlobalLoading: (loading: boolean) => void;
  // 重置UI状态（退出登录时使用）
  resetUIState: () => void;
}

export const useUIStore = create<UIState>((set, get) => ({
  // 初始状态
  sidebarOpen: true,
  globalLoading: false,

  // 设置侧边栏状态
  setSidebarOpen: (open: boolean) => {
    set({ sidebarOpen: open });
  },

  // 切换侧边栏
  toggleSidebar: () => {
    set({ sidebarOpen: !get().sidebarOpen });
  },

  // 设置全局加载状态
  setGlobalLoading: (loading: boolean) => {
    set({ globalLoading: loading });
  },

  // 重置UI状态（退出登录时使用）
  resetUIState: () => {
    set({
      sidebarOpen: true,
      globalLoading: false,
    });
  },
}));
