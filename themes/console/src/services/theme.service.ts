/**
 * 主题服务
 */

import { THEME_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  SwitchThemeRequest,
  SwitchThemeResponse,
  ListThemesRequest,
  ListThemesResponse,
  GetActiveThemeResponse,
} from "@/types";

class ThemeService {
  // ===== 主题管理 =====

  // 切换主题
  async switchTheme(request: SwitchThemeRequest): Promise<SwitchThemeResponse> {
    const response = await apiClient.post<ApiResponse<SwitchThemeResponse>>(
      THEME_ENDPOINTS.SWITCH_THEME,
      request
    );
    return response.data.data!;
  }

  // 获取当前主题
  async getActiveTheme(): Promise<GetActiveThemeResponse> {
    const response = await apiClient.get<ApiResponse<GetActiveThemeResponse>>(
      THEME_ENDPOINTS.GET_ACTIVE_THEME
    );
    return response.data.data!;
  }

  // 主题列表
  async listThemes(request: ListThemesRequest): Promise<ListThemesResponse> {
    const response = await apiClient.get<ApiResponse<ListThemesResponse>>(
      THEME_ENDPOINTS.LIST_THEMES,
      { params: request }
    );
    return response.data.data!;
  }
}

export const themeService = new ThemeService();
