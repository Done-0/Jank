/**
 * API 服务层
 */

import { POST_ENDPOINTS, CATEGORY_ENDPOINTS } from "../constants/index.js";

export class ApiService {
  /**
   * 通用 API 调用方法
   * @param {string} endpoint - API 端点
   * @param {Object} options - 请求选项
   * @returns {Promise<Object|null>} API 响应数据
   */
  async apiCall(endpoint, options = {}) {
    try {
      const response = await fetch(endpoint, {
        headers: {
          'Content-Type': 'application/json',
          ...options.headers
        },
        ...options
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const result = await response.json();
      return result.data || result;
    } catch (error) {
      console.error('API call failed:', error);
      return null;
    }
  }

  /**
   * 获取已发布文章列表
   * @param {number} pageNo - 页码
   * @param {number} pageSize - 每页数量
   * @param {number|null} categoryId - 分类ID
   * @returns {Promise<Object|null>} 文章列表数据
   */
  async getPublishedPosts(pageNo = 1, pageSize = 5, categoryId = null) {
    const params = new URLSearchParams({
      page_no: String(pageNo),
      page_size: String(pageSize)
    });
    
    if (categoryId) {
      params.append('category_id', String(categoryId));
    }
    
    return await this.apiCall(`${POST_ENDPOINTS.LIST_PUBLISHED}?${params}`);
  }

  /**
   * 获取分类列表
   * @param {number} pageNo - 页码
   * @param {number} pageSize - 每页数量
   * @param {boolean} isActive - 是否只获取启用的分类
   * @returns {Promise<Object|null>} 分类列表数据
   */
  async getCategories(pageNo = 1, pageSize = 100, isActive = true) {
    const params = new URLSearchParams({
      page_no: pageNo,
      page_size: pageSize
    });
    
    if (isActive !== null) {
      params.append('is_active', isActive);
    }
    
    return await this.apiCall(`${CATEGORY_ENDPOINTS.LIST_CATEGORIES}?${params}`);
  }

  /**
   * 获取文章详情
   * @param {string} id - 文章ID
   * @returns {Promise<Object|null>} 文章详情数据
   */
  async getPost(id) {
    const params = new URLSearchParams({ id: String(id) });
    return await this.apiCall(`${POST_ENDPOINTS.GET_POST}?${params}`);
  }
}
