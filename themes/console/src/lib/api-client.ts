/**
 * HTTP 客户端配置
 */

import { default as axios } from "axios";

import { USER_ENDPOINTS } from "@/api/endpoints";
import { AUTH_ROUTES } from "@/constants";
import { STORAGE_KEYS } from "@/constants/storage";
import { useAuthStore } from "@/stores/auth.store";
import { useUserStore } from "@/stores/user.store";

import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";

export interface ApiConfig {
  baseURL: string;
  timeout: number;
}

export class ApiClient {
  private readonly instance: AxiosInstance;
  private isRefreshing = false; // Token 刷新状态
  private failedQueue: Array<{
    // 失败请求队列
    resolve: (value?: unknown) => void;
    reject: (reason?: unknown) => void;
    config?: AxiosRequestConfig; // 保存原始请求配置
  }> = [];

  constructor(config: ApiConfig) {
    this.instance = axios.create(config);
    this.setupInterceptors();
  }

  // ===== 私有方法 =====

  // 处理失败请求队列
  private processQueue(error: unknown, newToken?: string) {
    this.failedQueue.forEach(({ resolve, reject, config }) => {
      if (error) {
        reject(error);
      } else {
        // 更新请求的认证头为新token
        if (newToken && config) {
          if (!config.headers) {
            config.headers = {};
          }
          config.headers.Authorization = `Bearer ${newToken}`;
        }
        resolve(config ? this.instance(config) : undefined);
      }
    });
    this.failedQueue = [];
  }

  // 设置拦截器
  private setupInterceptors() {
    // 请求拦截器 - 添加认证头
    this.instance.interceptors.request.use((config) => {
      const accessToken = localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN);

      if (accessToken) config.headers.Authorization = `Bearer ${accessToken}`;

      return config;
    });

    // 响应拦截器 - 处理认证错误
    this.instance.interceptors.response.use(
      (response: AxiosResponse) => {
        return response;
      },
      async (error) => {
        const originalRequest = error.config;

        // 处理 401 认证失败
        if (error.response?.status === 401 && !originalRequest._retry) {
          if (this.isRefreshing) {
            return new Promise((resolve, reject) => {
              this.failedQueue.push({
                resolve,
                reject,
                config: originalRequest,
              });
            });
          }

          originalRequest._retry = true;
          this.isRefreshing = true;

          const refreshToken = localStorage.getItem(STORAGE_KEYS.REFRESH_TOKEN);
          if (!refreshToken) {
            this.processQueue(error);
            this.redirectToLogin();
            return Promise.reject(error);
          }

          try {
            // 调用后端刷新token接口
            const refreshResponse = await this.instance.post(
              USER_ENDPOINTS.REFRESH_TOKEN,
              {
                refresh_token: refreshToken,
              }
            );
            const { access_token, refresh_token: newRefreshToken } =
              refreshResponse.data.data;
            useAuthStore
              .getState()
              .refreshTokens(access_token, newRefreshToken);
            originalRequest.headers.Authorization = `Bearer ${access_token}`;
            this.processQueue(null, access_token);
            return this.instance(originalRequest);
          } catch (refreshError) {
            this.processQueue(refreshError);
            this.redirectToLogin();
            return Promise.reject(refreshError);
          } finally {
            this.isRefreshing = false;
          }
        }

        return Promise.reject(error);
      }
    );
  }

  // 重定向到登录页
  private redirectToLogin() {
    // 清理认证状态（不重定向）
    useAuthStore.getState().clearAuth();
    // 清理用户信息
    useUserStore.getState().clearUser();
    // 重定向到登录页
    window.location.href = AUTH_ROUTES.LOGIN;
  }

  // ===== HTTP 方法 =====

  get<T>(url: string, config?: AxiosRequestConfig) {
    return this.instance.get<T>(url, config);
  }

  post<T>(url: string, data?: unknown, config?: AxiosRequestConfig) {
    return this.instance.post<T>(url, data, config);
  }

  put<T>(url: string, data?: unknown, config?: AxiosRequestConfig) {
    return this.instance.put<T>(url, data, config);
  }

  delete<T>(url: string, config?: AxiosRequestConfig) {
    return this.instance.delete<T>(url, config);
  }
}

// ===== 实例创建 =====

export const apiClient = new ApiClient({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000,
});
