/**
 * 应用状态管理 - Zustand
 * 管理 Jank 项目应用级别的状态
 */

import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface AppState {
  // 应用基础状态
  currentSessionId: string;
  isInitialized: boolean;

  // 用户偏好设置
  language: "zh" | "en";

  // 通知和提示
  notifications: AppNotification[];

  // 应用配置
  appConfig: AppConfig | null;

  // Actions
  setCurrentSessionId: (sessionId: string) => void;
  setIsInitialized: (initialized: boolean) => void;
  setLanguage: (language: "zh" | "en") => void;
  addNotification: (
    notification: Omit<AppNotification, "id" | "timestamp">
  ) => void;
  removeNotification: (id: string) => void;
  clearNotifications: () => void;
  setAppConfig: (config: AppConfig) => void;
  clearAllData: () => void;
}

// 通知类型
export interface AppNotification {
  id: string;
  type: "success" | "error" | "warning" | "info";
  title: string;
  message: string;
  timestamp: number;
  autoClose?: boolean;
}

// 应用配置类型
export interface AppConfig {
  siteName: string;
  version: string;
  features: {
    enableRegistration: boolean;
    enableComments: boolean;
    enablePlugins: boolean;
  };
}

export const useAppStore = create<AppState>()(
  persist(
    (set, get) => ({
      // 初始状态
      currentSessionId: "",
      isInitialized: false,
      language: "zh",
      notifications: [],
      appConfig: null,

      // Actions
      setCurrentSessionId: (sessionId: string) => {
        set({ currentSessionId: sessionId });
      },

      setIsInitialized: (initialized: boolean) => {
        set({ isInitialized: initialized });
      },

      setLanguage: (language: "zh" | "en") => {
        set({ language });
      },

      addNotification: (notification) => {
        const newNotification: AppNotification = {
          ...notification,
          id: Date.now().toString() + Math.random().toString(36).substr(2, 9),
          timestamp: Date.now(),
        };

        const { notifications } = get();
        set({ notifications: [...notifications, newNotification] });

        // 如果设置了自动关闭，5秒后自动移除
        if (notification.autoClose !== false) {
          setTimeout(() => {
            const { notifications: currentNotifications } = get();
            set({
              notifications: currentNotifications.filter(
                (n) => n.id !== newNotification.id
              ),
            });
          }, 5000);
        }
      },

      removeNotification: (id: string) => {
        const { notifications } = get();
        set({ notifications: notifications.filter((n) => n.id !== id) });
      },

      clearNotifications: () => {
        set({ notifications: [] });
      },

      setAppConfig: (config: AppConfig) => {
        set({ appConfig: config });
      },

      clearAllData: () => {
        set({
          currentSessionId: "",
          isInitialized: false,
          language: "zh",
          notifications: [],
          appConfig: null,
        });
      },
    }),
    // 持久化配置
    {
      name: "jank-app-store",
      partialize: (state) => ({
        currentSessionId: state.currentSessionId,
        language: state.language,
      }),
    }
  )
);
