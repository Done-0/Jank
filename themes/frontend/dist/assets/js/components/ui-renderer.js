/**
 * UI 渲染组件
 * 专门负责通用UI元素渲染，如加载状态、分页等
 */

export class UiRenderer {
  /**
   * 渲染加载状态
   * @param {string} containerId - 容器ID
   */
  renderLoading(containerId) {
    const container = document.getElementById(containerId);
    if (!container) return;

    container.innerHTML = `
      <div class="flex items-center justify-center py-12">
        <div class="flex items-center space-x-2 text-gray-600 dark:text-gray-400">
          <div class="w-4 h-4 border-2 border-current border-t-transparent rounded-full animate-spin"></div>
          <span>加载中...</span>
        </div>
      </div>
    `;
  }

  /**
   * 渲染错误状态
   * @param {string} containerId - 容器ID
   * @param {string} message - 错误消息
   */
  renderError(containerId, message) {
    const container = document.getElementById(containerId);
    if (!container) return;

    container.innerHTML = `
      <div class="flex items-center justify-center py-12">
        <div class="text-center">
          <div class="text-red-500 mb-2">⚠️</div>
          <p class="text-gray-600 dark:text-gray-400">${message}</p>
          <button onclick="location.reload()" class="mt-4 px-4 py-2 bg-gray-100 dark:bg-gray-800 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors">
            重试
          </button>
        </div>
      </div>
    `;
  }

  /**
   * 渲染分页
   * @param {Object} pagination - 分页对象
   * @param {string} containerId - 容器ID
   * @param {Function} onPageChange - 页码变化回调
   */
  renderPagination(pagination, containerId, onPageChange) {
    const { currentPage, total, pageSize } = pagination;
    const container = document.getElementById(containerId);
    if (!container) return;

    const totalPages = Math.ceil(total / pageSize);
    
    if (total === 0 || total <= pageSize) {
      container.innerHTML = '';
      return;
    }

    const prevDisabled = currentPage <= 1;
    const nextDisabled = currentPage >= totalPages;

    container.innerHTML = `
      <div class="pagination">
        <div class="pagination-info">
          显示第 ${(currentPage - 1) * pageSize + 1}-${Math.min(currentPage * pageSize, total)} 条，共 ${total} 条
        </div>
        
        <div class="pagination-controls">
          <button ${prevDisabled ? 'disabled' : ''} class="pagination-btn prev-btn" data-page="${currentPage - 1}">
            上一页
          </button>
          <span class="pagination-current">第 ${currentPage} 页，共 ${totalPages} 页</span>
          <button ${nextDisabled ? 'disabled' : ''} class="pagination-btn next-btn" data-page="${currentPage + 1}">
            下一页
          </button>
        </div>
      </div>
    `;

    const paginationBtns = container.querySelectorAll('.pagination-btn');
    paginationBtns.forEach(btn => {
      btn.addEventListener('click', (e) => {
        if (btn.disabled) return;
        const page = parseInt(btn.dataset.page);
        if (page && onPageChange) {
          onPageChange(page);
        }
      });
    });
  }

}
