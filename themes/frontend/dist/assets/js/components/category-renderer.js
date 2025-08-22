/**
 * 分类渲染组件
 */

export class CategoryRenderer {
  /**
   * 渲染分类按钮列表
   * @param {Array} categories - 分类数组
   * @param {string} containerId - 容器ID
   * @param {string} currentCategoryId - 当前选中的分类ID
   */
  renderCategoryButtons(categories, activeCategoryId = null) {
    const container = document.getElementById('category-buttons');
    if (!container) return;

    if (!categories || categories.length === 0) {
      container.innerHTML = '<p class="text-sm text-gray-500 dark:text-gray-400">No categories</p>';
      return;
    }

    // 添加"All"按钮
    const allButton = this.createCategoryButton({ id: null, name: 'All' }, activeCategoryId === null);
    
    // 渲染分类按钮
    const categoryButtons = categories.map(category => 
      this.createCategoryButton(category, category.id === activeCategoryId)
    ).join('');

    container.innerHTML = allButton + categoryButtons;
  }

  /**
   * 渲染分类列表
   * @param {Array} categories - 分类数组
   * @param {string} containerId - 容器ID
   */
  renderCategories(categories, containerId) {
    const container = document.getElementById(containerId);
    if (!container) return;

    if (!categories || categories.length === 0) {
      container.innerHTML = this.renderEmptyState('暂无分类');
      return;
    }

    const categoriesHtml = categories.map(category => this.createCategoryCard(category)).join('');
    container.innerHTML = `<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">${categoriesHtml}</div>`;
  }

  /**
   * 渲染分类过滤器
   * @param {Array} categories - 分类列表
   * @param {string} containerId - 容器ID
   * @param {string|null} currentCategory - 当前选中的分类ID
   */
  renderCategoryFilter(categories, containerId, currentCategory = null) {
    const container = document.getElementById(containerId);
    if (!container) return;

    if (!categories || categories.length === 0) {
      container.innerHTML = '';
      return;
    }

    const options = [
      `<option value="">全部分类</option>`,
      ...categories.map(category => 
        `<option value="${category.id}" ${category.id === currentCategory ? 'selected' : ''}>${category.name}</option>`
      )
    ].join('');

    container.innerHTML = `
      <select id="category-filter" class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        ${options}
      </select>
    `;
  }

  /**
   * 创建分类按钮
   * @param {Object} category - 分类对象
   * @param {boolean} isActive - 是否为当前选中状态
   * @returns {string} HTML字符串
   */
  createCategoryButton(category, isActive = false) {
    const activeClass = isActive ? ' active' : '';
    
    return `
      <button class="category-button${activeClass}" data-category-id="${category.id}">
        ${category.name}
      </button>
    `;
  }

  /**
   * 创建分类卡片
   * @param {Object} category - 分类对象
   * @returns {string} HTML字符串
   */
  createCategoryCard(category) {
    return `
      <div class="category-card bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600 transition-colors cursor-pointer p-6" data-category-id="${category.id}">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">
          ${category.name}
        </h3>
        ${category.description ? `<p class="text-gray-600 dark:text-gray-400 text-sm mb-4">${category.description}</p>` : ''}
        <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-500">
          <span>${category.post_count || 0} 篇文章</span>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </div>
      </div>
    `;
  }

  /**
   * 渲染空状态
   * @param {string} message - 空状态消息
   * @returns {string} 空状态HTML
   */
  renderEmptyState(message) {
    return `
      <div class="flex items-center justify-center py-12">
        <div class="text-center">
          <div class="text-gray-400 mb-2">🏷️</div>
          <p class="text-gray-600 dark:text-gray-400">${message}</p>
        </div>
      </div>
    `;
  }
}
